package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	authService "chopipay/internal/http/services/app/auth"
	"chopipay/internal/models/dto"
)

const logTag = "auth_controller | "

// AuthRoutes godoc
// @Tags Auth
// @Summary Login
// @Description Login with email and password
// @Accept  json
// @Produce  json
// @Param credentials body dto.Login true "Login data"
// @Success 200 {object} dto.Jwt
// @Router /auth/login [post]
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

// AuthRoutes godoc
// @Tags Auth
// @Summary Refresh token
// @Description Refresh token with refresh token
// @Param refreshToken body dto.RefreshToken true "Refresh token data"
// @Accept  json
// @Produce  json
// @Router /auth/refresh-token [post]
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