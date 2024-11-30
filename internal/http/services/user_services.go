package services

import (
	"chopipay/internal/repository/pg"
	"log"

	"chopipay/internal/http/security"
	"chopipay/internal/models/entities"
)

type UserService interface {
	Create(user *entities.User) error
	FindByID(id int) (*entities.User, error)
	Update(user *entities.User) error
	Delete(user *entities.User) error
	FindByUsername(username string) (*entities.User, error)
}

type UserServiceImpl struct {
	logToken       string
	userRepository pg.UserRepository
}

func NewUserService(userRepository pg.UserRepository) UserService {
	return &UserServiceImpl{
		logToken:       "UserService | ",
		userRepository: userRepository,
	}
}

func (s *UserServiceImpl) Create(user *entities.User) error {

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	user.Password = hashedPassword

	err = s.userRepository.Create(user)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("User created: ", user)
	return nil
}

func (s *UserServiceImpl) FindByID(id int) (*entities.User, error) {
	user, err := s.userRepository.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func (s *UserServiceImpl) Update(user *entities.User) error {
	err := s.userRepository.Update(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *UserServiceImpl) Delete(user *entities.User) error {
	err := s.userRepository.Delete(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *UserServiceImpl) FindByUsername(username string) (*entities.User, error) {
	user, err := s.userRepository.FindByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}
