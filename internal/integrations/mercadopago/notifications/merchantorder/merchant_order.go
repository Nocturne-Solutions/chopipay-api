package merchantorder

import (
	mpClientServices "chopipay/internal/integrations/mercadopago/client"
	"context"
	"errors"
	"log"

	"github.com/mercadopago/sdk-go/pkg/merchantorder"

	"chopipay/internal/models/entities"
)

const logTag = "MP_MerchantOrderProcessor | "

type Processor struct {
	NotificationId int
	Credentials    *entities.PersonalCredentials
}

func (p *Processor) Process() error {
	if p.Credentials == nil || p.NotificationId == 0 {
		log.Println(logTag + "Credentials or NotificationId is empty")
		return errors.New("Couldn't get merchant order: credentials or notification id is empty")
	}
	log.Println(logTag+"Processing merchant order: ", p.NotificationId)

	client := mpClientServices.GetClient(p.Credentials.AccessToken, "merchant_order")
	if err, ok := client.(error); ok {
		log.Println(logTag+"Error getting client: ", err.Error())
		return errors.New("error getting client: " + err.Error())
	}

	merchantorderResponse, err := client.(merchantorder.Client).Get(context.Background(), p.NotificationId)
	if err != nil {
		log.Println(logTag+"Error getting merchant order: ", err.Error())
		return errors.New("error getting merchant order: " + err.Error())
	}

	log.Println(logTag+"Merchant order found: ", merchantorderResponse)

	return nil
}
