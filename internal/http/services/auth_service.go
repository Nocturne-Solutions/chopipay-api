package services

import (
	"errors"
	"log"

	"chopipay/internal/http/security"
	"chopipay/internal/models/dto"
)

type AuthService interface {
	Login(username, password string) (*dto.Jwt, error)
	RefreshToken(refreshToken string) (*dto.Jwt, error)
}

type AuthServiceImpl struct {
	logToken    string
	userService UserService
}

func NewAuthService(userService UserService) AuthService {
	return &AuthServiceImpl{
		logToken:    "AuthService | ",
		userService: userService,
	}
}

func (s *AuthServiceImpl) Login(username, password string) (*dto.Jwt, error) {
	user, err := s.userService.FindByUsername(username)
	if err != nil {
		errorMessage := s.logToken + "Error getting user by username: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	if !security.ComparePasswords(user.Password, password) {
		errorMessage := s.logToken + "Invalid password"
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	token, err := security.CreateAccessToken(username)
	if err != nil {
		errorMessage := s.logToken + "Error creating access token: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return token, nil
}

func (s *AuthServiceImpl) RefreshToken(refreshToken string) (*dto.Jwt, error) {
	token, err := security.RefreshToken(refreshToken)
	if err != nil {
		errorMessage := s.logToken + "Error refreshing token: " + err.Error()
		log.Println(errorMessage)
		return nil, errors.New(errorMessage)
	}

	return token, nil
}
