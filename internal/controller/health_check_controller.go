package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type healthCheckController struct{}

func (c *healthCheckController) SetRoute(r gin.IRouter) {
	r.GET("/health", c.Handle)
}

func (c *healthCheckController) Handle(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "ok",
	})
}

func NewHealthCheckController() Controller {
	return &healthCheckController{}
}
