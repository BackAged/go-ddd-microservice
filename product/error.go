package product

import (
	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// available product errors
var (
	ErrInvalidStatus = cerror.NewDomainError("invalid product status", nil)
)
