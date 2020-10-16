package order_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BackAged/go-ddd-microservice/order"
)

var _ = Describe("Payment", func() {
	Describe("NewPaymentStatus", func() {
		Context("with valid payment status", func() {
			It("should return payment status", func() {
				_, err := order.NewPaymentStatus(string(order.PaymentStatusPaid))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid valid payment status", func() {
			It("should return error", func() {
				_, err := order.NewPaymentStatus(string("invalid"))
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("PaymentStatusStateMachine", func() {
		Context("paymenstaus can go from unpaid to paid", func() {
			It("should return  true", func() {
				ok := order.DefPaymentStatusStateMachine.CanGo(order.PaymentStatusUnPaid, order.PaymentStatusPaid)
				Expect(ok).To(BeTrue())
			})
		})

		Context("paymenstaus can not go from paid to unpaid", func() {
			It("should return  false", func() {
				ok := order.DefPaymentStatusStateMachine.CanGo(order.PaymentStatusPaid, order.PaymentStatusUnPaid)
				Expect(ok).To(BeFalse())
			})
		})
	})

	Describe("SetPaymentStatus", func() {
		Context("from unpaid to paid", func() {
			It("should be ok", func() {
				pymnt := order.Payment{
					Method: "bank",
				}

				Expect(pymnt.SetPaymentStatus(order.PaymentStatusPaid)).ToNot(HaveOccurred())
			})
		})

		Context("from paid to unpaid", func() {
			It("should not be ok", func() {
				pymnt := order.Payment{
					Method: "bank",
				}
				pymnt.SetPaymentStatus(order.PaymentStatusPaid)
				Expect(pymnt.SetPaymentStatus(order.PaymentStatusUnPaid)).To(HaveOccurred())
			})
		})
	})
})
