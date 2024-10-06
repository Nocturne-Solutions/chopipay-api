package personal

import (
	"github.com/gin-gonic/gin"

	personalController "chopipay/internal/http/controllers/personal"
)

func RegisterRoutes(router *gin.Engine) {
	router.POST("/personal", personalController.Create)
	router.GET("/personal/:id", personalController.GetByID)
	router.PUT("/personal/:id", personalController.Update)
	router.DELETE("/personal/:id", personalController.Delete)
	router.PUT("/personal/add-credential", personalController.AddPersonalCredential)
	router.GET("/personal/:id/shops", personalController.GetShopsByPersonalID)
}
