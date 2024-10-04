package personal

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chopipay/internal/models/entities"
	personalServices "chopipay/internal/http/services/user/app/personal"
	errorshandler "chopipay/internal/http/errors_handler"
)

const logTag = "personal_controller | "

func Create(c *gin.Context) {
	var personal entities.Personal
	err := c.BindJSON(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding personal")
		return
	}

	err = personalServices.Create(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error creating personal")
		return
	}

	c.JSON(http.StatusCreated, personal)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	personal, err := personalServices.GetByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error finding personal by id " + id)
		return
	}

	c.JSON(http.StatusOK, personal)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	personal, err := personalServices.GetByID(id_val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.BindJSON(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding personal")
		return
	}

	err = personalServices.Update(personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error updating personal")
		return
	}

	c.JSON(http.StatusOK, personal)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	personal, err := personalServices.GetByID(id_val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = personalServices.Delete(personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error deleting personal")
		return
	}

	c.JSON(http.StatusOK, personal)
}