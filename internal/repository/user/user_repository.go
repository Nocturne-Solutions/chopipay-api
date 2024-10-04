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
	return errors.New("Error creating user: " + err.Error())
}

func FindByID(id int) (*entities.User, error) {
	user := &entities.User{ID: id}
	err := pg.Db.Model(user).WherePK().Select(context.Background())
	return user, errors.New("Error finding user: " + err.Error())
}

func Update(user *entities.User) (error) {
	_, err := pg.Db.Model(user).WherePK().Update(context.Background())
	return errors.New("Error updating user: " + err.Error())
}

func Delete(user *entities.User) (error) {
	_, err := pg.Db.Model(user).WherePK().Delete(context.Background())
	return errors.New("Error deleting user: " + err.Error())
}