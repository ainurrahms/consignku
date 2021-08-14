package users

import (
	"consignku/bussiness/users"
	"time"
)

type Users struct {
	ID        int
	Username  string
	Password  string
	IDCity    int
	CityName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.ID,
		Username:  rec.Username,
		Password:  rec.Password,
		IDCity:    rec.IDCity,
		CityName:  rec.CityName,
		CreatedAt: rec.CreatedAt,
		UpdatedAt: rec.UpdatedAt,
	}
}

func fromDomain(userDomain users.Domain) *Users {
	return &Users{
		ID:        userDomain.Id,
		Username:  userDomain.Username,
		Password:  userDomain.Password,
		IDCity:    userDomain.IDCity,
		CityName:  userDomain.CityName,
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
