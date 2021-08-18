package transactions

import (
	TransactionsUsecase "consignku/bussiness/transactions"
	TransactionsItemDomain "consignku/bussiness/transactions_item"
	"consignku/drivers/databases/transactions_item"
	"consignku/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	ID               int
	TransactionItems []transactions_item.TransactionsItem
	UsersID          int
	Users            users.Users `gorm:"foreignKey:UsersID;references:ID"`
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func FromDomain(domain TransactionsUsecase.Domain) *Transactions {
	var items []transactions_item.TransactionsItem
	for _, item := range domain.TransactionItems {
		items = append(items, *transactions_item.FromDomain(&item))
	}
	return &Transactions{
		ID:               domain.ID,
		UsersID:          domain.UsersID,
		TransactionItems: items,
		CreatedAt:        domain.CreatedAt,
		UpdatedAt:        domain.UpdatedAt,
		DeletedAt:        domain.DeletedAt,
	}
}
func (rec *Transactions) ToDomain() TransactionsUsecase.Domain {
	var items []TransactionsItemDomain.Domain
	for _, item := range rec.TransactionItems {
		items = append(items, item.ToDomain())
	}
	return TransactionsUsecase.Domain{
		ID:               rec.ID,
		UsersID:          rec.UsersID,
		Usernames:        rec.Users.Username,
		TransactionItems: items,
		CreatedAt:        rec.CreatedAt,
		UpdatedAt:        rec.UpdatedAt,
		DeletedAt:        rec.DeletedAt,
	}
}
