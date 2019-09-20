package game

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"compelo/project"
)

const (
	IDParam = "gameId"

	// Key identifies the game inside the gin.Context
	Key = "game"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

type createGameParameter struct {
	Name string `json:"name" binding:"required"`
}

func (r *Router) Post(c *gin.Context) {
	p := c.MustGet(project.Key).(project.Project)

	var param createGameParameter
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	g, err := r.s.CreateGame(p.ID, param.Name)
	if err == nil {
		c.JSON(http.StatusCreated, g)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(project.Project)

	games, err := r.s.LoadGamesByProjectID(p.ID)
	if err == nil {
		c.JSON(http.StatusOK, games)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) Middleware(c *gin.Context) {
	p := c.MustGet(project.Key).(project.Project)

	gameID := c.Param(IDParam)
	id, err := strconv.Atoi(gameID)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		c.Abort()
	}

	g, err := r.s.LoadGameByID(uint(id))
	if err != nil {
		c.JSON(http.StatusNotFound, err)
		c.Abort()
	}
	if p.ID != g.ProjectID {
		c.JSON(http.StatusForbidden, gin.H{"message:": "not your game"})
		c.Abort()
	}

	c.Set(Key, g)
	c.Next()
}
