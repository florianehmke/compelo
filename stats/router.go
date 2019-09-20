package stats

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"compelo/game"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

func (r *Router) GetAll(c *gin.Context) {
	g := c.MustGet(game.Key).(game.Game)

	players, err := r.s.LoadPlayerStatsByGameID(g.ID)
	if err == nil {
		c.JSON(http.StatusOK, players)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
