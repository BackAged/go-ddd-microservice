package error_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

var _ = Describe("Notfound", func() {
	Describe("NewNotFoundError", func() {
		Context("when with message and error", func() {
			It("should has error and message", func() {
				ne := cerror.NewNotFoundError("message", []string{"shahin"})

				Expect(ne.GetMessage()).To(Equal("message"))
				Expect(ne.GetErrors()).To(HaveLen(1))
			})
		})
	})

	Describe("Add", func() {
		Context("when add error", func() {
			It("should has error", func() {
				ne := cerror.NewNotFoundError("message", []string{})
				ne.Add("keyval")

				Expect(ne.GetErrors()).To(HaveLen(1))
			})
		})

		Context("when given message", func() {
			It("should has message", func() {
				msg := "message"
				ne := cerror.NewNotFoundError(msg, []string{})

				Expect(ne.GetMessage()).To(Equal(msg))
			})
		})
	})
})
