package entities

type Personal struct {
	ID       	int    	`json:"id" pg:"id,pk"`
	User  	 	*User   `json:"user" pg:"rel:has-one, fk:user_id, join_fk:id"`
	UserID 		int    	`json:"user_id"`
	FirstName 	string 	`json:"first_name"`
	LastName  	string 	`json:"last_name"`
	Email    	string 	`json:"email"`
	Phone    	string 	`json:"phone"`
}