package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	Context("FLUSHALL", func() {
		It("Should not be able to FlushAll with a safe client", func() {
			err := safe.FlushAll()
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("FlushAll requires an Unsafe client. See wredis.NewUnsafe"))
		})

		It("Should be able to FlushAll with an unsafe client", func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})
	})

	Context("FLUSHDB", func() {
		It("Should not be able to FlushDb with a safe client", func() {
			err := safe.FlushDb()
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("FlushDb requires an Unsafe client. See wredis.NewUnsafe"))
		})

		It("Should be able to FlushDb with an unsafe client", func() {
			Ω(unsafe.FlushDb()).Should(Succeed())
		})
	})
})
