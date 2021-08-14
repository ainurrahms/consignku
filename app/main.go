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

	_productTypesUsecase "consignku/bussiness/product_types"
	_productTypesController "consignku/controller/product_types"
	_productTypesRepo "consignku/drivers/databases/product_types"

	_productUsedTimesUsecase "consignku/bussiness/product_used_times"
	_productUsedTimesController "consignku/controller/product_used_times"
	_productUsedTimesRepo "consignku/drivers/databases/product_used_times"

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

	userRepo := _userRepo.NewMySQLUserRepository(db)
	userUsecase := _userUsecase.NewUserUseCase(userRepo, &configJWT, timeoutContext)
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

	routesInit := _routes.RouteLists{
		JWTMiddleware:              configJWT.Init(),
		UserController:             *userCtrl,
		DiscountsController:        *dicountsCtrl,
		ProductTypesController:     *productTypesCtrl,
		ProductUsedTimesController: *productUsedTimesCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
