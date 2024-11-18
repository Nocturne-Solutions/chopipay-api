package pg

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type SaleProductsRepository interface {
	Create(saleProducts *entities.SaleProducts) (*entities.SaleProducts, error)
	FindBySaleId(saleId int) ([]entities.SaleProducts, error)
}

type SaleProductsRepositoryImpl struct {
	Db *pg.DB
}

func NewSaleProductsRepository(db *pg.DB) SaleProductsRepository {
	return &SaleProductsRepositoryImpl{
		Db: db,
	}
}

func (r *SaleProductsRepositoryImpl) Create(saleProducts *entities.SaleProducts) (*entities.SaleProducts, error) {
	_, err := r.Db.Model(saleProducts).Insert(context.Background())
	if err != nil {
		return nil, errors.New("error creating sale products. Cause: " + err.Error())
	}

	return saleProducts, nil
}

func (r *SaleProductsRepositoryImpl) FindBySaleId(saleId int) ([]entities.SaleProducts, error) {
	var saleProducts []entities.SaleProducts
	err := r.Db.Model(&saleProducts).Where("sale_id = ?", saleId).Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale products by sale id " + strconv.Itoa(saleId) + ". Cause: " + err.Error())
	}

	return saleProducts, nil
}
