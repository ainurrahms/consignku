package request

import "consignku/bussiness/discounts"

type Discounts struct {
	Code          string `json:"code"`
	DiscountValue int    `json:"discount_value"`
}

func (req *Discounts) ToDomain() *discounts.Domain {
	return &discounts.Domain{
		Code:          req.Code,
		DiscountValue: req.DiscountValue,
	}
}
