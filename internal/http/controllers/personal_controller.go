package controllers

import (
	"chopipay/internal/http/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	"chopipay/internal/models/entities"
)

type PersonalController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	AddPersonalCredential(c *gin.Context)
	GetShopsByPersonalID(c *gin.Context)
}

type PersonalControllerImpl struct {
	logToken            string
	personalService     services.PersonalService
	shopService         services.ShopService
	credentialsServices services.PersonalCredentialsService
}

func NewPersonalController(personalService services.PersonalService,
	shopService services.ShopService,
	credentialsServices services.PersonalCredentialsService) PersonalController {
	return &PersonalControllerImpl{
		logToken:            "PersonalController | ",
		personalService:     personalService,
		shopService:         shopService,
		credentialsServices: credentialsServices,
	}
}

func (pc *PersonalControllerImpl) Create(c *gin.Context) {
	var personal entities.Personal
	err := c.BindJSON(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error binding personal")
		return
	}

	err = pc.personalService.Create(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error creating personal")
		return
	}

	c.JSON(http.StatusCreated, personal)
}

func (pc *PersonalControllerImpl) GetByID(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	personal, err := pc.personalService.GetByID(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error finding personal by id "+id)
		return
	}

	c.JSON(http.StatusOK, personal)
}

func (pc *PersonalControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	personal, err := pc.personalService.GetByID(idVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = c.BindJSON(&personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error binding personal")
		return
	}

	err = pc.personalService.Update(personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error updating personal")
		return
	}

	c.JSON(http.StatusOK, personal)
}

func (pc *PersonalControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	personal, err := pc.personalService.GetByID(idVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	err = pc.personalService.Delete(personal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error deleting personal")
		return
	}

	c.JSON(http.StatusOK, personal)
}

func (pc *PersonalControllerImpl) AddPersonalCredential(c *gin.Context) {
	var personalCredential entities.PersonalCredentials
	err := c.BindJSON(&personalCredential)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error binding personal credential")
		return
	}

	err = pc.credentialsServices.AddPersonalCredential(&personalCredential)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error adding personal credential")
		return
	}

	c.JSON(http.StatusCreated, personalCredential)
}

func (pc *PersonalControllerImpl) GetShopsByPersonalID(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	shops, err := pc.shopService.GetAllByPersonalId(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error getting shops by personal id "+id)
		return
	}

	c.JSON(http.StatusOK, shops)
}
