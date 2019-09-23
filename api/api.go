package api

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"compelo/auth"
	"compelo/db"
	"compelo/frontend"
	"compelo/game"
	"compelo/match"
	"compelo/player"
	"compelo/project"
	"compelo/stats"
)

func Setup(dbPath string, secret string, dev bool) *gin.Engine {
	database := db.New(dbPath)

	projectService := project.NewService(database)
	projectRouter := project.NewRouter(projectService)

	playerService := player.NewService(database)
	playerRouter := player.NewRouter(playerService)

	gameService := game.NewService(database)
	gameRouter := game.NewRouter(gameService)

	matchService := match.NewService(database, playerService, gameService)
	matchRouter := match.NewRouter(matchService)

	statsService := stats.NewService(database, playerService)
	statsRouter := stats.NewRouter(statsService)

	authService := auth.NewService(auth.DefaultConfig().WithSecret(secret), projectService)
	authMW := authService.Middleware()

	if !dev {
		gin.SetMode(gin.ReleaseMode)
	}
	engine := gin.Default()
	if dev {
		engine.Use(createCORSMiddleware())
	}

	// Frontend
	engine.StaticFS("/app", frontend.Frontend)
	engine.NoRoute(frontendHandler)

	// API
	r := engine.Group("/api")
	r.GET("/refresh", authMW.RefreshHandler)

	// Projects
	r.POST("/create-project", projectRouter.CreateProject)
	r.POST("/select-project", authMW.LoginHandler)
	r.GET("/projects", projectRouter.GetAll)

	// Selected Project
	p := r.Group("/project")
	p.Use(authMW.MiddlewareFunc())
	p.POST("/players", playerRouter.Post)
	p.GET("/players", playerRouter.GetAll)
	p.POST("/games", gameRouter.Post)
	p.GET("/games", gameRouter.GetAll)

	// Selected Game
	g := p.Group("/games/:" + game.IDParam)
	g.Use(gameRouter.Middleware)
	g.POST("/matches", matchRouter.Post)
	g.GET("/matches", matchRouter.GetAll)
	g.GET("/players", statsRouter.GetAll)

	return engine
}

func frontendHandler(c *gin.Context) {
	c.Request.URL.Path = "/" // -> let frontend handle route
	http.FileServer(frontend.Frontend).ServeHTTP(c.Writer, c.Request)
}

func createCORSMiddleware() gin.HandlerFunc {
	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:4200"}
	config.AllowMethods = []string{"OPTIONS", "GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"}
	config.AllowHeaders = []string{"authorization", "content-type"}
	config.MaxAge = 12 * time.Hour
	return cors.New(config)
}
