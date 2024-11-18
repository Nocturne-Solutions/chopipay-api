package mpmerchantorder

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(mpMerchantOrder *entities.MpMerchantOrder) error {
	_, err := pg.Db.Model(mpMerchantOrder).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp merchant order: " + err.Error())

	}

	return nil
}
