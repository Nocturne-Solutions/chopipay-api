package credentials

import (
	"context"
	"errors"

	"chopipay/internal/models/entities"
	"chopipay/config/db/pg"
)

func AddPersonalCredential(personalCredential *entities.PersonalCredentials) error {
	_, err := pg.Db.Model(personalCredential).Insert(context.Background())
	if err != nil {
		return errors.New("error while inserting personal credential")
	}
	return nil
}

func GetPersonalCredentialByPersonalId(personalId int, credentialTypeId int) (*entities.PersonalCredentials, error) {
	personalCredential := &entities.PersonalCredentials{}
	
	err := pg.Db.Model(personalCredential).
					Where("personal_id = ?", personalId).
					Where("credential_type = ?", credentialTypeId).
					Select(context.Background())

	if err != nil {
		return nil, errors.New("error while fetching personal credential")
	}
	return personalCredential, nil
}