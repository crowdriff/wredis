package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Connection", func() {
	Context("SELECT", func() {
		It("Should successfully select a different database", func() {
			Ω(safe.Select(1)).Should(Succeed())
			Ω(safe.Select(0)).Should(Succeed())
		})

		It("Should return an error given a bad index", func() {
			err := safe.Select(-1)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("db index must be 0 or positive"))
		})
	})
})
