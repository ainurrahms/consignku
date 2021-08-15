package products

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	Id                  int
	Name                string
	Price               int
	Descriptions        string
	Status              bool
	ProductsTypeID      int
	ProductsUsedTimesID int
	Garansi             bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

type Usecase interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	Store(ctx context.Context, ProductsDomain *Domain) (Domain, error)
	Update(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	// Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindByID(id int) (Domain, error)
	Update(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Store(ctx context.Context, ProductsDomain *Domain) (Domain, error)
}
