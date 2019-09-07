package player

import (
	"net/http"

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
		Name      string `json:"name" binding:"required"`
		ProjectID uint   `json:"projectId" binding:"required"`
	}

	if err := c.Bind(&body); err != nil {
		c.JSON(http.StatusBadRequest, err)
	} else {
		p, err := r.s.CreatePlayer(body.ProjectID, body.Name)
		if err == nil {
			c.JSON(http.StatusOK, &p)
		} else {
			c.JSON(http.StatusBadRequest, err)
		}
	}
}

func (r *Router) GetAll(c *gin.Context) {
	c.JSON(http.StatusOK, r.s.LoadPlayers())
}
