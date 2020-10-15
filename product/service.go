package product

import (
	"context"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// Service provides Product port for application.
type Service interface {
	CreateProduct(context.Context, *Product) (*Product, error)
	GetProduct(context.Context, int64) (*Product, error)
	UpdateProduct(context.Context, *Product) (*Product, error)
	DeleteProduct(context.Context, string) error
}

// Service defines service
type service struct {
	repo Repo
}

// NewService returns a new eshop service.
func NewService(prdRepo Repo) Service {
	return &service{
		repo: prdRepo,
	}
}

func (s *service) CreateProduct(ctx context.Context, prd *Product) (*Product, error) {
	if ok, err := prd.IsValid(); !ok {
		return nil, err
	}

	shp, err := s.repo.Create(ctx, prd)
	if err != nil {
		return nil, err
	}

	return shp, nil
}

func (s *service) GetProduct(ctx context.Context, ProductID int64) (*Product, error) {
	shp, err := s.repo.GetByProductID(ctx, ProductID)
	if err != nil {
		return nil, err
	}
	if shp == nil {
		return nil, cerror.NewDomainError(
			"Product not found",
			cerror.NewNotFoundError("Product not found", []string{}),
		)
	}

	return shp, nil
}

func (s *service) UpdateProduct(ctx context.Context, prd *Product) (*Product, error) {
	panic("not implemented") // TODO: Implement
}

func (s *service) DeleteProduct(ctx context.Context, productID string) error {
	return s.repo.Delete(ctx, productID)
}
