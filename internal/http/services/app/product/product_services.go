package product

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"
	productRepository "chopipay/internal/repository/product"
)

const logTag = "product_services | "

func Create(product *entities.Product) error {
	err := productRepository.Create(product)
	if err != nil {
		errorMessage := logTag + "Error creating product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func Update(product *entities.Product, isPreference bool) error {
	err := productRepository.Update(product)
	if err != nil {
		errorMessage := logTag + "Error updating product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func Delete(id int) error {
	err := productRepository.Delete(id)
	if err != nil {
		errorMessage := logTag + "Error deleting product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func FindByID(id int) (*entities.Product, error) {
	product, err := productRepository.FindByID(id)
	if err != nil {
		errorMessage := logTag + "Error finding product by id " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return product, nil
}

func FindAllByShopID(shopID int) ([]entities.Product, error) {
	products, err := productRepository.FindAllByShopID(shopID)
	if err != nil {
		errorMessage := logTag + "Error finding products by shop id " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return products, nil
}