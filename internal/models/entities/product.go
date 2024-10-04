package entities

type Product struct {
	ID          int     `json:"id"`
	Name        string  `json:"name"`
	Price 	 float64 `json:"price"`
	Shop *Shop `pg:"rel:has-one" json:"shop"`
}