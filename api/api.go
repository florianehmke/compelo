package api

import (
	"compelo/project"
	"github.com/gin-gonic/gin"
	"log"
)

type Api struct {
	projectRouter *project.Router
}

func Serve(projectRouter *project.Router) {
	api := &Api{
		projectRouter: projectRouter,
	}

	r := gin.Default()
	r.POST("projects", api.projectRouter.Post)
	r.GET("projects", api.projectRouter.GetAll)

	log.Fatal(r.Run())
}
