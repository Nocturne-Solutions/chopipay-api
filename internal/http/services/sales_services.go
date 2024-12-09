package services

import (
	mpClientServices "chopipay/internal/integrations/mercadopago/client"
	mpPreferenceServices "chopipay/internal/integrations/mercadopago/preference"
	"chopipay/internal/repository/pg"
	"github.com/mercadopago/sdk-go/pkg/preference"
	"log"
	"time"

	dtos "chopipay/internal/models/dto"
	"chopipay/internal/models/entities"
	"chopipay/internal/utils/constants"
)

type SalesService interface {
	Create(sale *dtos.SaleDTO, currentUser string, isPreference bool) (*dtos.SaleDTO, error)
	Update(sale *entities.Sales) (*entities.Sales, error)
	FindByID(id int) (*entities.Sales, error)
}

type SalesServiceImpl struct {
	logToken           string
	salesRepository    pg.SalesRepository
	personalRepository pg.PersonalRepository
}

func (s SalesServiceImpl) Create(sale *dtos.SaleDTO, currentUser string, isPreference bool) (*dtos.SaleDTO, error) {
	newSale := &entities.Sales{
		CurrencyID:        sale.CurrencyID,
		PersonalID:        sale.PersonalID,
		DateCreated:       time.DateTime,
		LastModified:      time.DateTime,
		StatusID:          constants.SaleStatusCreatedVal,
		TotalPaidAmount:   sale.TotalPaidAmount,
		TransactionAmount: s.summarizeProducts(sale.ProductsDTO),
	}

	newSaleCreated, err := s.salesRepository.Create(newSale)
	if err != nil {
		return nil, err
	}

	sale.SaleID = newSaleCreated.ID
	sale.StatusID = newSaleCreated.StatusID

	go func() {
		if isPreference {
			personalCredentials, err := s.personalRepository.GetPersonalCredentialsByUsername(currentUser)
			if err != nil {
				log.Printf("%s Error getting personal credentials: %v", s.logToken, err)
				return
			}

			preferenceClient := mpClientServices.GetClient(personalCredentials.AccessToken, "preference").(preference.Client)

			productPreferenceDTO, err := mpPreferenceServices.CreatePreference(preferenceClient, sale.ProductsDTO, sale.PersonalID)
			if err != nil {
				log.Printf("%s Error creating preference: %v", s.logToken, err)
			}
			sale.PaymentPoint = productPreferenceDTO.PaymentPoint
			sale.SandboxPaymentPoint = productPreferenceDTO.SandboxPaymentPoint
		}
	}()

	return sale, nil
}

func (s SalesServiceImpl) Update(sale *entities.Sales) (*entities.Sales, error) {
	return s.salesRepository.Update(sale)
}

func (s SalesServiceImpl) FindByID(id int) (*entities.Sales, error) {
	return s.salesRepository.FindById(id)
}

func (s SalesServiceImpl) summarizeProducts(products []dtos.ProductDTO) float64 {
	var total float64
	for _, product := range products {
		total += product.Price
	}
	return total
}

func NewSalesService(salesRepository pg.SalesRepository, personalRepository pg.PersonalRepository) SalesService {
	return &SalesServiceImpl{
		logToken:           "[SalesService]",
		salesRepository:    salesRepository,
		personalRepository: personalRepository,
	}
}
