package match

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"

	"compelo"
	"compelo/game"
	"compelo/project"
)

type Router struct {
	s *Service
}

func NewRouter(s *Service) *Router {
	return &Router{s}
}

type CreateMatchParameter struct {
	Date   time.Time
	GameID uint

	Teams []struct {
		PlayerIDs []int `json:"playerIds" binding:"required"`
		Score     int   `json:"score" binding:"required"`
		Winner    bool  `json:"winner" binding:"required"`
	} `json:"teams" binding:"required"`
}

func (p *CreateMatchParameter) validate() error {
	return nil
}

func (r *Router) Post(c *gin.Context) {
	g := c.MustGet(game.Key).(compelo.Game)
	p := c.MustGet(project.Key).(compelo.Project)

	var param CreateMatchParameter
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

	m, err := r.s.CreateMatch(param, g)
	if err == nil {
		c.JSON(http.StatusCreated, m)
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
	}
}

func (r *Router) GetByID(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}

	m, err := r.s.LoadByID(uint(id))
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
