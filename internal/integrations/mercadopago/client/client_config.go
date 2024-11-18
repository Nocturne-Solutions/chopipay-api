package client

import (
	"errors"
	"log"

	"github.com/mercadopago/sdk-go/pkg/config"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"github.com/mercadopago/sdk-go/pkg/payment"
	"github.com/mercadopago/sdk-go/pkg/preference"
)

const logTag = "MP_ClientConfig | "

func InitClientConfig(accessToken string) (*config.Config, error) {
	log.Println(logTag + "Initializing MercadoPago preference client...")

	cfg, err := config.New(accessToken)
	if err != nil {
		log.Println("Error initializing MercadoPago")
		return nil, errors.New("error initializing MercadoPago config. Cause: " + err.Error())
	}

	return cfg, nil
}

func GetPreferenceClient(cfg *config.Config) *preference.Client {
	log.Println(logTag + "Getting MercadoPago preference client...")

	client := preference.NewClient(cfg)

	return &client
}

func GetClient(accessToken string, clientType string) interface{} {
	log.Println(logTag + "Getting MercadoPago client...")

	cfg, err := InitClientConfig(accessToken)
	if err != nil {
		log.Println(err)
		return nil
	}

	switch clientType {
	case "preference":
		return preference.NewClient(cfg)
	case "payment":
		return payment.NewClient(cfg)
	case "merchant_order":
		return merchantorder.NewClient(cfg)
	default:
		return errors.New("client type not found: " + clientType)
	}
}
