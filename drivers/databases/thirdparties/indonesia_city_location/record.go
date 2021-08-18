package indonesia_city_location

import (
	IndonesiaCityLocationUsecase "consignku/bussiness/indonesia_city_location"
)

type City struct {
	ID       int
	IDCity   int
	CityName string
}

func fromDomain(domain IndonesiaCityLocationUsecase.Domain) *City {
	return &City{
		ID:       domain.ID,
		IDCity:   domain.IDCity,
		CityName: domain.CityName,
	}
}
