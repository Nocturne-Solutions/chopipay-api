package auth

import (
	"chopipay/config/di"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine, init *di.Initialization) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", init.AuthController.Login)
		auth.POST("/refresh-token", init.AuthController.RefreshToken)
	}
}
