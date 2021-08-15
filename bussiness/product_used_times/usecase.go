package product_used_times

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"context"
	"time"
)

type ProductUsedTimesUsecase struct {
	ProductUsedTimesRepository Repository
	contextTimeout             time.Duration
}

func NewProductUsedTimesUsecase(nr Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &ProductUsedTimesUsecase{
		ProductUsedTimesRepository: nr,
		contextTimeout:             timeout,
	}
}

// func (nu *ProductUsedTimesUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
// 	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
// 	defer cancel()

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if perpage <= 0 {
// 		perpage = 25
// 	}

// 	res, total, err := nu.ProductUsedTimesRepository.Fetch(ctx, page, perpage)
// 	if err != nil {
// 		return []Domain{}, 0, err
// 	}

// 	return res, total, nil

// }

func (uc *ProductUsedTimesUsecase) Store(ctx context.Context, ProductUsedTimesDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.ProductUsedTimesRepository.Store(ctx, ProductUsedTimesDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cu *ProductUsedTimesUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.ProductUsedTimesRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *ProductUsedTimesUsecase) GetByID(ctx context.Context, id int) (Domain, error) {

	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.ProductUsedTimesRepository.FindByID(id)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (uc *ProductUsedTimesUsecase) Update(ctx context.Context, ProductUsedTimesDomain *Domain) (*Domain, error) {
	res, err := uc.ProductUsedTimesRepository.Update(ctx, ProductUsedTimesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (uc *ProductUsedTimesUsecase) Delete(ctx context.Context, ProductUsedTimesDomain *Domain) (*Domain, error) {
	res, err := uc.ProductUsedTimesRepository.Delete(ctx, ProductUsedTimesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}
