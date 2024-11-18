package merchantorder

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"
	merchantOrderPaymentsRepository "chopipay/internal/repository/mp_merchantorder_payments"
	mpPaymentRepository "chopipay/internal/repository/mp_payments"

	"github.com/mercadopago/sdk-go/pkg/merchantorder"
)

func ProcessMerchantOrder(mpMerchantOrderResponse *merchantorder.Response) error {
	log.Println("Processing merchant order...")
	if mpMerchantOrderResponse == nil {
		log.Println("Could not process merchant order, it's nil")
		return errors.New("error processing mercadopago merchant order: not found")
	}

	newMerchantOrder := entities.MpMerchantOrder{}
	newMerchantOrder.NewFromMpMerchantOrderResponse(mpMerchantOrderResponse)

	var newMpMerchantOrderPayments []entities.MpMerchantOrderPayments
	for _, mpPayment := range mpMerchantOrderResponse.Payments {
		mpPayment, err := mpPaymentRepository.GetByPaymentId(mpPayment.ID)
		if err != nil {
			log.Println("Error getting payment: ", err.Error())
			continue
		}
		newMpMerchantOrderPayment := entities.MpMerchantOrderPayments{
			MpPaymentID:       mpPayment.ID,
			MpMerchantOrderID: newMerchantOrder.ID,
		}
		newMpMerchantOrderPayments = append(newMpMerchantOrderPayments, newMpMerchantOrderPayment)
	}

	err := merchantOrderPaymentsRepository.SaveAll(newMpMerchantOrderPayments)
	if err != nil {
		log.Println("Error saving merchant order payments: ", err.Error())
		return errors.New("error saving merchant order payments: " + err.Error())
	}

	return nil
}
