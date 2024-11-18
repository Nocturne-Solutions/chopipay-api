package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type MpMerchantOrderPaymentsRepository interface {
	Create(mpMerchantOrderPayment *entities.MpMerchantOrderPayments) error
	SaveAll(mpMerchantOrderPayments []entities.MpMerchantOrderPayments) error
}

type MpMerchantOrderPaymentsRepositoryImpl struct {
	Db *pg.DB
}

func NewMpMerchantOrderPaymentsRepository(db *pg.DB) MpMerchantOrderPaymentsRepository {
	return &MpMerchantOrderPaymentsRepositoryImpl{Db: db}
}

func (r *MpMerchantOrderPaymentsRepositoryImpl) Create(mpMerchantOrderPayment *entities.MpMerchantOrderPayments) error {
	_, err := r.Db.Model(mpMerchantOrderPayment).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp merchant order payment: " + err.Error())

	}

	return nil
}

func (r *MpMerchantOrderPaymentsRepositoryImpl) SaveAll(mpMerchantOrderPayments []entities.MpMerchantOrderPayments) error {
	_, err := r.Db.Model(&mpMerchantOrderPayments).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp merchant order payments: " + err.Error())
	}

	return nil
}
