package ginkgo_sample_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestGinkgoSample(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "GinkgoSample Suite")
}
