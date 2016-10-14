package wredis_test

import (
	"time"

	. "github.com/crowdriff/wredis"

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

	Context("GET/SET", func() {
		It("should fail to GET given an empty key", func() {
			_, err := safe.Get("")
			Ω(err).Should(MatchError(EmptyKeyErr))
		})

		It("should fail to SET given an empty key", func() {
			err := safe.Set("", "value")
			Ω(err).Should(MatchError(EmptyKeyErr))
		})

		It("should SET and then GET a key correctly", func() {
			err := safe.Set(testKey, testVal)
			Ω(err).Should(BeNil())

			val, err := safe.Get(testKey)
			Ω(err).Should(BeNil())
			Ω(val).Should(Equal(testVal))
		})
	})

	Context("MGET", func() {
		It("should return an error when a key is empty", func() {
			_, err := safe.MGet([]string{"1", "2", ""})
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("keys cannot be empty"))
		})

		It("should return all values for the provided keys", func() {
			// insert keys into redis
			Ω(safe.Set("1", "one")).Should(Succeed())
			Ω(safe.Set("2", "two")).Should(Succeed())
			// get values
			vals, err := safe.MGet([]string{"1", "2", "3"})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(vals).Should(HaveLen(3))
			Ω(vals[0]).Should(Equal("one"))
			Ω(vals[1]).Should(Equal("two"))
			Ω(vals[2]).Should(Equal(""))
		})
	})

	Context("INCR", func() {
		It("should return an error with an empty key provided", func() {
			_, err := safe.Incr("")
			Ω(err).Should(MatchError(EmptyKeyErr))
		})

		It("should create and increment a new key", func() {
			n, err := safe.Incr(testKey)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(n).Should(Equal(int64(1)))
		})

		It("should increment a key up to 10", func() {
			for i := 0; i < 10; i++ {
				n, err := safe.Incr(testKey)
				Ω(err).ShouldNot(HaveOccurred())
				Ω(n).Should(Equal(int64(i + 1)))
			}
		})
	})

	Context("SETEX", func() {
		It("should fail when given an empty key", func() {
			err := safe.SetEx("", testVal, 1)
			Ω(err).Should(MatchError(EmptyKeyErr))
		})

		It("should fail when given 0 or negative seconds", func() {
			err := safe.SetEx(testKey, testVal, 0)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("seconds must be a postive integer"))
		})

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
	})
})
