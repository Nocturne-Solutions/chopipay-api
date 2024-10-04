package user

import (
	"log"

	"chopipay/internal/models/entities"
	userRepository "chopipay/internal/repository/user"
)

func Create(user *entities.User) (error) {
	err := userRepository.Create(user)
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

