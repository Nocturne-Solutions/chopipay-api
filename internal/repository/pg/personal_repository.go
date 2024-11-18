package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type PersonalRepository interface {
	Create(personal *entities.Personal) error
	GetByID(id int) (*entities.Personal, error)
	Update(personal *entities.Personal) error
	Delete(personal *entities.Personal) error
	GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error)
	GetPersonalCredentialsByShopID(shopID int) (*entities.PersonalCredentials, error)
}

type PersonalRepositoryImpl struct {
	Db *pg.DB
}

func NewPersonalRepository(db *pg.DB) PersonalRepository {
	return &PersonalRepositoryImpl{
		Db: db,
	}
}

func (r *PersonalRepositoryImpl) Create(personal *entities.Personal) error {
	_, err := r.Db.Model(personal).
		Relation("User").
		Insert(context.Background())
	if err != nil {
		return errors.New("Error creating personal: " + err.Error())
	}
	return nil
}

func (r *PersonalRepositoryImpl) GetByID(id int) (*entities.Personal, error) {
	personal := &entities.Personal{}
	err := r.Db.Model(personal).
		Relation("User").
		Where("personal.id = ?", id).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting personal by ID: " + err.Error())
	}
	return personal, nil
}

func (r *PersonalRepositoryImpl) Update(personal *entities.Personal) error {
	_, err := r.Db.Model(personal).WherePK().Update(context.Background())
	if err != nil {
		return errors.New("Error updating personal: " + err.Error())
	}
	return nil
}

func (r *PersonalRepositoryImpl) Delete(personal *entities.Personal) error {
	_, err := r.Db.Model(personal).WherePK().Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting personal: " + err.Error())
	}
	return nil
}

func (r *PersonalRepositoryImpl) GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error) {
	personalCredentials := &entities.PersonalCredentials{}
	err := r.Db.Model(personalCredentials).
		Join("JOIN personals AS personal ON personal.id = personal_credentials.personal_id").
		Join("JOIN users AS usr ON usr.id = personal.user_id").
		Where("usr.username = ?", username).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting personal credentials by username: " + err.Error())
	}
	return personalCredentials, nil
}

func (r *PersonalRepositoryImpl) GetPersonalCredentialsByShopID(shopID int) (*entities.PersonalCredentials, error) {
	personalCredentials := &entities.PersonalCredentials{}
	err := r.Db.Model(personalCredentials).
		Join("JOIN personals AS personal ON personal.id = personal_credentials.personal_id").
		Join("JOIN shops AS shop ON shop.personal_id = personal.id").
		Where("shop.id = ?", shopID).
		Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting personal credentials by shopID: " + err.Error())
	}
	return personalCredentials, nil
}
