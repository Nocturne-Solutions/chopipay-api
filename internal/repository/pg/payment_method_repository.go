package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type PaymentMethodRepository interface {
	Create(paymentMethod *entities.PaymentMethod) error
	FindByPaymentMethodId(paymentMethodId string) (*entities.PaymentMethod, error)
}

type PaymentMethodRepositoryImpl struct {
	Db *pg.DB
}

func NewPaymentMethodRepository(db *pg.DB) PaymentMethodRepository {
	return &PaymentMethodRepositoryImpl{
		Db: db,
	}
}

func (r *PaymentMethodRepositoryImpl) Create(paymentMethod *entities.PaymentMethod) error {
	_, err := r.Db.Model(paymentMethod).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating payment method: " + err.Error())

	}

	return nil
}

func (r *PaymentMethodRepositoryImpl) FindByPaymentMethodId(paymentMethodId string) (*entities.PaymentMethod, error) {
	paymentMethod := &entities.PaymentMethod{}

	err := r.Db.Model(paymentMethod).
		Where("payment_method_id = ?", paymentMethodId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error fetching payment method with id " + paymentMethodId + ": " + err.Error())
	}

	return paymentMethod, nil
}
