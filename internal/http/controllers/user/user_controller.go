package user

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chopipay/internal/models/entities"
	userServices "chopipay/internal/http/services/app/user"
	errorshandler "chopipay/internal/http/errors_handler"
)

const logTag = "user_controller | "

func Create(c *gin.Context) {
	var user entities.User
	err := c.BindJSON(&user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding user")
		return
	}

	err = userServices.Create(&user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error creating user")
		return
	}

	c.JSON(http.StatusCreated, user)
}

func FindByID(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	user, err := userServices.FindByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error finding user by id " + id)
		return
	}

	c.JSON(http.StatusOK, user)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	user, err := userServices.FindByID(id_val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var userToUpdate entities.User
	err = c.BindJSON(&userToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding user to update")
		return
	}
	userToUpdate.ID = id_val

	err = userServices.Update(&userToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error updating user by id " + id + " to update")
		return
	}

	c.JSON(http.StatusOK, user)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	user, err := userServices.FindByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error finding user by id " + id + " to delete")
		return
	}

	err = userServices.Delete(user)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error deleting user by id " + id)
		return
	}

	c.JSON(http.StatusOK, user)
}