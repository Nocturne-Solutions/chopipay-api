package entities

type Shop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Personal *Personal `pg:"rel:has-one" json:"personal"`
}