package mppayer

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(mpPayer *entities.MpPayer) error {
	_, err := pg.Db.Model(mpPayer).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp payer: " + err.Error())

	}

	return nil
}

func FindByPayerId(payerId int) (*entities.MpPayer, error) {
	mpPayer := &entities.MpPayer{}

	err := pg.Db.Model(mpPayer).
		Where("payer_id = ?", payerId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error fetching mp payer with id " + strconv.Itoa(payerId) + ": " + err.Error())
	}

	return mpPayer, nil
}
