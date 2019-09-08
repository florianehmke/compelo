package player

import (
	"compelo/project"
	"github.com/gin-gonic/gin"

	"compelo/models"
	"compelo/rest"
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

	var p *models.Player
	err := c.Bind(&body)
	if err == nil {
		p, err = r.s.CreatePlayer(uint(c.GetInt("projectID")), body.Name)
	}
	rest.WriteOkResponse(p, err, c)
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(models.Project)
	games, err := r.s.LoadPlayersByProjectID(p.ID)
	rest.WriteOkResponse(games, err, c)
}
