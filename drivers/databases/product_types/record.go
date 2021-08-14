package product_types

import (
	ProductTypesUsecase "consignku/bussiness/product_types"
	"time"

	"gorm.io/gorm"
)

type ProductTypes struct {
	Id        int
	Type      string
	Merk      string
	Code      string
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}

func fromDomain(domain ProductTypesUsecase.Domain) *ProductTypes {
	return &ProductTypes{
		Id:        domain.Id,
		Type:      domain.Type,
		Merk:      domain.Merk,
		Code:      domain.Code,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
		DeletedAt: domain.DeletedAt,
	}
}

func (rec *ProductTypes) toDomain() ProductTypesUsecase.Domain {
	return ProductTypesUsecase.Domain{
		Id:        rec.Id,
		Type:      rec.Type,
		Merk:      rec.Merk,
		Code:      rec.Code,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
		DeletedAt: rec.DeletedAt,
	}
}
