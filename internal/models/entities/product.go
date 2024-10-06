package entities

type Product struct {
	ID         int     `json:"id" pg:"id,pk"`
	Name       string  `json:"name" pg:"name"`
	Price      float64 `json:"price" pg:"price"`
	Shop       *Shop   `json:"shop" pg:"rel:has-one, fk:shop_id, join_fk:shop_id"`
	ShopID     int     `json:"shop_id" pg:"shop_id"`
	Description string `json:"description" pg:"description"`
	PreferenceID string `json:"preference_id" pg:"preference_id"`

}
