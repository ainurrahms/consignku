package indonesia_city_location

import (
	"consignku/bussiness/indonesia_city_location"
	"consignku/controller"
	"consignku/controller/indonesia_city_location/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type IndonesiaCityLocationController struct {
	IndonesiaCityLocationUsecase indonesia_city_location.Usecase
}

func NewIndonesiaCityController(uc indonesia_city_location.Usecase) *IndonesiaCityLocationController {
	return &IndonesiaCityLocationController{
		IndonesiaCityLocationUsecase: uc,
	}
}

func (ctrl *IndonesiaCityLocationController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.City{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.IndonesiaCityLocationUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
