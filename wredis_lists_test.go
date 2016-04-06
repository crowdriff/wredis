package wredis_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("WredisLists", func() {

	var testList = "wredis::test::list"

	Context("LPush", func() {
		BeforeEach(func() {
			unsafe.Del(testList)
		})

		It("should return an error when no key provided", func() {
			_, err := safe.LPush("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("key cannot be empty"))
		})

		It("should return an error when no items provided", func() {
			_, err := safe.LPush(testList)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("must provide at least one item"))
		})

		It("should return an error when an item is empty", func() {
			_, err := safe.LPush(testList, "test", "")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("an item cannot be empty"))
		})

		It("should push an item to a new list", func() {
			n, err := safe.LPush(testList, "testing")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(n).Should(Equal(int64(1)))
		})

		It("should push multiple items to a new list", func() {
			n, err := safe.LPush(testList, "1", "2")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(n).Should(Equal(int64(2)))

			n, err = safe.LPush(testList, "3", "4")
			Ω(err).ShouldNot(HaveOccurred())
			Ω(n).Should(Equal(int64(4)))
		})
	})
})
