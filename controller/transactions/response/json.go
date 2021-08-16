package response

import (
	"consignku/bussiness/transactions"
)

type Transactions struct {
	Price        int    `json:"price"`
	UsersID      int    `json:"users_id"`
	Usernames    string `json:"buyer_name"`
	DiscountsID  int    `json:"discounts_id"`
	DiscountsVal int    `json:"discounts_val"`
}

type Pagination struct {
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
}

func FromDomain(domain transactions.Domain) Transactions {
	return Transactions{
		Price:        domain.Price,
		UsersID:      domain.UsersID,
		Usernames:    domain.Usernames,
		DiscountsID:  domain.DiscountsID,
		DiscountsVal: domain.DiscountsVal,
	}
}
