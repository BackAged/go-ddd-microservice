package order

import (
	"time"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// Item deffines order items
type Item struct {
	ID        int64
	ProductID int64
	Name      string
	Image     string
	Slug      string
	Quantity  int64
	Price     float64
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

// IsValid checks if order is valid or not
func (i Item) IsValid() (bool, error) {
	ve := cerror.NewValidationError()

	if i.ProductID == 0 {
		ve.Add("product_id", "required")
	}
	if i.Price < 0 {
		ve.Add("price", "invalid")
	}
	if i.Quantity <= 0 {
		ve.Add("stock", "invalid")
	}

	if ve.HasErrors() {
		return false, cerror.NewDomainError("invalid order", ve)
	}

	return true, nil
}
