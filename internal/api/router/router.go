package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/cors"

	"compelo/frontend"
	"compelo/internal/api/handler"
	"compelo/internal/api/security"
)

func New(h *handler.Handler, s *security.JWT) http.Handler {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))
	r.Use(corsMiddleware().Handler)

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", s.Login)
		r.Post("/refresh", s.Refresh)
		r.Post("/projects", h.CreateProject)
		r.Get("/projects", h.GetAllProjects)
		r.Route("/projects/{"+handler.ProjectID+"}", func(r chi.Router) {
			r.Use(s.Verifier)
			r.Use(s.Authenticator)
			r.Use(h.ProjectCtx)
			r.Post("/players", h.CreatePlayer)
			r.Get("/players", h.GetAllPlayers)
			r.Post("/games", h.CreateGame)
			r.Get("/games", h.GetAllGames)
			r.Route("/games/{"+handler.GameID+"}", func(r chi.Router) {
				r.Use(h.GameCtx)
				r.Post("/matches", h.CreateMatch)
				r.Get("/matches", h.GetAllMatches)
				r.Get("/player-stats", h.GetAllPlayerStats)
			})
		})
	})
	r.HandleFunc("/*", frontendHandler)
	return r
}

func frontendHandler(w http.ResponseWriter, r *http.Request) {
	f, err := frontend.Frontend.Open(r.URL.Path)
	if err != nil {
		r.URL.Path = "/"
	} else {
		if err := f.Close(); err != nil {
			panic(err)
		}
	}
	http.FileServer(frontend.Frontend).ServeHTTP(w, r)
}

func corsMiddleware() *cors.Cors {
	return cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:4200"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: true,
		MaxAge:           300,
	})
}
