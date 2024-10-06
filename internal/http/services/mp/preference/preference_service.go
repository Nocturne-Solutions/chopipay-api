package preference

import (
	"log"
	"errors"
	"context"
	"strconv"
	"encoding/json"
	"time"

	"chopipay/config/server"
	"chopipay/internal/models/entities"
	"chopipay/internal/models/dto"

	"github.com/mercadopago/sdk-go/pkg/preference"
)

const logTag = "MP_PreferenceServices | "

func CreatePreference(preferenceCli preference.Client, product *entities.Product) (*dto.ProductPreferenceDTO, error) {
	log.Println(logTag + "Creating MercadoPago preference...")

	appURL := server.EnvVars["APP_URL"]
	if appURL == "" {
		errorMessage := logTag + "variable APP_URL not found"
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	profile := server.EnvVars["PROFILE"]
	if profile == "" {
		profile = "env"
	}

	var expires bool
	var expirationDateFrom time.Time
	var expirationDateTo time.Time
	if profile == "env" {
		expirationDateFrom = time.Now()
		expirationDateTo = time.Now().AddDate(0, 0, 1)
		expires = true
	} else {
		expirationDateFrom = time.Now()
		expirationDateTo = time.Now().AddDate(10, 0, 1)
		expires = false
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
		Expires: expires,
		ExpirationDateFrom: &expirationDateFrom,
		ExpirationDateTo: &expirationDateTo,
	}

	resource, err := preferenceCli.Create(context.Background(), request)
	if err != nil {
		errorMessage := logTag + "Error creating preference: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorMessage := logTag + "Error parsing resource to json: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	log.Println(logTag + "Preference created: " + string(jsonResource))

	return &dto.ProductPreferenceDTO{
		ID:                  product.ID,
		Name:                product.Name,
		Price:               product.Price,
		CurrencyID:          resource.Items[0].CurrencyID,
		ShopID:              product.ShopID,
		Description:         product.Description,
		PreferenceID:        resource.ID,
		PaymentPoint:        resource.InitPoint,
		SandboxPaymentPoint: resource.SandboxInitPoint,
		PictureURL:          resource.Items[0].PictureURL,
	}, nil
}

func GetPreference(preferenceCli preference.Client, preferenceID string) (*dto.ProductPreferenceDTO, error) {
	log.Println(logTag + "Getting MercadoPago preference...")

	resource, err := preferenceCli.Get(context.Background(), preferenceID)
	if err != nil {
		errorMessage := logTag + "Error getting preference: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorMessage := logTag + "Error parsing resource to json: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	log.Println(logTag + "Preference retrieved: " + string(jsonResource))

	return &dto.ProductPreferenceDTO{
		Name:                resource.Items[0].Title,
		Price:               resource.Items[0].UnitPrice,
		CurrencyID:          resource.Items[0].CurrencyID,
		Description:         resource.Items[0].Description,
		PreferenceID:        resource.ID,
		PaymentPoint:        resource.InitPoint,
		SandboxPaymentPoint: resource.SandboxInitPoint,
		PictureURL:          resource.Items[0].PictureURL,
	}, nil
}