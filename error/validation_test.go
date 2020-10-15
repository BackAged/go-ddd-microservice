package error_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	cerror "github.com/BackAged/go-ddd-microservice/error"
)

var _ = Describe("Validation", func() {
	Describe("Add", func() {
		Context("with add error", func() {
			It("should has error", func() {
				ve := cerror.NewValidationError()
				ve.Add("key", "val")

				Expect(ve.HasErrors()).To(BeTrue())
			})
		})

		Context("when empty", func() {
			It("should not has error", func() {
				ve := cerror.NewValidationError()

				Expect(ve.HasErrors()).To(BeFalse())
			})
		})
	})
})
