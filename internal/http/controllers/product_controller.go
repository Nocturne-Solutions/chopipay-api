package controllers

import (
	"chopipay/internal/http/services"
	mpClientServices "chopipay/internal/integrations/mercadopago/client"
	mpPreferenceServices "chopipay/internal/integrations/mercadopago/preference"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/preference"

	errorshandler "chopipay/internal/http/errors_handler"
	securityUtils "chopipay/internal/http/security/utils"
	_ "chopipay/internal/models/dto"
	"chopipay/internal/models/entities"
)

type ProductController interface {
	Create(c *gin.Context)
	FindByID(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
}

type ProductControllerImpl struct {
	logToken        string
	productServices services.ProductService
	personalService services.PersonalService
}

func NewProductController(productServices services.ProductService,
	personalService services.PersonalService) ProductController {
	return &ProductControllerImpl{
		logToken:        "ProductControllerImpl | ",
		productServices: productServices,
		personalService: personalService,
	}
}

func (pc *ProductControllerImpl) Create(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error binding product")
		return
	}

	err = pc.productServices.Create(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error creating product")
		return
	}

	c.JSON(http.StatusCreated, product)
}

func (pc *ProductControllerImpl) FindByID(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	isPreference := false
	isPreferenceParam := c.Query("isPreference")
	if isPreferenceParam != "" {
		isPreference, err = strconv.ParseBool(isPreferenceParam)
		if err != nil {
			errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting isPreference")
			return
		}
	}

	product, err := pc.productServices.FindByID(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error finding product by id "+id)
		return
	}

	if isPreference && product.PreferenceID != "" {
		preferenceClient, err := pc.getPreferenceClient(c)
		if err != nil {
			errorshandler.ErrorHandler(c, err, pc.logToken+"Error getting preference client")
			return
		}

		productPreferenceDTO, err := mpPreferenceServices.GetPreference(*preferenceClient, product.PreferenceID)
		if err != nil {
			errorshandler.ErrorHandler(c, err, pc.logToken+"Error getting MercadoPago preference")
			return
		}

		c.JSON(http.StatusOK, productPreferenceDTO)
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc *ProductControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	var product entities.Product
	err = c.BindJSON(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error binding product")
		return
	}

	product.ID = idVal

	err = pc.productServices.Update(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error updating product")
		return
	}

	c.JSON(http.StatusOK, product)
}

func (pc *ProductControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error converting id")
		return
	}

	err = pc.productServices.Delete(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, pc.logToken+"Error deleting product")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func (pc *ProductControllerImpl) getPreferenceClient(c *gin.Context) (*preference.Client, error) {
	currentUsername, err := securityUtils.GetCurrentUser(c)
	if err != nil {
		return nil, err
	}

	personalCredentials, err := pc.personalService.GetPersonalCredentialsByUsername(currentUsername)
	if err != nil {
		return nil, err
	}

	cfg, err := mpClientServices.InitClientConfig(personalCredentials.AccessToken)
	if err != nil {
		return nil, err
	}

	return mpClientServices.GetPreferenceClient(cfg), nil
}
