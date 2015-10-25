package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {
	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	It("should rename a key successfully", func() {
		Ω(unsafe.SAdd(key, set1)).Should(BeEquivalentTo(3))
		Ω(unsafe.SAdd(newKey, set2)).Should(BeEquivalentTo(4))

		Ω(unsafe.Rename(newKey, key)).Should(BeNil())
		Ω(unsafe.SCard(key)).Should(BeEquivalentTo(4))
	})

	Context("SET/GET", func() {
		It("should SET and then GET a key correctly", func() {
			testKey := "wredis::test::strings"
			testVal := "testvalue"
			err := safe.Set(testKey, testVal)
			Ω(err).Should(BeNil())

			val, err := safe.Get(testKey)
			Ω(err).Should(BeNil())
			Ω(val).Should(Equal(testVal))
		})
	})

})
