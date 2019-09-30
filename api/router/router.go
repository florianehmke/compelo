package router

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"

	"compelo/api/handler"
)

func New(h *handler.Handler) *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(60 * time.Second))

	r.Route("/api", func(r chi.Router) {
		r.Post("/projects", h.CreateProject)
		r.Get("/projects", h.GetAllProjects)
		r.Route("/projects/{projectID}", func(r chi.Router) {
			r.Use(h.ProjectCtx)
			r.Post("/players", h.CreatePlayer)
			r.Get("/players", h.GetAllPlayers)
			r.Post("/games", h.CreateGame)
			r.Get("/games", h.GetAllGames)
			r.Route("/games/{gameID}", func(r chi.Router) {
				r.Use(h.GameCtx)
				r.Post("/matches", h.CreateMatch)
				r.Get("/matches", h.GetAllMatches)
				r.Get("/player-stats", h.GetAllPlayerStats)
			})
		})
	})
	return r
}
