package response

import (
	"consignku/bussiness/product_used_times"
	"time"
)

type ProductUsedTimes struct {
	Id        int       `json:"id"`
	UsedTimes string    `json:"used_times"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain product_used_times.Domain) ProductUsedTimes {
	return ProductUsedTimes{
		Id:        domain.Id,
		UsedTimes: domain.UsedTimes,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
