package rabbitmq

import (
	"fmt"
  	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Ch *amqp.Channel

func InitRabbitMQ(envVars map[string]string) {
	host := envVars["RABBITMQ_HOST"]
	port := envVars["RABBITMQ_PORT"]
	user := envVars["RABBITMQ_USER"]
	password := envVars["RABBITMQ_PASSWORD"]
	vhost := envVars["RABBITMQ_VHOST"]

	if host == "" || port == "" || user == "" || password == "" || vhost == "" {
		log.Fatal("Missing environment variables for rabbitmq connection")
	}

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, vhost)
	
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func DeclareQueue(queueName string) {
	_, err := Ch.QueueDeclare(
		queueName, // name
		false,   // durable
		false,   // delete when unused
		false,   // exclusive
		false,   // no-wait
		nil,     // arguments
	)
	failOnError(err, "Failed to declare a queue")
}

func CloseRabbitMQChannel() {
	Ch.Close()
}

func failOnError(err error, msg string) {
	if err != nil {
	  log.Panicf("%s: %s", msg, err)
	}
}