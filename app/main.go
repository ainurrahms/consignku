package main

import (
	"log"
	"time"

	_dbDriverMysql "consignku/drivers/mysql"

	_userUsecase "consignku/bussiness/users"
	_userController "consignku/controller/users"
	_userRepo "consignku/drivers/databases/users"

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
	useCtrl := _userController.NewUserController(userUsecase)

	routesInit := _routes.RouteLists{
		UserController: *useCtrl,
	}

	routesInit.RouteRegister(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
