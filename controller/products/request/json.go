package request

import "consignku/bussiness/products"

type Products struct {
	Name               string `json:"name"`
	Price              int    `json:"price"`
	Descriptions       string `json:"descriptions"`
	Status             bool   `json:"status"`
	ProductTypesID     int    `json:"product_type_id"`
	ProductUsedTimesID int    `json:"product_used_id"`
	Garansi            bool   `json:"garansi"`
}

func (req *Products) ToDomain() *products.Domain {
	return &products.Domain{
		Name:               req.Name,
		Price:              req.Price,
		Descriptions:       req.Descriptions,
		Status:             req.Status,
		ProductTypesID:     req.ProductTypesID,
		ProductUsedTimesID: req.ProductUsedTimesID,
		Garansi:            req.Garansi,
	}
}
