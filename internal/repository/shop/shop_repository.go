package shop

import (
	"errors"
	"context"

	"chopipay/internal/models/entities"
	"chopipay/config/db/pg"
)

func GetByID(id int) (*entities.Shop, error) {
	shop := new(entities.Shop)
	err := pg.Db.Model(shop).
				Where("id = ?", id).
				Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting shop by ID: " + err.Error())
	}
	return shop, nil
}

func GetAllByPersonalId(personalId int) (*[]entities.Shop, error) {
	shops := new([]entities.Shop)
	err := pg.Db.Model(shops).
				Where("personal_id = ?", personalId).
				Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting shops by personal ID: " + err.Error())
	}
	return shops, nil
}

func Create(shop *entities.Shop) error {
	_, err := pg.Db.Model(shop).
				Returning("id").
				Insert(context.Background())

	if err != nil {
		return errors.New("Error creating shop: " + err.Error())
	}
	return nil
}

func Update(shop *entities.Shop) error {
	_, err := pg.Db.Model(shop).
				Where("id = ?", shop.ID).
				Update(context.Background())

	if err != nil {
		return errors.New("Error updating shop: " + err.Error())
	}
	return nil
}

func Delete(id int) error {
	_, err := pg.Db.Model(&entities.Shop{}).
				Where("id = ?", id).
				Delete(context.Background())

	if err != nil {
		return errors.New("Error deleting shop: " + err.Error())
	}
	return nil
}