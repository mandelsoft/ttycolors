package plain_test

import (
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
)

func TestConfig(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Plain Test Suite")
}

// withColors enables colors for the default context during a test.
func withColors() {
	var old bool
	BeforeEach(func() {
		old = ttycolors.NoColors
		ttycolors.NoColors = false
	})
	AfterEach(func() {
		ttycolors.NoColors = old
	})
}
