package entities

type PaymentMethod struct {
	ID              int    `json:"id" pg:"id,pk"`
	PaymentMethodID string `json:"payment_method_id" pg:"payment_method_id"`
	PaymentTypeID   string `json:"payment_type_id" pg:"payment_type_id"`
}
