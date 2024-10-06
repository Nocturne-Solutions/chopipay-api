package product
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chopipay/internal/models/entities"
	productServices "chopipay/internal/http/services/app/product"
	personalServices "chopipay/internal/http/services/app/personal"
	mpClientServices "chopipay/internal/http/services/mp/client" 
	mpPreferenceServices "chopipay/internal/http/services/mp/preference"
	errorshandler "chopipay/internal/http/errors_handler"
	securityUtils "chopipay/internal/http/security/utils"
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
		currentUsername, err := securityUtils.GetCurrentUser(c)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error getting current user")
			return
		}

		personalCredentials, err := personalServices.GetPersonalCredentialsByUsername(currentUsername)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error getting personal credentials by username")
			return
		}

		cfg, err := mpClientServices.InitClientConfig(personalCredentials.AccessToken)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error initializing MercadoPago")
			return
		}

		client := mpClientServices.GetPreferenceClient(cfg)

		err = mpPreferenceServices.CreatePreference(*client, &product)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error creating MercadoPago preference")
			return
		}

		err = productServices.Update(&product, true)
		if err != nil {
			errorshandler.ErrorHandler(c, err, logTag + "Error updating product")
			return
		}
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

	product, err := productServices.FindByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error finding product by id " + id)
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

	err = productServices.Update(&product, false)
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