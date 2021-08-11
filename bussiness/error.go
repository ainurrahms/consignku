package bussiness

import "errors"

var (
	ErrDuplicateData = errors.New("duplicate data")
	ErrInternalServer = errors.New("something gone wrong, contact administrator")
)