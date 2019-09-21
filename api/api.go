package api

import (
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"compelo/db"
	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
	"compelo/stats"
)

const frontendPath = "./frontend/compelo/dist/compelo/"

func Setup(dbPath string) *gin.Engine {
	database := db.New(dbPath)

	projectService := project.NewService(database)
	playerService := player.NewService(database)
	gameService := game.NewService(database)
	matchService := match.NewService(database, playerService, gameService)
	statsService := stats.NewService(database, playerService)

	return createRouter(
		project.NewRouter(projectService),
		player.NewRouter(playerService),
		match.NewRouter(matchService),
		game.NewRouter(gameService),
		stats.NewRouter(statsService),
	)
}

func createRouter(
	projectRouter *project.Router,
	playerRouter *player.Router,
	matchRouter *match.Router,
	gameRouter *game.Router,
	statsRouter *stats.Router,
) *gin.Engine {
	engine := gin.Default()
	engine.Static("/app", frontendPath)
	engine.NoRoute(func(c *gin.Context) { c.File(frontendPath + "index.html") })

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
	g.GET("/players", statsRouter.GetAll)

	return engine
}

func createCORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"authorization", "content-type"}
	config.MaxAge = 12 * time.Hour
	return cors.New(config)
}
