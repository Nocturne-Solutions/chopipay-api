package dto

type Login struct {
	Username string `json:"username" example:"username@gmail.com" description:"User email"`
	Password string `json:"password" example:"password" description:"User password"`
}