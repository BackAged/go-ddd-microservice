package order

import (
	"context"

	"github.com/BackAged/go-ddd-microservice/product"
)

// Service defines port for applicationi adapter
type Service interface {
	PlaceOrder(context.Context, *Order) (*Order, error)
	GetOrder(context.Context, string) (*Order, error)
	ListOrderByCustomerID(context.Context, int64, int64, int64) ([]*Order, error)
	CancelOrder(context.Context, string) (*Order, error)
}

// Service defines service
type service struct {
	prdSvc product.Service
	repo   Repo
}

// NewService returns a new order service.
func NewService(ordrRepo Repo, prdSvc product.Service) Service {
	return &service{
		repo:   ordrRepo,
		prdSvc: prdSvc,
	}
}

func (s *service) PlaceOrder(ctx context.Context, ordr *Order) (*Order, error) {
	return nil, nil
}

func (s *service) GetOrder(ctx context.Context, invoiceID string) (*Order, error) {
	return nil, nil
}

func (s *service) ListOrderByCustomerID(ctx context.Context, customerID int64, skip int64, limit int64) ([]*Order, error) {
	return []*Order{}, nil
}

func (s *service) CancelOrder(ctx context.Context, invoiceID string) (*Order, error) {
	return nil, nil
}
