package auth

import (
	"github.com/gin-gonic/gin"

	authController	"chopipay/internal/http/controllers/auth"
)

func RegisterRoutes(router *gin.Engine) {
	auth := router.Group("/auth")
	{
		auth.POST("/login", authController.Login)
		auth.POST("/refresh-token", authController.RefreshToken)
	}
}