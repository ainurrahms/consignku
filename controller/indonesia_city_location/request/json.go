package request

import (
	"consignku/bussiness/indonesia_city_location"
)

type City struct {
	IDCity   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

func (req *City) ToDomain() *indonesia_city_location.Domain {
	return &indonesia_city_location.Domain{
		IDCity:   req.IDCity,
		CityName: req.CityName,
	}
}
