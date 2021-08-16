package response

import (
	"consignku/bussiness/transactions"
	"consignku/drivers/databases/discounts"
	"consignku/drivers/databases/users"
)

type Transactions struct {
	Price          int                 `json:"price"`
	UsersID        int                 `json:"users_id"`
	Username       users.Users         `json:"username"`
	Usernames      string              `json:"buyer_name"`
	DiscountsID    int                 `json:"discounts_id"`
	DiscountsVal   int                 `json:"discounts_val"`
	DiscountsValue discounts.Discounts `json:"discounts_value"`
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
