package security

import (
	"errors"
	"net/http"

	"compelo/api/handler"
	"compelo/api/json"
)

var errGameForbidden = errors.New("game does not belong to your project")

func (sec *Security) GameSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		game := handler.MustLoadGameFromContext(r)
		claims := mustLoadClaimsFromContext(r)

		if game.ProjectGUID != claims.ProjectGUID {
			json.WriteErrorResponse(w, http.StatusForbidden, errGameForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
