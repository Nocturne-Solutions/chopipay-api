package user

import (
	"github.com/gin-gonic/gin"

	userController "chopipay/internal/http/controllers/user"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/user", userController.Create)
	router.GET("/user/:id", userController.FindByID)
	router.PUT("/user/:id", userController.Update)
	router.DELETE("/user/:id", userController.Delete)
}
