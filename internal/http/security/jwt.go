package security

import (
	"time"
	"errors"

	"chopipay/config/server"
	"chopipay/internal/models/dto"
	"github.com/golang-jwt/jwt/v5"
)

const accessTokenDuration = time.Minute * 120
const refreshTokenDuration = time.Minute * 130

func CreateAccessToken(username string) (*dto.Jwt, error) {
	accessToken, err := generateJwt(username, "access_token")
	if err != nil {
		return nil, err
	}

	refresToken, err := generateJwt(username, "refresh_token")
	if err != nil {
		return nil, err
	}

	iat := time.Now().Unix()
	expIn := accessTokenDuration.Seconds()

	return &dto.Jwt{
		AccessToken: accessToken,
		RefreshToken: refresToken,
		Iat: iat,
		ExpIn: expIn,
	}, nil
}

func RefreshToken(refreshToken string) (*dto.Jwt, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(refreshToken, 
							func(token *jwt.Token) (interface{}, error) {
								_, validMethod := token.Method.(*jwt.SigningMethodHMAC)
								
								if !validMethod {
									return nil, errors.New("invalid signing method")
								}
								
								return []byte(secretKey), nil
							})
	
	if err != nil {
		return nil, errors.New("invalid token")
	}

	tokenType := token.Claims.(jwt.MapClaims)["type"].(string)
	if tokenType != "refresh_token" {
		return nil, errors.New("invalid token type")
	}

	username := token.Claims.(jwt.MapClaims)["username"].(string)

	return CreateAccessToken(username)
}

func ValidateAcessToken(tokenString string) (bool, error) {
	secretKey, err := getSecretKey()
	if err != nil {
		return false, err
	}

	token, err := jwt.Parse(tokenString, 
							func(token *jwt.Token) (interface{}, error) {
								_, validMethod := token.Method.(*jwt.SigningMethodHMAC)
								
								if !validMethod {
									return nil, errors.New("invalid signing method")
								}
								
								return []byte(secretKey), nil
							})
	
	if err != nil {
		return false, errors.New("invalid token")
	}

	tokenType := token.Claims.(jwt.MapClaims)["type"].(string)
	if tokenType != "access_token" {
		return false, errors.New("invalid token type")
	}

	username := token.Claims.(jwt.MapClaims)["username"].(string)

	return (username != ""), err 
}

func generateJwt(username string, tokenType string) (string, error) {
	var exp int64
	if tokenType == "access_token" {
		exp = time.Now().Add(accessTokenDuration).Unix()
	} else {
		exp = time.Now().Add(refreshTokenDuration).Unix()
	}

	new_jwt := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"exp": exp,
		"iat": time.Now().Unix(),
		"type": tokenType,
	})

	secretKey, err := getSecretKey()
	if err != nil {
		return "", err
	}

	return new_jwt.SignedString([]byte(secretKey))
}

func getSecretKey() (string, error) {
	secretKey := server.EnvVars["JWT_SECRET_KEY"]
	if secretKey == "" {
		return "", errors.New("variable JWT_SECRET_KEY not found")
	}

	return secretKey, nil
}
