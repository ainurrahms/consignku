package response

import (
	"consignku/bussiness/products"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	Name                string         `json:"name"`
	Price               int            `json:"price"`
	Descriptions        string         `json:"descriptions"`
	Status              bool           `json:"status"`
	ProductsTypeID      int            `json:"product_type_id"`
	ProductsUsedTimesID int            `json:"product_used_id"`
	Garansi             bool           `json:"garansi"`
	CreatedAt           time.Time      `json:"created_at"`
	UpdatedAt           time.Time      `json:"updated_at"`
	DeletedAt           gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
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
