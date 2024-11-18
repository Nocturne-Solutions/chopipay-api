package dto

type ProductPreferenceDTO struct {
	ProductsDTO         []ProductDTO `json:"products"`
	CurrencyID          string       `json:"currency_id"`
	PreferenceID        string       `json:"preference_id"`
	PaymentPoint        string       `json:"payment_point"`
	SandboxPaymentPoint string       `json:"sandbox_payment_point"`
	PictureURL          string       `json:"picture_url"`
}
