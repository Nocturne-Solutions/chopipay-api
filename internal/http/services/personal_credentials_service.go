package services

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"
	pgRepo "chopipay/internal/repository/pg"
)

type PersonalCredentialsService interface {
	AddPersonalCredential(personalCredential *entities.PersonalCredentials) error
	GetPersonalCredentialByPersonalId(personalId int, credentialTypeId int) (*entities.PersonalCredentials, error)
}

type PersonalCredentialsServiceImpl struct {
	logToken              string
	personalRepository    pgRepo.PersonalRepository
	credentialsRepository pgRepo.PersonalCredentialsRepository
}

func NewPersonalCredentialsService(
	personalRepository pgRepo.PersonalRepository,
	credentialsRepository pgRepo.PersonalCredentialsRepository) PersonalCredentialsService {
	return &PersonalCredentialsServiceImpl{
		logToken:              "PersonalCredentialsService | ",
		personalRepository:    personalRepository,
		credentialsRepository: credentialsRepository,
	}
}

func (s *PersonalCredentialsServiceImpl) AddPersonalCredential(personalCredential *entities.PersonalCredentials) error {
	log.Println(s.logToken+"Adding personal credential: ", personalCredential)

	if personalCredential.PersonalID <= 0 {
		errorMsg := s.logToken + "Invalid personal ID"
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	_, err := s.personalRepository.GetByID(personalCredential.PersonalID)
	if err != nil {
		errorMsg := s.logToken + "Error getting personal by ID: " + err.Error()
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	err = s.credentialsRepository.AddPersonalCredential(personalCredential)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *PersonalCredentialsServiceImpl) GetPersonalCredentialByPersonalId(
	personalId int,
	credentialTypeId int) (*entities.PersonalCredentials, error) {
	log.Println(s.logToken+"Getting personal credential by personal ID: ", personalId)

	if personalId <= 0 {
		errorMsg := s.logToken + "Invalid personal ID"
		log.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	if credentialTypeId <= 0 {
		errorMsg := s.logToken + "Invalid credential type ID"
		log.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	personalCredential, err := s.credentialsRepository.GetPersonalCredentialByPersonalId(personalId, credentialTypeId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredential, nil
}
