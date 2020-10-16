package order

import (
	"errors"
	"fmt"
	"math/rand"
	"time"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// Status defines order status
type Status string

// Status const
const (
	StatusPending     Status = "PENDING"
	StatusCancelled   Status = "CANCELLED"
	StatusDelivered   Status = "DELIVERED"
	StatusProcesssing Status = "PROCESSING"
)

// ValidStatus holds all the valid order status
var ValidStatus = []Status{
	StatusPending, StatusCancelled, StatusProcesssing, StatusDelivered,
}

// IsValid checks if status is valid
func (status Status) IsValid() bool {
	for _, s := range ValidStatus {
		if status == s {
			return true
		}
	}

	return false
}

// NewStatus creates a status from string
func NewStatus(status string) (Status, error) {
	sts := Status(status)
	if !sts.IsValid() {
		return Status(""), ErrInvalidStatus
	}

	return sts, nil
}

// StatusStateMachine transform state
type StatusStateMachine map[Status][]Status

// CanGo check if the state transformation is valid
func (s StatusStateMachine) CanGo(formState Status, toState Status) bool {
	for _, ps := range s[formState] {
		if ps == toState {
			return true
		}
	}

	return false
}

// DefStatusStateMachine state
var DefStatusStateMachine = StatusStateMachine{
	StatusPending:     []Status{StatusCancelled, StatusProcesssing},
	StatusProcesssing: []Status{StatusCancelled, StatusDelivered},
	StatusDelivered:   []Status{},
	StatusCancelled:   []Status{},
}

// Order defines order
type Order struct {
	ID              int64
	invoiceNumber   string
	items           []*Item
	totalQuantity   int64
	status          Status
	paidAmount      float64
	totalAmount     float64
	subTotal        float64
	Customer        Customer
	PaymentDetails  Payment
	DeliveryDetails Delivery
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
}

// Cancel order
func (o *Order) Cancel() error {
	if ok := DefStatusStateMachine.CanGo(o.GetStatus(), StatusCancelled); !ok {
		return cerror.NewDomainError("invalid status set", nil)
	}

	o.status = StatusCancelled

	return nil
}

// Process order
func (o *Order) Process() error {
	if ok := DefStatusStateMachine.CanGo(o.GetStatus(), StatusProcesssing); !ok {
		return cerror.NewDomainError("invalid status set", nil)
	}

	o.status = StatusProcesssing

	return nil
}

// Deliver order
func (o *Order) Deliver() error {
	if ok := DefStatusStateMachine.CanGo(o.GetStatus(), StatusDelivered); !ok {
		return cerror.NewDomainError("invalid status set", nil)
	}

	o.status = StatusDelivered

	return nil
}

// GetStatus return order status
func (o *Order) GetStatus() Status {
	if o.status != "" {
		return o.status
	}

	return StatusPending
}

// IsEmpty checks whether the order is empty
func (o *Order) IsEmpty() bool {
	return len(o.items) == 0
}

// MakePayment pays order
func (o *Order) MakePayment(amount float64) error {
	if o.IsEmpty() {
		err := cerror.NewDomainError(fmt.Sprintf("can't pay empty order"), nil)
		return err
	}

	if amount <= 0 {
		err := cerror.NewDomainError(fmt.Sprintf("can't pay invalid amount"), nil)
		return err
	}

	if o.totalAmount == amount {
		if err := o.PaymentDetails.SetPaymentStatus(PaymentStatusPaid); err != nil {
			return err
		}
		o.paidAmount = o.totalAmount
	}

	if amount < o.totalAmount {
		if err := o.PaymentDetails.SetPaymentStatus(PaymentStatusPartial); err != nil {
			return err
		}
		o.paidAmount = amount
	}

	return nil
}

// GetPaymentStatus return order payment status
func (o *Order) GetPaymentStatus() PaymentStatus {
	return o.PaymentDetails.GetPaymentStatus()
}

// GetTotalQuantity return order total quantity
func (o *Order) GetTotalQuantity() int64 {
	return o.totalQuantity
}

// GetTotalAmount return order total amount
func (o *Order) GetTotalAmount() float64 {
	return o.totalAmount
}

// GetSubTotal return order sub total amount
func (o *Order) GetSubTotal() float64 {
	return o.subTotal
}

// GenerateInvoiceNumber generates a invoice number
func (o *Order) GenerateInvoiceNumber() string {
	random := func(min int, max int) int {
		rand.Seed(time.Now().UnixNano())
		return rand.Intn(max-min) + min
	}

	t := time.Now()
	return fmt.Sprintf("DDD%v%v", t.UnixNano(), random(111, 999))
}

// SetInvoiceNumber sets valid invoice number
func (o *Order) SetInvoiceNumber() error {
	if o.invoiceNumber == "" {
		o.invoiceNumber = o.GenerateInvoiceNumber()
		return nil
	}

	return errors.New("invoice number is immutable")
}

// GetInvoiceNumber return order invoice number
func (o *Order) GetInvoiceNumber() string {
	if o.invoiceNumber == "" {
		o.invoiceNumber = o.GenerateInvoiceNumber()
	}

	return o.invoiceNumber
}

// AddItem adds item to order
func (o *Order) AddItem(item *Item) error {
	if ok, err := item.IsValid(); !ok {
		return err
	}

	o.items = append(o.items, item)
	o.totalQuantity += item.Quantity
	o.totalAmount += (item.Price * float64(item.Quantity))
	o.subTotal += (item.Price * float64(item.Quantity))

	return nil
}

// AddItems add items to order
func (o *Order) AddItems(items []*Item) error {
	for _, v := range items {
		if err := o.AddItem(v); err != nil {
			return err
		}
	}

	return nil
}

// GetItems return order items
func (o *Order) GetItems() []*Item {
	return o.items
}

// GetDeliveryContactNumber returns provided delivery contact number or customer contact number
func (o *Order) GetDeliveryContactNumber() string {
	if o.DeliveryDetails.ContactNumber != "" {
		return o.DeliveryDetails.ContactNumber
	}

	return o.Customer.ContactNumber
}
