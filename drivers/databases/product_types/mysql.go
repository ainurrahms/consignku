package product_types

import (
	"consignku/bussiness/product_types"
	"context"

	"gorm.io/gorm"
)

type mysqlProductTypesRepository struct {
	Conn *gorm.DB
}

func NewMySQLProductTypesRepository(conn *gorm.DB) product_types.Repository {
	return &mysqlProductTypesRepository{
		Conn: conn,
	}
}

// func (nr *mysqlProductTypesRepository) Fetch(ctx context.Context, page, perpage int) ([]product_types.Domain, int, error) {
// 	rec := []ProductTypes{}

// 	offset := (page - 1) * perpage
// 	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
// 	if err != nil {
// 		return []product_types.Domain{}, 0, err
// 	}

// 	var totalData int64
// 	err = nr.Conn.Count(&totalData).Error
// 	if err != nil {
// 		return []product_types.Domain{}, 0, err
// 	}

// 	var domainProductTypes []product_types.Domain
// 	for _, value := range rec {
// 		domainProductTypes = append(domainProductTypes, value.toDomain())
// 	}
// 	return domainProductTypes, int(totalData), nil
// }

func (nr *mysqlProductTypesRepository) Store(ctx context.Context, userDomain *product_types.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *mysqlProductTypesRepository) Find(ctx context.Context) ([]product_types.Domain, error) {
	rec := []ProductTypes{}

	cr.Conn.Find(&rec)
	productTypesDomain := []product_types.Domain{}
	for _, value := range rec {
		productTypesDomain = append(productTypesDomain, value.toDomain())
	}

	return productTypesDomain, nil
}

func (nr *mysqlProductTypesRepository) FindByID(id int) (product_types.Domain, error) {
	rec := ProductTypes{}

	if err := nr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return product_types.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlProductTypesRepository) Update(ctx context.Context, ProductTypesDomain *product_types.Domain) (product_types.Domain, error) {
	rec := fromDomain(*ProductTypesDomain)

	result := nr.Conn.Updates(rec)

	if result.Error != nil {
		return product_types.Domain{}, result.Error
	}

	err := nr.Conn.Preload("ProductTypes").First(&rec, rec.Id).Error

	if err != nil {
		return product_types.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (nr *mysqlProductTypesRepository) Delete(ctx context.Context, ProductTypesDomain *product_types.Domain) (product_types.Domain, error) {
	rec := fromDomain(*ProductTypesDomain)

	result := nr.Conn.Delete(rec)

	if result.Error != nil {
		return product_types.Domain{}, result.Error
	}

	err := nr.Conn.Preload("ProductTypes").First(&rec, rec.Id).Error

	if err != nil {
		return product_types.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
