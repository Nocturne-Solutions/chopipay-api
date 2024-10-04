package entities

type Shop struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Personal 	*[]Personal `json:"personal" pg:"rel:has-many, fk:personal_id, join_fk:personal_id"`
	PersonalID  int    `json:"personal_id" pg:"personal_id"`
}