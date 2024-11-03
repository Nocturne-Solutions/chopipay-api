package main

import (
	"log"

	"github.com/gin-gonic/gin"
	_ "chopipay/docs"

	"chopipay/config/db/pg"
	"chopipay/config/server"
	"chopipay/config/rabbitmq"
	rmqConsumers "chopipay/internal/rabbitmq/consumer"
	"chopipay/internal/http/routes"
	rmq "chopipay/internal/rabbitmq"
)

// @title Chopipay API
// @version 1.0
// @description This is the Chopipay API documentation
// @termsOfService http://swagger.io/terms/

// @contact.name Nocturne Solutions
// @contact.url https://github.com/Nocturne-Solutions
// @contact.email nocturne.solutions@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
func main() {
	log.Println("Initializing server...")

	server.LoadEnvirontment()
	log.Println("Environment variables initialized")

	pg.InitConnection(server.EnvVars)
	log.Println("Database connection initialized")

	rabbitmq.InitRabbitMQ(server.EnvVars)
	log.Println("RabbitMQ connection initialized")
	rabbitmq.DeclareQueue(rmq.PreferenceNotificationQueue)
	// add more queues here
	log.Println("RabbitMQ queues declared successfully")
	go rmqConsumers.ConsumePaymentNotifications(rmq.PreferenceNotificationQueue)
	log.Printf("Consumming messages from RabbitMQ queue: %s", rmq.PreferenceNotificationQueue)
	log.Println("RabbitMQ connection initialized")
	
	router := gin.Default()
	routes.RegisterRoutes(router)
	
	log.Println("Server is running on port 8080")
	router.Run(":8080")

	defer pg.CloseConnection()
	defer rabbitmq.CloseRabbitMQChannel()
}