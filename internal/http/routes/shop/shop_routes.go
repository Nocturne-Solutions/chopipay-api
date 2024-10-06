package shop

import (
	shopController "chopipay/internal/http/controllers/shop"
	authMiddleware "chopipay/internal/http/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shop", authMiddleware.ValidateAuth(), shopController.Create)
	router.GET("/shop/:id", authMiddleware.ValidateAuth(), shopController.GetByID)
	router.GET("/shop/personal/:personal_id", authMiddleware.ValidateAuth(), shopController.GetAllByPersonalId)
	router.PUT("/shop/:id", authMiddleware.ValidateAuth(), shopController.Update)
	router.DELETE("/shop/:id", authMiddleware.ValidateAuth(), shopController.Delete)
	router.GET("/shop/:id/products", authMiddleware.ValidateAuth(), shopController.GetShopProducts)
}