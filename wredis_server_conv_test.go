package wredis_test

import (
	"math"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("ServerConv", func() {

	AfterEach(func() {
		Ω(safe.Select(0)).Should(Succeed())
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	Context("SelectAndFlushDb", func() {
		It("should fail when using a safe client", func() {
			err := safe.SelectAndFlushDb(0)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("SelectAndFlushDb requires an Unsafe client."))
		})

		It("should fail to select an out of range db", func() {
			err := unsafe.SelectAndFlushDb(math.MaxUint64)
			Ω(err).Should(HaveOccurred())
		})

		It("should succesfully select and flush a db", func() {
			testKey := "wredis::test::server::conv"
			Ω(safe.Select(1)).Should(Succeed())
			Ω(safe.Set(testKey, "test value"))
			Ω(safe.Exists(testKey)).Should(BeTrue())
			Ω(safe.Select(0)).Should(Succeed())

			err := unsafe.SelectAndFlushDb(1)
			Ω(err).ShouldNot(HaveOccurred())
			exists, err := safe.Exists(testKey)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(exists).Should(BeFalse())
		})
	})
})
