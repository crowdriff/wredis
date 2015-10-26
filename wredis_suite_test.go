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

// BeforeSuite
var _ = BeforeSuite(func() {
	var err error
	safe, err = NewDefaultPool()
	Ω(err).Should(BeNil())
	unsafe, err = NewUnsafe("localhost", 6379, 0)
	Ω(err).Should(BeNil())
})

// AfterSuite
var _ = AfterSuite(func() {
	safe.Close()
	unsafe.Close()
})
