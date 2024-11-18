package business

import (
	"chopipay/config/di"
	authMiddleware "chopipay/internal/http/routes/middlewares"
	"github.com/gin-gonic/gin"
)

func RegSalesRoutes(router *gin.Engine, init *di.Initialization) {
	router.POST("/sales", authMiddleware.ValidateAuth(), init.SalesCtrl.Create)
}
