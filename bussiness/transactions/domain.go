package transactions

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID           int
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
	Find(ctx context.Context, page, perpage int) ([]Domain, int, int, error)
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Update(ctx context.Context, TransactionsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
}

type Repository interface {
	Find(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Update(ctx context.Context, TransactionsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	FindByID(id int) (Domain, error)
}
