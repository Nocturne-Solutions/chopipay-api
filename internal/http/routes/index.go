package routes

import (
	"chopipay/config/di"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	authRouter "chopipay/internal/http/routes/auth"
	"chopipay/internal/http/routes/business"
	mpRouter "chopipay/internal/http/routes/mp"
	serverRouter "chopipay/internal/http/routes/server"
	userRouter "chopipay/internal/http/routes/user"
)

func InitRoutes(init *di.Initialization) *gin.Engine {
	log.Println("Registering routes...")
	router := gin.Default()

	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	userRouter.RegisterRoutes(router)
	RegPersonalRoutes(router, init)
	RegProductRoutes(router, init)
	RegShopRoutes(router, init)
	authRouter.RegisterRoutes(router)
	business.RegSalesRoutes(router, init)

	log.Println("Routes registered")
	return router
}
