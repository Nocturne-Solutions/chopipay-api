package personal

import (
	"github.com/gin-gonic/gin"

	personalController "chopipay/internal/http/controllers/personal"
	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/personal", authMiddleware.ValidateAuth(), personalController.Create)
	router.GET("/personal/:id", authMiddleware.ValidateAuth(), personalController.GetByID)
	router.PUT("/personal/:id", authMiddleware.ValidateAuth(), personalController.Update)
	router.DELETE("/personal/:id", authMiddleware.ValidateAuth(), personalController.Delete)
	router.PUT("/personal/add-credential", authMiddleware.ValidateAuth(), personalController.AddPersonalCredential)
	router.GET("/personal/:id/shops", authMiddleware.ValidateAuth(), personalController.GetShopsByPersonalID)
}
