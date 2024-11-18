package services

import (
	productRepository "chopipay/internal/repository/pg"
	"errors"
	"log"

	"chopipay/internal/models/entities"
)

type ProductService interface {
	Create(product *entities.Product) error
	Update(product *entities.Product) error
	Delete(id int) error
	FindByID(id int) (*entities.Product, error)
	FindAllByShopID(shopID int) ([]entities.Product, error)
}

type ProductServiceImpl struct {
	logToken          string
	productRepository productRepository.ProductRepository
}

func NewProductService(productRepository productRepository.ProductRepository) ProductService {
	return &ProductServiceImpl{
		logToken:          "ProductServiceImpl | ",
		productRepository: productRepository,
	}
}

func (s *ProductServiceImpl) Create(product *entities.Product) error {
	err := s.productRepository.Create(product)
	if err != nil {
		errorMessage := s.logToken + "Error creating product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func (s *ProductServiceImpl) Update(product *entities.Product) error {
	err := s.productRepository.Update(product)
	if err != nil {
		errorMessage := s.logToken + "Error updating product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func (s *ProductServiceImpl) Delete(id int) error {
	err := s.productRepository.Delete(id)
	if err != nil {
		errorMessage := s.logToken + "Error deleting product: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}

	return nil
}

func (s *ProductServiceImpl) FindByID(id int) (*entities.Product, error) {
	product, err := s.productRepository.FindByID(id)
	if err != nil {
		errorMessage := s.logToken + "Error finding product by id " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return product, nil
}

func (s *ProductServiceImpl) FindAllByShopID(shopID int) ([]entities.Product, error) {
	products, err := s.productRepository.FindAllByShopID(shopID)
	if err != nil {
		errorMessage := s.logToken + "Error finding products by shop id " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return products, nil
}
