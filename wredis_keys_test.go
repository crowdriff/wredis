package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	testKey := "wredis::test::keys"
	testVal := "testvalue"

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	Context("RENAME", func() {
		It("should rename a key successfully", func() {
			newKey := "wredis::test::new"
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Rename(testKey, newKey)).Should(Succeed())

			// test rename is successful
			Ω(safe.Exists(testKey)).Should(BeFalse())
			Ω(safe.Exists(newKey)).Should(BeTrue())
			Ω(safe.Get(newKey)).Should(Equal(testVal))
		})

		It("should fail if any of the keys are empty strings", func() {
			Ω(safe.Rename("", "")).ShouldNot(Succeed())
			Ω(safe.Rename("", "test")).ShouldNot(Succeed())
			Ω(safe.Rename("test", "")).ShouldNot(Succeed())
		})
	})

	Context("EXISTS", func() {
		It("should return true if a key exists", func() {
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Exists(testKey)).Should(BeTrue())
		})

		It("should return false if a key does not exist", func() {
			Ω(safe.Exists(testKey)).Should(BeFalse())
		})
	})
})
