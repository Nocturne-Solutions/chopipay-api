package entities

import (
	"github.com/mercadopago/sdk-go/pkg/payment"
	"time"
)

type MpPayment struct {
	ID                        int           `json:"id" pg:"id,pk"`
	PaymentID                 int           `json:"payment_id" pg:"payment_id"`
	OrderID                   string        `json:"order_id" pg:"order_id"`
	OrderType                 string        `json:"order_type" pg:"order_type"`
	AuthorizationCode         string        `json:"authorization_code" pg:"authorization_code"`
	CollectorID               int           `json:"collector_id" pg:"collector_id"`
	TransactionAmount         float64       `json:"transaction_amount" pg:"transaction_amount"`
	TotalPaidAmount           float64       `json:"total_paid_amount" pg:"total_paid_amount"`
	TransactionAmountRefunded float64       `json:"transaction_amount_refunded" pg:"transaction_amount_refunded"`
	CurrencyID                string        `json:"currency_id" pg:"currency_id"`
	Status                    string        `json:"status" pg:"status"`
	StatusDetail              string        `json:"status_detail" pg:"status_detail"`
	OperationType             string        `json:"operation_type" pg:"operation_type"`
	DateApproved              time.Time     `json:"date_approved" pg:"date_approved"`
	DateCreated               time.Time     `json:"date_created" pg:"date_created"`
	LastModified              time.Time     `json:"last_modified" pg:"last_modified"`
	AmountRefunded            float64       `json:"amount_refunded" pg:"amount_refunded"`
	Installments              int           `json:"installments" pg:"installments"`
	InstallmentAmount         float64       `json:"installment_amount" pg:"installment_amount"`
	NetReceivedAmount         float64       `json:"net_received_amount" pg:"net_received_amount"`
	OverpaidAmount            float64       `json:"overpaid_amount" pg:"overpaid_amount"`
	Description               string        `json:"description" pg:"description"`
	MpPayerID                 int           `json:"mp_payer_id" pg:"mp_payer_id"`
	MpPayer                   MpPayer       `json:"mp_payer" pg:"rel:has-one, fk:mp_payer_id, join_fk:id"`
	PaymentMethodID           int           `json:"payment_method_id" pg:"payment_method_id"`
	PaymentMethod             PaymentMethod `json:"payment_method" pg:"rel:has-one, fk:payment_method_id, join_fk:id"`
}

func (p *MpPayment) NewFromMpPaymentResponse(mpPaymentResponse *payment.Response, mpPayer MpPayer, paymentMethod PaymentMethod) {
	p.PaymentID = mpPaymentResponse.ID
	p.OrderID = mpPaymentResponse.Order.ID
	p.OrderType = mpPaymentResponse.Order.Type
	p.AuthorizationCode = mpPaymentResponse.AuthorizationCode
	p.CollectorID = mpPaymentResponse.CollectorID
	p.TransactionAmount = mpPaymentResponse.TransactionAmount
	p.TotalPaidAmount = mpPaymentResponse.TransactionDetails.TotalPaidAmount
	p.TransactionAmountRefunded = mpPaymentResponse.TransactionAmountRefunded
	p.CurrencyID = mpPaymentResponse.CurrencyID
	p.Status = mpPaymentResponse.Status
	p.StatusDetail = mpPaymentResponse.StatusDetail
	p.OperationType = mpPaymentResponse.OperationType
	p.DateApproved = mpPaymentResponse.DateApproved
	p.DateCreated = mpPaymentResponse.DateCreated
	p.LastModified = mpPaymentResponse.DateLastUpdated
	p.AmountRefunded = mpPaymentResponse.TransactionAmountRefunded
	p.Installments = mpPaymentResponse.Installments
	p.InstallmentAmount = mpPaymentResponse.TransactionDetails.InstallmentAmount
	p.NetReceivedAmount = mpPaymentResponse.TransactionDetails.NetReceivedAmount
	p.OverpaidAmount = mpPaymentResponse.TransactionDetails.OverpaidAmount
	p.Description = mpPaymentResponse.Description
	p.MpPayer = mpPayer
	p.MpPayerID = mpPayer.ID
	p.PaymentMethod = paymentMethod
	p.PaymentMethodID = paymentMethod.ID
}
