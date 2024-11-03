package merchantorder

import (
	"context"
	"log"
	"errors"

	"github.com/mercadopago/sdk-go/pkg/payment"

	mpClientServices "chopipay/internal/http/services/mp/client"
	"chopipay/internal/models/entities"

)

const logTag = "MP_PaymentProcessor | "

type Processor struct {
	NotificationId int
	Credentials *entities.PersonalCredentials
}

func (p *Processor) Process() error {
	if p.Credentials == nil || p.NotificationId == 0 {
		log.Println(logTag + "Credentials or NotificationId is empty")
		return errors.New("Couldn't get merchant order: credentials or notification id is empty")
	}
	log.Println(logTag + "Processing payment: ", p.NotificationId)
	
	client := mpClientServices.GetClient(p.Credentials.AccessToken, "payment")
	if err, ok := client.(error); ok {
		log.Println(logTag + "Error getting client: ", err.Error())
		return errors.New("error getting client: " + err.Error())
	}

	paymentResponse, err := client.(payment.Client).Get(context.Background(), p.NotificationId)
	if err != nil {
		log.Println(logTag + "Error getting payment: ", err.Error())
		return errors.New("error getting payment: " + err.Error())
	}

	log.Println(logTag + "Payment found: ", paymentResponse)

	return nil
}