package transactions

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"consignku/bussiness/discounts"
	"consignku/bussiness/users"
	"context"
	"time"
)

type transactionsUsecase struct {
	transactionsRepository Repository
	usersUsecase           users.Usecase
	discountsUsecase       discounts.Usecase
	contextTimeout         time.Duration
}

func NewTransactionsUsecase(nr Repository, usc users.Usecase, dc discounts.Usecase, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &transactionsUsecase{
		transactionsRepository: nr,
		usersUsecase:           usc,
		discountsUsecase:       dc,
		contextTimeout:         timeout,
	}
}

func (uc *transactionsUsecase) Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	_, err := uc.usersUsecase.GetByID(ctx, TransactionsDomain.UsersID)
	if err != nil {
		return Domain{}, bussiness.ErrProductsTypeIDNotFound
	}

	_, errs := uc.discountsUsecase.GetByID(ctx, TransactionsDomain.DiscountsID)
	if errs != nil {
		return Domain{}, bussiness.ErrProductsTypeIDNotFound
	}

	result, err := uc.transactionsRepository.Store(ctx, TransactionsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}
