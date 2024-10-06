package product
import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chopipay/internal/models/entities"
	productServices "chopipay/internal/http/services/app/product"
	errorshandler "chopipay/internal/http/errors_handler"
)

const logTag = "product_controller | "

func Create(c *gin.Context) {
	var product entities.Product
	err := c.BindJSON(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error binding product")
		return
	}

	err = productServices.Create(&product)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag + "Error creating product")
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