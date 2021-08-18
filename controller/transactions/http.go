package transactions

import (
	"consignku/app/middleware"
	"consignku/bussiness/transactions"
	"consignku/controller"
	"consignku/controller/transactions/request"
	"consignku/controller/transactions/response"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type TransactionsController struct {
	transactionsUsecase transactions.Usecase
}

func NewTransactionsController(transactionsUC transactions.Usecase) *TransactionsController {
	return &TransactionsController{
		transactionsUsecase: transactionsUC,
	}
}

func (ctrl *TransactionsController) Store(c echo.Context) error {
	ctx := c.Request().Context()
	req := request.Transactions{}

	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	res, err := ctrl.transactionsUsecase.Store(ctx, req.ToDomain())

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, res)
}

func (ctrl *TransactionsController) Find(c echo.Context) error {
	ctx := c.Request().Context()
	page := c.QueryParam("page")
	offset := c.QueryParam("limit")

	var varPage response.Pagination

	p, err := strconv.Atoi(page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	o, err := strconv.Atoi(offset)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	resp, count, lastPage, err := ctrl.transactionsUsecase.Find(ctx, p, o)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	varPage.CurrentPage = p
	varPage.LastPage = lastPage
	varPage.PerPage = o
	varPage.Total = count

	responseController := []response.Transactions{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}
	return controller.NewSuccessPaginationResponse(c, responseController, varPage)
}

func (ctrl *TransactionsController) GetAll(c echo.Context) error {
	ctx := c.Request().Context()

	resp, err := ctrl.transactionsUsecase.GetAll(ctx)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	responseController := []response.Transactions{}
	for _, value := range resp {
		responseController = append(responseController, response.FromDomain(value))
	}

	return controller.NewSuccessResponse(c, responseController)
}

func (ctrl *TransactionsController) GetByID(c echo.Context) error {
	ctx := c.Request().Context()

	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	userID := middleware.GetUser(c).ID

	transactions, err := ctrl.transactionsUsecase.GetByID(ctx, id, userID)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	return controller.NewSuccessResponse(c, response.FromDomain(transactions))
}

func (ctrl *TransactionsController) Update(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt

	_, err := ctrl.transactionsUsecase.Update(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully updated")
}

func (ctrl *TransactionsController) Delete(c echo.Context) error {
	ctx := c.Request().Context()

	id := c.Param("id")

	if strings.TrimSpace(id) == "" {
		return controller.NewErrorResponse(c, http.StatusBadRequest, errors.New("missing required id"))
	}

	req := request.Transactions{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domainReq := req.ToDomain()
	idInt, _ := strconv.Atoi(id)
	domainReq.ID = idInt

	_, err := ctrl.transactionsUsecase.Delete(ctx, domainReq)

	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully deleted")
}
