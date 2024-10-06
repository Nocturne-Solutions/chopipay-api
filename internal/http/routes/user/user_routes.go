package user

import (
	"github.com/gin-gonic/gin"

	userController "chopipay/internal/http/controllers/user"
	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/user", userController.Create)
	router.GET("/user/:id", authMiddleware.ValidateAuth(), userController.FindByID)
	router.PUT("/user/:id", authMiddleware.ValidateAuth(), userController.Update)
	router.DELETE("/user/:id", authMiddleware.ValidateAuth(), userController.Delete)
}
