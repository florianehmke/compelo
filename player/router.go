package player

import (
	"compelo"
	"compelo/project"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

func (r *Router) Post(c *gin.Context) {
	var body struct {
		Name string `json:"name" binding:"required"`
	}
	err := c.Bind(&body)
	var player *compelo.Player
	if err == nil {
		p := c.MustGet(project.Key).(compelo.Project)
		player, err = r.s.CreatePlayer(p.ID, body.Name)
	}

	if err == nil {
		c.JSON(http.StatusCreated, player)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(compelo.Project)
	players, err := r.s.LoadPlayersByProjectID(p.ID)

	if err == nil {
		c.JSON(http.StatusOK, players)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
