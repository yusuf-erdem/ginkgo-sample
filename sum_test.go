package ginkgo_sample_test

import (
	. "ginkgo-sample"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sum", func() {
	var (
		p, q, m, n, sum1, sum2 int
	)
	BeforeEach(func() {
		p, q, sum1 = 5, 6, 11
		// Putting wrong value of sum2 intentionally
		m, n, sum2 = 8, 7, 16
	})
	Context("Addition of two digits", func() {
		It("should return sum of the two digits", func() {
			addition_of_two_digits := Sum(p, q)
			Expect(addition_of_two_digits).Should(Equal(sum1))
		})
		It("should not return the sum provided", func() {
			addition_of_two_digits := Sum(m, n)
			Expect(addition_of_two_digits).ShouldNot(Equal(sum2))
		})
	})
})
