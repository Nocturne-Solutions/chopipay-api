package utils

import (
	"errors"
	"strings"

	"chopipay/config/server"
	"github.com/golang-jwt/jwt/v5"
	"github.com/gin-gonic/gin"
)

func GetCurrentUser(c *gin.Context) (string, error) {
	claims, err := getClaims(c)
	if err != nil {
		return "", err
	}

	username := claims["username"].(string)
	if username == "" {
		return "", errors.New("username not found in claims")
	}

	return username, nil
}

func getClaims(c *gin.Context) (jwt.MapClaims, error) {
	authHeader := c.GetHeader("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	secretKey, err := getSecretKey()
	if err != nil {
		return nil, errors.New("error getting secret key")
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
		return nil, errors.New("invalid token")
	}

	claims := token.Claims.(jwt.MapClaims)
	return claims, nil
}

func getSecretKey() (string, error) {
	secretKey := server.EnvVars["JWT_SECRET_KEY"]
	if secretKey == "" {
		return "", errors.New("variable JWT_SECRET_KEY not found")
	}

	return secretKey, nil
}
