package product

import (
	"github.com/gin-gonic/gin"

	productController "chopipay/internal/http/controllers/product"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/product", productController.Create)
	router.GET("/product/:id", productController.FindByID)
	router.PUT("/product/:id", productController.Update)
	router.DELETE("/product/:id", productController.Delete)
}
