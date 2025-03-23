package ttycolors_test

import (
	"github.com/mandelsoft/ttycolors/ansi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
)

var _ = Describe("Test Environment", func() {
	It("disabled", func() {
		ctx := ttycolors.NewContext(false) // provide a context always disabled

		s := ctx.Blue("a blue text with some ", ctx.Italic("italic"), " word")
		Expect(s.String()).To(Equal("a blue text with some italic word"))
	})

	It("enabled", func() {
		ctx := ttycolors.NewContext(true) // provide a context always enabled

		s := ctx.Blue("a blue text with some ", ctx.Italic("italic"), " word")
		Expect(s.String()).To(Equal(
			ansi.Mode(ansi.COLOR_FG_BLUE) + "a blue text with some " +
				ansi.Mode(ansi.MODE_ITALIC) + "italic" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + " word" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})
})
