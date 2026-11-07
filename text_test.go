package ttycolors_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Text", func() {
	s := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"

	Context("SplitAt", func() {
		It("provides self-contained parts", func() {
			first, second := SplitAt(s, 2)
			Expect(first).To(Equal(ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hä" + ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
			Expect(second).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD) + "llo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"))
		})

		It("keeps explicit resets", func() {
			first, second := SplitAt(s, 5)
			Expect(first).To(Equal(ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
			Expect(second).To(Equal(ansi.Mode(ansi.MODE_RESET_ALL) + " world"))
		})

		It("is the base of splitting Strings", func() {
			ctx := NewContext(true)
			l, r := ctx.Sequence(s).SplitAt(7)
			first, second := SplitAt(s, 7)
			Expect(l.String()).To(Equal(first))
			Expect(r.String()).To(Equal(second))
		})

		It("keeps the formats of the parts in an outer format", func() {
			ctx := NewContext(true)
			first, second := SplitAt(ctx.Blue("abcd").String(), 1)
			Expect(ctx.Sequence(ctx.Reverse(first), second).String()).To(Equal(ansi.Mode(ansi.MODE_REVERSE) +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "a" + ansi.Mode(ansi.COLOR_FG_DEFAULT) + ansi.Mode(ansi.MODE_RESET_REVERSE) +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bcd" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("panics for a negative index", func() {
			Expect(func() { SplitAt(s, -1) }).To(Panic())
		})
	})

	Context("SubString", func() {
		It("provides self-contained substrings", func() {
			Expect(SubString(s, 1, 4)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD) + "äll" + ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
			Expect(SubString(s, 6, 8)).To(Equal(ansi.Mode(ansi.MODE_RESET_ALL) + "wo"))
		})

		It("keeps untouched formats untouched", func() {
			Expect(SubString("ab"+ansi.Mode(ansi.MODE_BOLD)+"cd", 0, 2)).To(Equal("ab"))
			Expect(SubString("ab"+ansi.Mode(ansi.MODE_BOLD)+"cd", 1, 3)).To(Equal("b" + ansi.Mode(ansi.MODE_BOLD) + "c" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("provides empty substrings", func() {
			for i := 0; i <= ansi.CharLen(s); i++ {
				Expect(SubString(s, i, i)).To(Equal(""))
			}
		})

		It("panics for invalid ranges", func() {
			Expect(func() { SubString(s, -1, 2) }).To(Panic())
			Expect(func() { SubString(s, 3, 1) }).To(Panic())
		})
	})

	Context("StripIfDisabled", func() {
		var old bool
		BeforeEach(func() {
			old = NoColors
		})
		AfterEach(func() {
			NoColors = old
		})

		It("uses the default context", func() {
			NoColors = false
			Expect(StripIfDisabled(s)).To(Equal(s))
			NoColors = true
			Expect(StripIfDisabled(s)).To(Equal("Hällo world"))
		})
	})
})
