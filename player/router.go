package player

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"compelo/project"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

type createPlayerParameter struct {
	Name string `json:"name" binding:"required"`
}

func (r *Router) Post(c *gin.Context) {
	p := c.MustGet(project.Key).(project.Project)

	var param createPlayerParameter
	err := c.Bind(&param)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}

	player, err := r.s.CreatePlayer(p.ID, param.Name)
	if err == nil {
		c.JSON(http.StatusCreated, player)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(project.Project)

	players, err := r.s.LoadPlayersByProjectID(p.ID)
	if err == nil {
		c.JSON(http.StatusOK, players)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
