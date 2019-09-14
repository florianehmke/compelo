package match

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

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
	var body struct {
		Teams         uint          `json:"teams" binding:"required"`
		WinningTeam   uint          `json:"winningTeam" binding:"required"`
		TeamPlayerMap map[uint]uint `json:"teamPlayerMap" binding:"required"`
		TeamScoreMap  map[uint]int  `json:"teamScoreMap" binding:"required"`
	}

	// TODO verify that player IDs belong to project.
	g := c.MustGet(game.Key).(compelo.Game)

	var m *compelo.Match
	err := c.Bind(&body)
	if err == nil {
		m, err = r.s.CreateMatch(CreateMatchParameter{
			Date:          time.Now(),
			GameID:        g.ID,
			Teams:         body.Teams,
			TeamPlayerMap: body.TeamPlayerMap,
			TeamScoreMap:  body.TeamScoreMap,
			WinningTeam:   body.WinningTeam,
		})
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

	var m CompleteMatch
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
	c.JSON(http.StatusOK, r.s.LoadMatches())
}
