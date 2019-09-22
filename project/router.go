package project

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

const (
	// Key identifies the project inside the gin.Context
	Key = "project"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

type createProjectParameter struct {
	Name     string `json:"name" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func (r *Router) CreateProject(c *gin.Context) {
	var param createProjectParameter
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	p, err := r.s.CreateProject(param.Name, param.Password)
	if err == nil {
		c.JSON(http.StatusCreated, p)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, r.s.LoadProjects())
}
