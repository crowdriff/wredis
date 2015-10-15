package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sets", func() {
	BeforeEach(func() {
		Ω(unsafe.SAdd(key, set1)).Should(BeEquivalentTo(3))
		Ω(unsafe.SAdd(newKey, set2)).Should(BeEquivalentTo(4))
	})

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	It("should add members to a set and return the number inserted", func() {
		Ω(unsafe.SAdd(key, set2)).Should(BeEquivalentTo(2))
	})

	It("should fail if an empty set is passed to sadd", func() {
		_, err := unsafe.SAdd(key, []string{})
		Ω(err).ShouldNot(BeNil())
	})

	It("should successfully store the difference of two sets", func() {
		gained, _ := unsafe.SDiffStore(gainedKey, newKey, key)
		Ω(gained).Should(BeEquivalentTo(2))

		lost, _ := unsafe.SDiffStore(lostKey, key, newKey)
		Ω(lost).Should(BeEquivalentTo(1))
	})

	It("should fail if empty parameters are passed to sdiffstore", func() {
		_, err := unsafe.SDiffStore("", "", "")
		Ω(err).ShouldNot(BeNil())
	})
})
