package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"chopipay/config/db/pg"

	"chopipay/internal/http/routes"
)

func main() {
	log.Println("Initializing server...")

	log.Println("Initializing database connection...")
	err := pg.InitConnection()
	if err != nil {
		log.Fatalf("Error initializing database connection: %v", err)
	}
	log.Println("Database connection initialized")
	
	router := gin.Default()
	routes.RegisterRoutes(router)
	
	log.Println("Server is running on port 8080")
	router.Run(":8080")
}