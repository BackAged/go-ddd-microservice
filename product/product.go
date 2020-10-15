package product

import (
	"time"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// Status defines shopitem status
type Status string

// Valid product status
const (
	StatusActive   = Status("ACTIVE")
	StatusPending  = Status("PENDING")
	StatusInActive = Status("INACTIVE")
)

// NewStatus returns status from  a string
func NewStatus(sts string) (Status, error) {
	stsFrmStr := Status(sts)

	if ok, err := stsFrmStr.IsValid(); !ok {
		return Status(""), err
	}

	return stsFrmStr, nil
}

// IsValid returns if status is valid or not
func (s Status) IsValid() (bool, error) {
	if s != StatusActive && s != StatusInActive && s != StatusPending {
		return false, ErrInvalidStatus
	}

	return true, nil
}

// Product defines product type
type Product struct {
	ID              int64
	ProductID       int64
	Image           string
	Name            string
	Slug            string
	Price           float64
	DiscountedPrice float64
	Stock           int64
	Status          Status
	Version         int64
	CreatedAt       time.Time
	UpdatedAt       time.Time
}

// IsValid checks if product is valid or not
func (p Product) IsValid() (bool, error) {
	ve := cerror.NewValidationError()

	if p.Image == "" {
		ve.Add("image", "required")
	}
	if p.Name == "" {
		ve.Add("name", "required")
	}
	if p.Slug == "" {
		ve.Add("slug", "required")
	}
	if p.Price < 0 {
		ve.Add("price", "invalid")
	}
	if p.Stock < 0 {
		ve.Add("stock", "invalid")
	}
	if p.DiscountedPrice > p.Price || p.DiscountedPrice < 0 {
		ve.Add("discounted_price", "invalid")
	}
	if ok, _ := p.Status.IsValid(); !ok {
		ve.Add("status", "invalid")
	}

	if ve.HasErrors() {
		return false, cerror.NewDomainError("invalid product", ve)
	}

	return true, nil
}
