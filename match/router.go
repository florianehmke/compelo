package match

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"time"

	"compelo"
	"compelo/game"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

func (r *Router) Post(c *gin.Context) {
	var param CreateMatchParameter

	// TODO verify that player IDs belong to project.
	g := c.MustGet(game.Key).(compelo.Game)

	var m compelo.Match
	err := c.Bind(&param)
	if err == nil {
		param.GameID = g.ID
		param.Date = time.Now()
		m, err = r.s.CreateMatch(param)
	}

	if err == nil {
		c.JSON(http.StatusCreated, m)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	var m Match
	if err == nil {
		m, err = r.s.LoadByID(uint(id))
	}

	if err == nil {
		c.JSON(http.StatusOK, m)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	g := c.MustGet(game.Key).(compelo.Game)

	matches, err := r.s.LoadByGameID(g.ID)
	if err == nil {
		c.JSON(http.StatusOK, matches)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
