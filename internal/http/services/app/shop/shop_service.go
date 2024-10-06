package shop

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"
	shopRepository "chopipay/internal/repository/shop"
)

const logTag = "ShopService | "

func Create(shop *entities.Shop) error {
	log.Println(logTag + "Creating shop: ", shop)

	err := shopRepository.Create(shop)
	if err != nil {
		errorMessage := logTag + "Error creating shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func GetByID(id int) (*entities.Shop, error) {
	log.Println(logTag + "Getting shop by ID: ", id)

	shop, err := shopRepository.GetByID(id)
	if err != nil {
		errorMessage := logTag + "Error getting shop by ID: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	return shop, nil
}

func GetAllByPersonalId(personalId int) (*[]entities.Shop, error) {
	log.Println(logTag + "Getting shops by personal ID: ", personalId)

	shops, err := shopRepository.GetAllByPersonalId(personalId)
	if err != nil {
		errorMessage := logTag + "Error getting shops by personal ID: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}
	return shops, nil
}

func Update(shop *entities.Shop) error {
	log.Println(logTag + "Updating shop: ", shop)

	err := shopRepository.Update(shop)
	if err != nil {
		errorMessage := logTag + "Error updating shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}

func Delete(id int) error {
	log.Println(logTag + "Deleting shop by ID: ", id)

	err := shopRepository.Delete(id)
	if err != nil {
		errorMessage := logTag + "Error deleting shop: " + err.Error()
		log.Println(errorMessage)
		return errors.New(errorMessage)
	}
	return nil
}