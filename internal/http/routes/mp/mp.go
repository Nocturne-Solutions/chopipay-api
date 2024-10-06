package mp

import (
	"github.com/gin-gonic/gin"

	mpController "chopipay/internal/http/controllers/mp"
	_ "chopipay/internal/models/dto/mp"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/mp/payment/notification", mpController.PaymentNotification)
}