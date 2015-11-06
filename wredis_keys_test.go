package wredis_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	testKey := "wredis::test::keys"
	testVal := "testvalue"

	Describe("EXISTS", func() {
		AfterEach(func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})

		It("should return true if a key exists", func() {
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Exists(testKey)).Should(BeTrue())
		})

		It("should return false if a key does not exist", func() {
			Ω(safe.Exists(testKey)).Should(BeFalse())
		})
	})

	Describe("KEYS", func() {
		BeforeEach(func() {
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Set(fmt.Sprintf("%s::second", testKey), testVal)).Should(
				Succeed())
			Ω(safe.Set(fmt.Sprintf("%s::third", testKey), testVal)).Should(
				Succeed())
		})

		AfterEach(func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})

		It("should fetch all 3 keys with the general pattern", func() {
			pattern := "wredis::test::*"
			keys, err := safe.Keys(pattern)
			Ω(err).Should(BeNil())
			Ω(len(keys)).Should(Equal(3))
		})

		It("should fetch 2 keys with the specific pattern", func() {
			pattern := "wredis::test::keys::*"
			keys, err := safe.Keys(pattern)
			Ω(err).Should(BeNil())
			Ω(len(keys)).Should(Equal(2))
		})

		It("should be able to handle patterns that return no keys", func() {
			pattern := "redis::test::keys::*"
			keys, err := safe.Keys(pattern)
			Ω(err).Should(BeNil())
			Ω(len(keys)).Should(Equal(0))
		})
	})

	Describe("RENAME", func() {
		AfterEach(func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})

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

})
