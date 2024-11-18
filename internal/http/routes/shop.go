package routes

import (
	"chopipay/config/di"
	authMiddleware "chopipay/internal/http/routes/middlewares"

	"github.com/gin-gonic/gin"
)

func RegShopRoutes(router *gin.Engine, init *di.Initialization) {
	router.POST("/shop", authMiddleware.ValidateAuth(), init.ShopCtrl.Create)
	router.GET("/shop/:id", authMiddleware.ValidateAuth(), init.ShopCtrl.GetByID)
	router.GET("/shop/personal/:personal_id", authMiddleware.ValidateAuth(), init.ShopCtrl.GetAllByPersonalId)
	router.PUT("/shop/:id", authMiddleware.ValidateAuth(), init.ShopCtrl.Update)
	router.DELETE("/shop/:id", authMiddleware.ValidateAuth(), init.ShopCtrl.Delete)
	router.GET("/shop/:id/products", authMiddleware.ValidateAuth(), init.ShopCtrl.GetShopProducts)
}
