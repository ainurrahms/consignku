package product_types

import (
	"consignku/bussiness/product_types"
	"consignku/controller"
	"consignku/controller/product_types/request"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type ProductTypesController struct {
	ProductTypesUseCase product_types.Usecase
}

func NewProductTypesController(uc product_types.Usecase) *ProductTypesController {
	return &ProductTypesController{
		ProductTypesUseCase: uc,
	}
}

// func (ctrl *ProductTypesController) Fetch(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	page := c.QueryParam("page")
// 	perpage := c.QueryParam("perpage")

// 	convPage, _ := strconv.Atoi(page)
// 	convPerPage, _ := strconv.Atoi(perpage)

// 	res, _, err := ctrl.ProductTypesUseCase.Fetch(ctx, convPage, convPerPage)

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controller.NewSuccessResponse(c, res)
// }

func (ctrl *ProductTypesController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.ProductTypes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.ProductTypesUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *ProductTypesController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	discounts, err := ctrl.ProductTypesUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, discounts)
}

func (ctrl *ProductTypesController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.ProductTypes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.ProductTypesUseCase.Update(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *ProductTypesController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.ProductTypes{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.ProductTypesUseCase.Delete(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully deleted")
}
