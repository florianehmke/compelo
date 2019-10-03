package security

import (
	"errors"
	"net/http"

	"compelo/internal/api/handler"
	"compelo/pkg/json"
)

var projectForbidden = errors.New("not your project")

func (sec *Security) ProjectSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := handler.MustLoadProjectFromContext(r)
		claims := mustLoadClaimsFromContext(r)

		if project.ID != claims.ProjectID {
			json.Error(w, http.StatusForbidden, projectForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
