package transactions_item

import (
	"consignku/bussiness/transactions_item"
	"consignku/drivers/databases/discounts"
	"consignku/drivers/databases/products"
	"time"
)

type TransactionsItem struct {
	ID             int
	TransactionsID int
	ProductsID     int
	Products       products.Products `gorm:"foreignKey:ProductsID;references:ID"`
	ProductName    string
	DiscountsID    int
	Discounts      discounts.Discounts `gorm:"foreignKey:DiscountsID;references:ID"`
	DiscountsVal   int
	ProductsPrice  int
	TotalPrice     int
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

func FromDomain(domain *transactions_item.Domain) *TransactionsItem {
	return &TransactionsItem{
		TransactionsID: domain.TransactionsID,
		ProductsID:     domain.ProductsID,
		ProductName:    domain.ProductName,
		DiscountsID:    domain.DiscountsID,
		DiscountsVal:   domain.DiscountsVal,
		ProductsPrice:  domain.ProductsPrice,
		TotalPrice:     domain.TotalPrice,
	}
}

func (rec *TransactionsItem) ToDomain() transactions_item.Domain {
	return transactions_item.Domain{
		ID:             rec.ID,
		TransactionsID: rec.TransactionsID,
		ProductsID:     rec.ProductsID,
		ProductsPrice:  rec.ProductsPrice,
		ProductName:    rec.Products.Name,
		TotalPrice:     rec.ProductsPrice - rec.DiscountsVal,
		DiscountsID:    rec.DiscountsID,
		DiscountsVal:   rec.Discounts.DiscountValue,
	}
}
