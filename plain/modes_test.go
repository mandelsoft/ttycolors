package plain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Modes", func() {
	withColors()

	It("provides formatted strings", func() {
		Expect(plain.Bold("b")).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "b" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		Expect(plain.Italic("i")).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "i" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
	})

	It("is identical to the strings of package ttycolors", func() {
		Expect(plain.StrikeThrough("x", "y")).To(Equal(ttycolors.StrikeThrough("x", "y").String()))
		Expect(plain.CurlyUnderline("x")).To(Equal(ttycolors.CurlyUnderline("x").String()))
	})

	It("supports nesting", func() {
		Expect(plain.Italic("a", plain.NonItalic("b"), "c")).To(Equal(
			ansi.Mode(ansi.MODE_ITALIC) + "a" + ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" +
				ansi.Mode(ansi.MODE_ITALIC) + "c" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
	})

	It("omits formatting without text", func() {
		Expect(plain.Bold()).To(Equal(""))
	})

	It("uses the features of the default context", func() {
		old := ttycolors.EnabledFeatures()
		defer ttycolors.SetFeatures(old)

		Expect(plain.DashedUnderline("x")).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		ttycolors.EnableFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(plain.DashedUnderline("x")).To(Equal(ansi.Mode(ansi.UNDERLINE_DASHED) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})
})

var _ = Describe("Modes without colors", func() {
	It("provides plain text", func() {
		old := ttycolors.NoColors
		defer func() { ttycolors.NoColors = old }()
		ttycolors.NoColors = true
		Expect(plain.Bold("a", plain.Italic("b"))).To(Equal("ab"))
	})
})
