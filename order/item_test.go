package order_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"github.com/BackAged/go-ddd-microservice/order"
)

var _ = Describe("Item", func() {
	Describe("Item", func() {
		Context("invalid item IsValid", func() {
			It("should return error", func() {
				itm := order.Item{}
				_, err := itm.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("valid item IsValid", func() {
			It("should not return error", func() {
				itm := order.Item{
					ProductID: 2,
					Price:     100,
					Quantity:  2,
				}
				_, err := itm.IsValid()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with price less than 0 item IsValid", func() {
			It("should not return error", func() {
				itm := order.Item{
					ProductID: 2,
					Price:     -100,
					Quantity:  2,
				}
				_, err := itm.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with stock less than 0 product IsValid", func() {
			It("should not return error", func() {
				itm := order.Item{
					ProductID: 2,
					Price:     -100,
					Quantity:  2,
				}
				_, err := itm.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
