package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"

	"chopipay/internal/models/entities"
)

type ShopRepository interface {
	GetByID(id int) (*entities.Shop, error)
	GetAllByPersonalId(personalId int) (*[]entities.Shop, error)
	Create(shop *entities.Shop) error
	Update(shop *entities.Shop) error
	Delete(id int) error
}

type ShopRepositoryImpl struct {
	Db *pg.DB
}

func NewShopRepository(db *pg.DB) ShopRepository {
	return &ShopRepositoryImpl{
		Db: db,
	}
}

func (r *ShopRepositoryImpl) GetByID(id int) (*entities.Shop, error) {
	shop := new(entities.Shop)
	err := r.Db.Model(shop).
		Where("id = ?", id).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting shop by ID: " + err.Error())
	}
	return shop, nil
}

func (r *ShopRepositoryImpl) GetAllByPersonalId(personalId int) (*[]entities.Shop, error) {
	shops := new([]entities.Shop)
	err := r.Db.Model(shops).
		Where("personal_id = ?", personalId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting shops by personal ID: " + err.Error())
	}
	return shops, nil
}

func (r *ShopRepositoryImpl) Create(shop *entities.Shop) error {
	_, err := r.Db.Model(shop).
		Returning("id").
		Insert(context.Background())

	if err != nil {
		return errors.New("Error creating shop: " + err.Error())
	}
	return nil
}

func (r *ShopRepositoryImpl) Update(shop *entities.Shop) error {
	_, err := r.Db.Model(shop).
		Where("id = ?", shop.ID).
		Update(context.Background())

	if err != nil {
		return errors.New("Error updating shop: " + err.Error())
	}
	return nil
}

func (r *ShopRepositoryImpl) Delete(id int) error {
	_, err := r.Db.Model(&entities.Shop{}).
		Where("id = ?", id).
		Delete(context.Background())

	if err != nil {
		return errors.New("Error deleting shop: " + err.Error())
	}
	return nil
}
