package mp

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
	_ "github.com/mercadopago/sdk-go/pkg/payment"
	_ "github.com/mercadopago/sdk-go/pkg/preference"

	rmqPublisher "chopipay/internal/rabbitmq/publisher"
	personalServices "chopipay/internal/http/services/app/personal"
	productServices "chopipay/internal/http/services/app/product"
	mpCientServices "chopipay/internal/http/services/mp/client"
	mpDtos "chopipay/internal/models/dto/mp"
)

func PaymentNotification(c *gin.Context) {
	log.Println("Receiving payment notification...")

	param_productId := c.Request.URL.Query().Get("productId")
	param_id := c.Request.URL.Query().Get("id")
	param_topic := c.Request.URL.Query().Get("topic")

	log.Printf("Params: id=%s, productId=%s, topic=%s", param_id, param_productId, param_topic)

	if param_id == "" || param_productId == "" || param_topic == "" {
		log.Printf("id(%s),productId(%s) or topic(%s) is empty", param_id, param_productId, param_topic)
		returnSuccess(c)
	}

	id, err := strconv.Atoi(param_id)
	if err != nil || id == 0 {
		log.Printf("error converting id %s. Cause: %s", param_id, err.Error())
		returnSuccess(c)
	}

	productId, err := strconv.Atoi(param_productId)
	if err != nil || productId == 0 {
		log.Printf("error converting productId %s. Cause: %s", param_productId, err.Error())
		returnSuccess(c)
	}

	log.Printf("Values: id=%d, productId=%d, topic=%s", id, productId, param_topic)

	product, err := productServices.FindByID(productId)
	if err != nil {
		log.Println("Error finding product by id: ", err.Error())
		returnSuccess(c)
	}
	log.Println("Product found: ", product.ID)

	credentials, err := personalServices.GetPersonalCredentialsByShopID(product.ShopID)
	if err != nil {
		log.Println("Error getting personal credentials by shop id: ", err.Error())
		returnSuccess(c)
	}

	if param_topic == "merchant_order" {
		log.Println("processing merchant_order: ", id)
		client := mpCientServices.GetClient(credentials.AccessToken, param_topic)
		if err, ok := client.(error); ok {
			log.Println("Error getting client: ", err.Error())
			returnSuccess(c)
		}
		client.(merchantorder.Client).Get(context.Background(), id)
		log.Println("Merchant order received: ", id)
		paymentBody, err := c.GetRawData()
		if err != nil {
			log.Println("Error getting body: ", err.Error())
			returnSuccess(c)
		}

		var notification mpDtos.NotificationsDTO
		err = json.Unmarshal(paymentBody, &notification)
		if err != nil {
			log.Println("Error parsing body: ", err.Error())
			returnSuccess(c)
		}
		
		rmqPublisher.PublishMessage("mp_payment_notification", string(paymentBody))
	}

	returnSuccess(c)
}

func returnSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"notification": "received",
	})
}
