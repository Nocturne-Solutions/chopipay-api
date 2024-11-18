package pg

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type SalesStatusRepository interface {
	FindByStatus(status string) (*entities.SaleStatus, error)
	FindById(id int) (*entities.SaleStatus, error)
}

type SalesStatusRepositoryImpl struct {
	Db *pg.DB
}

func NewSalesStatusRepository(db *pg.DB) SalesStatusRepository {
	return &SalesStatusRepositoryImpl{
		Db: db,
	}
}

func (r *SalesStatusRepositoryImpl) FindByStatus(status string) (*entities.SaleStatus, error) {
	saleStatus := &entities.SaleStatus{Status: status}
	err := r.Db.Model(saleStatus).Where("status = ?", status).Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale status by status " + status + ". Cause: " + err.Error())
	}

	return saleStatus, nil
}

func (r *SalesStatusRepositoryImpl) FindById(id int) (*entities.SaleStatus, error) {
	saleStatus := &entities.SaleStatus{ID: id}
	err := r.Db.Model(saleStatus).WherePK().Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale status by id " + strconv.Itoa(id) + ". Cause: " + err.Error())
	}

	return saleStatus, nil
}
