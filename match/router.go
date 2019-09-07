package match

import (
	"compelo/models"
	"compelo/rest"
	"net/http"
	"strconv"
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

	// TODO verify that game ID belongs to project.
	// TODO verify that player IDs belong to project.

	var m *models.Match
	err := c.Bind(&body)
	if err == nil {
		m, err = r.s.CreateMatch(CreateMatchParameter{
			Date:            time.Now(),
			GameID:          body.GameID,
			Teams:           body.Teams,
			PlayerTeamMap:   body.PlayerTeamMap,
			TeamScoreMap:    body.TeamScoreMap,
			WinnerMatchTeam: body.WinnerMatchTeam,
		})
	}
	rest.WriteOkResponse(m, err, c)
}

func (r *Router) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)

	var m CompleteMatch
	if err == nil {
		m, err = r.s.LoadByID(uint(id))
	}
	rest.WriteOkResponse(m, err, c)
}

func (r *Router) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, r.s.LoadMatches())
}
