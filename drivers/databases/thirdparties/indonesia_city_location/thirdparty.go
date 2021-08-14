package indonesia_city_location

import (
	"consignku/bussiness/indonesia_city_location"
	"context"
	"encoding/json"
	"net/http"
)

type IndonesiaCityLocation struct {
	httpClient http.Client
}

func NewIndonesiaCityLocation() indonesia_city_location.Repository {
	return &IndonesiaCityLocation{
		httpClient: http.Client{},
	}
}

func (ipl *IndonesiaCityLocation) GetCityLocation(ctx context.Context, id int) (indonesia_city_location.Domain, error) {
	req, _ := http.NewRequest("GET", "https://dev.farizdotid.com/api/daerahindonesia/kota"+id, nil)
	resp, err := ipl.httpClient.Do(req)
	if err != nil {
		return indonesia_city_location.Domain{}, err
	}

	defer resp.Body.Close()

	data := Response{}
	err = json.NewDecoder(resp.Body).Decode(&data)
	if err != nil {
		return indonesia_city_location.Domain{}, err
	}

	return data.toDomain(), nil
}
