package rabbitmq

import (
	amqp "github.com/rabbitmq/amqp091-go"

	"chopipay/internal/http/services"
)

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
