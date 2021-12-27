package security

import (
	"errors"
	"net/http"

	"compelo/api/handler"
	"compelo/api/json"
)

var projectForbidden = errors.New("not your project")

func (sec *Security) ProjectSecurity(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		project := handler.MustLoadProjectFromContext(r)
		claims := mustLoadClaimsFromContext(r)

		if project.GUID != claims.ProjectGUID {
			json.WriteErrorResponse(w, http.StatusForbidden, projectForbidden)
			return
		}

		next.ServeHTTP(w, r)
	})
}
