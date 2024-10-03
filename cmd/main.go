package main

import (
	"log"

	"github.com/gin-gonic/gin"

	serverRouter "chopipay/internal/http/routes/server"
	mpRouter "chopipay/internal/http/routes/mp"
)

func main() {
	log.Println("Initializing server...")

	router := gin.Default()

	log.Println("Registering routes...")
	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}