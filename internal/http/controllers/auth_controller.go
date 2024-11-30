package controllers

import (
	authService "chopipay/internal/http/services"
	"net/http"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	"chopipay/internal/models/dto"
)

type AuthController interface {
	Login(c *gin.Context)
	RefreshToken(c *gin.Context)
}

type AuthControllerImpl struct {
	logToken    string
	authService authService.AuthService
}

func NewAuthController(authService authService.AuthService) AuthController {
	return &AuthControllerImpl{
		logToken:    "AuthController | ",
		authService: authService,
	}
}

// Login AuthRoutes godoc
// @Tags Auth
// @Summary Login
// @Description Login with email and password
// @Accept  json
// @Produce  json
// @Param credentials body dto.Login true "Login data"
// @Success 200 {object} dto.Jwt
// @Router /auth/login [post]
func (ac *AuthControllerImpl) Login(c *gin.Context) {
	var credentials dto.Login
	err := c.BindJSON(&credentials)
	if err != nil {
		errorshandler.ErrorHandler(c, err, ac.logToken+"Error binding credentials")
		return
	}

	token, err := ac.authService.Login(credentials.Username, credentials.Password)
	if err != nil {
		errorshandler.ErrorHandler(c, err, ac.logToken+"Error logging in")
		return
	}

	c.JSON(http.StatusOK, token)
}

// RefreshToken AuthRoutes godoc
// @Tags Auth
// @Summary Refresh token
// @Description Refresh token with refresh token
// @Param refreshToken body dto.RefreshToken true "Refresh token data"
// @Accept  json
// @Produce  json
// @Router /auth/refresh-token [post]
func (ac *AuthControllerImpl) RefreshToken(c *gin.Context) {
	// get the refresh token from the body request as "refresh_token"
	var refreshToken dto.RefreshToken
	err := c.BindJSON(&refreshToken)
	if err != nil {
		errorshandler.ErrorHandler(c, err, ac.logToken+"Error binding token")
		return
	}

	newToken, err := ac.authService.RefreshToken(refreshToken.Token)
	if err != nil {
		errorshandler.ErrorHandler(c, err, ac.logToken+"Error refreshing token")
		return
	}

	c.JSON(http.StatusOK, newToken)
}
