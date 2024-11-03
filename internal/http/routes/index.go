package routes

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	shopRouter "chopipay/internal/http/routes/shop"
	mpRouter "chopipay/internal/http/routes/mp"
	personalRouter "chopipay/internal/http/routes/personal"
	productRouter "chopipay/internal/http/routes/product"
	serverRouter "chopipay/internal/http/routes/server"
	userRouter "chopipay/internal/http/routes/user"
	authRouter "chopipay/internal/http/routes/auth"
)

func RegisterRoutes(router *gin.Engine) {
	log.Println("Registering routes...")
	
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	
	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	userRouter.RegisterRoutes(router)
	personalRouter.RegisterRoutes(router)
	productRouter.RegisterRoutes(router)
	shopRouter.RegisterRoutes(router)
	authRouter.RegisterRoutes(router)

	log.Println("Routes registered")
}
