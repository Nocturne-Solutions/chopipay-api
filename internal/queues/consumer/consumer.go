package rabbitmq

import (
	"chopipay/internal/exceptions"
	"encoding/json"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"

	"chopipay/config/rabbitmq"
	"chopipay/internal/http/services"
	notStrategy "chopipay/internal/integrations/mercadopago/notifications/strategies"
	dtos "chopipay/internal/models/dto"
)

const logTag = "[RabbitMQ_Consumer:mp_payment_notification] "

type Consumer interface {
	PaymentNotifications(queueName string)
}

type consumer struct {
	ch              *amqp.Channel
	productService  services.ProductService
	personalService services.PersonalService
}

func NewConsumer(ch *amqp.Channel,
	productService services.ProductService,
	personalService services.PersonalService) Consumer {
	return &consumer{
		ch:              ch,
		productService:  productService,
		personalService: personalService,
	}
}

func (c *consumer) PaymentNotifications(queueName string) {
	log.Printf("Consumming messages from RabbitMQ queue: %s", queueName)
	msgs, err := rabbitmq.Ch.Consume(
		queueName, // queue
		"",        // consumer
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	exceptions.PanicException(err, "failed to register a consumer")

	for d := range msgs {

		log.Printf(logTag+"Processing preference notification: %s", d.Body)
		var productPayment dtos.ProductMpPaymentDTO
		err = json.Unmarshal(d.Body, &productPayment)
		if err != nil {
			log.Println(logTag+"Error parsing body: ", err.Error())
			break
		}

		product, err := c.productService.FindByID(productPayment.ProductID)
		if err != nil {
			log.Println(logTag+"Error finding product by id: ", err.Error())
			break
		}
		log.Println(logTag+"Product found: ", product.ID)

		credentials, err := c.personalService.GetPersonalCredentialsByShopID(product.ShopID)
		if err != nil {
			log.Println(logTag+"Error getting personal credentials by shop id: ", err.Error())
			break
		}

		processor := &notStrategy.NotificationProcessor{}
		err = processor.Process(productPayment.Topic, productPayment.ID, credentials)
		if err != nil {
			log.Println(logTag+"Error processing notification: ", err.Error())
			break
		}

		println("Preference notification processed successfully")
	}
}
