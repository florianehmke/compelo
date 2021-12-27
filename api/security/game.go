package security

// var gameForbidden = errors.New("game does not belong to your project")

// func (sec *Security) GameSecurity(next http.Handler) http.Handler {
// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// 		game := handler.MustLoadGameFromContext(r)
// 		claims := mustLoadClaimsFromContext(r)

// 		if game.ProjectID != claims.ProjectID {
// 			json.Error(w, http.StatusForbidden, gameForbidden)
// 			return
// 		}

// 		next.ServeHTTP(w, r)
// 	})
// }
