package main

import (
	"log"

	"github.com/labstack/echo/v4"
	"github.com/spf13/viper"
)

func init(){
	viper.SetConfigFile(`config.json`)
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	if viper.GetBool(`debug`) {
		log.Println("Service RUN on DEBUG mode")
	}
}

func main() {
	// configdb := _dbHelper.ConfigDB{
	// 	DB_USERNAME: viper.GetString(`database.username`),
	// 	DB_PASSWORD: viper.GetString(`database.pass`),
	// 	DB_HOST: viper.GetString(`database.host`),
	// 	DB_PORT: viper.GetString(`database.port`),
	// 	DB_DATABASE: viper.GetString(`database.db`),
	// }

	// db := configdb.InitialDB()
	// timeoutContext := time.Duration(viper.GetInt("context.timeout")) * time.Second
	e := echo.New()

	// fmt.Println(e)

	log.Fatal(e.Start(viper.GetString("server.address")))
}
