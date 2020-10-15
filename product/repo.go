package product

import "context"

// Repo defines product storage interface
type Repo interface {
	Create(context.Context, *Product) (*Product, error)
	GetByProductID(context.Context, int64) (*Product, error)
	Update(context.Context, *Product) (*Product, error)
	Delete(context.Context, string) error
}
