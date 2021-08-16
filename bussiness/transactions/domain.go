package transactions

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id           int
	Price        int
	UsersID      int
	Usernames    string
	DiscountsID  int
	DiscountsVal int
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt
}

type Usecase interface {
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
}

type Repository interface {
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
}
