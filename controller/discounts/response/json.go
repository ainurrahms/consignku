package response

import (
	"consignku/bussiness/discounts"
	"time"
)

type Discounts struct {
	ID            int       `json:"id"`
	Code          string    `json:"code"`
	DiscountValue int       `json:"discount_value"`
	CreatedAt     time.Time `json:"created_at"`
	UpdatedAt     time.Time `json:"updated_at"`
}

func FromDomain(domain discounts.Domain) Discounts {
	return Discounts{
		ID:            domain.ID,
		Code:          domain.Code,
		DiscountValue: domain.DiscountValue,
		CreatedAt:     domain.CreatedAt,
		UpdatedAt:     domain.UpdatedAt,
	}
}
