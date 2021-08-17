package products

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"consignku/bussiness/product_types"
	"consignku/bussiness/product_used_times"
	"context"
	"time"
)

type productsUsecase struct {
	productsRepository      Repository
	productTypesUsecase     product_types.Usecase
	productUsedTimesUsecase product_used_times.Usecase
	contextTimeout          time.Duration
}

func NewProductsUsecase(nr Repository, pt product_types.Usecase, pu product_used_times.Usecase, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &productsUsecase{
		productsRepository:      nr,
		productTypesUsecase:     pt,
		productUsedTimesUsecase: pu,
		contextTimeout:          timeout,
	}
}

func (uc *productsUsecase) Store(ctx context.Context, ProductTypesDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	_, err := uc.productTypesUsecase.GetByID(ctx, ProductTypesDomain.ProductTypesID)
	if err != nil {
		return Domain{}, bussiness.ErrProductsTypeIDNotFound
	}

	_, errs := uc.productUsedTimesUsecase.GetByID(ctx, ProductTypesDomain.ProductUsedTimesID)
	if errs != nil {
		return Domain{}, bussiness.ErrProductsUsedTimesIDNotFound
	}

	result, err := uc.productsRepository.Store(ctx, ProductTypesDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil

}

func (uc *productsUsecase) Find(ctx context.Context, page, perpage int) ([]Domain, int, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 10
	}
	res, count, err := uc.productsRepository.Find(ctx, page, perpage)

	lastPage := count / perpage

	if count%perpage > 0 {
		lastPage += 1
	}

	if err != nil {
		return []Domain{}, 0, 1, err
	}
	return res, count, lastPage, nil
}

func (uc *productsUsecase) Update(ctx context.Context, ProductTypesDomain *Domain) (*Domain, error) {
	_, err := uc.productTypesUsecase.GetByID(ctx, ProductTypesDomain.ProductTypesID)
	if err != nil {
		return &Domain{}, bussiness.ErrProductsTypeIDNotFound
	}

	_, errs := uc.productUsedTimesUsecase.GetByID(ctx, ProductTypesDomain.ProductUsedTimesID)
	if errs != nil {
		return &Domain{}, bussiness.ErrProductsUsedTimesIDNotFound
	}
	res, err := uc.productsRepository.Update(ctx, ProductTypesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (cu *productsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.productsRepository.FindAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *productsUsecase) GetByID(ctx context.Context, id int) (Domain, error) {

	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.productsRepository.FindByID(id)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (uc *productsUsecase) Delete(ctx context.Context, ProductUsedTimesDomain *Domain) (*Domain, error) {
	res, err := uc.productsRepository.Delete(ctx, ProductUsedTimesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}
