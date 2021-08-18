package request

import (
	"consignku/bussiness/product_used_times"
)

type ProductUsedTimes struct {
	UsedTimes string `json:"used_times"`
}

func (req *ProductUsedTimes) ToDomain() *product_used_times.Domain {
	return &product_used_times.Domain{
		UsedTimes: req.UsedTimes,
	}
}
