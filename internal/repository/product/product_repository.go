package product

import (
	"context"
	"errors"

	"chopipay/internal/models/entities"
	"chopipay/config/db/pg"
)

func Create(product *entities.Product) error {
	_, err := pg.Db.Model(product).
					Insert(context.Background())
	if err != nil {
		return errors.New("Error creating product" + err.Error())
	}

	return nil
}

func FindByID(id int) (*entities.Product, error) {
	product := new(entities.Product)
	err := pg.Db.Model(product).
					Where("id = ?", id).
					Select(context.Background())
	if err != nil {
		return nil, errors.New("Error finding product by id " + err.Error())
	}

	return product, nil
}

func Update(product *entities.Product) error {
	_, err := pg.Db.Model(product).
					WherePK().
					Update(context.Background())
	if err != nil {
		return errors.New("Error updating product" + err.Error())
	}

	return nil
}

func Delete(id int) error {
	product := new(entities.Product)
	_, err := pg.Db.Model(product).
					Where("id = ?", id).
					Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting product" + err.Error())
	}

	return nil
}

func FindAllByShopID(shopID int) ([]entities.Product, error) {
	var products []entities.Product
	err := pg.Db.Model(&products).
					Where("shop_id = ?", shopID).
					Select(context.Background())
	if err != nil {
		return nil, errors.New("Error finding all products by shop id " + err.Error())
	}

	return products, nil
}