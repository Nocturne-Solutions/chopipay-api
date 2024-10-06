package credentials

import (
	"errors"
	"log"

	"chopipay/internal/models/entities"
	credentialsRepository "chopipay/internal/repository/credentials"
	personalRepisotory "chopipay/internal/repository/personal"
)

const logTag = "PersonalCredentialsService | "

func AddPersonalCredential(personalCredential *entities.PersonalCredentials) error {
	log.Println(logTag + "Adding personal credential: ", personalCredential)

	if personalCredential.PersonalID <= 0 {
		errorMsg := logTag + "Invalid personal ID"
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	_, err := personalRepisotory.GetByID(personalCredential.PersonalID)
	if err != nil {
		errorMsg := logTag + "Error getting personal by ID: " + err.Error()
		log.Println(errorMsg)
		return errors.New(errorMsg)
	}

	err = credentialsRepository.AddPersonalCredential(personalCredential)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetPersonalCredentialByPersonalId(personalId int, credentialTypeId int) (*entities.PersonalCredentials, error) {
	log.Println(logTag + "Getting personal credential by personal ID: ", personalId)

	if personalId <= 0 {
		errorMsg := logTag + "Invalid personal ID"
		log.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	if credentialTypeId <= 0 {
		errorMsg := logTag + "Invalid credential type ID"
		log.Println(errorMsg)
		return nil, errors.New(errorMsg)
	}

	personalCredential, err := credentialsRepository.GetPersonalCredentialByPersonalId(personalId, credentialTypeId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return personalCredential, nil
}

