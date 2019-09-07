package api

import (
	"log"

	"github.com/gin-gonic/gin"

	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
)

type Api struct {
	projectRouter *project.Router
	playerRouter  *player.Router
	matchRouter   *match.Router
	gameRouter    *game.Router
}

func Serve(projectRouter *project.Router, playerRouter *player.Router, matchRouter *match.Router, gameRouter *game.Router) {
	api := &Api{
		projectRouter: projectRouter,
		playerRouter:  playerRouter,
		matchRouter:   matchRouter,
		gameRouter:    gameRouter,
	}

	r := gin.Default()

	r.POST("projects", api.projectRouter.Post)
	r.GET("projects", api.projectRouter.GetAll)

	// Sub-router for project specific activities.
	p := r.Group("project")
	p.Use(ProjectMiddleware)

	p.POST("players", api.playerRouter.Post)
	p.GET("players", api.playerRouter.GetAll)
	p.POST("matches", api.matchRouter.Post)
	p.GET("matches", api.matchRouter.GetAll)
	p.GET("matches/:id", api.matchRouter.GetByID)
	p.POST("games", api.gameRouter.Post)
	p.GET("games", api.gameRouter.GetAll)
	log.Fatal(r.Run())
}

func ProjectMiddleware(c *gin.Context) {
	c.Set("projectID", 1)

	// TODO
	// - read from token
	// - ensure it exists, respond with 404 else
}
