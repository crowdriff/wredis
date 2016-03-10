package wredis_test

import (
	. "github.com/crowdriff/wredis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Wredis", func() {

	var pool *Wredis
	var err error

	AfterEach(func() {
		if pool != nil {
			Ω(pool.Close()).Should(Succeed())
		}
		pool = nil
	})

	It("Should create a new default pool", func() {
		pool, err = NewDefaultPool()
		Ω(err).ShouldNot(HaveOccurred())
		Ω(pool).ShouldNot(BeNil())
	})

	It("Should fail to create a new pool given an empty host", func() {
		_, err = NewPool("", 6379, 0)
		Ω(err).Should(HaveOccurred())
		Ω(err.Error()).Should(Equal("host cannot be empty"))
	})

	It("Should fail to create a new pool given an invalid port", func() {
		_, err = NewPool("localhost", 0, 0)
		Ω(err).Should(HaveOccurred())
		Ω(err.Error()).Should(Equal("port cannot be 0"))
	})

	It("Should create an unsafe pool successfully", func() {
		pool, err = NewUnsafe("localhost", 6379, 0)
		Ω(err).ShouldNot(HaveOccurred())
		Ω(pool).ShouldNot(BeNil())
		Ω(pool.FlushAll()).Should(Succeed())
	})

	It("should fail to create an unsafe pool given an empty host", func() {
		_, err := NewUnsafe("", 6379, 0)
		Ω(err).Should(HaveOccurred())
		Ω(err.Error()).Should(Equal("host cannot be empty"))
	})

	It("should fail to create an unsafe pool given an invalid port", func() {
		_, err := NewUnsafe("localhost", 0, 0)
		Ω(err).Should(HaveOccurred())
		Ω(err.Error()).Should(Equal("port cannot be 0"))
	})
})
