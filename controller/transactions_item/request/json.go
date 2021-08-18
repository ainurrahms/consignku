package request

import transctions_item "consignku/bussiness/transactions_item"

type TransactionsItem struct {
	ProductsID    int    `json:"products_id"`
	ProductsPrice int    `json:"products_price"`
	ProductName   string `json:"products_name"`
	DiscountsID   int    `json:"discounts_id"`
	DiscountsVal  int    `json:"discounts_val"`
}

func (req *TransactionsItem) ToDomain() *transctions_item.Domain {
	return &transctions_item.Domain{
		ProductsID:    req.ProductsID,
		DiscountsID:   req.DiscountsID,
		DiscountsVal:  req.DiscountsVal,
		ProductsPrice: req.ProductsPrice,
		ProductName:   req.ProductName,
	}
}
