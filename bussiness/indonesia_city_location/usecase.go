package indonesia_city_location

import (
	"context"
	"time"
)

type IndonesiaCityLocationUsecase struct {
	IndonesiaCityLocationRepository Repository
	contextTimeout                  time.Duration
}

func NewIndonesiaCityLocationusecase(nr Repository, timeout time.Duration) Usecase {
	return &IndonesiaCityLocationUsecase{
		IndonesiaCityLocationRepository: nr,
		contextTimeout:                  timeout,
	}
}

func (uc *IndonesiaCityLocationUsecase) Store(ctx context.Context, IndonesiaCityLocationDomain *Domain) error {
	ctx, cancel := context.WithTimeout(ctx, uc.contextTimeout)
	defer cancel()

	err := uc.IndonesiaCityLocationRepository.Store(ctx, IndonesiaCityLocationDomain)
	if err != nil {
		return err
	}

	return nil
}
