package discounts_test

import (
	"consignku/app/middleware"
	bussiness "consignku/bussiness"
	discounts "consignku/bussiness/discounts"
	discountsMock "consignku/bussiness/discounts/mocks"
	"context"
	"errors"
	"os"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

var (
	discountsRepository discountsMock.Repository
	discountsUsecase    discounts.Usecase
)

func setup() {
	configJWT := middleware.ConfigJWT{
		SecretJWT:       viper.GetString(`jwt.secret`),
		ExpiresDuration: viper.GetInt(`jwt.expired`),
	}
	discountsUsecase = discounts.NewDiscountsUsecase(&discountsRepository, &configJWT, 2)
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}

func TestAddDiscounts(t *testing.T) {
	t.Run("success test", func(t *testing.T) {

		var domains = discounts.Domain{
			Code:          "INIVOUCHER",
			DiscountValue: 0,
		}

		discountsRepository.On("Store", mock.Anything, mock.AnythingOfType("*discounts.Domain")).Return(nil).Once()
		err := discountsUsecase.Store(context.Background(), &domains)
		assert.NoError(t, err)
	})
}

func TestGetByIDDiscounts(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {

		var domains = discounts.Domain{
			Code:          "INIVOUCHER",
			DiscountValue: 0,
		}

		discountsRepository.On("FindByID", context.Background(), mock.AnythingOfType("int")).Return(domains, nil).Once()
		result, err := discountsUsecase.GetByID(context.Background(), 1)
		assert.NoError(t, err)
		assert.Equal(t, domains.Code, result.Code)
		assert.Equal(t, domains.DiscountValue, result.DiscountValue)
	})

	t.Run("test case 2, id not found", func(t *testing.T) {
		result, err := discountsUsecase.GetByID(context.Background(), -1)

		assert.Equal(t, result, discounts.Domain{})
		assert.Equal(t, err, bussiness.ErrIDNotFound)
	})
}

func TestGetAllDiscounts(t *testing.T) {
	t.Run("success test", func(t *testing.T) {
		domain := []discounts.Domain{
			{
				ID:            1,
				Code:          "INIVOUCHER",
				DiscountValue: 0,
			},
			{
				ID:            2,
				Code:          "INIVOUCHER",
				DiscountValue: 0,
			},
		}
		discountsRepository.On("Find", context.Background()).Return(domain, nil).Once()

		result, err := discountsUsecase.GetAll(context.Background())

		assert.Equal(t, 2, len(result))
		assert.Nil(t, err)
	})

	t.Run("test case 2, repository error", func(t *testing.T) {
		errRepository := errors.New("mysql not running")
		discountsRepository.On("Find", context.Background()).Return([]discounts.Domain{}, errRepository).Once()

		result, err := discountsUsecase.GetAll(context.Background())

		assert.Equal(t, 0, len(result))
		assert.Equal(t, errRepository, err)
	})
}

func TestUpdateDiscounts(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {

		var domains = discounts.Domain{
			Code:          "INIVOUCHERS",
			DiscountValue: 0,
		}

		discountsRepository.On("Update", context.Background(), mock.AnythingOfType("*discounts.Domain")).Return(domains, nil).Once()
		_, err := discountsUsecase.Update(context.Background(), &domains)
		assert.NoError(t, err)
	})
}

func TestDeleteDiscounts(t *testing.T) {
	t.Run("test case 1, success test", func(t *testing.T) {

		var domains = discounts.Domain{
			Code:          "INIVOUCHERS",
			DiscountValue: 0,
		}

		discountsRepository.On("Delete", context.Background(), mock.AnythingOfType("*discounts.Domain")).Return(domains, nil).Once()
		_, err := discountsUsecase.Delete(context.Background(), &domains)
		assert.NoError(t, err)
	})
}
