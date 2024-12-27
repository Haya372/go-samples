package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	SetRoute(r gin.IRouter)
}

func SetRoutes(r gin.IRouter, controllers []Controller) {
	for _, c := range controllers {
		c.SetRoute(r)
	}
}
