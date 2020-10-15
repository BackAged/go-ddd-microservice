package product_test

import (
	"github.com/BackAged/go-ddd-microservice/product"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Product", func() {
	Describe("NewStatus", func() {
		Context("with invalid status string", func() {
			It("should occur error", func() {
				_, err := product.NewStatus("invalid")
				Expect(err).To(HaveOccurred())
			})
		})

		Context("With valid status string", func() {
			It("should get a valid status", func() {
				sts, err := product.NewStatus("ACTIVE")
				Expect(err).ToNot(HaveOccurred())

				Expect(sts).To(Equal(product.StatusActive))
			})
		})
	})

	Describe("Status", func() {
		Context("invalid status IsValid", func() {
			It("should return error", func() {
				sts := product.Status("invalid")
				_, err := sts.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("valid status IsValid", func() {
			It("should not return error", func() {
				sts := product.Status("ACTIVE")
				_, err := sts.IsValid()
				Expect(err).ToNot(HaveOccurred())
			})
		})
	})

	Describe("Product", func() {
		Context("invalid product IsValid", func() {
			It("should return error", func() {
				prd := product.Product{}
				_, err := prd.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})

		Context("valid product IsValid", func() {
			It("should not return error", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: 2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				_, err := prd.IsValid()
				Expect(err).ToNot(HaveOccurred())
			})
		})

		Context("with price less than 0 product IsValid", func() {
			It("should not return error", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: 2,
					Price:           -2,
					Name:            "name",
					Slug:            "slug",
					Stock:           0,
					Status:          product.StatusActive,
					Version:         0,
				}
				_, err := prd.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})
		Context("with stock less than 0 product IsValid", func() {
			It("should not return error", func() {
				prd := product.Product{
					Image:           "image",
					DiscountedPrice: 2,
					Price:           2,
					Name:            "name",
					Slug:            "slug",
					Stock:           -1,
					Status:          product.StatusActive,
					Version:         0,
				}
				_, err := prd.IsValid()
				Expect(err).To(HaveOccurred())
			})
		})
	})
})
