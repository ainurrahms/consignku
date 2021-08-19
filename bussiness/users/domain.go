package users

import (
	"context"
	"time"
)

type Domain struct {
	Id        int
	Name      string
	Username  string
	Password  string
	Token     string
	IDCity    int
	CityName  string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Usecase interface {
	Register(ctx context.Context, data *Domain) error
	Login(ctx context.Context, username, password string) (string, error)
	GetByID(ctx context.Context, id int) (Domain, error)
}

type Repository interface {
	GetByUsername(ctx context.Context, username string) (Domain, error)
	FindByID(id int) (Domain, error)
	GetByUsernamePassword(ctx context.Context, username, password string) (Domain, error)
	Store(ctx context.Context, data *Domain) error
}
