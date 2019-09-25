package match

import (
	"errors"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"

	"compelo/game"
	"compelo/project"
)

const IDParam = "matchId"

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

type createMatchParameter struct {
	gameID uint
	date   time.Time

	Teams []struct {
		PlayerIDs []int `json:"playerIds" binding:"required"`
		Score     int   `json:"score" binding:"required"`

		result      string
		ratingDelta int
	} `json:"teams" binding:"required"`
}

func (p *createMatchParameter) validate() error {
	if len(p.Teams) < 2 {
		return errors.New("at least two teams required")
	}
	teamSize := len(p.Teams[0].PlayerIDs)
	for _, t := range p.Teams {
		if len(t.PlayerIDs) != teamSize {
			return errors.New("all teams need the same amount of players")
		}
	}
	playerMap := map[int]bool{}
	for _, t := range p.Teams {
		for _, pid := range t.PlayerIDs {
			if _, ok := playerMap[pid]; ok {
				return errors.New("player can only be in one team")
			}
			playerMap[pid] = true
		}
	}
	return nil
}

func (r *Router) Post(c *gin.Context) {
	g := c.MustGet(game.Key).(game.Game)
	p := c.MustGet(project.Key).(project.Project)

	var param createMatchParameter
	if err := c.Bind(&param); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	if err := param.validate(); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	// Make sure that players exist + belong to current project.
	for _, t := range param.Teams {
		for _, pid := range t.PlayerIDs {
			player, err := r.s.playerService.LoadPlayerByID(uint(pid))
			if err != nil || player.ProjectID != p.ID {
				c.JSON(http.StatusForbidden, gin.H{"message": "not your player"})
				return
			}
		}
	}

	m, err := r.s.createMatch(param, g)
	if err == nil {
		c.JSON(http.StatusCreated, m)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetAll(c *gin.Context) {
	g := c.MustGet(game.Key).(game.Game)

	matches, err := r.s.LoadMatchesByGameID(g.ID)
	if err == nil {
		c.JSON(http.StatusOK, matches)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}
