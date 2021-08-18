package request

import (
	"consignku/bussiness/product_types"
)

type ProductTypes struct {
	Type string `json:"type"`
	Merk string `json:"merk"`
	Code string `json:"code"`
}

func (req *ProductTypes) ToDomain() *product_types.Domain {
	return &product_types.Domain{
		Code: req.Code,
		Type: req.Type,
		Merk: req.Merk,
	}
}
