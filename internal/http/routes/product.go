package routes

import (
	"chopipay/config/di"
	"github.com/gin-gonic/gin"

	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegProductRoutes(router *gin.Engine, init *di.Initialization) {
	router.POST("/product", authMiddleware.ValidateAuth(), init.ProductCtrl.Create)
	router.GET("/product/:id", authMiddleware.ValidateAuth(), init.ProductCtrl.FindByID)
	router.PUT("/product/:id", authMiddleware.ValidateAuth(), init.ProductCtrl.Update)
	router.DELETE("/product/:id", authMiddleware.ValidateAuth(), init.ProductCtrl.Delete)
}
