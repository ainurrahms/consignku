package request

import "consignku/bussiness/users"

type Users struct {
	Username string `json:"username"`
	Password string `json:"password"`
	IDCity   int    `json:"city_id"`
}

func (req *Users) ToDomain() *users.Domain {
	return &users.Domain{
		Username: req.Username,
		Password: req.Password,
		IDCity:   req.IDCity,
	}
}
