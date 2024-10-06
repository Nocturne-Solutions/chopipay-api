package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	shopRouter "chopipay/internal/http/routes/shop"
	mpRouter "chopipay/internal/http/routes/mp"
	personalRouter "chopipay/internal/http/routes/personal"
	serverRouter "chopipay/internal/http/routes/server"
	userRouter "chopipay/internal/http/routes/user"
	authRouter "chopipay/internal/http/routes/auth"
)

func RegisterRoutes(router *gin.Engine) {
	log.Println("Registering routes...")
	
	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	userRouter.RegisterRoutes(router)
	personalRouter.RegisterRoutes(router)
	shopRouter.RegisterRoutes(router)
	authRouter.RegisterRoutes(router)

	log.Println("Routes registered")
}
