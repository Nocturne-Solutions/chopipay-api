package rabbitmq

import (
	"log"

	"chopipay/config/rabbitmq"
)

func ConsumeMessages(queueName string) {
	msgs, err := rabbitmq.Ch.Consume(
		queueName, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	failOnError(err, "Failed to register a consumer")

	for d := range msgs {
		log.Printf("Received a message: %s", d.Body)
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}