// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	indonesia_city_location "consignku/bussiness/indonesia_city_location"
	context "context"

	mock "github.com/stretchr/testify/mock"
)

// Repository is an autogenerated mock type for the Repository type
type Repository struct {
	mock.Mock
}

// GetCityLocation provides a mock function with given fields: ctx, idcity
func (_m *Repository) GetCityLocation(ctx context.Context, idcity int) (indonesia_city_location.Domain, error) {
	ret := _m.Called(ctx, idcity)

	var r0 indonesia_city_location.Domain
	if rf, ok := ret.Get(0).(func(context.Context, int) indonesia_city_location.Domain); ok {
		r0 = rf(ctx, idcity)
	} else {
		r0 = ret.Get(0).(indonesia_city_location.Domain)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, idcity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Store provides a mock function with given fields: ctx, cityDomain
func (_m *Repository) Store(ctx context.Context, cityDomain *indonesia_city_location.Domain) error {
	ret := _m.Called(ctx, cityDomain)

	var r0 error
	if rf, ok := ret.Get(0).(func(context.Context, *indonesia_city_location.Domain) error); ok {
		r0 = rf(ctx, cityDomain)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
