package main

import (
	"log"
	"time"

	_dbDriverMysql "consignku/drivers/mysql"

	_userUsecase "consignku/bussiness/users"
	_userController "consignku/controller/users"
	_userRepo "consignku/drivers/databases/users"

	_discountsUsecase "consignku/bussiness/discounts"
	_discountsController "consignku/controller/discounts"
	_discountsRepo "consignku/drivers/databases/discounts"
	_indonesiaCityLocation "consignku/drivers/databases/thirdparties/indonesia_city_location"

	_productTypesUsecase "consignku/bussiness/product_types"
	_productTypesController "consignku/controller/product_types"
	_productTypesRepo "consignku/drivers/databases/product_types"

	_productUsedTimesUsecase "consignku/bussiness/product_used_times"
	_productUsedTimesController "consignku/controller/product_used_times"
	_productUsedTimesRepo "consignku/drivers/databases/product_used_times"

	_productsUsecase "consignku/bussiness/products"
	_productsController "consignku/controller/products"
	_productsRepo "consignku/drivers/databases/products"

	_transactionsUsecase "consignku/bussiness/transactions"
	_transactionsController "consignku/controller/transactions"
	_transactionsRepo "consignku/drivers/databases/transactions"

	_indonesiaCityLocationUsecase "consignku/bussiness/indonesia_city_location"
	_indonesiaCityLocationController "consignku/controller/indonesia_city_location"

	_routes "consignku/app/routes"

	_middleware "consignku/app/middleware"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init() {
	viper.SetConfigFile(`app/config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	configdb := _dbDriverMysql.ConfigDB{
		DB_USERNAME: viper.GetString(`database.username`),
		DB_PASSWORD: viper.GetString(`database.pass`),
		DB_HOST:     viper.GetString(`database.host`),
		DB_PORT:     viper.GetString(`database.port`),
		DB_DATABASE: viper.GetString(`database.db`),
	}

	db := configdb.InitialDB()

	configJWT := _middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}

	timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	indonesiaCityLocationRepo := _indonesiaCityLocation.NewIndonesiaCityLocation(db)
	indonesiaCityLocationUsecase := _indonesiaCityLocationUsecase.NewIndonesiaCityLocationusecase(indonesiaCityLocationRepo, timeoutContext)
	indonesiaCityLocationController := _indonesiaCityLocationController.NewIndonesiaCityController(indonesiaCityLocationUsecase)

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUseCase(userRepo, &configJWT, indonesiaCityLocationRepo, timeoutContext)
	userCtrl := _userController.NewUserController(userUsecase)

	discountsRepo := _discountsRepo.NewMySQLDiscountsRepository(db)
	discountsUsecase := _discountsUsecase.NewDiscountsUsecase(discountsRepo, &configJWT, timeoutContext)
	dicountsCtrl := _discountsController.NewDiscountsController(discountsUsecase)

	productTypesRepo := _productTypesRepo.NewMySQLProductTypesRepository(db)
	productTypesUsecase := _productTypesUsecase.NewProductTypesUsecase(productTypesRepo, &configJWT, timeoutContext)
	productTypesCtrl := _productTypesController.NewProductTypesController(productTypesUsecase)

	productUsedTimesRepo := _productUsedTimesRepo.NewMySQLProductUsedTimesRepository(db)
	productUsedTimesUsecase := _productUsedTimesUsecase.NewProductUsedTimesUsecase(productUsedTimesRepo, &configJWT, timeoutContext)
	productUsedTimesCtrl := _productUsedTimesController.NewProductUsedTimesController(productUsedTimesUsecase)

	productsRepo := _productsRepo.NewMysqlProductsRepository(db)
	productsUsecase := _productsUsecase.NewProductsUsecase(productsRepo, productTypesUsecase, productUsedTimesUsecase, &configJWT, timeoutContext)
	productsCtrl := _productsController.NewProductsController(productsUsecase)

	transactionsRepo := _transactionsRepo.NewMysqlProductsRepository(db)
	transcationsUsecase := _transactionsUsecase.NewTransactionsUsecase(transactionsRepo, userUsecase, discountsUsecase, productsUsecase, &configJWT, timeoutContext)
	transcationsCtrl := _transactionsController.NewTransactionsController(transcationsUsecase)

	routesInit := _routes.RouteLists{
		JWTMiddleware:                   configJWT.Init(),
		UserController:                  *userCtrl,
		DiscountsController:             *dicountsCtrl,
		ProductTypesController:          *productTypesCtrl,
		ProductUsedTimesController:      *productUsedTimesCtrl,
		ProductsController:              *productsCtrl,
		TransactionsController:          *transcationsCtrl,
		IndonesiaCityLocationController: *indonesiaCityLocationController,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
