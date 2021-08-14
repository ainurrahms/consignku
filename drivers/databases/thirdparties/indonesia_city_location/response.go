package indonesia_city_location

import "consignku/bussiness/indonesia_city_location"

type Response struct {
	IDCity     int    `json:"id_city"`
	IDProvinsi string `json:"id_provinsi"`
	City       string `json:"city"`
}

func (resp *Response) toDomain() indonesia_city_location.Domain {
	return indonesia_city_location.Domain{
		IDCity:     resp.IDCity,
		IDProvinsi: resp.IDProvinsi,
		City:       resp.City,
	}
}
