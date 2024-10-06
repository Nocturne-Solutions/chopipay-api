package user

import (
	"context"
	"errors"

	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/config/db/pg"
	"chopipay/internal/models/entities"
)

func Create(user *entities.User) (error) {
	_, err := pg.Db.Model(user).Insert(context.Background())
	if err == nil {
		return nil
	}
	return errors.New("Error creating user: " + err.Error())
}

func FindByID(id int) (*entities.User, error) {
	user := &entities.User{ID: id}
	err := pg.Db.Model(user).WherePK().Select(context.Background())
	if err != nil {
		return nil, errors.New("Error getting user by ID: " + err.Error())
	}
	return user, nil
}

func Update(user *entities.User) (error) {
	_, err := pg.Db.Model(user).WherePK().Update(context.Background())
	if err != nil {
		return errors.New("Error updating user: " + err.Error())
	}
	return nil
}

func Delete(user *entities.User) (error) {
	_, err := pg.Db.Model(user).WherePK().Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting user: " + err.Error())
	}
	return nil
}

func FindByUsername(username string) (*entities.User, error) {
	user := &entities.User{Username: username}
	err := pg.Db.Model(user).
				Where("username = ?", username).
				Select(context.Background())
	if err != nil {
		return nil, errors.New("Error getting user by username: " + err.Error())
	}
	return user, nil
}