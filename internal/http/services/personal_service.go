package services

import (
	personalRepository "chopipay/internal/repository/pg"
	"log"

	"chopipay/internal/models/entities"
)

type PersonalService interface {
	Create(personal *entities.Personal) error
	GetByID(id int) (*entities.Personal, error)
	Update(personal *entities.Personal) error
	Delete(personal *entities.Personal) error
	GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error)
	GetPersonalCredentialsByShopID(shopID int) (*entities.PersonalCredentials, error)
	GetPersonalCredentialsByPersonalId(personalId int) (*entities.PersonalCredentials, error)
}

type PersonalServiceImpl struct {
	logToken           string
	personalRepository personalRepository.PersonalRepository
}

func NewPersonalService(personalRepository personalRepository.PersonalRepository) PersonalService {
	return &PersonalServiceImpl{
		logToken:           "PersonalService | ",
		personalRepository: personalRepository,
	}
}

func (s *PersonalServiceImpl) Create(personal *entities.Personal) error {
	log.Println("Creating personal: ", personal)
	err := s.personalRepository.Create(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Personal created: ", personal)
	return nil
}

func (s *PersonalServiceImpl) GetByID(id int) (*entities.Personal, error) {
	personal, err := s.personalRepository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personal, nil
}

func (s *PersonalServiceImpl) Update(personal *entities.Personal) error {
	err := s.personalRepository.Update(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *PersonalServiceImpl) Delete(personal *entities.Personal) error {
	err := s.personalRepository.Delete(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *PersonalServiceImpl) GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error) {
	personalCredentials, err := s.personalRepository.GetPersonalCredentialsByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredentials, nil
}

func (s *PersonalServiceImpl) GetPersonalCredentialsByShopID(shopID int) (*entities.PersonalCredentials, error) {
	personalCredentials, err := s.personalRepository.GetPersonalCredentialsByShopID(shopID)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredentials, nil
}

func (s *PersonalServiceImpl) GetPersonalCredentialsByPersonalId(personalId int) (*entities.PersonalCredentials, error) {
	personalCredentials, err := s.personalRepository.GetPersonalCredentialsByPersonalId(personalId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredentials, nil
}
