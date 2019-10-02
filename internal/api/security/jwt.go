package security

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/brianvoe/sjwt"

	"compelo/internal"
	"compelo/pkg/json"
)

const ClaimsKey = "claims"

type JWT struct {
	svc *compelo.Service

	timeout    time.Duration
	maxRefresh int64
	secretKey  []byte
}

func NewJWT(svc *compelo.Service, timeoutSec int, secretKey string) *JWT {
	return &JWT{
		svc:        svc,
		timeout:    time.Second * time.Duration(timeoutSec),
		maxRefresh: 60 * 7 * 24,
		secretKey:  []byte(secretKey),
	}
}

type Claims struct {
	ProjectID uint `json:"projectId"`
}

type AuthRequest struct {
	ProjectName string `json:"projectName"`
	Password    string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (j *JWT) Login(w http.ResponseWriter, r *http.Request) {
	var login AuthRequest
	err := json.Unmarshal(r.Body, &login)
	if err != nil {
		json.Error(w, http.StatusBadRequest, err)
		return
	}
	project, err := j.svc.AuthenticateProject(login.ProjectName, login.Password)
	if err != nil {
		json.Error(w, http.StatusUnauthorized, err)
		return
	}

	claims := sjwt.New()
	claims.Set("projectId", project.ID)

	now := time.Now()
	claims.SetExpiresAt(now.Add(j.timeout))
	claims.SetIssuedAt(now)

	json.Write(w, http.StatusOK, AuthResponse{
		Token: claims.Generate(j.secretKey),
	})
}

func (j *JWT) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenStr := tokenFromHeader(r)
	if valid := sjwt.Verify(tokenStr, j.secretKey); !valid {
		json.Error(w, http.StatusUnauthorized, sjwt.ErrTokenInvalid)
		return
	}
	rawClaims, err := sjwt.Parse(tokenStr)
	if err != nil {
		json.Error(w, http.StatusUnauthorized, err)
		return
	}
	issuedAt, err := rawClaims.GetIssuedAt()
	if err != nil {
		json.Error(w, http.StatusUnauthorized, err)
		return
	}
	if (time.Now().Unix() - issuedAt) > j.maxRefresh {
		json.Error(w, http.StatusUnauthorized, errors.New("max refresh time exceeded"))
		return
	}
	rawClaims.SetExpiresAt(time.Now().Add(j.timeout))
	json.Write(w, http.StatusOK, AuthResponse{
		Token: rawClaims.Generate(j.secretKey),
	})
}

// VerifyToken verifies, parses and validates the jwt.
//
// 1. Extract bearer token from request headers.
// 2. Verify that the token signature matches.
// 3. Parse the token's claims.
// 4. Validate the token's claims (checks for expiration).
// 5. Populate claims struct and put it into request context.
func (j *JWT) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := tokenFromHeader(r)
		if valid := sjwt.Verify(tokenStr, j.secretKey); !valid {
			json.Error(w, http.StatusUnauthorized, sjwt.ErrTokenInvalid)
			return
		}
		rawClaims, err := sjwt.Parse(tokenStr)
		if err != nil {
			json.Error(w, http.StatusUnauthorized, err)
			return
		}
		if err := rawClaims.Validate(); err != nil {
			json.Error(w, http.StatusUnauthorized, err)
			return
		}
		var claims Claims
		if err := rawClaims.ToStruct(&claims); err != nil {
			json.Error(w, http.StatusUnauthorized, err)
			return
		}
		ctx := context.WithValue(r.Context(), ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func tokenFromHeader(r *http.Request) string {
	bearer := r.Header.Get("Authorization")
	if len(bearer) > 7 && strings.ToUpper(bearer[0:6]) == "BEARER" {
		return bearer[7:]
	}
	return ""
}
