package controller

import (
	"github.com/gin-gonic/gin"

	jwt "github.com/appleboy/gin-jwt/v2"
)

type loginController struct {
	jwtMiddleware *jwt.GinJWTMiddleware
}

func (c *loginController) SetRoute(r gin.IRouter) {
	r.POST("/login", c.jwtMiddleware.LoginHandler)
}

func NewLoginController(jwtMiddleware *jwt.GinJWTMiddleware) Controller {
	return &loginController{
		jwtMiddleware: jwtMiddleware,
	}
}
