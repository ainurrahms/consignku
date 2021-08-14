package product_used_times

import (
	"consignku/bussiness/product_used_times"
	"consignku/controller"
	"consignku/controller/product_used_times/request"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductUsedTimesController struct {
	ProductUsedTimesUseCase product_used_times.Usecase
}

func NewProductUsedTimesController(uc product_used_times.Usecase) *ProductUsedTimesController {
	return &ProductUsedTimesController{
		ProductUsedTimesUseCase: uc,
	}
}

// func (ctrl *ProductUsedTimesController) Fetch(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	page := c.QueryParam("page")
// 	perpage := c.QueryParam("perpage")

// 	convPage, _ := strconv.Atoi(page)
// 	convPerPage, _ := strconv.Atoi(perpage)

// 	res, _, err := ctrl.ProductUsedTimesUseCase.Fetch(ctx, convPage, convPerPage)

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controller.NewSuccessResponse(c, res)
// }

func (ctrl *ProductUsedTimesController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.ProductUsedTimes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.ProductUsedTimesUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *ProductUsedTimesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	discounts, err := ctrl.ProductUsedTimesUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, discounts)
}

func (ctrl *ProductUsedTimesController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.ProductUsedTimes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.ProductUsedTimesUseCase.Update(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *ProductUsedTimesController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.ProductUsedTimes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.ProductUsedTimesUseCase.Delete(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully deleted")
}
