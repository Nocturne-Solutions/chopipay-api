package routes

import (
	"chopipay/config/di"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authRouter "chopipay/internal/http/routes/auth"
	"chopipay/internal/http/routes/business"
	mpRouter "chopipay/internal/http/routes/mp"
	serverRouter "chopipay/internal/http/routes/server"
)

func InitRoutes(init *di.Initialization) *gin.Engine {
	log.Println("Registering routes...")
	if profile := os.Getenv("PROFILE"); profile == "prod" {
		gin.SetMode(gin.ReleaseMode)
	}
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	RegUserRoutes(router, init)
	RegPersonalRoutes(router, init)
	RegProductRoutes(router, init)
	RegShopRoutes(router, init)
	authRouter.RegisterRoutes(router, init)
	business.RegSalesRoutes(router, init)

	log.Println("Routes registered")
	return router
}
