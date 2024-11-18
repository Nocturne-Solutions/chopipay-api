package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"

	"chopipay/internal/models/entities"
)

type ProductRepository interface {
	Create(product *entities.Product) error
	FindByID(id int) (*entities.Product, error)
	Update(product *entities.Product) error
	Delete(id int) error
	FindAllByShopID(shopID int) ([]entities.Product, error)
}

type ProductRepositoryImpl struct {
	Db *pg.DB
}

func NewProductRepository(db *pg.DB) ProductRepository {
	return &ProductRepositoryImpl{
		Db: db,
	}
}

func (r *ProductRepositoryImpl) Create(product *entities.Product) error {
	_, err := r.Db.Model(product).
		Insert(context.Background())
	if err != nil {
		return errors.New("Error creating product" + err.Error())
	}

	return nil
}

func (r *ProductRepositoryImpl) FindByID(id int) (*entities.Product, error) {
	product := new(entities.Product)
	err := r.Db.Model(product).
		Where("id = ?", id).
		Select(context.Background())
	if err != nil {
		return nil, errors.New("Error finding product by id " + err.Error())
	}

	return product, nil
}

func (r *ProductRepositoryImpl) Update(product *entities.Product) error {
	_, err := r.Db.Model(product).
		WherePK().
		Update(context.Background())
	if err != nil {
		return errors.New("Error updating product" + err.Error())
	}

	return nil
}

func (r *ProductRepositoryImpl) Delete(id int) error {
	product := new(entities.Product)
	_, err := r.Db.Model(product).
		Where("id = ?", id).
		Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting product" + err.Error())
	}

	return nil
}

func (r *ProductRepositoryImpl) FindAllByShopID(shopID int) ([]entities.Product, error) {
	var products []entities.Product
	err := r.Db.Model(&products).
		Where("shop_id = ?", shopID).
		Select(context.Background())
	if err != nil {
		return nil, errors.New("Error finding all products by shop id " + err.Error())
	}

	return products, nil
}
