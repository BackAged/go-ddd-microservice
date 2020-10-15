package error_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestError(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Error Suite")
}
