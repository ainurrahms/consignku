package transactions

import (
	"consignku/bussiness/transactions"
	"consignku/controller"
	"consignku/controller/transactions/request"
	"net/http"

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

	_, err := ctrl.transactionsUsecase.Store(ctx, req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	return controller.NewSuccessResponse(c, "Successfully inserted")
}
