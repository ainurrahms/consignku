package products

import (
	ProductsUsecase "consignku/bussiness/products"
	"consignku/drivers/databases/product_types"
	"consignku/drivers/databases/product_used_times"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID                  int
	Name                string
	Price               int
	Descriptions        string
	Status              bool
	ProductsTypeID      int
	ProductsTypes       product_types.ProductTypes `gorm:"foreignKey:ProductsTypeID;references:Id"`
	ProductsUsedTimesID int
	ProductsUsedTimes   product_used_times.ProductUsedTimes `gorm:"foreignKey:ProductsUsedTimesID;references:Id"`
	Garansi             bool
	CreatedAt           time.Time
	UpdatedAt           time.Time
	DeletedAt           gorm.DeletedAt
}

func fromDomain(domain ProductsUsecase.Domain) *Products {
	return &Products{
		ID:                  domain.ID,
		Name:                domain.Name,
		Price:               domain.Price,
		Descriptions:        domain.Descriptions,
		Status:              domain.Status,
		ProductsTypeID:      domain.ProductsTypeID,
		ProductsUsedTimesID: domain.ProductsUsedTimesID,
		Garansi:             domain.Garansi,
		CreatedAt:           domain.CreatedAt,
		UpdatedAt:           domain.UpdatedAt,
		DeletedAt:           domain.DeletedAt,
	}
}

func (rec *Products) toDomain() ProductsUsecase.Domain {
	return ProductsUsecase.Domain{
		ID:                  rec.ID,
		Name:                rec.Name,
		Price:               rec.Price,
		Descriptions:        rec.Descriptions,
		Status:              rec.Status,
		ProductsTypeID:      rec.ProductsTypeID,
		ProductsUsedTimesID: rec.ProductsUsedTimesID,
		Garansi:             rec.Garansi,
		CreatedAt:           rec.CreatedAt,
		UpdatedAt:           rec.UpdatedAt,
		DeletedAt:           rec.DeletedAt,
	}
}
