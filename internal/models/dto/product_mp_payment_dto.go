package dto

import (
	"encoding/json"
)

type ProductMpPaymentDTO struct {
	ID         int    `json:"id"`
	PersonalID int    `json:"personal_id"`
	Topic      string `json:"topic"`
}

func (p *ProductMpPaymentDTO) ToByte() ([]byte, error) {
	return json.Marshal(p)
}
