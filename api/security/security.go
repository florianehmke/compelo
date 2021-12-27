package security

import (
	"context"
	"errors"
	"net/http"
	"strings"
	"time"

	"github.com/brianvoe/sjwt"
	"golang.org/x/crypto/bcrypt"

	"compelo/api/json"
	"compelo/query"
)

const ClaimsKey = "claims"

type Security struct {
	q *query.Compelo

	timeout    time.Duration
	maxRefresh int64
	secretKey  []byte
}

func New(q *query.Compelo, timeoutSec int, secretKey string) *Security {
	return &Security{
		q:          q,
		timeout:    time.Second * time.Duration(timeoutSec),
		maxRefresh: 60 * 7 * 24,
		secretKey:  []byte(secretKey),
	}
}

type Claims struct {
	ProjectGUID string `json:"projectGuid"`
}

type AuthRequest struct {
	ProjectGUID string `json:"projectGuid"`
	Password    string `json:"password"`
}

type AuthResponse struct {
	Token string `json:"token"`
}

func (sec *Security) Login(w http.ResponseWriter, r *http.Request) {
	var login AuthRequest
	err := json.Unmarshal(r.Body, &login)
	if err != nil {
		json.WriteError(w, http.StatusBadRequest, err)
		return
	}

	project, err := sec.q.GetProjectBy(login.ProjectGUID)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	err = bcrypt.CompareHashAndPassword(project.PasswordHash, []byte(login.Password))
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}

	claims := sjwt.New()
	claims.Set("projectGuid", project.GUID)

	now := time.Now()
	claims.SetExpiresAt(now.Add(sec.timeout))
	claims.SetIssuedAt(now)

	json.Write(w, http.StatusOK, AuthResponse{
		Token: claims.Generate(sec.secretKey),
	})
}

func (sec *Security) Refresh(w http.ResponseWriter, r *http.Request) {
	tokenStr := tokenFromHeader(r)
	if valid := sjwt.Verify(tokenStr, sec.secretKey); !valid {
		json.WriteError(w, http.StatusUnauthorized, sjwt.ErrTokenInvalid)
		return
	}
	rawClaims, err := sjwt.Parse(tokenStr)
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	issuedAt, err := rawClaims.GetIssuedAt()
	if err != nil {
		json.WriteError(w, http.StatusUnauthorized, err)
		return
	}
	if (time.Now().Unix() - issuedAt) > sec.maxRefresh {
		json.WriteError(w, http.StatusUnauthorized, errors.New("max refresh time exceeded"))
		return
	}
	rawClaims.SetExpiresAt(time.Now().Add(sec.timeout))
	json.Write(w, http.StatusOK, AuthResponse{
		Token: rawClaims.Generate(sec.secretKey),
	})
}

// VerifyToken verifies, parses and validates the jwt.
//
// 1. Extract bearer token from request headers.
// 2. Verify that the token signature matches.
// 3. Parse the token's claims.
// 4. Validate the token's claims (checks for expiration).
// 5. Populate claims struct and put it into request context.
func (sec *Security) VerifyToken(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := tokenFromHeader(r)
		if valid := sjwt.Verify(tokenStr, sec.secretKey); !valid {
			json.WriteError(w, http.StatusUnauthorized, sjwt.ErrTokenInvalid)
			return
		}
		rawClaims, err := sjwt.Parse(tokenStr)
		if err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		if err := rawClaims.Validate(); err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		var claims Claims
		if err := rawClaims.ToStruct(&claims); err != nil {
			json.WriteError(w, http.StatusUnauthorized, err)
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

func mustLoadClaimsFromContext(r *http.Request) Claims {
	claims, ok := r.Context().Value(ClaimsKey).(Claims)
	if !ok {
		panic("claims must be set in context")
	}
	return claims
}
