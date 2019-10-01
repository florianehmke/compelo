package handler

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/go-chi/jwtauth"
)

var (
	jwtKey    = []byte("my_secret_key")
	tokenAuth = jwtauth.New("HS256", jwtKey, nil)

	Verifier      = jwtauth.Verifier(tokenAuth)
	Authenticator = jwtauth.Authenticator
)

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

func (h *Handler) Login(w http.ResponseWriter, r *http.Request) {
	var login LoginRequest
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&login)
	if err != nil {
		writeError(w, http.StatusBadRequest, err)
		return
	}

	project, err := h.svc.AuthorizeProject(login.ProjectName, login.Password)
	if err != nil {
		writeError(w, http.StatusUnauthorized, err)
		return
	}

	expirationTime := time.Now().Add(5 * time.Minute)
	claims := &Claims{
		ProjectID: project.ID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err)
		return
	}

	writeJSON(w, http.StatusOK, LoginResponse{
		Token:  tokenString,
		Expire: expirationTime.Format(time.RFC3339),
	})
}
