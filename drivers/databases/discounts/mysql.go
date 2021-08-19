package discounts

import (
	"consignku/bussiness/discounts"
	"context"

	"gorm.io/gorm"
)

type mysqlDiscountsRepository struct {
	Conn *gorm.DB
}

func NewMySQLDiscountsRepository(conn *gorm.DB) discounts.Repository {
	return &mysqlDiscountsRepository{
		Conn: conn,
	}
}

// func (nr *mysqlDiscountsRepository) Fetch(ctx context.Context, page, perpage int) ([]discounts.Domain, int, error) {
// 	rec := []Discounts{}

// 	offset := (page - 1) * perpage
// 	err := nr.Conn.Offset(offset).Limit(perpage).Find(&rec).Error
// 	if err != nil {
// 		return []discounts.Domain{}, 0, err
// 	}

// 	var totalData int64
// 	err = nr.Conn.Count(&totalData).Error
// 	fmt.Println(err)
// 	if err != nil {
// 		return []discounts.Domain{}, 0, err
// 	}

// 	var domainDiscounts []discounts.Domain
// 	for _, value := range rec {
// 		domainDiscounts = append(domainDiscounts, value.toDomain())
// 	}
// 	return domainDiscounts, int(totalData), nil
// }

func (nr *mysqlDiscountsRepository) Store(ctx context.Context, userDomain *discounts.Domain) error {
	rec := fromDomain(*userDomain)

	result := nr.Conn.Create(rec)
	if result.Error != nil {
		return result.Error
	}

	return nil
}

func (cr *mysqlDiscountsRepository) Find(ctx context.Context) ([]discounts.Domain, error) {
	rec := []Discounts{}

	cr.Conn.Find(&rec)
	discoutnsDomain := []discounts.Domain{}
	for _, value := range rec {
		discoutnsDomain = append(discoutnsDomain, value.toDomain())
	}

	return discoutnsDomain, nil
}

func (nr *mysqlDiscountsRepository) FindByID(ctx context.Context, id int) (discounts.Domain, error) {
	rec := Discounts{}

	if err := nr.Conn.Where("id = ?", id).First(&rec).Error; err != nil {
		return discounts.Domain{}, err
	}
	return rec.toDomain(), nil
}

func (nr *mysqlDiscountsRepository) Update(ctx context.Context, discountsDomain *discounts.Domain) (discounts.Domain, error) {
	rec := fromDomain(*discountsDomain)

	result := nr.Conn.Updates(rec)

	if result.Error != nil {
		return discounts.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Discounts").First(&rec, rec.ID).Error

	if err != nil {
		return discounts.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}

func (nr *mysqlDiscountsRepository) Delete(ctx context.Context, discountsDomain *discounts.Domain) (discounts.Domain, error) {
	rec := fromDomain(*discountsDomain)

	result := nr.Conn.Delete(rec)

	if result.Error != nil {
		return discounts.Domain{}, result.Error
	}

	err := nr.Conn.Preload("Discounts").First(&rec, rec.ID).Error

	if err != nil {
		return discounts.Domain{}, result.Error
	}
	return rec.toDomain(), nil
}
