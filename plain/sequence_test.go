package plain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Sequence", func() {
	withColors()

	It("provides a formatted string", func() {
		Expect(plain.Sequence("a", plain.Red("b"), 1)).To(Equal("a" + ansi.Mode(ansi.COLOR_FG_RED) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT) + "1"))
	})

	It("resets formats left set by embedded strings", func() {
		Expect(plain.Sequence("\x1b[1mbold")).To(Equal("\x1b[1mbold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
	})

	It("provides a formatted string for a context", func() {
		p := plain.For(ttycolors.NewContext(true))
		Expect(p.Sequence("a", p.Red("b"))).To(Equal("a" + ansi.Mode(ansi.COLOR_FG_RED) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.For(ttycolors.NewContext(false)).Sequence("a", 1)).To(Equal("a1"))
	})
})
