package product_used_times

import (
	ProductUsedTimesUsecase "consignku/bussiness/product_used_times"
	"time"

	"gorm.io/gorm"
)

type ProductUsedTimes struct {
	Id        int
	UsedTimes string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func fromDomain(domain ProductUsedTimesUsecase.Domain) *ProductUsedTimes {
	return &ProductUsedTimes{
		Id:        domain.Id,
		UsedTimes: domain.UsedTimes,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *ProductUsedTimes) toDomain() ProductUsedTimesUsecase.Domain {
	return ProductUsedTimesUsecase.Domain{
		Id:        rec.Id,
		UsedTimes: rec.UsedTimes,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
