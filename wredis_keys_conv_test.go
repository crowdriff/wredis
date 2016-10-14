package wredis_test

import (
	"fmt"

	. "github.com/crowdriff/wredis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("KeysConv", func() {

	testKey := "wredis::test::keys::conv"
	testVal := "testvalue"

	AfterEach(func() {
		Ω(unsafe.FlushAll()).Should(Succeed())
	})

	Context("DelWithPattern", func() {

		testPattern := "wredis::test::keys::conv*"

		BeforeEach(func() {
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Set(fmt.Sprintf("%s::second", testKey), testVal)).Should(
				Succeed())
			Ω(safe.Set(fmt.Sprintf("%s::third", testKey), testVal)).Should(
				Succeed())
		})

		It("should return an error when using a safe client", func() {
			_, err := safe.DelWithPattern("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(ContainSubstring("DelWithPattern requires an Unsafe client."))
		})

		It("should fail given an empty pattern", func() {
			_, err := unsafe.DelWithPattern("")
			Ω(err).Should(MatchError(EmptyPatternErr))
		})

		It("should return an error if no keys matched the pattern", func() {
			badPattern := "bad::pattern"
			expectedErr := "no keys found with pattern: " + badPattern

			del, err := unsafe.DelWithPattern(badPattern)
			Ω(del).Should(Equal(int64(-1)))
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal(expectedErr))
		})

		It("should delete all keys corresponding to the pattern", func() {
			del, err := unsafe.DelWithPattern(testPattern)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(del).Should(Equal(int64(3)))
		})
	})
})
