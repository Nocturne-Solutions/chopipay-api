package business

import (
	"chopipay/internal/http/controllers/utils"
	errorshandler "chopipay/internal/http/errors_handler"
	securityUtils "chopipay/internal/http/security/utils"
	dtos "chopipay/internal/models/dto"
	"github.com/gin-gonic/gin"
	"net/http"

	services "chopipay/internal/http/services"
)

const logTag = "SalesControllerImpl | "

type SalesController interface {
	Create(c *gin.Context)
}

func (sc *SalesControllerImpl) Create(c *gin.Context) {
	var newSale dtos.SaleDTO
	err := c.BindJSON(&newSale)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error binding sale")
		return
	}

	currentUser, err := securityUtils.GetCurrentUser(c)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error getting current user")
		return
	}

	isPreference, err := utils.GetBooleanFromString(c.Query("isPreference"))
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error converting isPreference")
		return
	}

	saleDTO, err := sc.salesService.Create(&newSale, currentUser, isPreference)
	if err != nil {
		errorshandler.ErrorHandler(c, err, logTag+"Error creating sale")
		return
	}

	c.JSON(http.StatusCreated, saleDTO)
}

type SalesControllerImpl struct {
	salesService services.SalesService
}

func NewSalesController(salesService services.SalesService) SalesController {
	return &SalesControllerImpl{
		salesService: salesService,
	}
}
