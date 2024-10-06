package preference

import (
	"log"
	"errors"
	"context"
	"strconv"
	"encoding/json"

	"chopipay/config/server"
	"chopipay/internal/models/entities"

	"github.com/mercadopago/sdk-go/pkg/preference"
)

const logTag = "MP_PreferenceServices | "

func CreatePreference(preferenceCli preference.Client, product *entities.Product) error {
	log.Println(logTag + "Creating MercadoPago preference...")

	appURL := server.EnvVars["APP_URL"]
	if appURL == "" {
		errorMessage := logTag + "variable APP_URL not found"
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	notificationURL := appURL + "/mp/payment/notification?productId=" + strconv.Itoa(product.ID)

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:       product.Name,
				Quantity:    1,
				UnitPrice:   product.Price,
				Description: product.Description,
			},
		},
		NotificationURL: notificationURL,
		ExternalReference: "chopipay",
	}

	resource, err := preferenceCli.Create(context.Background(), request)
	if err != nil {
		errorMessage := logTag + "Error creating preference: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	product.PreferenceID = resource.ID

	// parse resource to json
	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorMessage := logTag + "Error parsing resource to json: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	log.Println(logTag + "Preference created: " + string(jsonResource))

	return nil
}