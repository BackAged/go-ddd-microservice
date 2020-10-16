package order_test

import (
	"github.com/BackAged/go-ddd-microservice/order"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Order", func() {
	Describe("Order NewStatus", func() {
		Context("with valid status", func() {
			It("should return  status", func() {
				_, err := order.NewStatus(string(order.StatusPending))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with invalid status", func() {
			It("should return error", func() {
				_, err := order.NewStatus("error")
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Order GetStatus", func() {
		Context("when empty order", func() {
			It("should return pending", func() {
				ordr := order.Order{}
				Expect(ordr.GetStatus()).To(Equal(order.StatusPending))
			})
		})
	})

	Describe("Order Cancel", func() {
		Context("when order is pending", func() {
			It("should be processable", func() {
				ordr := order.Order{}
				err := ordr.Process()
				Expect(err).ToNot(HaveOccurred())
			})
		})
		Context("when order is pending", func() {
			It("should be cancellable", func() {
				ordr := order.Order{}
				err := ordr.Cancel()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when order is pending", func() {
			It("should not be deliverable", func() {
				ordr := order.Order{}
				err := ordr.Deliver()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when order is processing", func() {
			It("should be deliverable", func() {
				ordr := order.Order{}
				err := ordr.Process()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Deliver()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when order is processing", func() {
			It("should be cancelable", func() {
				ordr := order.Order{}
				err := ordr.Process()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Cancel()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("when order is cancelled", func() {
			It("should not be deliverable", func() {
				ordr := order.Order{}
				err := ordr.Cancel()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Deliver()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when order is cancelled", func() {
			It("should not be processable", func() {
				ordr := order.Order{}
				err := ordr.Cancel()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Deliver()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("when order is delivered", func() {
			It("should not be cancellable", func() {
				ordr := order.Order{}
				err := ordr.Process()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Deliver()
				Expect(err).ToNot(HaveOccurred())

				err = ordr.Cancel()
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Order Empty", func() {
		Context("with no items", func() {
			It("should be empty order", func() {
				ordr := order.Order{}
				Expect(ordr.IsEmpty()).To(BeTrue())
			})
		})

		Context("with item", func() {
			It("should not be empty", func() {
				ordr := order.Order{}
				ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     2,
					Quantity:  2,
				})
				Expect(ordr.IsEmpty()).To(BeFalse())
			})
		})
	})

	Describe("Order Add Item", func() {
		Context("with valid item", func() {
			It("should not return error", func() {
				ordr := order.Order{}
				err := ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     2,
					Quantity:  2,
				})
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with valid item", func() {
			It("should update totalamount", func() {
				ordr := order.Order{}
				err := ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     2,
					Quantity:  2,
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetTotalAmount()).To(Equal(float64(4)))
			})
		})

		Context("with valid item", func() {
			It("should update quantity", func() {
				ordr := order.Order{}
				err := ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     2,
					Quantity:  2,
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetTotalQuantity()).To(Equal(int64(2)))
			})
		})

		Context("with valid item", func() {
			It("should update subtotal", func() {
				ordr := order.Order{}
				err := ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     2,
					Quantity:  2,
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetSubTotal()).To(Equal(float64(4)))
			})
		})

		Context("with invalid item", func() {
			It("should return error", func() {
				ordr := order.Order{}
				err := ordr.AddItem(&order.Item{
					ProductID: 2,
					Price:     -2,
					Quantity:  2,
				})
				Expect(err).To(HaveOccurred())
			})
		})
	})

	Describe("Order Add Items", func() {
		Context("with valid items", func() {
			It("should not return error", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with ivalid items", func() {
			It("should return error", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     -2,
						Quantity:  2,
					},
				})
				Expect(err).To(HaveOccurred())
			})
		})

		Context("with valid items", func() {
			It("should update totalamount", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetTotalAmount()).To(Equal(float64(8)))
			})
		})

		Context("with valid items", func() {
			It("should update subtotal", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetSubTotal()).To(Equal(float64(8)))
			})
		})

		Context("with valid items", func() {
			It("should update totalquantity", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetTotalAmount()).To(Equal(float64(8)))
			})
		})
	})

	Describe("Order GetInvoice", func() {
		Context("", func() {
			It("should not be empty", func() {
				ordr := order.Order{}
				Expect(ordr.GetInvoiceNumber()).ToNot(BeEmpty())
			})
		})
	})

	Describe("Order SetInvoice", func() {
		Context("when invoice number empty", func() {
			It("should be set", func() {
				ordr := order.Order{}
				Expect(ordr.SetInvoiceNumber()).ToNot(HaveOccurred())
			})
		})

		Context("when invoice number not empty", func() {
			It("should not be set", func() {
				ordr := order.Order{}
				ordr.SetInvoiceNumber()
				Expect(ordr.SetInvoiceNumber()).To(HaveOccurred())
			})
		})
	})

	Describe("Order GenerateInvoiceNumber", func() {
		Context("", func() {
			It("should return invoice", func() {
				ordr := order.Order{}
				Expect(ordr.GenerateInvoiceNumber()).ToNot(BeEmpty())
			})
		})
	})

	Describe("Order MakePayment", func() {
		Context("with payment equal of totalamount", func() {
			It("should not return error", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				err = ordr.MakePayment(float64(8))
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with payment equal of totalamount", func() {
			It("order payment status should be paid", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				err = ordr.MakePayment(float64(8))
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetPaymentStatus()).To(Equal(order.PaymentStatusPaid))
			})
		})

		Context("with payment less than totalamount", func() {
			It("order payment status should be partial", func() {
				ordr := order.Order{}
				err := ordr.AddItems([]*order.Item{
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
					&order.Item{
						ProductID: 2,
						Price:     2,
						Quantity:  2,
					},
				})
				err = ordr.MakePayment(float64(4))
				Expect(err).ToNot(HaveOccurred())
				Expect(ordr.GetPaymentStatus()).To(Equal(order.PaymentStatusPartial))
			})
		})

		Context("with empty order", func() {
			It("should return error", func() {
				ordr := order.Order{}
				err := ordr.MakePayment(float64(4))
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with invalid amount", func() {
			It("should return error", func() {
				ordr := order.Order{}
				err := ordr.MakePayment(float64(-4))
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
