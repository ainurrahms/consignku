package routes

import (
	"consignku/controller/users"

	"github.com/labstack/echo/v4"
)

type RouteLists struct {
	UserController users.UserController
}

func (r *RouteLists) RouteRegister(e *echo.Echo) {
	users := e.Group("v1/api/auth")
	users.POST("/register", r.UserController.Register)
	users.POST("/login", r.UserController.Login)
}
