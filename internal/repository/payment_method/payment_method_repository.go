package paymentmethod

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(paymentMethod *entities.PaymentMethod) error {
	_, err := pg.Db.Model(paymentMethod).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating payment method: " + err.Error())

	}

	return nil
}

func FindByPaymentMethodId(paymentMethodId string) (*entities.PaymentMethod, error) {
	paymentMethod := &entities.PaymentMethod{}

	err := pg.Db.Model(paymentMethod).
		Where("payment_method_id = ?", paymentMethodId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error fetching payment method with id " + paymentMethodId + ": " + err.Error())
	}

	return paymentMethod, nil
}
