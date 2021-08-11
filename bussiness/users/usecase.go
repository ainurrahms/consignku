package users

import (
	"consignku/bussiness"
	"consignku/helper/encrypt"
	"context"
	"strings"
	"time"
)

type userUsecase struct {
	userRepository Repository
	contextTimeout time.Duration
}

func NewUserUseCase(ur Repository, timeout time.Duration) Usecase {
	return &userUsecase{
		userRepository: ur,
		contextTimeout: timeout,
	}
}

func (uc *userUsecase) Register(ctx context.Context, userDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	existedUser, err := uc.userRepository.GetByUsername(ctx, userDomain.Username)

	if err != nil {
		if !strings.Contains(err.Error(),"not found"){
			return  err
		}
	}

	if existedUser != (Domain{}) {
		return bussiness.ErrDuplicateData
	}

	userDomain.Password, err = encrypt.Hash(userDomain.Password)

	if err != nil {
		return bussiness.ErrInternalServer
	}

	err = uc.userRepository.Store(ctx, userDomain)
	if err != nil {
		return err
	}
	
	return nil
}
