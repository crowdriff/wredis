package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Server", func() {
	It("Should not be able to FlushAll with a safe client", func() {
		err := safe.FlushAll()
		Ω(err).ShouldNot(BeNil())
		Ω(err.Error()).Should(Equal("FlushAll requires an Unsafe client. See wredis.NewUnsafe."))
	})

	It("Should be able to FlushAll with an unsafe client", func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})
})
