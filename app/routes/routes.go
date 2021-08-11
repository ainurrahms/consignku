package routes

import (
	"consignku/controller/users"

	"github.com/labstack/echo/v4"
)

type RouteLists struct {
	UserController users.UserController
}

func 	(r *RouteLists) RouteRegister(e *echo.Echo){
	users := e.Group("auth")
	users.POST("/register",r.UserController.Register)
}