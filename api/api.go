package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
)

func Serve(
	projectRouter *project.Router,
	playerRouter *player.Router,
	matchRouter *match.Router,
	gameRouter *game.Router,
) {
	r := gin.Default()

	r.POST("create-project", projectRouter.CreateProject)
	r.POST("select-project", projectRouter.SelectProject)
	r.GET("projects", projectRouter.GetAll)

	// Sub-router for project specific activities.
	p := r.Group("project")
	p.Use(projectRouter.Middleware())

	p.POST("players", playerRouter.Post)
	p.GET("players", playerRouter.GetAll)
	p.POST("matches", matchRouter.Post)
	p.GET("matches", matchRouter.GetAll)
	p.GET("matches/:id", matchRouter.GetByID)
	p.POST("games", gameRouter.Post)
	p.GET("games", gameRouter.GetAll)
	log.Fatal(r.Run())
}
