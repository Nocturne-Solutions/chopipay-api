package shop

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"chopipay/internal/models/entities"
	ShopServices "chopipay/internal/http/services/app/shop"
	productServices "chopipay/internal/http/services/app/product"
	errorshandler "chopipay/internal/http/errors_handler"
)

const logTag = "shop_controller | "

func Create(c *gin.Context) {
	var shop entities.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error binding shop")
		return
	}

	err = ShopServices.Create(&shop)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error creating shop")
		return
	}

	c.JSON(http.StatusCreated, shop)
}

func GetByID(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting id")
		return
	}

	shop, err := ShopServices.GetByID(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error finding shop by id "+id)
		return
	}

	c.JSON(http.StatusOK, shop)
}

func GetAllByPersonalId(c *gin.Context) {
	personalId := c.Param("personal_id")
	personalId_val, err := strconv.Atoi(personalId)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting personal_id")
		return
	}

	shops, err := ShopServices.GetAllByPersonalId(personalId_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error finding shops by personal_id "+personalId)
		return
	}

	c.JSON(http.StatusOK, shops)
}

func Update(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting id")
		return
	}

	shop, err := ShopServices.GetByID(id_val)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var shopToUpdate entities.Shop
	err = c.BindJSON(&shopToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error binding shop to update")
		return
	}

	shopToUpdate.ID = shop.ID
	err = ShopServices.Update(&shopToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error updating shop")
		return
	}

	c.JSON(http.StatusOK, shopToUpdate)
}

func Delete(c *gin.Context) {
	id := c.Param("id")
	id_val, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting id")
		return
	}

	err = ShopServices.Delete(id_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error deleting shop by id "+id)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shop deleted"})
}

func GetShopProducts(c *gin.Context) {
	shopId := c.Param("id")
	shopId_val, err := strconv.Atoi(shopId)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting shop id")
		return
	}

	products, err := productServices.FindAllByShopID(shopId_val)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error finding products by shop id "+ shopId)
		return
	}

	c.JSON(http.StatusOK, products)
}