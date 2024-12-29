package preference

import (
	"context"
	"encoding/json"
	"errors"
	"log"
	"strconv"
	"time"

	"chopipay/config/server"
	"chopipay/internal/models/dto"

	"github.com/mercadopago/sdk-go/pkg/preference"
)

const logTag = "[MP_PreferenceServices] "

func CreatePreference(preferenceCli preference.Client, products []dto.ProductDTO, personalId int) (*dto.ProductPreferenceDTO, error) {
	log.Println(logTag + "Creating MercadoPago preference...")

	appURL := server.GetEnvironment().AppUrl
	if appURL == "" {
		errorMessage := logTag + "variable APP_URL not found"
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	profile := server.GetEnvironment().Profile
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

	notificationURL := appURL + "/mp/payment/notification?personalId=" + strconv.Itoa(personalId)
	var items []preference.ItemRequest
	for _, product := range products {
		items = append(items, preference.ItemRequest{
			ID:          strconv.Itoa(product.ProductID),
			Title:       product.Name,
			Quantity:    product.Quantity,
			UnitPrice:   product.Price,
			Description: product.Description,
		})
	}

	request := preference.Request{
		Items:              items,
		NotificationURL:    notificationURL,
		ExternalReference:  "chopipay",
		Expires:            expires,
		ExpirationDateFrom: &expirationDateFrom,
		ExpirationDateTo:   &expirationDateTo,
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
		ProductsDTO:         products,
		PreferenceID:        resource.ID,
		PaymentPoint:        resource.InitPoint,
		SandboxPaymentPoint: resource.SandboxInitPoint,
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

	var products []dto.ProductDTO
	for _, item := range resource.Items {
		productId, err := strconv.Atoi(item.ID)
		if err != nil {
			log.Println(logTag + "Error parsing product id: " + err.Error())
			continue
		}
		products = append(products, dto.ProductDTO{
			ProductID:   productId,
			Name:        item.Title,
			Price:       item.UnitPrice,
			Quantity:    item.Quantity,
			Description: item.Description,
		})
	}

	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorMessage := logTag + "Error parsing resource to json: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	log.Println(logTag + "Preference retrieved: " + string(jsonResource))

	return &dto.ProductPreferenceDTO{
		ProductsDTO:         products,
		PreferenceID:        resource.ID,
		PaymentPoint:        resource.InitPoint,
		SandboxPaymentPoint: resource.SandboxInitPoint,
	}, nil
}
