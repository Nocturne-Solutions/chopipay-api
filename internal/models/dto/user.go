package dto

import "chopipay/internal/models/entities"

type UserDTO struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}

func (u *UserDTO) FromUser(user *entities.User) *UserDTO {
	return &UserDTO{
		ID:       u.ID,
		Username: u.Username,
	}
}

