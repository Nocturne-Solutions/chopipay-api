package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"

	"chopipay/internal/models/entities"
)

type PersonalCredentialsRepository interface {
	AddPersonalCredential(personalCredential *entities.PersonalCredentials) error
	GetPersonalCredentialByPersonalId(personalId int, credentialTypeId int) (*entities.PersonalCredentials, error)
}

type PersonalCredentialsRepositoryImpl struct {
	Db *pg.DB
}

func NewPersonalCredentialsRepository(db *pg.DB) PersonalCredentialsRepository {
	return &PersonalCredentialsRepositoryImpl{
		Db: db,
	}
}

func (r *PersonalCredentialsRepositoryImpl) AddPersonalCredential(personalCredential *entities.PersonalCredentials) error {
	_, err := r.Db.Model(personalCredential).Insert(context.Background())
	if err != nil {
		return errors.New("error while inserting personal credential")
	}
	return nil
}

func (r *PersonalCredentialsRepositoryImpl) GetPersonalCredentialByPersonalId(personalId int, credentialTypeId int) (*entities.PersonalCredentials, error) {
	personalCredential := &entities.PersonalCredentials{}

	err := r.Db.Model(personalCredential).
		Where("personal_id = ?", personalId).
		Where("credential_type = ?", credentialTypeId).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("error while fetching personal credential")
	}
	return personalCredential, nil
}
