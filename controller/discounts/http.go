package discounts

import (
	"consignku/bussiness/discounts"
	"consignku/controller"
	"consignku/controller/discounts/request"
	"consignku/controller/discounts/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type DiscountsController struct {
	discountsUseCase discounts.Usecase
}

func NewDiscountsController(uc discounts.Usecase) *DiscountsController {
	return &DiscountsController{
		discountsUseCase: uc,
	}
}

// func (ctrl *DiscountsController) Fetch(c echo.Context) error {
// 	ctx := c.Request().Context()
// 	page := c.QueryParam("page")
// 	perpage := c.QueryParam("perpage")

// 	convPage, _ := strconv.Atoi(page)
// 	convPerPage, _ := strconv.Atoi(perpage)

// 	res, _, err := ctrl.discountsUseCase.Fetch(ctx, convPage, convPerPage)

// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}

// 	return controller.NewSuccessResponse(c, res)
// }

func (ctrl *DiscountsController) Store(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.Discounts{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	err := ctrl.discountsUseCase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}

func (ctrl *DiscountsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.discountsUseCase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Discounts{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *DiscountsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	discounts, err := ctrl.discountsUseCase.GetByID(ctx, id)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, discounts)
}

func (ctrl *DiscountsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Discounts{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.discountsUseCase.Update(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *DiscountsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Discounts{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.Id = idInt

	_, err := ctrl.discountsUseCase.Delete(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully deleted")
}
