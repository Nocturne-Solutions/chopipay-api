package pg

import (
	"context"
	"errors"

	"github.com/go-pg/pg/v11"
	_ "github.com/go-pg/pg/v11/orm"

	"chopipay/internal/models/entities"
)

type UserRepository interface {
	Create(user *entities.User) error
	FindByID(id int) (*entities.User, error)
	Update(user *entities.User) error
	Delete(user *entities.User) error
	FindByUsername(username string) (*entities.User, error)
}

type UserRepositoryImpl struct {
	Db *pg.DB
}

func NewUserRepository(db *pg.DB) UserRepository {
	return &UserRepositoryImpl{
		Db: db,
	}
}

func (r *UserRepositoryImpl) Create(user *entities.User) error {
	_, err := r.Db.Model(user).Insert(context.Background())
	if err == nil {
		return nil
	}
	return errors.New("Error creating user: " + err.Error())
}

func (r *UserRepositoryImpl) FindByID(id int) (*entities.User, error) {
	user := &entities.User{ID: id}
	err := r.Db.Model(user).WherePK().Select(context.Background())
	if err != nil {
		return nil, errors.New("Error getting user by ID: " + err.Error())
	}
	return user, nil
}

func (r *UserRepositoryImpl) Update(user *entities.User) error {
	_, err := r.Db.Model(user).WherePK().Update(context.Background())
	if err != nil {
		return errors.New("Error updating user: " + err.Error())
	}
	return nil
}

func (r *UserRepositoryImpl) Delete(user *entities.User) error {
	_, err := r.Db.Model(user).WherePK().Delete(context.Background())
	if err != nil {
		return errors.New("Error deleting user: " + err.Error())
	}
	return nil
}

func (r *UserRepositoryImpl) FindByUsername(username string) (*entities.User, error) {
	user := &entities.User{Username: username}
	err := r.Db.Model(user).
		Where("username = ?", username).
		Select(context.Background())
	if err != nil {
		return nil, errors.New("Error getting user by username: " + err.Error())
	}
	return user, nil
}
