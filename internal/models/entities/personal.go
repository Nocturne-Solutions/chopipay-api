package entities

type Personal struct {
	ID       int    `json:"id"`
	User  	*User    `pg:"rel:has-one" json:"user"`
	FirtsName string `json:"firts_name"`
	LastName string `json:"last_name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}