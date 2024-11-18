package routes

import (
	"chopipay/config/di"
	"github.com/gin-gonic/gin"

	authMiddleware "chopipay/internal/http/routes/middlewares"
)

func RegPersonalRoutes(router *gin.Engine, init *di.Initialization) {
	router.POST("/personal", authMiddleware.ValidateAuth(), init.PersonalCtrl.Create)
	router.GET("/personal/:id", authMiddleware.ValidateAuth(), init.PersonalCtrl.GetByID)
	router.PUT("/personal/:id", authMiddleware.ValidateAuth(), init.PersonalCtrl.Update)
	router.DELETE("/personal/:id", authMiddleware.ValidateAuth(), init.PersonalCtrl.Delete)
	router.PUT("/personal/add-credential", authMiddleware.ValidateAuth(), init.PersonalCtrl.AddPersonalCredential)
	router.GET("/personal/:id/shops", authMiddleware.ValidateAuth(), init.PersonalCtrl.GetShopsByPersonalID)
}
