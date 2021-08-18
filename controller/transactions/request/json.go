package request

import (
	"consignku/bussiness/transactions"
	TransactionsItemDomain "consignku/bussiness/transactions_item"
	"consignku/controller/transactions_item/request"
)

type Transactions struct {
	Items   []request.TransactionsItem `json:"items"`
	UsersID int                        `json:"users_id"`
}

func (req *Transactions) ToDomain() *transactions.Domain {

	var items []TransactionsItemDomain.Domain

	for _, item := range req.Items {
		items = append(items, *item.ToDomain())
	}

	return &transactions.Domain{
		UsersID:          req.UsersID,
		TransactionItems: items,
	}
}
