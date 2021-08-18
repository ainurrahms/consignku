package response

import (
	"consignku/bussiness/products"
	"time"

	"gorm.io/gorm"
)

type Products struct {
	ID                 int            `json:"id"`
	Name               string         `json:"name"`
	Price              int            `json:"price"`
	Descriptions       string         `json:"descriptions"`
	Status             bool           `json:"status"`
	ProductTypesID     int            `json:"product_type_id"`
	ProductType        string         `json:"product_type"`
	ProductMerk        string         `json:"product_merk"`
	ProductCode        string         `json:"product_code"`
	ProductUsedTimesID int            `json:"product_used_id"`
	ProductsUsedTimes  string         `json:"product_used_time"`
	Garansi            bool           `json:"garansi"`
	CreatedAt          time.Time      `json:"created_at"`
	UpdatedAt          time.Time      `json:"updated_at"`
	DeletedAt          gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain products.Domain) Products {
	return Products{
		ID:                 domain.ID,
		Name:               domain.Name,
		Price:              domain.Price,
		Descriptions:       domain.Descriptions,
		Status:             domain.Status,
		ProductTypesID:     domain.ProductTypesID,
		ProductType:        domain.ProductType,
		ProductMerk:        domain.ProductMerk,
		ProductCode:        domain.ProductCode,
		ProductUsedTimesID: domain.ProductUsedTimesID,
		ProductsUsedTimes:  domain.ProductsUsedTimes,
		Garansi:            domain.Garansi,
		CreatedAt:          domain.CreatedAt,
		UpdatedAt:          domain.UpdatedAt,
		DeletedAt:          domain.DeletedAt,
	}
}
