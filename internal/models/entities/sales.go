package entities

type Sales struct {
	ID                int      `json:"id" pg:"id"`
	DateCreated       string   `json:"date_created" pg:"date_created"`
	LastModified      string   `json:"last_modified" pg:"last_modified"`
	TransactionAmount float64  `json:"transaction_amount" pg:"transaction_amount"`
	TotalPaidAmount   float64  `json:"total_paid_amount" pg:"total_paid_amount"`
	CurrencyID        string   `json:"currency_id" pg:"currency_id"`
	StatusID          int      `json:"status_id" pg:"status_id"`
	PersonalID        int      `json:"personal_id" pg:"personal_id"`
	Personal          Personal `json:"personal" pg:"-"`
}
