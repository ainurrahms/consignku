package indonesia_city_location

import "context"

type Domain struct {
	IDCity     int
	IDProvinsi string
	City       string
}

type Repository interface {
	GetCityLocation(ctx context.Context, idcity int) (Domain, error)
}
