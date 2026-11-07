package plain_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Text", func() {
	withColors()

	It("splits like package ttycolors", func() {
		s := plain.Blue("ab", plain.Bold("cd"), "ef")
		for n := 0; n <= ansi.CharLen(s)+1; n++ {
			f, r := plain.SplitAt(s, n)
			ef, er := ttycolors.SplitAt(s, n)
			Expect(f).To(Equal(ef))
			Expect(r).To(Equal(er))
		}
	})

	It("provides self-contained parts", func() {
		f, r := plain.SplitAt(plain.Blue("abcd"), 1)
		Expect(plain.Reverse(f) + r).To(Equal(ansi.Mode(ansi.MODE_REVERSE) +
			ansi.Mode(ansi.COLOR_FG_BLUE) + "a" + ansi.Mode(ansi.COLOR_FG_DEFAULT) + ansi.Mode(ansi.MODE_RESET_REVERSE) +
			ansi.Mode(ansi.COLOR_FG_BLUE) + "bcd" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("extracts substrings like package ttycolors", func() {
		s := plain.Blue("ab", plain.Bold("cd"), "ef")
		Expect(plain.SubString(s, 1, 4)).To(Equal(ttycolors.SubString(s, 1, 4)))
		Expect(plain.SubString(s, 1, 4)).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "b" +
			ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "cd" + ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
	})

	Context("StripIfDisabled", func() {
		s := ansi.Mode(ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)

		It("keeps strings if enabled", func() {
			Expect(plain.StripIfDisabled(s)).To(Equal(s))
		})

		It("removes output formats if disabled", func() {
			ttycolors.NoColors = true
			Expect(plain.StripIfDisabled(s)).To(Equal("bold"))
		})

		It("is required to remove the formats of arguments", func() {
			ttycolors.NoColors = true
			Expect(plain.Blue(s)).To(Equal(s))
			Expect(plain.Blue(plain.StripIfDisabled(s))).To(Equal("bold"))
		})
	})

	It("strips for a context", func() {
		s := ansi.Mode(ansi.MODE_BOLD) + "bold"
		Expect(plain.For(ttycolors.NewContext(true)).StripIfDisabled(s)).To(Equal(s))
		Expect(plain.For(ttycolors.NewContext(false)).StripIfDisabled(s)).To(Equal("bold"))
	})
})
