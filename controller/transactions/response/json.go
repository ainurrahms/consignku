package response

import (
	"consignku/bussiness/transactions"
	"consignku/controller/transactions_item/response"
)

type Transactions struct {
	ID               int                         `json:"id"`
	UsersID          int                         `json:"users_id"`
	Usernames        string                      `json:"buyer_name"`
	TotalPrice       int                         `json:"total_price"`
	TransactionItems []response.TransactionsItem `json:"transaction_items"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func FromDomain(domain transactions.Domain) Transactions {
	var items []response.TransactionsItem
	var totalPrice int
	for _, item := range domain.TransactionItems {
		items = append(items, response.FromDomain(item))
		totalPrice = item.ProductsPrice - item.DiscountsVal
	}
	return Transactions{
		ID:               domain.ID,
		UsersID:          domain.UsersID,
		Usernames:        domain.Usernames,
		TotalPrice:       totalPrice,
		TransactionItems: items,
	}
}
