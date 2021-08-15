package products

import (
	"consignku/bussiness/products"
	"context"

	"gorm.io/gorm"
)

type mysqlProductsRepository struct {
	Conn *gorm.DB
}

func NewMysqlProductsRepository(conn *gorm.DB) products.Repository {
	return &mysqlProductsRepository{
		Conn: conn,
	}
}

func (nr *mysqlProductsRepository) Store(ctx context.Context, productsDomain *products.Domain) (products.Domain, error) {
	rec := fromDomain(*productsDomain)

	result := nr.Conn.Create(&rec)
	if result.Error != nil {
		return products.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Products").First(&rec, rec.Id).Error
	if err != nil {
		return products.Domain{}, result.Error
	}

	return rec.toDomain(), nil
}

func (nr *mysqlProductsRepository) Update(ctx context.Context, ProductTypesDomain *products.Domain) (products.Domain, error) {
	rec := fromDomain(*ProductTypesDomain)

	result := nr.Conn.Updates(rec)

	if result.Error != nil {
		return products.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Products").First(&rec, rec.Id).Error

	if err != nil {
		return products.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}

func (nr *mysqlProductsRepository) FindByID(id int) (products.Domain, error) {
	rec := Products{}

	if err := nr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return products.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlProductsRepository) Delete(ctx context.Context, ProductUsedTimes *products.Domain) (products.Domain, error) {
	rec := fromDomain(*ProductUsedTimes)

	result := nr.Conn.Delete(rec)

	if result.Error != nil {
		return products.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Products").First(&rec, rec.Id).Error

	if err != nil {
		return products.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
