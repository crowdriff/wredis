package wredis_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("StringsConv", func() {

	var (
		testKey = "wredis::test::strings"
		testVal = "testvalue"
	)

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	Context("SetExDuration", func() {
		It("should fail when given a small druation", func() {
			err := safe.SetExDuration(testKey, testVal, 500*time.Millisecond)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("duration must be at least 1 second"))
		})

		It("should succeed in setting a key with an expiry", func() {
			err := safe.SetExDuration(testKey, testVal, time.Second)
			Ω(err).ShouldNot(HaveOccurred())

			Eventually(func() (bool, error) {
				return safe.Exists(testKey)
			}, 2*time.Second, 100*time.Millisecond).Should(BeFalse())
		})
	})
})
