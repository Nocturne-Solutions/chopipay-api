package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type MpPaymentRepository interface {
	Create(mpPayment *entities.MpPayment) error
	GetByPaymentId(paymentId int) (*entities.MpPayment, error)
}

type MpPaymentRepositoryImpl struct {
	Db *pg.DB
}

func NewMpPaymentRepository(db *pg.DB) MpPaymentRepository {
	return &MpPaymentRepositoryImpl{Db: db}
}

func (r *MpPaymentRepositoryImpl) Create(mpPayment *entities.MpPayment) error {
	_, err := r.Db.Model(mpPayment).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp payment: " + err.Error())

	}

	return nil
}

func (r *MpPaymentRepositoryImpl) GetByPaymentId(paymentId int) (*entities.MpPayment, error) {
	mpPayment := &entities.MpPayment{PaymentID: paymentId}
	err := r.Db.Model(mpPayment).
		Where("payment_id = ?", paymentId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("error getting mp payment: " + err.Error())
	}

	return mpPayment, nil
}
