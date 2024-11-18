package mpmerchantorderpayments

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(mpMerchantOrderPayment *entities.MpMerchantOrderPayments) error {
	_, err := pg.Db.Model(mpMerchantOrderPayment).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp merchant order payment: " + err.Error())

	}

	return nil
}

func SaveAll(mpMerchantOrderPayments []entities.MpMerchantOrderPayments) error {
	_, err := pg.Db.Model(&mpMerchantOrderPayments).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp merchant order payments: " + err.Error())
	}

	return nil
}
