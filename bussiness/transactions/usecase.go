package transactions

import (
	"consignku/app/middleware"
	"consignku/bussiness"
	"consignku/bussiness/discounts"
	"consignku/bussiness/products"
	"consignku/bussiness/users"
	"context"
	"time"
)

type transactionsUsecase struct {
	transactionsRepository Repository
	usersUsecase           users.Usecase
	productsUsecase        products.Usecase
	discountsUsecase       discounts.Usecase
	contextTimeout         time.Duration
}

func NewTransactionsUsecase(nr Repository, usc users.Usecase, dc discounts.Usecase, pu products.Usecase, jwtauth *middleware.ConfigJWT, timeout time.Duration) Usecase {
	return &transactionsUsecase{
		transactionsRepository: nr,
		usersUsecase:           usc,
		discountsUsecase:       dc,
		contextTimeout:         timeout,
	}
}

func (uc *transactionsUsecase) Find(ctx context.Context, page, perpage int) ([]Domain, int, int, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	if page <= 0 {
		page = 1
	}
	if perpage <= 0 {
		perpage = 10
	}
	res, count, err := uc.transactionsRepository.Find(ctx, page, perpage)

	lastPage := count / perpage

	if count%perpage > 0 {
		lastPage += 1
	}

	if err != nil {
		return []Domain{}, 0, 1, err
	}
	return res, count, lastPage, nil
}

func (uc *transactionsUsecase) Store(ctx context.Context, TransactionsDomain *Domain) (Domain, error) {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	_, err := uc.usersUsecase.GetByID(ctx, TransactionsDomain.UsersID)
	if err != nil {
		return Domain{}, bussiness.ErrProductsTypeIDNotFound
	}

	result, err := uc.transactionsRepository.Store(ctx, TransactionsDomain)
	if err != nil {
		return Domain{}, err
	}

	return result, nil
}

func (cu *transactionsUsecase) GetAll(ctx context.Context) ([]Domain, error) {
	resp, err := cu.transactionsRepository.FindAll(ctx)
	if err != nil {
		return []Domain{}, err
	}
	return resp, nil
}

func (uc *transactionsUsecase) GetByID(ctx context.Context, id, userID int) (Domain, error) {
	if id <= 0 {
		return Domain{}, bussiness.ErrIDNotFound
	}

	resp, err := uc.transactionsRepository.FindByID(id, userID)

	if err != nil {
		return Domain{}, err
	}
	return resp, nil
}

func (uc *transactionsUsecase) Update(ctx context.Context, TransactionsDomain *Domain) (*Domain, error) {
	_, err := uc.usersUsecase.GetByID(ctx, TransactionsDomain.UsersID)
	if err != nil {
		return &Domain{}, bussiness.ErrUsersIDNotFound
	}

	res, err := uc.transactionsRepository.Update(ctx, TransactionsDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}

func (uc *transactionsUsecase) Delete(ctx context.Context, ProductUsedTimesDomain *Domain) (*Domain, error) {
	res, err := uc.transactionsRepository.Delete(ctx, ProductUsedTimesDomain)

	if err != nil {
		return &Domain{}, err
	}

	return &res, nil
}
