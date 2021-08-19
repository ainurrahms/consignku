package response

import (
	"consignku/bussiness/product_types"
	"time"
)

type ProductTypes struct {
	Id        int       `json:"id"`
	Code      string    `json:"code"`
	Type      string    `json:"type"`
	Merk      string    `json:"merk"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain product_types.Domain) ProductTypes {
	return ProductTypes{
		Id:        domain.Id,
		Code:      domain.Code,
		Type:      domain.Type,
		Merk:      domain.Merk,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
