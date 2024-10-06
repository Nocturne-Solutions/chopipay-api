package personal

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(personal *entities.Personal) error {
	_, err := pg.Db.Model(personal).
			Relation("User").
			Insert(context.Background())
	if err != nil {
		return errors.New("Error creating personal: " + err.Error())

	}
	return nil
}

func GetByID(id int) (*entities.Personal, error) {
	personal := &entities.Personal{}
	err := pg.Db.Model(personal).
				Relation("User").
				Where("personal.id = ?", id).
				Select(context.Background())
	
				if err != nil {
		return nil, errors.New("Error getting personal by ID: " + err.Error())
	}
	return personal, nil
}

func Update(personal *entities.Personal) error {
	_, err := pg.Db.Model(personal).WherePK().Update(context.Background())
	if err != nil {
		return errors.New("Error updating personal: " + err.Error())
	}
	return nil
}

func Delete(personal *entities.Personal) error {
	_, err := pg.Db.Model(personal).WherePK().Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting personal: " + err.Error())
	}
	return nil
}

func GetPersonalCredentialsByUsername(username string) (*entities.PersonalCredentials, error) {
	personalCredentials := &entities.PersonalCredentials{}
	err := pg.Db.Model(personalCredentials).
				Join("JOIN personals AS personal ON personal.id = personal_credentials.personal_id").
				Join("JOIN users AS usr ON usr.id = personal.user_id").
				Where("usr.username = ?", username).
				Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting personal credentials by username: " + err.Error())
	}
	return personalCredentials, nil
}

func GetPersonalCredentialsByShopID(shopID int) (*entities.PersonalCredentials, error) {
	personalCredentials := &entities.PersonalCredentials{}
	err := pg.Db.Model(personalCredentials).
				Join("JOIN personals AS personal ON personal.id = personal_credentials.personal_id").
				Join("JOIN shops AS shop ON shop.personal_id = personal.id").
				Where("shop.id = ?", shopID).
				Select(context.Background())

	if err != nil {
		return nil, errors.New("Error getting personal credentials by shopID: " + err.Error())
	}
	return personalCredentials, nil
}