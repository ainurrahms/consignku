package transactions

import (
	"context"
	"time"

	transactionItem "consignku/bussiness/transactions_item"

	"gorm.io/gorm"
)

type Domain struct {
	ID               int
	DiscountsID      int
	UsersID          int
	Usernames        string
	ProductsID       int
	ProductsPrice    int
	TransactionItems []transactionItem.Domain
	CreatedAt        time.Time
	UpdatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

type Usecase interface {
	Find(ctx context.Context, page, perpage int) ([]Domain, int, int, error)
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Update(ctx context.Context, TransactionsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	GetByID(ctx context.Context, id, userID int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Update(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	FindByID(id, userID int) (Domain, error)
}
