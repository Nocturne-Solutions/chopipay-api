package controllers

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	errorshandler "chopipay/internal/http/errors_handler"
	"chopipay/internal/http/services"
	"chopipay/internal/models/entities"
)

type ShopController interface {
	Create(c *gin.Context)
	GetByID(c *gin.Context)
	GetAllByPersonalId(c *gin.Context)
	Update(c *gin.Context)
	Delete(c *gin.Context)
	GetShopProducts(c *gin.Context)
}

type ShopControllerImpl struct {
	logToken       string
	shopService    services.ShopService
	productService services.ProductService
}

func NewShopController(shopService services.ShopService,
	productService services.ProductService) ShopController {
	return &ShopControllerImpl{
		logToken:       "ShopController | ",
		shopService:    shopService,
		productService: productService,
	}
}

func (sc *ShopControllerImpl) Create(c *gin.Context) {
	var shop entities.Shop
	err := c.BindJSON(&shop)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error binding shop")
		return
	}

	err = sc.shopService.Create(&shop)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error creating shop")
		return
	}

	c.JSON(http.StatusCreated, shop)
}

func (sc *ShopControllerImpl) GetByID(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error converting id")
		return
	}

	shop, err := sc.shopService.GetByID(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error finding shop by id "+id)
		return
	}

	c.JSON(http.StatusOK, shop)
}

func (sc *ShopControllerImpl) GetAllByPersonalId(c *gin.Context) {
	personalId := c.Param("personal_id")
	personalIdVal, err := strconv.Atoi(personalId)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error converting personal_id")
		return
	}

	shops, err := sc.shopService.GetAllByPersonalId(personalIdVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error finding shops by personal_id "+personalId)
		return
	}

	c.JSON(http.StatusOK, shops)
}

func (sc *ShopControllerImpl) Update(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error converting id")
		return
	}

	shop, err := sc.shopService.GetByID(idVal)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var shopToUpdate entities.Shop
	err = c.BindJSON(&shopToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error binding shop to update")
		return
	}

	shopToUpdate.ID = shop.ID
	err = sc.shopService.Update(&shopToUpdate)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error updating shop")
		return
	}

	c.JSON(http.StatusOK, shopToUpdate)
}

func (sc *ShopControllerImpl) Delete(c *gin.Context) {
	id := c.Param("id")
	idVal, err := strconv.Atoi(id)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error converting id")
		return
	}

	err = sc.shopService.Delete(idVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error deleting shop by id "+id)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Shop deleted"})
}

func (sc *ShopControllerImpl) GetShopProducts(c *gin.Context) {
	shopId := c.Param("id")
	shopIdVal, err := strconv.Atoi(shopId)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error converting shop id")
		return
	}

	products, err := sc.productService.FindAllByShopID(shopIdVal)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logToken+"Error finding products by shop id "+shopId)
		return
	}

	c.JSON(http.StatusOK, products)
}
