package entities

type User struct {
	ID       int    `json:"id" pg:"id,pk"`	
	Username string `json:"username" pg:"username"`
	Password string `json:"password" pg:"password"`
}