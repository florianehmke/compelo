package api

import (
	"compelo/game"
	"github.com/gin-gonic/gin"
	"log"

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

	r.POST("players", api.playerRouter.Post)
	r.GET("players", api.playerRouter.GetAll)

	r.POST("matches", api.matchRouter.Post)
	r.GET("matches", api.matchRouter.GetAll)

	r.POST("games", api.gameRouter.Post)
	r.GET("games", api.gameRouter.GetAll)
	log.Fatal(r.Run())
}
