package security

import (
	"errors"
	"net/http"

	"compelo/api"
	"compelo/api/handler"
)

var projectForbidden = errors.New("not your project")

func (sec *Security) ProjectSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := handler.MustLoadProjectFromContext(r)
		claims := mustLoadClaimsFromContext(r)

		if project.GUID != claims.ProjectGUID {
			api.
				api.Error(w, http.StatusForbidden, projectForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
