package main

import (
	"log"

	"chopipay/config/db/pg"
	"chopipay/config/di"
	"chopipay/config/rabbitmq"
	"chopipay/config/server"
	_ "chopipay/docs"
	"chopipay/internal/http/routes"
	rmq "chopipay/internal/queues"
	rmqConsumers "chopipay/internal/queues/consumer"
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

	server.LoadEnvironment()
	log.Println("Environment variables initialized")

	pg.InitConnection(server.EnvVars)
	defer pg.CloseConnection()
	log.Println("Database connection initialized")

	init := di.Init()

	rabbitmq.InitRabbitMQ(server.EnvVars)
	defer rabbitmq.CloseRabbitMQChannel()
	log.Println("RabbitMQ connection initialized")
	rabbitmq.DeclareQueue(rmq.PreferenceNotificationQueue)
	// add more queues here
	log.Println("RabbitMQ queues declared successfully")
	consumers := rmqConsumers.NewConsumer(rabbitmq.Ch, init.ProductServices, init.PersonalService)
	go consumers.PaymentNotifications(rmq.PreferenceNotificationQueue)
	log.Println("RabbitMQ connection initialized")

	app := routes.InitRoutes(init)

	log.Println("Server is running on port 8080")
	err := app.Run(":8080")
	if err != nil {
		log.Println("Error starting server: ", err)
		return
	}
}
