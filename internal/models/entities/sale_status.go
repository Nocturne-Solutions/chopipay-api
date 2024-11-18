package entities

type SaleStatus struct {
	ID     int    `json:"id" pg:"id"`
	Status string `json:"status" pg:"status"`
}
