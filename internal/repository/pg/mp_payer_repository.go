package pg

import (
	"context"
	"errors"
	"strconv"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type MpPayerRepository interface {
	Create(mpPayer *entities.MpPayer) error
	FindByPayerId(payerId int) (*entities.MpPayer, error)
}

type MpPayerRepositoryImpl struct {
	Db *pg.DB
}

func NewMpPayerRepository(db *pg.DB) MpPayerRepository {
	return &MpPayerRepositoryImpl{Db: db}
}

func (r *MpPayerRepositoryImpl) Create(mpPayer *entities.MpPayer) error {
	_, err := r.Db.Model(mpPayer).Insert(context.Background())
	if err != nil {
		return errors.New("Error creating mp payer: " + err.Error())

	}

	return nil
}

func (r *MpPayerRepositoryImpl) FindByPayerId(payerId int) (*entities.MpPayer, error) {
	mpPayer := &entities.MpPayer{}

	err := r.Db.Model(mpPayer).
		Where("payer_id = ?", payerId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error fetching mp payer with id " + strconv.Itoa(payerId) + ": " + err.Error())
	}

	return mpPayer, nil
}
