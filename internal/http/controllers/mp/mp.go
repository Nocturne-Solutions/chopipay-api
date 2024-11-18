package mp

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	_ "github.com/mercadopago/sdk-go/pkg/payment"
	_ "github.com/mercadopago/sdk-go/pkg/preference"

	dtos "chopipay/internal/models/dto"
	rmq "chopipay/internal/queues"
	rmqPublisher "chopipay/internal/queues/publisher"
)

const logTag = "[MP_PaymentNotification] "

func PaymentNotification(c *gin.Context) {
	log.Println(logTag + "Receiving payment notification...")

	paramPersonalId := c.Request.URL.Query().Get("personalId")
	param_id := c.Request.URL.Query().Get("id")
	param_data_id := c.Request.URL.Query().Get("data.id")
	param_topic := c.Request.URL.Query().Get("topic")
	param_type := c.Request.URL.Query().Get("type")

	if paramPersonalId == "" {
		log.Printf(logTag + "product id is empty")
		returnSuccess(c)
		return
	}

	if param_topic == "" {
		param_topic = param_type
	}

	if param_id == "" {
		param_id = param_data_id
	}

	log.Printf(logTag+"Params: id=%s, productId=%s, topic=%s", param_id, paramPersonalId, param_topic)
	if param_id == "" || param_topic == "" {
		log.Printf(logTag+"id(%s) or topic(%s) is empty", param_id, param_topic)
		returnSuccess(c)
		return
	}

	id, err := strconv.Atoi(param_id)
	if err != nil || id == 0 {
		log.Printf(logTag+"error converting id %s. Cause: %s", param_id, err.Error())
		returnSuccess(c)
		return
	}

	productId, err := strconv.Atoi(paramPersonalId)
	if err != nil || productId == 0 {
		log.Printf(logTag+"error converting productId %s. Cause: %s", paramPersonalId, err.Error())
		returnSuccess(c)
		return
	}

	log.Printf(logTag+"Values: id=%d, productId=%d, topic=%s", id, productId, param_topic)

	log.Printf("%s processing %s: %d", logTag, param_topic, id)
	productPayment := dtos.ProductMpPaymentDTO{
		ProductID: productId,
		ID:        id,
		Topic:     param_topic,
	}

	productPaymentByte, err := productPayment.ToByte()
	if err != nil {
		log.Println(logTag+"Error parsing productPayment to byte: ", err.Error())
		returnSuccess(c)
		return
	}
	log.Printf(logTag+"Publishing message to RabbitMQ on topic <%s>", rmq.PreferenceNotificationQueue)
	rmqPublisher.PublishMessage(rmq.PreferenceNotificationQueue, string(productPaymentByte))
	returnSuccess(c)
}

func returnSuccess(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"notification": "received",
	})
}
