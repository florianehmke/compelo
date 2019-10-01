package security

import (
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"

	"compelo/internal"
	"compelo/pkg/json"
)

type JWT struct {
	svc *compelo.Service
	cfg *jwtConfig
}

func NewJWT(svc *compelo.Service, timeoutSec int, secretKey string) *JWT {
	cfg := &jwtConfig{
		timeout:    time.Second * time.Duration(timeoutSec),
		maxRefresh: time.Hour * 7 * 24,
		secretKey:  []byte(secretKey),
	}
	cfg.initialize()
	return &JWT{
		svc: svc,
		cfg: cfg,
	}
}

type jwtConfig struct {
	timeout    time.Duration
	maxRefresh time.Duration
	secretKey  []byte

	jwtAuth  *jwtauth.JWTAuth
	verifier func(http.Handler) http.Handler
}

func (c *jwtConfig) initialize() {
	c.jwtAuth = jwtauth.New("HS256", c.secretKey, nil)
	c.verifier = jwtauth.Verifier(c.jwtAuth)
}

type Claims struct {
	ProjectID uint `json:"projectId"`
	jwt.StandardClaims
}

type LoginRequest struct {
	ProjectName string `json:"projectName"`
	Password    string `json:"password"`
}

type LoginResponse struct {
	Token  string `json:"token"`
	Expire string `json:"expire"`
}

func (j *JWT) Authenticator(handler http.Handler) http.Handler {
	return jwtauth.Authenticator(handler)
}

func (j *JWT) Verifier(handler http.Handler) http.Handler {
	return j.cfg.verifier(handler)
}

func (j *JWT) Login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	// Get the JSON body and decode into credentials
	err := json.Unmarshal(r.Body, &login)
	if err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}

	project, err := j.svc.AuthorizeProject(login.ProjectName, login.Password)
	if err != nil {
		json.Error(w, http.StatusUnauthorized, err)
		return
	}

	expirationTime := time.Now().Add(j.cfg.timeout)
	claims := &Claims{
		ProjectID: project.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.cfg.secretKey)
	if err != nil {
		json.Error(w, http.StatusInternalServerError, err)
		return
	}

	json.Write(w, http.StatusOK, LoginResponse{
		Token:  tokenString,
		Expire: expirationTime.Format(time.RFC3339),
	})
}

func (j *JWT) Refresh(w http.ResponseWriter, r *http.Request) {
	tknStr := jwtauth.TokenFromHeader(r)
	if tknStr == "" {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
		return j.cfg.secretKey, nil
	})
	if !tkn.Valid {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	_, err = j.svc.LoadProjectByID(claims.ProjectID)
	if err != nil {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}

	if time.Now().Sub(time.Unix(claims.IssuedAt, 0)) > j.cfg.maxRefresh {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	expirationTime := time.Now().Add(j.cfg.timeout)
	claims.ExpiresAt = expirationTime.Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(j.cfg.secretKey)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	json.Write(w, http.StatusOK, LoginResponse{
		Token:  tokenString,
		Expire: expirationTime.Format(time.RFC3339),
	})
}
