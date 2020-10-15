package product_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestProduct(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Product Suite")
}
