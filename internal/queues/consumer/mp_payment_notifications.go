package rabbitmq

import (
	"encoding/json"
	"fmt"
	"log"

	"chopipay/config/rabbitmq"
	"chopipay/internal/exceptions"
	notStrategy "chopipay/internal/integrations/mercadopago/notifications/strategies"
	dtos "chopipay/internal/models/dto"
)

func (c *consumer) PaymentNotifications(queueName string) {
	logTag := fmt.Sprintf("[RabbitMQ_Consumer:%s]", queueName)
	log.Printf("%s Consumming messages from RabbitMQ queue: %s", logTag, queueName)
	messages, err := rabbitmq.Ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	exceptions.PanicException(err, logTag+" failed to register a consumer")

	for d := range messages {

		log.Printf("%s Processing preference notification: %s", logTag, d.Body)
		var productPayment dtos.ProductMpPaymentDTO
		err = json.Unmarshal(d.Body, &productPayment)
		if err != nil {
			log.Println(logTag+" Error parsing body: ", err.Error())
			break
		}

		personalCredentials, err := c.personalService.GetPersonalCredentialsByPersonalId(productPayment.PersonalID)
		if err != nil {
			log.Println(logTag+" Error getting personal credentials by personal id: ", err.Error())
			break
		}

		processor := &notStrategy.NotificationProcessor{}
		err = processor.Process(productPayment.Topic, productPayment.ID, personalCredentials)
		if err != nil {
			log.Println(logTag+" Error processing notification: ", err.Error())
			break
		}

		println(logTag + " Preference notification processed successfully")
	}
}
