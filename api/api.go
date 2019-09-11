package api

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"compelo/db"
	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
)

func Setup(dbPath string) *gin.Engine {
	database := db.New(dbPath)

	projectService := project.NewService(database)
	playerService := player.NewService(database)
	gameService := game.NewService(database)
	matchService := match.NewService(database, playerService, gameService)

	return createRouter(
		project.NewRouter(projectService),
		player.NewRouter(playerService),
		match.NewRouter(matchService),
		game.NewRouter(gameService),
	)
}

func createRouter(
	projectRouter *project.Router,
	playerRouter *player.Router,
	matchRouter *match.Router,
	gameRouter *game.Router,
) *gin.Engine {
	engine := gin.Default()
	engine.Use(createCORSMiddleware())

	r := engine.Group("/api")
	r.POST("/create-project", projectRouter.CreateProject)
	r.POST("/select-project", projectRouter.SelectProject)
	r.GET("/projects", projectRouter.GetAll)

	// Sub-router for project specific activities.
	p := r.Group("/project")
	p.Use(projectRouter.Middleware())

	p.POST("/players", playerRouter.Post)
	p.GET("/players", playerRouter.GetAll)
	p.POST("/games", gameRouter.Post)
	p.GET("/games", gameRouter.GetAll)

	g := p.Group("/games/:" + game.IDParam)
	g.Use(gameRouter.Middleware)
	g.POST("/matches", matchRouter.Post)
	g.GET("/matches", matchRouter.GetAll)
	g.GET("/matches/:id", matchRouter.GetByID)

	return engine
}

func createCORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	return cors.New(config)
}
