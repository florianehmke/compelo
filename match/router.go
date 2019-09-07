package match

import (
	"net/http"
	"time"

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
		GameID          uint          `json:"gameId" binding:"required"`
		Teams           uint          `json:"teams" binding:"required"`
		PlayerTeamMap   map[uint]uint `json:"playerTeamMap" binding:"required"`
		TeamScoreMap    map[uint]int  `json:"teamScoreMap" binding:"required"`
		WinnerMatchTeam uint          `json:"winnerMatchTeamId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		p, err := r.s.CreateMatch(CreateMatchParameter{
			Date:            time.Now(),
			GameID:          body.GameID,
			Teams:           body.Teams,
			PlayerTeamMap:   body.PlayerTeamMap,
			TeamScoreMap:    body.TeamScoreMap,
			WinnerMatchTeam: body.WinnerMatchTeam,
		})
		if err == nil {
			c.JSON(http.StatusOK, &p)
		} else {
			c.JSON(http.StatusBadRequest, err)
		}
	}
}

func (r *Router) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, r.s.LoadMatches())
}
