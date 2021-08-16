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

	err := nr.Conn.Preload("Products").First(&rec, rec.ID).Error
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

	err := nr.Conn.Preload("Products").First(&rec, rec.ID).Error

	if err != nil {
		return products.Domain{}, result.Error
	}
	return rec.toDomain(), nil

}

func (nr *mysqlProductsRepository) Find(ctx context.Context, page, perpage int) ([]products.Domain, int, error) {
	rec := []Products{}
	offset := (page - 1) * perpage

	err := nr.Conn.Joins("ProductUsedTimes").Joins("ProductTypes").Find(&rec).Offset(offset).Limit(perpage).Error

	if err != nil {
		return []products.Domain{}, 0, err
	}
	count := nr.Conn.Joins("ProductUsedTimes").Joins("ProductTypes").Find(&rec).Offset(offset).Limit(perpage).RowsAffected

	if err != nil {
		return []products.Domain{}, 0, err
	}

	var res []products.Domain
	for _, value := range rec {
		res = append(res, value.toDomain())
	}

	return res, int(count), nil
}

func (cr *mysqlProductsRepository) FindAll(ctx context.Context) ([]products.Domain, error) {
	rec := []Products{}

	cr.Conn.Find(&rec)
	productUsedTimesDomain := []products.Domain{}
	for _, value := range rec {
		productUsedTimesDomain = append(productUsedTimesDomain, value.toDomain())
	}

	return productUsedTimesDomain, nil
}

func (nr *mysqlProductsRepository) FindByID(id int) (products.Domain, error) {
	rec := Products{}

	if err := nr.Conn.Joins("ProductTypes").Joins("ProductUsedTimes").Where("id = ?", id).First(&rec).Error; err != nil {
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

	err := nr.Conn.Preload("Products").First(&rec, rec.ID).Error

	if err != nil {
		return products.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
