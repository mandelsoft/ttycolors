package plain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Colors", func() {
	withColors()

	It("provides formatted strings", func() {
		Expect(plain.Blue("b")).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.BgBrightRed("r")).To(Equal(ansi.Mode(ansi.COLOR_BG_BRIGHT_RED) + "r" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		Expect(plain.RGBColor(1, 2, 3, "x")).To(Equal(ansi.Mode(ansi.Code_FG_RGB(1, 2, 3)) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.BgIdColor(4, "x")).To(Equal(ansi.Mode(ansi.Code_BG_ID(4)) + "x" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
	})

	It("is identical to the strings of package ttycolors", func() {
		Expect(plain.BgRGBColor(1, 2, 3, "x")).To(Equal(ttycolors.BgRGBColor(1, 2, 3, "x").String()))
		Expect(plain.UnderlineIdColor(1, "x")).To(Equal(ttycolors.UnderlineIdColor(1, "x").String()))
	})

	It("restores the outer color after nested strings", func() {
		Expect(plain.Blue("a ", plain.Bold("bold"), " blue")).To(Equal(
			ansi.Mode(ansi.COLOR_FG_BLUE) + "a " +
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD) +
				" blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.Red("a", plain.Blue("b"), "c")).To(Equal(
			ansi.Mode(ansi.COLOR_FG_RED) + "a" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT) +
				ansi.Mode(ansi.COLOR_FG_RED) + "c" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("provides formatted strings for a context", func() {
		p := plain.For(ttycolors.NewContext(true))
		Expect(p.Red("r")).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "r" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(p.BgIdColor(4, "x")).To(Equal(ansi.Mode(ansi.Code_BG_ID(4)) + "x" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		Expect(plain.For(ttycolors.NewContext(false)).Red("r")).To(Equal("r"))
	})
})
