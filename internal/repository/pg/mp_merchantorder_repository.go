package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type MpMerchantOrderRepository interface {
	Create(mpMerchantOrder *entities.MpMerchantOrder) error
}

type MpMerchantOrderRepositoryImpl struct {
	Db *pg.DB
}

func NewMpMerchantOrderRepository(db *pg.DB) MpMerchantOrderRepository {
	return &MpMerchantOrderRepositoryImpl{
		Db: db,
	}
}

func (r *MpMerchantOrderRepositoryImpl) Create(mpMerchantOrder *entities.MpMerchantOrder) error {
	_, err := r.Db.Model(mpMerchantOrder).Insert(context.Background())
	if err != nil {
		return errors.New("error creating mp merchant order: " + err.Error())
	}

	return nil
}
