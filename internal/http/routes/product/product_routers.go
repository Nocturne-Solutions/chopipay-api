package product

import (
	"github.com/gin-gonic/gin"

	productController "chopipay/internal/http/controllers/product"
	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/product", authMiddleware.ValidateAuth(), productController.Create)
	router.GET("/product/:id", authMiddleware.ValidateAuth(), productController.FindByID)
	router.PUT("/product/:id", authMiddleware.ValidateAuth(), productController.Update)
	router.DELETE("/product/:id", authMiddleware.ValidateAuth(), productController.Delete)
}
