package product_types

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id        int
	Type      string
	Merk      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Store(ctx context.Context, ProductTypesDomain *Domain) error
	Update(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Find(ctx context.Context) ([]Domain, error)
	FindByID(id int) (Domain, error)
	Update(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Store(ctx context.Context, ProductTypesDomain *Domain) error
}
