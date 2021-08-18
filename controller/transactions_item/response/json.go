package response

import (
	"consignku/bussiness/transactions_item"
	"time"
)

type TransactionsItem struct {
	ID             int       `json:"id"`
	TransactionsID int       `json:"transaction_id"`
	ProductsID     int       `json:"products_id"`
	ProductName    string    `json:"product_name"`
	ProductsPrice  int       `json:"product_price"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

func FromDomain(domain transactions_item.Domain) TransactionsItem {
	return TransactionsItem{
		ID:             domain.ID,
		TransactionsID: domain.TransactionsID,
		ProductsID:     domain.ProductsID,
		ProductName:    domain.ProductName,
		ProductsPrice:  domain.ProductsPrice,
		CreatedAt:      domain.CreatedAt,
		UpdatedAt:      domain.UpdatedAt,
	}
}
