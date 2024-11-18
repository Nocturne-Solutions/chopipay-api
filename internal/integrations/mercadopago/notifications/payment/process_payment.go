package merchantorder

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"

	"github.com/mercadopago/sdk-go/pkg/payment"
)

func ProcessPayment(mpPaymentResponse *payment.Response) error {
	log.Println("Processing payment...")
	if mpPaymentResponse == nil {
		log.Println("Could not process payment, it's nil")
		return errors.New("error processing mercadopago payment: not found")
	}

	// TODO crete asyncronous process to save paymentMethod and payer
	paymentMethod := entities.PaymentMethod{
		PaymentMethodID: mpPaymentResponse.PaymentMethodID,
		PaymentTypeID:   mpPaymentResponse.PaymentTypeID,
	}

	payer := entities.MpPayer{}
	payer.NewFromMpPayerResponse(mpPaymentResponse.Payer)

	// TODO then save the payment

	newMpPayment := &entities.MpPayment{}
	newMpPayment.NewFromMpPaymentResponse(mpPaymentResponse, payer, paymentMethod)

	return nil
}
