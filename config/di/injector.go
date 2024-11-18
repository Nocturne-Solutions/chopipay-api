// go:build wireinject
//go:build wireinject
// +build wireinject

package di

import (
	"chopipay/config/db/pg"
	"chopipay/config/server"
	"chopipay/internal/http/controllers"
	"chopipay/internal/http/controllers/business"
	"chopipay/internal/http/services"
	pgRepo "chopipay/internal/repository/pg"
	"github.com/google/wire"
)

var envVars = wire.NewSet(server.LoadEnvironment)

var pgDb = wire.NewSet(pg.InitConnection)

/*	Repositories */
var salesRepoSet = wire.NewSet(pgRepo.NewSalesRepository)
var productRepoSet = wire.NewSet(pgRepo.NewProductRepository)
var shopRepoSet = wire.NewSet(pgRepo.NewShopRepository)
var personalRepoSet = wire.NewSet(pgRepo.NewPersonalRepository)
var personalCredentialsRepoSet = wire.NewSet(pgRepo.NewPersonalCredentialsRepository)
var mpMerchantOrderPaymentsRepoSet = wire.NewSet(pgRepo.NewMpMerchantOrderPaymentsRepository)
var mpMerchantOrderRepoSet = wire.NewSet(pgRepo.NewMpMerchantOrderRepository)
var mpPayerRepoSet = wire.NewSet(pgRepo.NewMpPayerRepository)
var mpPaymentRepoSet = wire.NewSet(pgRepo.NewMpPaymentRepository)
var paymentMethodRepoSet = wire.NewSet(pgRepo.NewPaymentMethodRepository)
var saleProductsRepoSet = wire.NewSet(pgRepo.NewSaleProductsRepository)
var saleStatusRepoSet = wire.NewSet(pgRepo.NewSalesStatusRepository)
var userRepositorySet = wire.NewSet(pgRepo.NewUserRepository)

/*	Services */
var salesServiceSet = wire.NewSet(services.NewSalesService)
var productServiceSet = wire.NewSet(services.NewProductService)
var shopServiceSet = wire.NewSet(services.NewShopService)
var personalServiceSet = wire.NewSet(services.NewPersonalService)
var personalCredentialsServiceSet = wire.NewSet(services.NewPersonalCredentialsService)

/*	Controllers */
var salesCtrlSet = wire.NewSet(business.NewSalesController)
var productCtrlSet = wire.NewSet(controllers.NewProductController)
var shopCtrlSet = wire.NewSet(controllers.NewShopController)
var personalCtrlSet = wire.NewSet(controllers.NewPersonalController)

func Init() *Initialization {
	wire.Build(NewDiInit,
		envVars, pgDb,
		/*	Controllers */
		salesCtrlSet, productCtrlSet, shopCtrlSet, personalCtrlSet,
		/*	Services */
		salesServiceSet, productServiceSet, shopServiceSet, personalServiceSet, personalCredentialsServiceSet,
		/*	Repositories */
		salesRepoSet, productRepoSet, shopRepoSet, personalRepoSet, personalCredentialsRepoSet, mpMerchantOrderPaymentsRepoSet,
		mpMerchantOrderRepoSet, mpPayerRepoSet, mpPaymentRepoSet, paymentMethodRepoSet, saleProductsRepoSet, saleStatusRepoSet,
		userRepositorySet,
	)

	return nil
}
