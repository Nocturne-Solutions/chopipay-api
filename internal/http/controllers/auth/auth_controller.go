package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	authService "chopipay/internal/http/services/app/auth"
	"chopipay/internal/models/dto"
)

const logTag = "auth_controller | "

func Login(c *gin.Context) {
	var credentials dto.Login
	err := c.BindJSON(&credentials)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error binding credentials")
		return
	}

	token, err := authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error logging in")
		return
	}

	c.JSON(http.StatusOK, token)
}

func RefreshToken(c *gin.Context) {
	// get the refresh token from the body request as "refresh_token"
	var refreshToken dto.RefreshToken
	err := c.BindJSON(&refreshToken)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error binding token")
		return
	}

	newToken, err := authService.RefreshToken(refreshToken.Token)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error refreshing token")
		return
	}

	c.JSON(http.StatusOK, newToken)
}