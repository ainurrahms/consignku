package response

import (
	"consignku/bussiness/discounts"
	"time"

	"gorm.io/gorm"
)

type Discounts struct {
	Id            int            `json:"id"`
	Code          string         `json:"code"`
	DiscountValue int            `json:"discount_value"`
	CreatedAt     time.Time      `json:"created_at"`
	UpdatedAt     time.Time      `json:"updated_at"`
	DeletedAt     gorm.DeletedAt `json:"deleted_at"`
}

func FromDomain(domain discounts.Domain) Discounts {
	return Discounts{
		Id:            domain.Id,
		Code:          domain.Code,
		DiscountValue: domain.DiscountValue,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
