package main

import (
	"log"

	"github.com/gin-gonic/gin"

	"chopipay/config/db/pg"
	"chopipay/config/server"
	"chopipay/config/rabbitmq"
	rmqConsumers "chopipay/internal/rabbitmq/consumer"
	"chopipay/internal/http/routes"
)

func main() {
	log.Println("Initializing server...")

	server.LoadEnvirontment()
	log.Println("Environment variables initialized")

	pg.InitConnection(server.EnvVars)
	log.Println("Database connection initialized")

	rabbitmq.InitRabbitMQ(server.EnvVars)
	log.Println("RabbitMQ connection initialized")
	rabbitmq.DeclareQueue("mp_payment_notification")
	// add more queues here
	log.Println("RabbitMQ queues declared successfully")
	go rmqConsumers.ConsumeMessages("mp_payment_notification")
	log.Println("Consumming messages from RabbitMQ queue: mp_payment_notification")
	log.Println("RabbitMQ connection initialized")
	
	router := gin.Default()
	routes.RegisterRoutes(router)
	
	log.Println("Server is running on port 8080")
	router.Run(":8080")

	defer pg.CloseConnection()
	defer rabbitmq.CloseRabbitMQChannel()
}