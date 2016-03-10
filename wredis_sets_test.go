package wredis_test

import (
	. "github.com/crowdriff/wredis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sets", func() {

	var (
		testKey  = "wredis::test::sets::one"
		otherKey = "wredis::test::sets::two"

		testSet  = []string{"a", "b", "c"}
		otherSet = []string{"a", "b", "d", "e"}
	)

	BeforeEach(func() {
		Ω(unsafe.SAdd(testKey, testSet)).Should(BeEquivalentTo(3))
		Ω(unsafe.SAdd(otherKey, otherSet)).Should(BeEquivalentTo(4))
	})

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	Context("SADD", func() {
		It("should Add members to an existing set successfully", func() {
			Ω(safe.SAdd(testKey, otherSet)).Should(BeEquivalentTo(2))
		})

		It("should fail if an empty slice is passed to SAdd", func() {
			_, err := safe.SAdd(testKey, []string{})
			Ω(err).ShouldNot(BeNil())
		})
	})

	Context("SCARD", func() {
		It("should return the correct count of members in a set", func() {
			Ω(safe.SCard(testKey)).Should(BeEquivalentTo(3))
			Ω(safe.SCard(otherKey)).Should(BeEquivalentTo(4))
		})

		It("should fail given an empty key", func() {
			_, err := safe.SCard("")
			Ω(err).Should(MatchError(EmptyKeyErr))
		})
	})

	Context("SDIFFSTORE", func() {
		var diffKey = "wredis::test::sets::diff"

		It("should successfully store the difference of two sets correctly", func() {
			diff, err := safe.SDiffStore(diffKey, otherKey, testKey)
			Ω(err).Should(BeNil())
			Ω(diff).Should(BeEquivalentTo(2))
		})

		It("should successfully store the difference of two sets correctly", func() {
			diff, err := safe.SDiffStore(diffKey, testKey, otherKey)
			Ω(err).Should(BeNil())
			Ω(diff).Should(BeEquivalentTo(1))
		})

		It("should fail if empty dest is passed", func() {
			_, err := safe.SDiffStore("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("dest cannot be an empty string"))
		})

		It("should fail if no set keys are passed", func() {
			_, err := safe.SDiffStore(diffKey)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("SDiffStore requires at least 1 input set"))
		})

		It("should fail if not set keys are passed", func() {
			_, err := safe.SDiffStore(diffKey, "key", "", "otherKey")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("set keys cannot be empty strings"))
		})
	})

	Context("SMEMBERS", func() {
		It("should returns the members of a set successfully", func() {
			members, err := safe.SMembers(testKey)
			Ω(err).Should(BeNil())
			Ω(members).Should(HaveLen(3))
			Ω(members).Should(ConsistOf(testSet))
		})

		It("should return an error if key passed is empty", func() {
			_, err := safe.SMembers("")
			Ω(err).Should(MatchError(EmptyKeyErr))
		})
	})

	Context("SUNIONSTORE", func() {
		var unionKey = "wredis::test::sets::union"

		It("should successfully store the union of two sets correctly", func() {
			union, err := safe.SUnionStore(unionKey, otherKey, testKey)
			Ω(err).Should(BeNil())
			Ω(union).Should(BeEquivalentTo(5))
		})

		It("should successfully store the union of two sets correctly", func() {
			union, err := safe.SUnionStore(unionKey, testKey, otherKey)
			Ω(err).Should(BeNil())
			Ω(union).Should(BeEquivalentTo(5))
		})

		It("should fail if empty dest is passed", func() {
			_, err := safe.SUnionStore("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("dest cannot be an empty string"))
		})

		It("should fail if no set keys are passed", func() {
			_, err := safe.SUnionStore(unionKey)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("SUnionStore requires at least 1 input set"))
		})

		It("should fail if not set keys are passed", func() {
			_, err := safe.SUnionStore(unionKey, "key", "", "otherKey")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("set keys cannot be empty strings"))
		})
	})
})
