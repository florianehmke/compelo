package security

import (
	"errors"
	"net/http"

	"github.com/go-chi/chi"

	"compelo/internal/api/handler"
	"compelo/pkg/json"
)

var projectForbidden = errors.New("not your project")

func (sec *Security) ProjectSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		param := chi.URLParam(r, handler.ProjectID)
		claim := r.Context().Value(ClaimsKey).(Claims).ProjectID

		if param != claim {
			json.Error(w, http.StatusForbidden, projectForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
