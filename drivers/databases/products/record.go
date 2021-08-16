package products

import (
	ProductsUsecase "consignku/bussiness/products"
	"time"

	"gorm.io/gorm"
)

type Products struct {
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

func fromDomain(domain ProductsUsecase.Domain) *Products {
	return &Products{
		Id:                  domain.Id,
		Name:                domain.Name,
		Price:               domain.Price,
		Descriptions:        domain.Descriptions,
		Status:              domain.Status,
		ProductsTypeID:      domain.ProductsTypeID,
		ProductsUsedTimesID: domain.ProductsTypeID,
		Garansi:             domain.Garansi,
		CreatedAt:           domain.CreatedAt,
		UpdatedAt:           domain.UpdatedAt,
		DeletedAt:           domain.DeletedAt,
	}
}

func (rec *Products) toDomain() ProductsUsecase.Domain {
	return ProductsUsecase.Domain{
		Id:                  rec.Id,
		Name:                rec.Name,
		Price:               rec.Price,
		Descriptions:        rec.Descriptions,
		Status:              rec.Status,
		ProductsTypeID:      rec.ProductsTypeID,
		ProductsUsedTimesID: rec.ProductsTypeID,
		Garansi:             rec.Garansi,
		CreatedAt:           rec.CreatedAt,
		UpdatedAt:           rec.UpdatedAt,
		DeletedAt:           rec.DeletedAt,
	}
}
