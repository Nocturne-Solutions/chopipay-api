package auth

import (
	"errors"
	"log"

	"chopipay/internal/http/security"
	"chopipay/internal/http/services/app/user"
	"chopipay/internal/models/dto"
)

const logTag = "AuthService"

func Login(username, password string) (*dto.Jwt, error) {
	user, err := user.FindByUsername(username)
	if err != nil {
		errorMessage := logTag + "Error getting user by username: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	if !security.ComparePasswords(user.Password, password) {
		errorMessage := logTag + "Invalid password"
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	token, err := security.CreateAccessToken(username)
	if err != nil {
		errorMessage := logTag + "Error creating access token: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return token, nil
}

func RefreshToken(refreshToken string) (*dto.Jwt, error) {
	token, err := security.RefreshToken(refreshToken)
	if err != nil {
		errorMessage := logTag + "Error refreshing token: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return token, nil
}