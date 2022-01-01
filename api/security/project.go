package security

import (
	"errors"
	"net/http"

	"compelo/api/handler"
	"compelo/api/json"

	"github.com/go-chi/chi"
)

var errProjectForbidden = errors.New("not your project")

func (sec *Security) ProjectSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		claims := mustLoadClaimsFromContext(r)

		if chi.URLParam(r, handler.ProjectGUID) != claims.ProjectGUID {
			json.WriteErrorResponse(w, http.StatusForbidden, errProjectForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
