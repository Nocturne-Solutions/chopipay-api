package entities

type Product struct {
	ID         int     `json:"id" pg:"id,pk"`
	Name       string  `json:"name" pg:"name"`
	Price      float64 `json:"price" pg:"price"`
	Shop       *Shop   `json:"shop" pg:"rel:has-one, fk:shop_id, join_fk:shop_id"`
	WebhookURL string  `json:"webhook_url" pg:"webhook_url"`
}
