package order

import cerror "github.com/BackAged/go-ddd-microservice/error"

// available product errors
var (
	ErrInvalidStatus        = cerror.NewDomainError("invalid order status", nil)
	ErrInvalidPaymentStatus = cerror.NewDomainError("invalid order payment status", nil)
)
