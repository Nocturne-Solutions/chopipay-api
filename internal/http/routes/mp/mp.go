package mp

import (
	"github.com/gin-gonic/gin"

	mpController "chopipay/internal/http/controllers/mp"
	authMiddleware "chopipay/internal/http/routes/middlewares"
	_ "chopipay/internal/models/dto/mp"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/mp/preference", authMiddleware.ValidateAuth(), mpController.CreatePreference)
	router.GET("/mp/preference", authMiddleware.ValidateAuth(), mpController.GetPreference)
	router.POST("/mp/payment/notification", mpController.PaymentNotification)
}