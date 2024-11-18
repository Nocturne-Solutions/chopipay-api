package sale_products

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(saleProducts *entities.SaleProducts) (*entities.SaleProducts, error) {
	_, err := pg.Db.Model(saleProducts).Insert(context.Background())
	if err != nil {
		return nil, errors.New("error creating sale products. Cause: " + err.Error())
	}

	return saleProducts, nil
}

func FindBySaleId(saleId int) ([]entities.SaleProducts, error) {
	var saleProducts []entities.SaleProducts
	err := pg.Db.Model(&saleProducts).Where("sale_id = ?", saleId).Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale products by sale id " + strconv.Itoa(saleId) + ". Cause: " + err.Error())
	}

	return saleProducts, nil
}
