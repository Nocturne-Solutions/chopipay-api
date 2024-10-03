package mp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/preference"
	_ "github.com/mercadopago/sdk-go/pkg/payment"

	mpConfig "chopipay/config/mercadopago"
	errorshandler "chopipay/internal/http/errors_handler"
)

func CreatePreference(c *gin.Context) {
	log.Println("Creating MercadoPago preference...")

	log.Println("Initializing MercadoPago config...")

	cfg, err := mpConfig.Initialize()
	if err != nil {
		errorshandler.ErrorHandler(c, err, "Error initializing MercadoPago")
		return
	}

	log.Printf("Config: %+v", cfg)

	client := preference.NewClient(cfg.Cfg)

	request := preference.Request{
		Items: []preference.ItemRequest{
			{
				Title:       "My product 3",
				Quantity:    1,
				UnitPrice:   123,
			},
		},
		NotificationURL: "https://06b7-186-138-229-135.ngrok-free.app/mp/payment/notification?userId=12345",
		ExternalReference: "vendedor_carlos",
	}

	resource, err := client.Create(context.Background(), request)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resource)

	// Return resource as JSON
	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorshandler.ErrorHandler(c, err, "Error marshalling resource")
		return
	}

	log.Println("json Resource: ", string(jsonResource))

	c.JSON(http.StatusOK, gin.H{
		"resource": resource,
	})
}

func GetPreference(c *gin.Context) {
	log.Println("Getting MercadoPago preference...")

	log.Println("Initializing MercadoPago config...")

	cfg, err := mpConfig.Initialize()
	if err != nil {
		errorshandler.ErrorHandler(c, err, "Error initializing MercadoPago")
		return
	}

	log.Printf("Config: %+v", cfg)

	client := preference.NewClient(cfg.Cfg)
	// items := preference.Get(context.Background(), )
	resource, err := client.Get(context.Background(), "2013036600-05e83294-ed4c-41af-884b-108ab9e4cb1f")
	if err != nil {
		log.Println(err)
		return
	}
	log.Println(resource)

	// Return resource as JSON
	jsonResource, err := json.Marshal(resource)
	if err != nil {
		errorshandler.ErrorHandler(c, err, "Error marshalling resource")
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"resource": jsonResource,
	})
}

func PaymentNotification(c *gin.Context) {
	log.Println("Receiving payment notification...")

	/* id := c.Query("id")
	topic := c.Query("topic")

	switch topic {
		case "payment":
			payment := payment.NewClient()
			payment.Get(context.Background(), 123)
			log.Println("Payment received: ", id)
		case "merchant_order":
			log.Println("Merchant order received: ", id)
		default:
			log.Println("Unknown topic: ", topic)
	} */
	
	paymentBody, err := c.GetRawData()
	if err != nil {
		errorshandler.ErrorHandler(c, err, "Error getting raw data")
		return
	}

	log.Println("body:\n ", string(paymentBody))

	params := c.Request.URL.Query()
	log.Println("params:\n ", params)

	c.JSON(http.StatusOK, gin.H{
		"notification": "paymente received",
	})
} 