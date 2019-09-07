package game

import (
	"compelo/models"
	"compelo/rest"
	"github.com/gin-gonic/gin"
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

	var p *models.Game
	err := c.Bind(&body)
	if err == nil {
		p, err = r.s.CreateGame(uint(c.GetInt("projectID")), body.Name)
	}
	rest.WriteOkResponse(p, err, c)
}

func (r *Router) GetAll(c *gin.Context) {
	projectID := c.GetInt("projectID")
	games, err := r.s.LoadGamesByProjectID(uint(projectID))
	rest.WriteOkResponse(games, err, c)
}
