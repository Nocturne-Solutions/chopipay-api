package business

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"chopipay/internal/http/controllers/utils"
	errorshandler "chopipay/internal/http/errors_handler"
	securityUtils "chopipay/internal/http/security/utils"
	services "chopipay/internal/http/services"
	dtos "chopipay/internal/models/dto"
)

type SalesController interface {
	Create(c *gin.Context)
}

func (sc *SalesControllerImpl) Create(c *gin.Context) {
	var newSale dtos.SaleDTO
	err := c.BindJSON(&newSale)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logTag+"Error binding sale")
		return
	}

	currentUser, err := securityUtils.GetCurrentUser(c)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logTag+"Error getting current user")
		return
	}

	isPreference, err := utils.GetBooleanFromString(c.Query("isPreference"))
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logTag+"Error converting isPreference")
		return
	}

	saleDTO, err := sc.salesService.Create(&newSale, currentUser, isPreference)
	if err != nil {
		errorshandler.ErrorHandler(c, err, sc.logTag+"Error creating sale")
		return
	}

	c.JSON(http.StatusCreated, saleDTO)
}

type SalesControllerImpl struct {
	logTag       string
	salesService services.SalesService
}

func NewSalesController(salesService services.SalesService) SalesController {
	return &SalesControllerImpl{
		logTag:       "[SalesControllerImpl] ",
		salesService: salesService,
	}
}
