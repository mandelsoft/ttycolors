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

	Context("scanned formats", func() {
		pre := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "ab" + ansi.Mode(ansi.MODE_RESET_ITALIC)

		It("is identical to the strings of package ttycolors", func() {
			Expect(plain.ScanFormat(pre)("x")).To(Equal(ttycolors.ScanFormat(pre).String("x").String()))
		})

		It("continues a string", func() {
			Expect(plain.ScanFormat(pre)("x")).To(Equal(
				ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD, ansi.MODE_RESET_ITALIC) + "x" +
					ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})

		It("keeps explicit resets in an outer format", func() {
			Expect(plain.Italic("a", plain.ScanFormat(ansi.Mode(ansi.MODE_RESET_ITALIC))("b"))).To(Equal(
				ansi.Mode(ansi.MODE_ITALIC) + "a" + ansi.Mode(ansi.MODE_RESET_ITALIC) + "b"))
		})
	})

	Context("for a context", func() {
		It("provides composed formats", func() {
			p := plain.For(ttycolors.NewContext(true))
			Expect(p.New(ttycolors.FmtRed)("x")).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
			Expect(plain.For(ttycolors.NewContext(false)).New(ttycolors.FmtRed)("x")).To(Equal("x"))
		})

		It("takes the settings when the function is created", func() {
			ctx := ttycolors.NewContext(true)
			f := plain.For(ctx).New(ttycolors.FmtRed)
			ctx.Enable(false)
			Expect(f("x")).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("provides scanned formats", func() {
			p := plain.For(ttycolors.NewContext(true))
			Expect(p.ScanFormat(ansi.Mode(ansi.COLOR_FG_RED))("x")).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
			Expect(plain.For(ttycolors.NewContext(false)).ScanFormat(ansi.Mode(ansi.COLOR_FG_RED))("x")).To(Equal("x"))
		})
	})
})
