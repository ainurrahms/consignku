package indonesia_city_location

import (
	"context"
	"time"

	"gorm.io/gorm"
)

type Domain struct {
	ID        int
	IDCity    int
	CityName  string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

type Usecase interface {
	Store(ctx context.Context, cityDomain *Domain) error
}
type Repository interface {
	GetCityLocation(ctx context.Context, idcity int) (Domain, error)
	Store(ctx context.Context, cityDomain *Domain) error
}
