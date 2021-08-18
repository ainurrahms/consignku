package transactions_item

import "time"

type Domain struct {
	ID             int
	TransactionsID int
	ProductsID     int
	ProductName    string
	DiscountsID    int
	DiscountsVal   int
	ProductsPrice  int
	TotalPrice     int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}
