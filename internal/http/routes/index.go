package routes

import (
	"log"

	"github.com/gin-gonic/gin"

	serverRouter "chopipay/internal/http/routes/server"
	mpRouter "chopipay/internal/http/routes/mp"
	userRouter "chopipay/internal/http/routes/user"
	personalRouter "chopipay/internal/http/routes/personal"
)

func RegisterRoutes(router *gin.Engine) {
	log.Println("Registering routes...")
	
	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	userRouter.RegisterRoutes(router)
	personalRouter.RegisterRoutes(router)

	log.Println("Routes registered")
}
