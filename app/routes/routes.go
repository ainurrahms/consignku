package routes

import (
	"consignku/controller/discounts"
	"consignku/controller/indonesia_city_location"
	"consignku/controller/product_types"
	"consignku/controller/product_used_times"
	"consignku/controller/products"
	"consignku/controller/transactions"
	"consignku/controller/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteLists struct {
	JWTMiddleware                   middleware.JWTConfig
	UserController                  users.UserController
	DiscountsController             discounts.DiscountsController
	ProductsController              products.ProductsController
	ProductTypesController          product_types.ProductTypesController
	ProductUsedTimesController      product_used_times.ProductUsedTimesController
	TransactionsController          transactions.TransactionsController
	IndonesiaCityLocationController indonesia_city_location.IndonesiaCityLocationController
}

func (r *RouteLists) RouteRegister(e *echo.Echo) {
	users := e.Group("v1/api/auth")
	users.POST("/register", r.UserController.Register)
	users.POST("/login", r.UserController.Login)

	discounts := e.Group("v1/api/discounts")
	//get all belum :(
	// discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.POST("/create", r.DiscountsController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.GET("", r.DiscountsController.GetAll, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.GET("/id/:id", r.DiscountsController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.PUT("/update/:id", r.DiscountsController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	discounts.DELETE("/delete/:id", r.DiscountsController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	productType := e.Group("v1/api/product/type")
	//get all belum :(
	// discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))

	productType.POST("/create", r.ProductTypesController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.GET("", r.ProductTypesController.GetAll, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.GET("/id/:id", r.ProductTypesController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.PUT("/update/:id", r.ProductTypesController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	productType.DELETE("/delete/:id", r.ProductTypesController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	productUsedTimes := e.Group("v1/api/product/used")
	//get all belum :(
	// discounts.GET("/all", r.DiscountsController.Fetch, middleware.JWTWithConfig(r.JWTMiddleware))

	productUsedTimes.POST("/create", r.ProductUsedTimesController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.GET("", r.ProductUsedTimesController.GetAll, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.GET("/id/:id", r.ProductUsedTimesController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.PUT("/update/:id", r.ProductUsedTimesController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	productUsedTimes.DELETE("/delete/:id", r.ProductUsedTimesController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	productsTimes := e.Group("v1/api/product")
	productsTimes.POST("/create", r.ProductsController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	productsTimes.GET("", r.ProductsController.GetAll, middleware.JWTWithConfig(r.JWTMiddleware))
	productsTimes.GET("/id/:id", r.ProductsController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	productsTimes.PUT("/update/:id", r.ProductsController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	productsTimes.DELETE("/delete/:id", r.ProductsController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	transactionsTimes := e.Group("v1/api/transactions")
	transactionsTimes.POST("/create", r.TransactionsController.Store, middleware.JWTWithConfig(r.JWTMiddleware))
	transactionsTimes.GET("", r.TransactionsController.Find, middleware.JWTWithConfig(r.JWTMiddleware))
	transactionsTimes.GET("/all", r.TransactionsController.GetAll, middleware.JWTWithConfig(r.JWTMiddleware))
	transactionsTimes.GET("/id/:id", r.TransactionsController.GetByID, middleware.JWTWithConfig(r.JWTMiddleware))
	transactionsTimes.PUT("/update/:id", r.TransactionsController.Update, middleware.JWTWithConfig(r.JWTMiddleware))
	transactionsTimes.DELETE("/delete/:id", r.TransactionsController.Delete, middleware.JWTWithConfig(r.JWTMiddleware))

	cityTimes := e.Group("v1/api/city")
	cityTimes.POST("/create", r.IndonesiaCityLocationController.Store)

}
