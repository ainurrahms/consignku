package discounts

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID            int
	Code          string
	DiscountValue int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     gorm.DeletedAt
}

type Usecase interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	GetAll(ctx context.Context) ([]Domain, error)
	Store(ctx context.Context, discountsDomain *Domain) error
	Update(ctx context.Context, discountsDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, discountsDomain *Domain) (*Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	FindByID(id int) (Domain, error)
	Update(ctx context.Context, discountsDomain *Domain) (Domain, error)
	Delete(ctx context.Context, discountsDomain *Domain) (Domain, error)
	Store(ctx context.Context, discountsDomain *Domain) error
}
