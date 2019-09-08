package rest

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WriteOkResponse(ok interface{}, err error, c *gin.Context) {
	if err == nil {
		c.JSON(http.StatusOK, ok)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}

func WriteCreatedResponse(ok interface{}, err error, c *gin.Context) {
	if err == nil {
		c.JSON(http.StatusCreated, ok)
	} else {
		c.JSON(http.StatusBadRequest, err)
	}
}
