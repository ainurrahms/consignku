package product_types

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"context"
	"time"
)

type ProductTypesUsecase struct {
	ProductTypesRepository Repository
	contextTimeout         time.Duration
}

func NewProductTypesUsecase(nr Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ProductTypesUsecase{
		ProductTypesRepository: nr,
		contextTimeout:         timeout,
	}
}

// func (nu *ProductTypesUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
// 	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
// 	defer cancel()

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if perpage <= 0 {
// 		perpage = 25
// 	}

// 	res, total, err := nu.ProductTypesRepository.Fetch(ctx, page, perpage)
// 	if err != nil {
// 		return []Domain{}, 0, err
// 	}

// 	return res, total, nil

// }

func (uc *ProductTypesUsecase) Store(ctx context.Context, ProductTypesDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.ProductTypesRepository.Store(ctx, ProductTypesDomain)
	if err != nil {
		return err
	}

	return nil
}

func (uc *ProductTypesUsecase) GetByID(ctx context.Context, id int) (Domain, error) {

	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.ProductTypesRepository.FindByID(id)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (uc *ProductTypesUsecase) Update(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error) {
	res, err := uc.ProductTypesRepository.Update(ctx, ProductTypesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (uc *ProductTypesUsecase) Delete(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error) {
	res, err := uc.ProductTypesRepository.Delete(ctx, ProductTypesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}
