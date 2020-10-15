package error_test

import (
	"errors"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

var _ = Describe("Domain", func() {
	Describe("NewDomainError", func() {
		Context("when message & error", func() {
			It("should has error string of message and error", func() {
				de := cerror.NewDomainError("message", errors.New("error"))

				Expect(de.Error()).To(Equal("message: error"))
			})
		})

		Context("when only message", func() {
			It("should has error string of message only", func() {
				de := cerror.NewDomainError("message", nil)

				Expect(de.Error()).To(Equal("message"))
			})
		})
	})

	Describe("GetMessage", func() {
		Context("with message", func() {
			It("should return message", func() {
				de := cerror.NewDomainError("message", errors.New("error"))

				Expect(de.GetMessage()).To(Equal("message"))
			})
		})
	})

	Describe("GetError", func() {
		Context("with error", func() {
			It("should return error", func() {
				de := cerror.NewDomainError("message", errors.New("error"))

				Expect(de.GetError()).To(HaveOccurred())
			})
		})

		Context("without error", func() {
			It("should return nil", func() {
				de := cerror.NewDomainError("message", nil)

				Expect(de.GetError()).ToNot(HaveOccurred())
			})
		})
	})
})
