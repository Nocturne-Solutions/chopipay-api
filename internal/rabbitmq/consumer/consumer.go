package rabbitmq

import (
	"log"

	"encoding/json"

	"chopipay/config/rabbitmq"
	personalServices "chopipay/internal/http/services/app/personal"
	productServices "chopipay/internal/http/services/app/product"
	notStrategy "chopipay/internal/http/services/mp/strategies/interfaces"
	dtos "chopipay/internal/models/dto"
)

const logTag = "[RabbitMQ_Consumer:mp_payment_notification] "

func ConsumePaymentNotifications(queueName string) {
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

		log.Printf(logTag + "Processing preference notification: %s", d.Body)
		var productPayment dtos.ProductMpPaymentDTO
		err = json.Unmarshal(d.Body, &productPayment)
		if err != nil {
			log.Println(logTag + "Error parsing body: ", err.Error())
			break
		}
		
		product, err := productServices.FindByID(productPayment.ProductID)
		if err != nil {
			log.Println(logTag + "Error finding product by id: ", err.Error())
			break
		}
		log.Println(logTag + "Product found: ", product.ID)

		credentials, err := personalServices.GetPersonalCredentialsByShopID(product.ShopID)
		if err != nil {
			log.Println(logTag + "Error getting personal credentials by shop id: ", err.Error())
			break
		}

		processor := &notStrategy.NotificationProcessor{}
		err = processor.Process(productPayment.Topic, productPayment.ID, credentials)
		if err != nil {
			log.Println(logTag + "Error processing notification: ", err.Error())
			break
		}
		
		println("Preference notification processed successfully")
	}
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
	}
}