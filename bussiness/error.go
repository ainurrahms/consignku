package bussiness

import "errors"

var (
	ErrDuplicateData               = errors.New("duplicate data")
	ErrInternalServer              = errors.New("something gone wrong, contact administrator")
	ErrUsernamePasswordNotFound    = errors.New("(Username) or (Password) empty")
	ErrIDNotFound                  = errors.New("id not found")
	ErrProductsTypeIDNotFound      = errors.New("product not found")
	ErrProductsUsedTimesIDNotFound = errors.New("used times not found")
	ErrUsersIDNotFound             = errors.New("user not found")
	ErrDiscountsIDNotFound         = errors.New("discounts not found")
)
