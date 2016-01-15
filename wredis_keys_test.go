package wredis_test

import (
	"fmt"
	"time"

	"github.com/garyburd/redigo/redis"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Keys", func() {

	testKey := "wredis::test::keys"
	testVal := "testvalue"

	Describe("DEL", func() {
		BeforeEach(func() {
			Ω(safe.Set(testKey, testVal)).Should(Succeed())
			Ω(safe.Exists(testKey)).Should(BeTrue())
		})

		AfterEach(func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})

		It("should delete a key successfully", func() {
			Ω(safe.Del(testKey)).Should(BeEquivalentTo(1))
			Ω(safe.Exists(testKey)).Should(BeFalse())
		})

		It("should fail if not given any keys", func() {
			_, err := safe.Del()
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("must provide at least 1 key"))

			_, err = safe.Del([]string{}...)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("must provide at least 1 key"))
		})

		It("should fail if any of the keys are empty", func() {
			_, err := safe.Del("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("keys cannot be empty strings"))

			_, err = safe.Del([]string{""}...)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("keys cannot be empty strings"))
		})
	})

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

		It("should fail if given an empty key", func() {
			_, err := safe.Exists("")
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("key cannot be empty"))
		})
	})

	Describe("EXPIRE", func() {
		AfterEach(func() {
			Ω(unsafe.FlushAll()).Should(Succeed())
		})

		It("should return an error if a blank key is provided", func() {
			_, err := safe.Expire("", 0)
			Ω(err).Should(HaveOccurred())
			Ω(err.Error()).Should(Equal("key cannot be an empty string"))
		})

		It("should return false when expire called on a non-existing key", func() {
			ok, err := safe.Expire(testKey, 10)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(ok).Should(BeFalse())
		})

		It("should set an expire value", func() {
			Ω(safe.Set(testKey, testVal)).ShouldNot(HaveOccurred())
			ok, err := safe.Expire(testKey, 10)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(ok).Should(BeTrue())
			n, err := safe.ExecInt64(func(conn redis.Conn) (int64, error) {
				return redis.Int64(conn.Do("TTL", testKey))
			})
			Ω(err).ShouldNot(HaveOccurred())
			Ω(n).Should(BeNumerically(">", int64(0)))
		})

		It("should set an expire value of 1 second", func() {
			Ω(safe.Set(testKey, testVal)).ShouldNot(HaveOccurred())
			ok, err := safe.Expire(testKey, 1)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(ok).Should(BeTrue())
			Eventually(func() (bool, error) {
				return safe.Exists(testKey)
			}, 2*time.Second, 100*time.Millisecond).Should(BeFalse())
		})

		It("should expire a key immediately", func() {
			Ω(safe.Set(testKey, testVal)).ShouldNot(HaveOccurred())
			ok, err := safe.Expire(testKey, 0)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(ok).Should(BeTrue())
			ok, err = safe.Exists(testKey)
			Ω(err).ShouldNot(HaveOccurred())
			Ω(ok).Should(BeFalse())
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
