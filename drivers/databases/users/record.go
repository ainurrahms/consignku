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
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (rec *Users) toDomain() users.Domain {
	return users.Domain{
		Id:        rec.ID,
		Username:  rec.Username,
		Password:  rec.Password,
		IDCity:    rec.IDCity,
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
		CreatedAt: userDomain.CreatedAt,
		UpdatedAt: userDomain.UpdatedAt,
	}
}
