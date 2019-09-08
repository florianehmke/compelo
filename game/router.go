package game

import (
	"compelo"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"compelo/project"
	"compelo/rest"
)

const IDParam = "gameId"

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
	var g *compelo.Game
	if err == nil {
		p := c.MustGet(project.Key).(compelo.Project)
		g, err = r.s.CreateGame(p.ID, body.Name)
	}
	rest.WriteCreatedResponse(g, err, c)
}

func (r *Router) GetAll(c *gin.Context) {
	p := c.MustGet(project.Key).(compelo.Project)
	games, err := r.s.LoadGamesByProjectID(p.ID)
	rest.WriteOkResponse(games, err, c)
}

func (r *Router) Middleware(c *gin.Context) {
	gameID := c.Param(IDParam)
	id, err := strconv.Atoi(gameID)

	if err == nil {
		g, err := r.s.LoadGameByID(uint(id))
		if err != nil {
			c.JSON(http.StatusNotFound, err)
			c.Abort()
		}

		p := c.MustGet(project.Key).(compelo.Project)
		if p.ID != g.ProjectID {
			c.JSON(http.StatusForbidden, gin.H{"message:": "not your game"})
			c.Abort()
		}
		c.Next()
	}
}
