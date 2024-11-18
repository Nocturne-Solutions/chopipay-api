package services

import (
	shopRepository "chopipay/internal/repository/pg"
	"errors"
	"log"

	"chopipay/internal/models/entities"
)

type ShopService interface {
	Create(shop *entities.Shop) error
	GetByID(id int) (*entities.Shop, error)
	GetAllByPersonalId(personalId int) (*[]entities.Shop, error)
	Update(shop *entities.Shop) error
	Delete(id int) error
}

type ShopServiceImpl struct {
	logToken       string
	shopRepository shopRepository.ShopRepository
}

func NewShopService(shopRepository shopRepository.ShopRepository) ShopService {
	return &ShopServiceImpl{
		logToken:       "ShopService | ",
		shopRepository: shopRepository,
	}
}

func (s *ShopServiceImpl) Create(shop *entities.Shop) error {
	log.Println(s.logToken+"Creating shop: ", shop)

	err := s.shopRepository.Create(shop)
	if err != nil {
		errorMessage := s.logToken + "Error creating shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func (s *ShopServiceImpl) GetByID(id int) (*entities.Shop, error) {
	log.Println(s.logToken+"Getting shop by ID: ", id)

	shop, err := s.shopRepository.GetByID(id)
	if err != nil {
		errorMessage := s.logToken + "Error getting shop by ID: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	return shop, nil
}

func (s *ShopServiceImpl) GetAllByPersonalId(personalId int) (*[]entities.Shop, error) {
	log.Println(s.logToken+"Getting shops by personal ID: ", personalId)

	shops, err := s.shopRepository.GetAllByPersonalId(personalId)
	if err != nil {
		errorMessage := s.logToken + "Error getting shops by personal ID: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	return shops, nil
}

func (s *ShopServiceImpl) Update(shop *entities.Shop) error {
	log.Println(s.logToken+"Updating shop: ", shop)

	err := s.shopRepository.Update(shop)
	if err != nil {
		errorMessage := s.logToken + "Error updating shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func (s *ShopServiceImpl) Delete(id int) error {
	log.Println(s.logToken+"Deleting shop by ID: ", id)

	err := s.shopRepository.Delete(id)
	if err != nil {
		errorMessage := s.logToken + "Error deleting shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}
