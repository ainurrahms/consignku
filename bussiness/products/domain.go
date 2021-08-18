package products

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID                 int
	Name               string
	Price              int
	Descriptions       string
	Status             bool
	ProductTypesID     int
	ProductUsedTimesID int
	ProductType        string
	ProductMerk        string
	ProductCode        string
	ProductsUsedTimes  string
	Garansi            bool
	IDCity             int
	CreatedAt          time.Time
	UpdatedAt          time.Time
	DeletedAt          gorm.DeletedAt
}

type Usecase interface {
	Find(ctx context.Context, page, perpage int) ([]Domain, int, int, error)
	Store(ctx context.Context, ProductsDomain *Domain) (Domain, error)
	Update(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error)
	GetAll(ctx context.Context) ([]Domain, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	Find(ctx context.Context, page, perpage int) ([]Domain, int, error)
	FindAll(ctx context.Context) ([]Domain, error)
	FindByID(id int) (Domain, error)
	Update(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Delete(ctx context.Context, ProductTypesDomain *Domain) (Domain, error)
	Store(ctx context.Context, ProductsDomain *Domain) (Domain, error)
}
