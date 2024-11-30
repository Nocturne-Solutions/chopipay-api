package controllers

import (
	"chopipay/internal/http/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	"chopipay/internal/models/dto"
	"chopipay/internal/models/entities"
)

type UserController interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type UserControllerImpl struct {
	logToken    string
	userService services.UserService
}

func NewUserController(userService services.UserService) UserController {
	return &UserControllerImpl{
		logToken:    "UserController | ",
		userService: userService,
	}
}

func (uc *UserControllerImpl) Create(c *gin.Context) {
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error binding user")
		return
	}

	err = uc.userService.Create(&user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error creating user")
		return
	}

	c.JSON(http.StatusCreated, dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
	})
}

func (uc *UserControllerImpl) FindByID(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error converting id")
		return
	}

	user, err := uc.userService.FindByID(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error finding user by id "+id)
		return
	}

	c.JSON(http.StatusOK, dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
	})
}

func (uc *UserControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error converting id")
		return
	}

	user, err := uc.userService.FindByID(idVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userToUpdate entities.User
	err = c.BindJSON(&userToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error binding user to update")
		return
	}
	userToUpdate.ID = idVal
	userToUpdate.Username = user.Username

	err = uc.userService.Update(&userToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error updating user by id "+id+" to update")
		return
	}

	c.JSON(http.StatusOK, dto.UserDTO{
		ID:       userToUpdate.ID,
		Username: userToUpdate.Username,
	})
}

func (uc *UserControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error converting id")
		return
	}

	user, err := uc.userService.FindByID(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error finding user by id "+id+" to delete")
		return
	}

	err = uc.userService.Delete(user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, uc.logToken+"Error deleting user by id "+id)
		return
	}

	c.JSON(http.StatusOK, dto.UserDTO{
		ID:       user.ID,
		Username: user.Username,
	})
}
