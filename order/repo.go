package order

import "context"

// Repo defines interface for storage
type Repo interface {
	Create(context.Context, *Order) (*Order, error)
	Get(context.Context, string) (*Order, error)
	ListByCustomerID(context.Context, int64, int64, int64) ([]*Order, error)
	Update(context.Context, *Order) (*Order, error)
}
