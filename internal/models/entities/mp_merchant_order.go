package entities

import (
	"github.com/mercadopago/sdk-go/pkg/merchantorder"
	"time"
)

type MpMerchantOrder struct {
	ID                int       `json:"id" pg:"id,pk"`
	MerchantOrderID   int       `json:"merchant_order_id" pg:"merchant_order_id"`
	Status            string    `json:"status" pg:"status"`
	ExternalReference string    `json:"external_reference" pg:"external_reference"`
	PreferenceID      string    `json:"preference_id" pg:"preference_id"`
	CollectorID       int       `json:"collector_id" pg:"collector_id"`
	DateCreated       time.Time `json:"date_created" pg:"date_created"`
	LastUpdated       time.Time `json:"last_updated" pg:"last_updated"`
	TotalAmount       float64   `json:"total_amount" pg:"total_amount"`
	PaidAmount        float64   `json:"paid_amount" pg:"paid_amount"`
	RefundedAmount    float64   `json:"refunded_amount" pg:"refunded_amount"`
	PayerID           int       `json:"payer_id" pg:"payer_id"`
	Cancelled         bool      `json:"cancelled" pg:"cancelled"`
	IsTest            bool      `json:"is_test" pg:"is_test"`
	OrderStatus       string    `json:"order_status" pg:"order_status"`
}

func (m *MpMerchantOrder) NewFromMpMerchantOrderResponse(mpMerchantOrderResponse *merchantorder.Response) {
	m.MerchantOrderID = mpMerchantOrderResponse.ID
	m.Status = mpMerchantOrderResponse.Status
	m.ExternalReference = mpMerchantOrderResponse.ExternalReference
	m.PreferenceID = mpMerchantOrderResponse.PreferenceID
	m.CollectorID = mpMerchantOrderResponse.Collector.ID
	m.DateCreated = mpMerchantOrderResponse.DateCreated
	m.LastUpdated = mpMerchantOrderResponse.LastUpdated
	m.TotalAmount = mpMerchantOrderResponse.TotalAmount
	m.PaidAmount = mpMerchantOrderResponse.PaidAmount
	m.RefundedAmount = mpMerchantOrderResponse.RefundedAmount
	m.PayerID = mpMerchantOrderResponse.Payer.ID
	m.Cancelled = mpMerchantOrderResponse.Cancelled
	m.OrderStatus = mpMerchantOrderResponse.OrderStatus
}
