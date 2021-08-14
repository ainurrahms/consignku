package product_used_times

import (
	"consignku/bussiness/product_used_times"
	"context"

	"gorm.io/gorm"
)

type mysqlProductUsedTimesRepository struct {
	Conn *gorm.DB
}

func NewMySQLProductUsedTimesRepository(conn *gorm.DB) product_used_times.Repository {
	return &mysqlProductUsedTimesRepository{
		Conn: conn,
	}
}

// func (nr *mysqlProductUsedTimesRepository) Fetch(ctx context.Context, page, perpage int) ([]product_types.Domain, int, error) {
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

func (nr *mysqlProductUsedTimesRepository) Store(ctx context.Context, userDomain *product_used_times.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (nr *mysqlProductUsedTimesRepository) FindByID(id int) (product_used_times.Domain, error) {
	rec := ProductUsedTimes{}

	if err := nr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return product_used_times.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlProductUsedTimesRepository) Update(ctx context.Context, ProductUsedTimes *product_used_times.Domain) (product_used_times.Domain, error) {
	rec := fromDomain(*ProductUsedTimes)

	result := nr.Conn.Updates(rec)

	if result.Error != nil {
		return product_used_times.Domain{}, result.Error
	}

	err := nr.Conn.Preload("ProductTypes").First(&rec, rec.Id).Error

	if err != nil {
		return product_used_times.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (nr *mysqlProductUsedTimesRepository) Delete(ctx context.Context, ProductUsedTimes *product_used_times.Domain) (product_used_times.Domain, error) {
	rec := fromDomain(*ProductUsedTimes)

	result := nr.Conn.Delete(rec)

	if result.Error != nil {
		return product_used_times.Domain{}, result.Error
	}

	err := nr.Conn.Preload("ProductTypes").First(&rec, rec.Id).Error

	if err != nil {
		return product_used_times.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
