package entities

import "github.com/mercadopago/sdk-go/pkg/payment"

type MpPayer struct {
	ID                   int    `json:"id" pg:"id,pk"`
	PayerID              string `json:"payer_id" pg:"payer_id"`
	Email                string `json:"email" pg:"email"`
	FirstName            string `json:"first_name" pg:"first_name"`
	LastName             string `json:"last_name" pg:"last_name"`
	IdentificationType   string `json:"identification_type" pg:"identification_type"`
	IdentificationNumber string `json:"identification_number" pg:"identification_number"`
}

func (p *MpPayer) NewFromMpPayerResponse(mpPayerResponse payment.PayerResponse) {
	p.PayerID = mpPayerResponse.ID
	p.Email = mpPayerResponse.Email
	p.FirstName = mpPayerResponse.FirstName
	p.LastName = mpPayerResponse.LastName
	p.IdentificationType = mpPayerResponse.Identification.Type
	p.IdentificationNumber = mpPayerResponse.Identification.Number
}
