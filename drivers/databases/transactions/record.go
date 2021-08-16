package transactions

import (
	TransactionsUsecase "consignku/bussiness/transactions"
	"consignku/drivers/databases/discounts"
	"consignku/drivers/databases/users"
	"time"

	"gorm.io/gorm"
)

type Transactions struct {
	Id          int
	Price       int
	UsersID     int
	Users       users.Users `gorm:"foreignKey:UsersID;references:ID"`
	DiscountsID int
	Discounts   discounts.Discounts `gorm:"foreignKey:DiscountsID;references:Id"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt
}

func fromDomain(domain TransactionsUsecase.Domain) *Transactions {
	return &Transactions{
		Id:          domain.Id,
		Price:       domain.Price,
		UsersID:     domain.UsersID,
		DiscountsID: domain.DiscountsID,
		CreatedAt:   domain.CreatedAt,
		UpdatedAt:   domain.UpdatedAt,
		DeletedAt:   domain.DeletedAt,
	}
}
func (rec *Transactions) toDomain() TransactionsUsecase.Domain {
	return TransactionsUsecase.Domain{
		Id:           rec.Id,
		Price:        rec.Price,
		UsersID:      rec.UsersID,
		Usernames:    rec.Users.Username,
		DiscountsID:  rec.DiscountsID,
		DiscountsVal: rec.Discounts.DiscountValue,
		CreatedAt:    rec.CreatedAt,
		UpdatedAt:    rec.UpdatedAt,
		DeletedAt:    rec.DeletedAt,
	}
}
