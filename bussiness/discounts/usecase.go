package discounts

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"context"
	"time"
)

type discountsUsecase struct {
	discountsRepository Repository
	contextTimeout      time.Duration
}

func NewDiscountsUsecase(nr Repository, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &discountsUsecase{
		discountsRepository: nr,
		contextTimeout:      timeout,
	}
}

// func (nu *discountsUsecase) Fetch(ctx context.Context, page, perpage int) ([]Domain, int, error) {
// 	ctx, cancel := context.WithTimeout(ctx, nu.contextTimeout)
// 	defer cancel()

// 	if page <= 0 {
// 		page = 1
// 	}
// 	if perpage <= 0 {
// 		perpage = 25
// 	}

// 	res, total, err := nu.discountsRepository.Fetch(ctx, page, perpage)
// 	if err != nil {
// 		return []Domain{}, 0, err
// 	}

// 	return res, total, nil

// }

func (uc *discountsUsecase) Store(ctx context.Context, discountsDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.discountsRepository.Store(ctx, discountsDomain)
	if err != nil {
		return err
	}

	return nil
}

func (cu *discountsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.discountsRepository.Find(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *discountsUsecase) GetByID(ctx context.Context, id int) (Domain, error) {

	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.discountsRepository.FindByID(id)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (uc *discountsUsecase) Update(ctx context.Context, discountsDomain *Domain) (*Domain, error) {
	res, err := uc.discountsRepository.Update(ctx, discountsDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (uc *discountsUsecase) Delete(ctx context.Context, discountsDomain *Domain) (*Domain, error) {
	res, err := uc.discountsRepository.Delete(ctx, discountsDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}
