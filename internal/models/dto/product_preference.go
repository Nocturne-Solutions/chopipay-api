package dto

type ProductPreferenceDTO struct {
	ID                  int     `json:"id"`
	Name                string  `json:"name"`
	Price               float64 `json:"price"`
	CurrencyID		  string  `json:"currency_id"`
	ShopID              int     `json:"shop_id"`
	Description         string  `json:"description"`
	PreferenceID        string  `json:"preference_id"`
	PaymentPoint        string  `json:"payment_point"`
	SandboxPaymentPoint string  `json:"sandbox_payment_point"`
	PictureURL          string  `json:"picture_url"`
}
