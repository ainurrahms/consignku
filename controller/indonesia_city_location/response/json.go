package response

import (
	"consignku/bussiness/indonesia_city_location"
	"time"
)

type City struct {
	ID        int       `json:"id"`
	IDCity    int       `json:"city_id"`
	CityName  string    `json:"city_name"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func FromDomain(domain indonesia_city_location.Domain) City {
	return City{
		ID:        domain.ID,
		IDCity:    domain.IDCity,
		CityName:  domain.CityName,
		CreatedAt: domain.CreatedAt,
		UpdatedAt: domain.UpdatedAt,
	}
}
