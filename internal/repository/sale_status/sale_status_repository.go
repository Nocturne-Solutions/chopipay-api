package sale_status

import (
	"context"
	"errors"
	"strconv"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func FindByStatus(status string) (*entities.SaleStatus, error) {
	saleStatus := &entities.SaleStatus{Status: status}
	err := pg.Db.Model(saleStatus).Where("status = ?", status).Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale status by status " + status + ". Cause: " + err.Error())
	}

	return saleStatus, nil
}

func FindById(id int) (*entities.SaleStatus, error) {
	saleStatus := &entities.SaleStatus{ID: id}
	err := pg.Db.Model(saleStatus).WherePK().Select(context.Background())
	if err != nil {
		return nil, errors.New("error finding sale status by id " + strconv.Itoa(id) + ". Cause: " + err.Error())
	}

	return saleStatus, nil
}
