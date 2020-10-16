package order

import (
	"fmt"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

// PaymentStatus defines order payment
type PaymentStatus string

// Payment status const
const (
	PaymentStatusPaid    PaymentStatus = "PAID"
	PaymentStatusPartial PaymentStatus = "PARTIAL"
	PaymentStatusUnPaid  PaymentStatus = "UNPAID"
)

// ValidPaymentStatus holds all the valid payment status
var ValidPaymentStatus = []PaymentStatus{
	PaymentStatusPaid, PaymentStatusPartial, PaymentStatusUnPaid,
}

// IsValid checks if payment status is valid
func (pymntSts PaymentStatus) IsValid() bool {
	for _, s := range ValidPaymentStatus {
		if pymntSts == s {
			return true
		}
	}

	return false
}

// NewPaymentStatus creates a payment status from string
func NewPaymentStatus(status string) (PaymentStatus, error) {
	pymntSts := PaymentStatus(status)
	if !pymntSts.IsValid() {
		return PaymentStatus(""), ErrInvalidPaymentStatus
	}

	return pymntSts, nil
}

// PaymentStatusStateMachine transform state
type PaymentStatusStateMachine map[PaymentStatus][]PaymentStatus

// CanGo check if the state transformation is valid
func (psm PaymentStatusStateMachine) CanGo(formState PaymentStatus, toState PaymentStatus) bool {
	for _, ps := range psm[formState] {
		if ps == toState {
			return true
		}
	}

	return false
}

// DefPaymentStatusStateMachine state
var DefPaymentStatusStateMachine = PaymentStatusStateMachine{
	PaymentStatusUnPaid:  []PaymentStatus{PaymentStatusPaid, PaymentStatusPartial},
	PaymentStatusPartial: []PaymentStatus{PaymentStatusPaid},
	PaymentStatusPaid:    []PaymentStatus{PaymentStatusPaid},
}

// Payment defines payment details
type Payment struct {
	status PaymentStatus
	Method string
}

// SetPaymentStatus sets valid payment status for valid state
func (p *Payment) SetPaymentStatus(pymntSts PaymentStatus) error {
	if !pymntSts.IsValid() {
		return cerror.NewDomainError(fmt.Sprintf("can't set order payment status %s", pymntSts), nil)
	}

	if DefPaymentStatusStateMachine.CanGo(p.GetPaymentStatus(), pymntSts) {
		p.status = pymntSts
	} else {
		return cerror.NewDomainError(fmt.Sprintf("can't set order payment status %s", pymntSts), nil)
	}

	return nil
}

// GetPaymentStatus return payment status
func (p *Payment) GetPaymentStatus() PaymentStatus {
	if p.status != "" {
		return p.status
	}

	return PaymentStatusUnPaid
}
