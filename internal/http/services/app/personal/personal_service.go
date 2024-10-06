package personal

import (
	"log"

	"chopipay/internal/models/entities"
	personalRepository "chopipay/internal/repository/personal"
)

func Create(personal *entities.Personal) error {
	log.Println("Creating personal: ", personal)
	err := personalRepository.Create(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("Personal created: ", personal)
	return nil
}

func GetByID(id int) (*entities.Personal, error) {
	personal, err := personalRepository.GetByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personal, nil
}

func Update(personal *entities.Personal) error {
	err := personalRepository.Update(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Delete(personal *entities.Personal) error {
	err := personalRepository.Delete(personal)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error) {
	personalCredentials, err := personalRepository.GetPersonalCredentialsByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredentials, nil
}