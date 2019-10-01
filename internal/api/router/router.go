package router

import (
	"net/http"
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

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

	r.Route("/api", func(r chi.Router) {
		r.Post("/login", s.Login)
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
	return r
}
