package pg

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type SalesRepository interface {
	Create(sale *entities.Sales) (*entities.Sales, error)
	FindById(id int) (*entities.Sales, error)
	Update(sale *entities.Sales) (*entities.Sales, error)
}

type SalesRepositoryImpl struct {
	Db *pg.DB
}

func NewSalesRepository(db *pg.DB) SalesRepository {
	return &SalesRepositoryImpl{
		Db: db,
	}
}

func (sr *SalesRepositoryImpl) Create(sale *entities.Sales) (*entities.Sales, error) {
	_, err := sr.Db.Model(sale).Insert(context.Background())
	if err != nil {
		return nil, errors.New("error creating sale. Cause: " + err.Error())
	}

	return sale, nil
}

func (sr *SalesRepositoryImpl) FindById(id int) (*entities.Sales, error) {
	sale := &entities.Sales{ID: id}
	err := sr.Db.Model(sale).WherePK().Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale by id " + strconv.Itoa(id) + ". Cause: " + err.Error())
	}

	return sale, nil
}

func (sr *SalesRepositoryImpl) Update(sale *entities.Sales) (*entities.Sales, error) {
	_, err := sr.Db.Model(sale).WherePK().Update(context.Background())
	if err != nil {
		return nil, errors.New("error updating sale. Cause: " + err.Error())
	}

	return sale, nil
}
