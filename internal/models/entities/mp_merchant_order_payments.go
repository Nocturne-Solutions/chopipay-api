package entities

type MpMerchantOrderPayments struct {
	ID                int             `json:"id" pg:"id,pk"`
	MpMerchantOrderID int             `json:"mp_merchant_order_id" pg:"mp_merchant_order_id"`
	MpMerchantOrder   MpMerchantOrder `json:"mp_merchant_order" pg:"rel:has-one, fk:mp_merchant_order_id, join_fk:id"`
	MpPaymentID       int             `json:"mp_payment_id" pg:"mp_payment_id"`
	MpPayment         MpPayment       `json:"mp_payment" pg:"rel:has-one, fk:mp_payment_id, join_fk:id"`
}
