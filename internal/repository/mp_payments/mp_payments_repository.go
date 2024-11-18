package mppayments

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(mpPayment *entities.MpPayment) error {
	_, err := pg.Db.Model(mpPayment).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp payment: " + err.Error())

	}

	return nil
}

func GetByPaymentId(paymentId int) (*entities.MpPayment, error) {
	mpPayment := &entities.MpPayment{PaymentID: paymentId}
	err := pg.Db.Model(mpPayment).
		Where("payment_id = ?", paymentId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("error getting mp payment: " + err.Error())
	}

	return mpPayment, nil
}
