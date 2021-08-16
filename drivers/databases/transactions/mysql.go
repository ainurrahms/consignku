package transactions

import (
	"consignku/bussiness/transactions"
	"context"

	"gorm.io/gorm"
)

type mysqlTransactionsRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductsRepository(conn *gorm.DB) transactions.Repository {
	return &mysqlTransactionsRepository{
		Conn: conn,
	}
}

func (nr *mysqlTransactionsRepository) Store(ctx context.Context, transactionsDomain *transactions.Domain) (transactions.Domain, error) {
	rec := fromDomain(*transactionsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Transactions").First(&rec, rec.Id).Error
	if err != nil {
		return transactions.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}
