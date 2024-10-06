package shop

import (
	shopController "chopipay/internal/http/controllers/shop"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/shop", shopController.Create)
	router.GET("/shop/:id", shopController.GetByID)
	router.GET("/shop/personal/:personal_id", shopController.GetAllByPersonalId)
	router.PUT("/shop/:id", shopController.Update)
	router.DELETE("/shop/:id", shopController.Delete)
	router.GET("/shop/:id/products", shopController.GetShopProducts)
}