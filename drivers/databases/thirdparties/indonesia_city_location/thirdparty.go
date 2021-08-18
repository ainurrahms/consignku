package indonesia_city_location

import (
	"consignku/bussiness/indonesia_city_location"
	"context"
	"encoding/json"
	"net/http"
	"strconv"

	"gorm.io/gorm"
)

type IndonesiaCityLocation struct {
	httpClient http.Client
	Conn       *gorm.DB
}

func NewIndonesiaCityLocation(conn *gorm.DB) indonesia_city_location.Repository {
	return &IndonesiaCityLocation{
		httpClient: http.Client{},
		Conn:       conn,
	}
}

func (ipl *IndonesiaCityLocation) GetCityLocation(ctx context.Context, id int) (indonesia_city_location.Domain, error) {
	req, _ := http.NewRequest("GET", "https://dev.farizdotid.com/api/daerahindonesia/kota/"+strconv.Itoa(id), nil)
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

func (nr *IndonesiaCityLocation) Store(ctx context.Context, cityDomain *indonesia_city_location.Domain) error {
	rec := fromDomain(*cityDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}
