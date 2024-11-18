package di

import (
	"chopipay/internal/http/controllers"
	"chopipay/internal/http/controllers/business"
	"chopipay/internal/http/services"
	"chopipay/internal/repository/pg"
)

type Initialization struct {
	/*	Repositories */
	salesRepo                   pg.SalesRepository
	productRepo                 pg.ProductRepository
	shopRepo                    pg.ShopRepository
	personalRepo                pg.PersonalRepository
	personalCredentialsRepo     pg.PersonalCredentialsRepository
	mpMerchantOrderPaymentsRepo pg.MpMerchantOrderPaymentsRepository
	mpMerchantOrderRepo         pg.MpMerchantOrderRepository
	mpPayerRepo                 pg.MpPayerRepository
	mpPaymentRepo               pg.MpPaymentRepository
	paymentMethodRepo           pg.PaymentMethodRepository
	saleProductsRepo            pg.SaleProductsRepository
	saleStatusRepo              pg.SalesStatusRepository
	userRepo                    pg.UserRepository
	/*	Services */
	salesServices              services.SalesService
	ProductServices            services.ProductService
	shopService                services.ShopService
	PersonalService            services.PersonalService
	personalCredentialsService services.PersonalCredentialsService
	/*	Controllers */
	SalesCtrl    business.SalesController
	ProductCtrl  controllers.ProductController
	ShopCtrl     controllers.ShopController
	PersonalCtrl controllers.PersonalController
}

func NewDiInit(salesRepo pg.SalesRepository,
	productRepo pg.ProductRepository,
	shopRepo pg.ShopRepository,
	personalRepo pg.PersonalRepository,
	personalCredentialsRepo pg.PersonalCredentialsRepository,
	mpMerchantOrderPaymentsRepo pg.MpMerchantOrderPaymentsRepository,
	mpMerchantOrderRepo pg.MpMerchantOrderRepository,
	mpPayerRepo pg.MpPayerRepository,
	mpPaymentRepo pg.MpPaymentRepository,
	paymentMethodRepo pg.PaymentMethodRepository,
	saleProductsRepo pg.SaleProductsRepository,
	saleStatusRepo pg.SalesStatusRepository,
	userRepo pg.UserRepository,
	/*	Services */
	salesServices services.SalesService,
	productServices services.ProductService,
	shopService services.ShopService,
	personalService services.PersonalService,
	personalCredentialsService services.PersonalCredentialsService,
	/*	Controllers */
	salesCtrl business.SalesController,
	productCtrl controllers.ProductController,
	shopCtrl controllers.ShopController,
	personalCtrl controllers.PersonalController) *Initialization {
	return &Initialization{
		/*	Repositories */
		salesRepo:                   salesRepo,
		productRepo:                 productRepo,
		shopRepo:                    shopRepo,
		personalRepo:                personalRepo,
		personalCredentialsRepo:     personalCredentialsRepo,
		mpMerchantOrderPaymentsRepo: mpMerchantOrderPaymentsRepo,
		mpMerchantOrderRepo:         mpMerchantOrderRepo,
		mpPayerRepo:                 mpPayerRepo,
		mpPaymentRepo:               mpPaymentRepo,
		paymentMethodRepo:           paymentMethodRepo,
		saleProductsRepo:            saleProductsRepo,
		saleStatusRepo:              saleStatusRepo,
		userRepo:                    userRepo,
		/*	Services */
		salesServices:              salesServices,
		ProductServices:            productServices,
		shopService:                shopService,
		PersonalService:            personalService,
		personalCredentialsService: personalCredentialsService,
		/*	Controllers */
		SalesCtrl:    salesCtrl,
		ProductCtrl:  productCtrl,
		ShopCtrl:     shopCtrl,
		PersonalCtrl: personalCtrl,
	}
}
