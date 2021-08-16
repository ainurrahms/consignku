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

	err := nr.Conn.Preload("Transactions").First(&rec, rec.ID).Error
	if err != nil {
		return transactions.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlTransactionsRepository) Find(ctx context.Context, page, perpage int) ([]transactions.Domain, int, error) {
	rec := []Transactions{}
	offset := (page - 1) * perpage

	err := nr.Conn.Joins("Users").Joins("Discounts").Find(&rec).Offset(offset).Limit(perpage).Error

	if err != nil {
		return []transactions.Domain{}, 0, err
	}
	count := nr.Conn.Joins("Users").Joins("Discounts").Find(&rec).Offset(offset).Limit(perpage).RowsAffected

	if err != nil {
		return []transactions.Domain{}, 0, err
	}

	var res []transactions.Domain
	for _, value := range rec {
		res = append(res, value.toDomain())
	}

	return res, int(count), nil
}

func (cr *mysqlTransactionsRepository) FindAll(ctx context.Context) ([]transactions.Domain, error) {
	rec := []Transactions{}

	cr.Conn.Joins("Users").Joins("Discounts").Find(&rec)
	productUsedTimesDomain := []transactions.Domain{}
	for _, value := range rec {
		productUsedTimesDomain = append(productUsedTimesDomain, value.toDomain())
	}

	return productUsedTimesDomain, nil
}

func (nr *mysqlTransactionsRepository) FindByID(id int) (transactions.Domain, error) {
	rec := Transactions{}

	if err := nr.Conn.Joins("Users").Joins("Discounts").Where("transactions.id = ?", id).First(&rec).Error; err != nil {
		return transactions.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlTransactionsRepository) Update(ctx context.Context, TransactionsDomain *transactions.Domain) (transactions.Domain, error) {
	rec := fromDomain(*TransactionsDomain)

	result := nr.Conn.Updates(rec)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Transactions").First(&rec, rec.ID).Error

	if err != nil {
		return transactions.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}

func (nr *mysqlTransactionsRepository) Delete(ctx context.Context, TransactionsDomain *transactions.Domain) (transactions.Domain, error) {
	rec := fromDomain(*TransactionsDomain)

	result := nr.Conn.Delete(rec)

	if result.Error != nil {
		return transactions.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Transactions").First(&rec, rec.ID).Error

	if err != nil {
		return transactions.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
