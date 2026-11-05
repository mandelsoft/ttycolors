package plain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Format", func() {
	withColors()

	It("provides formatted strings for composed formats", func() {
		redItalic := plain.New(ttycolors.FmtRed, ttycolors.FmtItalic)
		Expect(redItalic("a", plain.Bold("b"))).To(Equal(
			ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_ITALIC) + "a" +
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "b" + ansi.Mode(ansi.MODE_RESET_BOLD) +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_ITALIC)))
	})

	It("is identical to the strings of package ttycolors", func() {
		Expect(plain.New(ttycolors.FmtBgWhite, ttycolors.FmtBold)("x")).
			To(Equal(ttycolors.New(ttycolors.FmtBgWhite, ttycolors.FmtBold).String("x").String()))
	})

	It("is compatible with the format functions", func() {
		funcs := []plain.Func{plain.Blue, plain.New(ttycolors.FmtBlue)}
		Expect(funcs[0]("x")).To(Equal(funcs[1]("x")))
	})

	It("can be nested", func() {
		red := plain.New(ttycolors.FmtRed)
		Expect(plain.Blue("a", red("b"), "c")).To(Equal(
			ansi.Mode(ansi.COLOR_FG_BLUE) + "a" +
				ansi.Mode(ansi.COLOR_FG_RED) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT) +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "c" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("copies the settings when created", func() {
		red := plain.New(ttycolors.FmtRed)
		ttycolors.NoColors = true
		Expect(red("x")).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.New(ttycolors.FmtRed)("x")).To(Equal("x"))
	})
})
