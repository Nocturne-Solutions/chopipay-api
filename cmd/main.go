package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"chopipay/config/db/pg"

	serverRouter "chopipay/internal/http/routes/server"
	mpRouter "chopipay/internal/http/routes/mp"
	userRouter "chopipay/internal/http/routes/user"
)

func main() {
	log.Println("Initializing server...")

	log.Println("Initializing database connection...")
	err := pg.InitConnection()
	if err != nil {
		log.Println("Error initializing database connection: ", err)
		log.Fatalf("Error initializing database connection: %v", err)
	}
	log.Println("Database connection initialized")
	
	log.Println("Registering routes...")
	router := gin.Default()
	serverRouter.RegisterRoutes(router)
	mpRouter.RegisterRoutes(router)
	userRouter.RegisterRoutes(router)
	log.Println("Routes registered")

	log.Println("Server is running on port 8080")
	router.Run(":8080")
}