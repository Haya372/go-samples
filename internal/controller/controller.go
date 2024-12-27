package controller

import "github.com/gin-gonic/gin"

type Controller interface {
	SetRoute(r gin.IRouter)
}
