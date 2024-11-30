package routes

import (
	"chopipay/config/di"
	"github.com/gin-gonic/gin"

	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegUserRoutes(router *gin.Engine, init *di.Initialization) {
	router.POST("/user", init.UserCtrl.Create)
	router.GET("/user/:id", authMiddleware.ValidateAuth(), init.UserCtrl.FindByID)
	router.PUT("/user/:id", authMiddleware.ValidateAuth(), init.UserCtrl.Update)
	router.DELETE("/user/:id", authMiddleware.ValidateAuth(), init.UserCtrl.Delete)
}
