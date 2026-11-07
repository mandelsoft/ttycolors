package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Text", func() {
	bold := ansi.Mode(ansi.MODE_BOLD)
	nobold := ansi.Mode(ansi.MODE_RESET_BOLD)
	italic := ansi.Mode(ansi.MODE_ITALIC)
	noitalic := ansi.Mode(ansi.MODE_RESET_ITALIC)
	resetAll := ansi.Mode(ansi.MODE_RESET_ALL)

	DescribeTable("SplitText",
		func(s string, n int, first, second string) {
			f, r := SplitText(s, n)
			Expect(f).To(Equal(first))
			Expect(r).To(Equal(second))
			Expect(ansi.StripAll(f + r)).To(Equal(ansi.StripAll(s)))
		},
		Entry("plain text", "abcd", 2, "ab", "cd"),
		Entry("multi byte runes", "äöüß", 1, "ä", "öüß"),
		Entry("closes and reopens a set group", bold+"abcd"+nobold, 2,
			bold+"ab"+nobold, bold+"cd"+nobold),
		Entry("keeps untouched groups untouched", "ab"+bold+"cd", 1,
			"a", "b"+bold+"cd"),
		Entry("reopens explicitly reset groups", noitalic+"abcd", 2,
			noitalic+"ab", noitalic+"cd"),
		Entry("reopens all groups reset", bold+"ab"+resetAll+"cd", 3,
			bold+"ab"+resetAll+"c", resetAll+"d"),
		Entry("combines reset and set groups", italic+bold+"ab"+noitalic+"cd", 3,
			italic+bold+"ab"+noitalic+"c"+nobold, ansi.Mode(ansi.MODE_BOLD, ansi.MODE_RESET_ITALIC)+"d"),
		Entry("keeps the last code of a group", ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_FG_BLUE)+"ab", 1,
			ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_FG_BLUE)+"a"+ansi.Mode(ansi.COLOR_FG_DEFAULT), ansi.Mode(ansi.COLOR_FG_BLUE)+"b"),
		Entry("keeps extended colors", ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(1, 2, 3))+"ab", 1,
			ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(1, 2, 3))+"a"+ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.COLOR_BG_DEFAULT),
			ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(1, 2, 3))+"b"),
		Entry("keeps sub parameter codes in separate sequences", ansi.Modes(ansi.UNDERLINE_CURLY, ansi.COLOR_FG_RED)+"ab", 1,
			ansi.Modes(ansi.UNDERLINE_CURLY, ansi.COLOR_FG_RED)+"a"+ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_UNDERLINE),
			ansi.Modes(ansi.COLOR_FG_RED, ansi.UNDERLINE_CURLY)+"b"),
		Entry("replaces resets at the split position", bold+"ab"+nobold+"cd", 2,
			bold+"ab"+nobold, nobold+"cd"),
		Entry("replaces sets at the split position", "ab"+bold+"cd"+nobold, 2,
			"ab", bold+"cd"+nobold),
		Entry("replaces mixed sequences at the split position", bold+"ab"+nobold+italic+"cd", 2,
			bold+"ab"+nobold, ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_ITALIC)+"cd"),
		Entry("replaces sets at the split position after a set group", bold+"ab"+italic+"cd", 2,
			bold+"ab"+nobold, ansi.Mode(ansi.MODE_BOLD, ansi.MODE_ITALIC)+"cd"),
		Entry("puts control sequences at the split position into the first part", "\x1b[2Kab\x1b]0;title\a"+bold+"cd", 2,
			"\x1b[2Kab\x1b]0;title\a", bold+"cd"),
		Entry("puts untracked codes at the split position into the second part", "ab\x1b[11mcd", 2,
			"ab", "\x1b[11mcd"),
		Entry("start", bold+"ab", 0, "", bold+"ab"),
		Entry("end", bold+"ab", 2, bold+"ab"+nobold, ""),
		Entry("beyond end", bold+"ab"+nobold, 5, bold+"ab"+nobold, ""),
		Entry("only sequences", bold, 0, bold+nobold, ""),
		Entry("empty", "", 0, "", ""),
	)

	It("panics for a negative index", func() {
		Expect(func() { SplitText("ab", -1) }).To(Panic())
	})

	Context("ScanFormat", func() {
		It("provides the codes of touched groups", func() {
			f := ScanFormat(bold + "a" + ansi.Mode(ansi.COLOR_FG_RED) + "b" + noitalic + "c")
			Expect(f).To(Equal(map[*Group]ansi.CodeSequence{
				Highlight: ansi.MODE_BOLD,
				FG:        ansi.COLOR_FG_RED,
				Italic:    ansi.MODE_RESET_ITALIC,
			}))
		})

		It("provides the reset codes for a reset of all modes", func() {
			f := ScanFormat(bold + "a" + resetAll + "b")
			Expect(f).To(HaveLen(len(groups)))
			for _, g := range groups {
				Expect(f[g]).To(Equal(g.reset))
			}
		})

		It("is empty for plain text", func() {
			Expect(ScanFormat("abc")).To(BeEmpty())
		})
	})
})
