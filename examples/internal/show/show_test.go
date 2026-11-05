package show_test

import (
	"bytes"
	"testing"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/examples/internal/show"
)

func TestShow(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Show Test Suite")
}

var _ = Describe("Printer", func() {
	It("quotes byte slice types", func() {
		type bytes []byte
		Expect(show.Format(bytes("\x1b"))).To(Equal(`"\x1b"`))
		Expect(show.Format(1)).To(Equal("1"))
	})

	It("aligns calls and results", func() {
		var buf bytes.Buffer
		p := show.NewFor(&buf)
		p.Title("title")
		p.Call("a()", "x\x1b[1m")
		p.Call("longer()", show.R2(1, []byte("b"))...)
		p.Code("code")
		p.Flush()
		Expect(buf.String()).To(Equal("\n// title\n" +
			"a()      // \"x\\x1b[1m\"\n" +
			"longer() // 1, \"b\"\n" +
			"code\n"))
	})
})
