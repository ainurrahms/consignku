package request

import "consignku/bussiness/transactions"

type Transactions struct {
	Price       int `json:"price"`
	UsersID     int `json:"users_id"`
	DiscountsID int `json:"discounts_id"`
}

func (req *Transactions) ToDomain() *transactions.Domain {
	return &transactions.Domain{
		Price:       req.Price,
		UsersID:     req.UsersID,
		DiscountsID: req.DiscountsID,
	}
}
