package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/mercadopago/sdk-go/pkg/preference"

	errorshandler "chopipay/internal/http/errors_handler"
	securityUtils "chopipay/internal/http/security/utils"
	personalServices "chopipay/internal/http/services/app/personal"
	productServices "chopipay/internal/http/services/app/product"
	mpClientServices "chopipay/internal/http/services/mp/client"
	mpPreferenceServices "chopipay/internal/http/services/mp/preference"
	_ "chopipay/internal/models/dto"
	"chopipay/internal/models/entities"
)

const logTag = "product_controller | "

func Create(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding product")
		return
	}
	
	isPreference := false
	isPreferenceParam := c.Query("isPreference")
	if isPreferenceParam != "" {
		isPreference, err = strconv.ParseBool(isPreferenceParam)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error converting isPreference")
			return
		}
	}


	err = productServices.Create(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error creating product")
		return
	}

	if isPreference {
		preferenceClient, err := getPreferenceClient(c)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error getting preference client")
			return
		}

		productPreferenceDTO, err := mpPreferenceServices.CreatePreference(*preferenceClient, &product)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error creating MercadoPago preference")
			return
		}

		product.PreferenceID = productPreferenceDTO.PreferenceID

		err = productServices.Update(&product)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error updating product")
			return
		}
	
		c.JSON(http.StatusCreated, productPreferenceDTO)
		return
	}
	
	c.JSON(http.StatusCreated, product)
}

func FindByID(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	isPreference := false
	isPreferenceParam := c.Query("isPreference")
	if isPreferenceParam != "" {
		isPreference, err = strconv.ParseBool(isPreferenceParam)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error converting isPreference")
			return
		}
	}

	product, err := productServices.FindByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error finding product by id " + id)
		return
	}

	if isPreference && product.PreferenceID != "" {
		preferenceClient, err := getPreferenceClient(c)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error getting preference client")
			return
		}

		productPreferenceDTO, err := mpPreferenceServices.GetPreference(*preferenceClient, product.PreferenceID)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error getting MercadoPago preference")
			return
		}

		c.JSON(http.StatusOK, productPreferenceDTO)
		return
	}

	c.JSON(http.StatusOK, product)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	var product entities.Product
	err = c.BindJSON(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding product")
		return
	}

	product.ID = id_val

	err = productServices.Update(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error updating product")
		return
	}

	c.JSON(http.StatusOK, product)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error converting id")
		return
	}

	err = productServices.Delete(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error deleting product")
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Product deleted"})
}

func getPreferenceClient(c *gin.Context) (*preference.Client, error) {
	currentUsername, err := securityUtils.GetCurrentUser(c)
	if err != nil {
		return nil, err
	}

	personalCredentials, err := personalServices.GetPersonalCredentialsByUsername(currentUsername)
	if err != nil {
		return nil, err
	}

	cfg, err := mpClientServices.InitClientConfig(personalCredentials.AccessToken)
	if err != nil {
		return nil, err
	}

	return mpClientServices.GetPreferenceClient(cfg), nil
}