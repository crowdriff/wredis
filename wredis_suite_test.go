package wredis_test

import (
	. "github.com/crowdriff/wredis"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"testing"
)

// TestProcess is the root test process
func TestProcess(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Wredis Suite")
}

// safe and unsafe are global pointer to Wredis
// objects used for testing
var safe *Wredis
var unsafe *Wredis

// other variables for testing
var (
	// redis.Pool
	pool = MakeDefaultPool()

	// test keys
	key       = "wredis::test::set"
	newKey    = "wredis::test::set::new"
	gainedKey = "wredis::test::set::gained"
	lostKey   = "wredis::test::set::lost"

	// sets
	set1 = []string{"a", "b", "c"}
	set2 = []string{"a", "b", "d", "e"}
)

// BeforeSuite
var _ = BeforeSuite(func() {
	Ω(pool).ShouldNot(BeNil())
	Ω(pool.ActiveCount()).Should(Equal(0))

	safe = NewWredis(pool)
	unsafe = NewUnsafeWredis(pool)
})

// AfterSuite
var _ = AfterSuite(func() {
	// TODO Ω(pool.ActiveCount()).Should(Equal(0))
	pool.Close()
})
