package entities

type SaleProducts struct {
	ID        int `json:"id" pg:"id"`
	SaleID    int `json:"sale_id" pg:"sale_id"`
	ProductID int `json:"product_id" pg:"product_id"`
	Quantity  int `json:"quantity" pg:"quantity"`

	Product Product `json:"product" pg:"-"`
	Sale    Sales   `json:"sales" pg:"-"`
}
