package publisher

import (
	"log"
	"context"
	"time"

	"chopipay/config/rabbitmq"
	amqp "github.com/rabbitmq/amqp091-go"
)

func PublishMessage(queueName string, body string) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	err := rabbitmq.Ch.PublishWithContext(ctx,
		"",     // exchange
		queueName, // routing key
		false,  // mandatory
		false,  // immediate
		amqp.Publishing {
		  ContentType: "text/plain",
		  Body:        []byte(body),
		})
	failOnError(err, "Failed to publish a message: " + body)
	log.Printf("amqp: %s | Message published successfully", queueName)
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Printf("%s: %s", msg, err)
	}
}