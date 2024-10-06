package user

import (
	"log"

	"chopipay/internal/http/security"
	"chopipay/internal/models/entities"
	userRepository "chopipay/internal/repository/user"
)

func Create(user *entities.User) (error) {

	hashedPassword, err := security.HashPassword(user.Password)
	if err != nil {
		log.Println(err)
		return err
	}

	user.Password = hashedPassword

	err = userRepository.Create(user)
	if err != nil {
		log.Println(err)
		return err
	}
	log.Println("User created: ", user)
	return nil
}

func FindByID(id int) (*entities.User, error) {
	user, err := userRepository.FindByID(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

func Update(user *entities.User) (error) {
	err := userRepository.Update(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Delete(user *entities.User) (error) {
	err := userRepository.Delete(user)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func FindByUsername(username string) (*entities.User, error) {
	user, err := userRepository.FindByUsername(username)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return user, nil
}

