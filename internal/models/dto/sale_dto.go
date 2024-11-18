package dto

import "chopipay/internal/models/entities"

type SaleDTO struct {
	SaleID              int          `json:"sale_id"`
	CurrencyID          string       `json:"currency_id"`
	PersonalID          int          `json:"personal_id"`
	ProductsDTO         []ProductDTO `json:"products"`
	TotalPaidAmount     float64      `json:"total_paid_amount"`
	PaymentPoint        string       `json:"payment_point"`
	SandboxPaymentPoint string       `json:"sandbox_payment_point"`
	StatusID            int          `json:"status_id"`
}

func (s *SaleDTO) newFromSale(sale *entities.Sales) *SaleDTO {
	return &SaleDTO{
		SaleID:          sale.ID,
		CurrencyID:      sale.CurrencyID,
		PersonalID:      sale.PersonalID,
		TotalPaidAmount: sale.TotalPaidAmount,
		StatusID:        sale.StatusID,
	}
}
