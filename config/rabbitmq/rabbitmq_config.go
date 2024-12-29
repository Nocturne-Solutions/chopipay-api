package rabbitmq

import (
	"chopipay/config/server"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

var Ch *amqp.Channel

func InitRabbitMQ() {
	host := server.GetEnvironment().RabbitmqHost
	port := server.GetEnvironment().RabbitmqPort
	user := server.GetEnvironment().RabbitmqUser
	password := server.GetEnvironment().RabbitmqPassword
	vhost := server.GetEnvironment().RabbitmqVhost

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, vhost)

	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to connect to RabbitMQ")

	Ch, err = conn.Channel()
	failOnError(err, "Failed to open a channel")
}

func DeclareQueue(queueName string) {
	_, err := Ch.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		false,     // exclusive
		false,     // no-wait
		nil,       // arguments
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
