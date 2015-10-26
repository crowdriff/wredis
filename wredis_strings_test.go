package wredis_test

import (
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Strings", func() {

	var (
		testKey = "wredis::test::strings"
		testVal = "testvalue"
	)

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	It("should SET and then GET a key correctly", func() {
		err := safe.Set(testKey, testVal)
		Ω(err).Should(BeNil())

		val, err := safe.Get(testKey)
		Ω(err).Should(BeNil())
		Ω(val).Should(Equal(testVal))
	})

	Context("SETEX", func() {
		It("should set a key, which expires successfully", func() {
			err := safe.SetEx(testKey, testVal, 1)
			Ω(err).Should(BeNil())

			exists, err := safe.Exists(testKey)
			Ω(err).Should(BeNil())
			Ω(exists).Should(BeTrue())

			Eventually(func() (bool, error) {
				return safe.Exists(testKey)
			}, 2*time.Second, 100*time.Millisecond).Should(BeFalse())
		})

		It("should fail when given an empty key", func() {
			err := safe.SetEx("", testVal, 1)
			Ω(err).ShouldNot(BeNil())
			Ω(err.Error()).Should(Equal("key cannot be an empty string"))
		})

		It("should fail when given a small druation", func() {
			err := safe.SetExDuration(testKey, testVal, 500*time.Millisecond)
			Ω(err).ShouldNot(BeNil())
			Ω(err.Error()).Should(Equal("duration must be at least 1 second"))
		})
	})

})
