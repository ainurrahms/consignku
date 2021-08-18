package indonesia_city_location

import "consignku/bussiness/indonesia_city_location"

type Response struct {
	IDCity   int    `json:"city_id"`
	CityName string `json:"city_name"`
}

func (resp *Response) toDomain() indonesia_city_location.Domain {
	return indonesia_city_location.Domain{
		IDCity:   resp.IDCity,
		CityName: resp.CityName,
	}
}
