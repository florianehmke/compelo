package player

import (
	"compelo"
	"compelo/project"
	"github.com/gin-gonic/gin"

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
	err := c.Bind(&body)
	var player *compelo.Player
	if err == nil {
		p := c.MustGet(project.Key).(compelo.Project)
		player, err = r.s.CreatePlayer(p.ID, body.Name)
	}
	rest.WriteCreatedResponse(player, err, c)
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(compelo.Project)
	games, err := r.s.LoadPlayersByProjectID(p.ID)
	rest.WriteOkResponse(games, err, c)
}
