package ansi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("ScanSequence", func() {
	scan := func(s string) (int, ansi.Covered) {
		return ansi.ScanSequence([]byte(s))
	}

	Context("no sequence", func() {
		DescribeTable("returns 0",
			func(s string) {
				l, c := scan(s)
				Expect(l).To(Equal(0))
				Expect(c).To(Equal(ansi.S_NONE))
			},
			Entry("empty", ""),
			Entry("text", "text"),
			Entry("escape only", "\x1b"),
			Entry("no CSI", "\x1bXm"),
			Entry("incomplete", "\x1b[1;3"),
			Entry("invalid byte", "\x1b[1\x01m"),
			Entry("not leading", "a\x1b[1m"),
			Entry("erase line", "\x1b[2K"),
			Entry("cursor up", "\x1b[A"),
			Entry("cursor position", "\x1b[10;20H"),
			Entry("private mode", "\x1b[?25l"),
			Entry("intermediate byte", "\x1b[1 m"),
		)
	})

	Context("SGR sequences", func() {
		DescribeTable("covered areas",
			func(s string, covered ansi.Covered) {
				l, c := scan(s + "rest")
				Expect(l).To(Equal(len(s)))
				Expect(c).To(Equal(covered))
			},
			Entry("empty reset", "\x1b[m", ansi.S_ALL),
			Entry("reset", "\x1b[0m", ansi.S_ALL),
			Entry("empty param is reset", "\x1b[;1m", ansi.S_ALL),
			Entry("bold", "\x1b[1m", ansi.S_HIGHLIGHT),
			Entry("dim", "\x1b[2m", ansi.S_HIGHLIGHT),
			Entry("reset bold", "\x1b[22m", ansi.S_HIGHLIGHT),
			Entry("italic", "\x1b[3m", ansi.S_ITALIC),
			Entry("reset italic", "\x1b[23m", ansi.S_ITALIC),
			Entry("underline", "\x1b[4m", ansi.S_UNDERLINE),
			Entry("double underline", "\x1b[21m", ansi.S_UNDERLINE),
			Entry("curly underline", "\x1b[4:3m", ansi.S_UNDERLINE),
			Entry("blink", "\x1b[5m", ansi.S_BLINK),
			Entry("rapid blink", "\x1b[6m", ansi.S_BLINK),
			Entry("reverse", "\x1b[7m", ansi.S_REVERSE),
			Entry("invisible", "\x1b[8m", ansi.S_INVIS),
			Entry("strike", "\x1b[9m", ansi.S_STRIKE),
			Entry("reset strike", "\x1b[29m", ansi.S_STRIKE),
			Entry("overline", "\x1b[53m", ansi.S_OVERLINE),
			Entry("reset overline", "\x1b[55m", ansi.S_OVERLINE),
			Entry("fg color", "\x1b[31m", ansi.S_FG_COLOR),
			Entry("fg default", "\x1b[39m", ansi.S_FG_COLOR),
			Entry("bg color", "\x1b[44m", ansi.S_BG_COLOR),
			Entry("bg default", "\x1b[49m", ansi.S_BG_COLOR),
			Entry("bright fg", "\x1b[95m", ansi.S_FG_COLOR),
			Entry("bright bg", "\x1b[107m", ansi.S_BG_COLOR),
			Entry("ul default", "\x1b[59m", ansi.S_UL_COLOR),
			Entry("fg and bg", "\x1b[31;44m", ansi.S_FG_COLOR|ansi.S_BG_COLOR),
			Entry("unknown code", "\x1b[60m", ansi.S_NONE),
			Entry("combined", "\x1b[1;3;31m", ansi.S_HIGHLIGHT|ansi.S_ITALIC|ansi.S_FG_COLOR),
		)

		DescribeTable("extended colors skip their arguments",
			func(s string, covered ansi.Covered) {
				l, c := scan(s)
				Expect(l).To(Equal(len(s)))
				Expect(c).To(Equal(covered))
			},
			Entry("fg id", "\x1b[38;5;1m", ansi.S_FG_COLOR),
			Entry("bg id", "\x1b[48;5;9m", ansi.S_BG_COLOR),
			Entry("fg rgb", "\x1b[38;2;1;3;4m", ansi.S_FG_COLOR),
			Entry("bg rgb", "\x1b[48;2;7;8;9m", ansi.S_BG_COLOR),
			Entry("fg rgb colon", "\x1b[38:2::1:3:4m", ansi.S_FG_COLOR),
			Entry("bg id colon", "\x1b[48:5:1m", ansi.S_BG_COLOR),
			Entry("underline id", "\x1b[58;5;4m", ansi.S_UL_COLOR),
			Entry("underline rgb", "\x1b[58;2;4;5;6m", ansi.S_UL_COLOR),
		)

		DescribeTable("complex sequences",
			func(s string, covered ansi.Covered) {
				l, c := scan(s + "rest")
				Expect(l).To(Equal(len(s)))
				Expect(c).To(Equal(covered))
			},
			Entry("fg id and bg rgb", "\x1b[38;5;1;48;2;1;2;3m", ansi.S_FG_COLOR|ansi.S_BG_COLOR),
			Entry("bg rgb and ul id", "\x1b[48;2;1;2;3;58;5;7m", ansi.S_BG_COLOR|ansi.S_UL_COLOR),
			Entry("extended colors and modes", "\x1b[1;38;5;9;3;48;2;4;5;6;7m",
				ansi.S_HIGHLIGHT|ansi.S_FG_COLOR|ansi.S_ITALIC|ansi.S_BG_COLOR|ansi.S_REVERSE),
			Entry("colon color followed by code", "\x1b[38:5:1;4m", ansi.S_FG_COLOR|ansi.S_UNDERLINE),
			Entry("code followed by colon color", "\x1b[4;48:2::1:2:3m", ansi.S_UNDERLINE|ansi.S_BG_COLOR),
			Entry("colon and semicolon colors", "\x1b[38:2::1:2:3;48;5;2m", ansi.S_FG_COLOR|ansi.S_BG_COLOR),
			Entry("reset in the middle", "\x1b[1;0;31m", ansi.S_ALL),
			Entry("reset at the end", "\x1b[31;m", ansi.S_ALL),
			Entry("truncated id color", "\x1b[38;5m", ansi.S_FG_COLOR),
			Entry("truncated rgb color", "\x1b[48;2;1m", ansi.S_BG_COLOR),
			Entry("extended color without kind", "\x1b[38m", ansi.S_FG_COLOR),
			Entry("unknown color kind", "\x1b[38;7;1m", ansi.S_FG_COLOR),
			Entry("unknown color kind after code", "\x1b[3;48;4;9m", ansi.S_ITALIC|ansi.S_BG_COLOR),
			Entry("empty color kind", "\x1b[38;;1m", ansi.S_FG_COLOR),
			Entry("code out of byte range", "\x1b[300m", ansi.S_NONE),
			Entry("code out of byte range with code", "\x1b[300;1m", ansi.S_HIGHLIGHT),
			Entry("overflowing code", "\x1b[99999999999999999999999;3m", ansi.S_ITALIC),
		)

		It("maps 21 to underline (limitation)", func() {
			_, c := scan("\x1b[21m")
			Expect(c).To(Equal(ansi.S_UNDERLINE))
		})

		It("handles the constants for rapid blink and double underline", func() {
			_, c := scan(ansi.Mode(ansi.MODE_RAPID_BLINK, ansi.MODE_DOUBLE_UNDERLINE))
			Expect(c).To(Equal(ansi.S_BLINK | ansi.S_UNDERLINE))
			Expect(ansi.Mode(ansi.MODE_RAPID_BLINK, ansi.MODE_DOUBLE_UNDERLINE)).To(Equal("\x1b[6;21m"))
		})

		It("handles codes following an extended color", func() {
			s := "\x1b[38;2;1;3;4;1m"
			l, c := scan(s + "rest")
			Expect(l).To(Equal(len(s)))
			Expect(c).To(Equal(ansi.S_FG_COLOR | ansi.S_HIGHLIGHT))
		})

		It("handles sequences generated by Mode", func() {
			s := ansi.Mode(ansi.Code_BG_RGB(4, 5, 7), ansi.MODE_ITALIC)
			l, c := scan(s)
			Expect(l).To(Equal(len(s)))
			Expect(c).To(Equal(ansi.S_BG_COLOR | ansi.S_ITALIC))
		})

		It("handles underline color constants", func() {
			s := ansi.Mode(ansi.NewCodeSequence(ansi.COLOR_UL, ansi.COLOR_ID, 4))
			l, c := scan(s)
			Expect(l).To(Equal(len(s)))
			Expect(c).To(Equal(ansi.S_UL_COLOR))

			s = ansi.Mode(ansi.COLOR_UL_DEFAULT)
			l, c = scan(s)
			Expect(l).To(Equal(len(s)))
			Expect(c).To(Equal(ansi.S_UL_COLOR))
		})
	})

	Context("Covered names", func() {
		It("provides names", func() {
			Expect(ansi.S_NONE.String()).To(Equal("S_NONE"))
			Expect(ansi.S_ALL.String()).To(Equal("S_ALL"))
			Expect(ansi.S_ITALIC.String()).To(Equal("S_ITALIC"))
			Expect((ansi.S_FG_COLOR | ansi.S_HIGHLIGHT).String()).To(Equal("S_HIGHLIGHT|S_FG_COLOR"))
			Expect(ansi.S_COLOR.String()).To(Equal("S_FG_COLOR|S_BG_COLOR|S_UL_COLOR"))
			Expect((ansi.S_STRIKE | ansi.Covered(1<<20)).String()).To(Equal("S_STRIKE|0x100000"))
		})
	})

	Context("Covered", func() {
		It("checks areas", func() {
			c := ansi.S_FG_COLOR | ansi.S_ITALIC
			Expect(c.Has(ansi.S_FG_COLOR)).To(BeTrue())
			Expect(c.Has(ansi.S_FG_COLOR | ansi.S_ITALIC)).To(BeTrue())
			Expect(c.Has(ansi.S_COLOR)).To(BeFalse())
			Expect(ansi.S_COLOR.Has(ansi.S_FG_COLOR | ansi.S_BG_COLOR | ansi.S_UL_COLOR)).To(BeTrue())
			Expect(c.Has(ansi.S_BLINK)).To(BeFalse())
			Expect(ansi.S_ALL.Has(c)).To(BeTrue())
		})
	})
})

var _ = Describe("Scan", func() {
	DescribeTable("covered areas of a string",
		func(s string, covered ansi.Covered) {
			Expect(ansi.Scan(s)).To(Equal(covered))
		},
		Entry("empty", "", ansi.S_NONE),
		Entry("plain text", "text", ansi.S_NONE),
		Entry("single sequence", "\x1b[1mbold", ansi.S_HIGHLIGHT),
		Entry("multiple sequences", "a\x1b[1mb\x1b[31mc\x1b[44md", ansi.S_HIGHLIGHT|ansi.S_FG_COLOR|ansi.S_BG_COLOR),
		Entry("adjacent sequences", "\x1b[3m\x1b[38;5;1mx", ansi.S_ITALIC|ansi.S_FG_COLOR),
		Entry("sequence at the end", "text\x1b[4m", ansi.S_UNDERLINE),
		Entry("reset", "a\x1b[1mb\x1b[0m", ansi.S_ALL),
		Entry("multi byte runes", "ä\x1b[9möü\x1b[29mß", ansi.S_STRIKE),
		Entry("other control sequences", "\x1b[2Ka\x1b[10;20H", ansi.S_NONE),
		Entry("other control sequence before format", "\x1b[2K\x1b[7m", ansi.S_REVERSE),
		Entry("lone escape", "a\x1bb\x1b[5m", ansi.S_BLINK),
		Entry("escape before sequence", "\x1b\x1b[8m", ansi.S_INVIS),
		Entry("incomplete sequence at the end", "\x1b[1mtext\x1b[31", ansi.S_HIGHLIGHT),
	)

	It("handles sequences generated by Mode", func() {
		s := ansi.Mode(ansi.Code_FG_RGB(9, 7, 4)) + "text" + ansi.Mode(ansi.COLOR_BG_ID, ansi.Param(3)) + "more"
		Expect(ansi.Scan(s)).To(Equal(ansi.S_FG_COLOR | ansi.S_BG_COLOR))
	})
})

var _ = Describe("Effect", func() {
	effect := func(s string) ansi.Effect {
		l, e := ansi.ScanSequenceEffect([]byte(s))
		Expect(l).To(Equal(len(s)))
		return e
	}

	It("provides names", func() {
		Expect(ansi.Effect{Set: ansi.S_ITALIC}.String()).To(Equal("{Set: S_ITALIC, Reset: S_NONE}"))
	})

	It("combines effects in order", func() {
		e := ansi.Effect{Set: ansi.S_HIGHLIGHT, Reset: ansi.S_FG_COLOR}
		Expect(e.Covered()).To(Equal(ansi.S_HIGHLIGHT | ansi.S_FG_COLOR))
		Expect(e.Then(ansi.Effect{Reset: ansi.S_HIGHLIGHT, Set: ansi.S_FG_COLOR})).
			To(Equal(ansi.Effect{Set: ansi.S_FG_COLOR, Reset: ansi.S_HIGHLIGHT}))
	})

	DescribeTable("sequence effects",
		func(s string, set, reset ansi.Covered) {
			Expect(effect(s)).To(Equal(ansi.Effect{Set: set, Reset: reset}))
		},
		Entry("set", "\x1b[1;31m", ansi.S_HIGHLIGHT|ansi.S_FG_COLOR, ansi.S_NONE),
		Entry("reset", "\x1b[22;39m", ansi.S_NONE, ansi.S_HIGHLIGHT|ansi.S_FG_COLOR),
		Entry("reset all", "\x1b[0m", ansi.S_NONE, ansi.S_ALL),
		Entry("empty reset all", "\x1b[m", ansi.S_NONE, ansi.S_ALL),
		Entry("set after reset all", "\x1b[0;1m", ansi.S_HIGHLIGHT, ansi.S_ALL&^ansi.S_HIGHLIGHT),
		Entry("reset all after set", "\x1b[1;0m", ansi.S_NONE, ansi.S_ALL),
		Entry("bold resets highlight first", "\x1b[22;1m", ansi.S_HIGHLIGHT, ansi.S_NONE),
		Entry("extended color", "\x1b[38;5;0;48;2;0;0;0m", ansi.S_FG_COLOR|ansi.S_BG_COLOR, ansi.S_NONE),
		Entry("underline style and color", "\x1b[4:3;58:5:1m", ansi.S_UNDERLINE|ansi.S_UL_COLOR, ansi.S_NONE),
		Entry("underline color reset", "\x1b[59m", ansi.S_NONE, ansi.S_UL_COLOR),
		Entry("overline", "\x1b[53m", ansi.S_OVERLINE, ansi.S_NONE),
		Entry("overline reset", "\x1b[55m", ansi.S_NONE, ansi.S_OVERLINE),
	)

	It("ignores invalid sequences", func() {
		l, e := ansi.ScanSequenceEffect([]byte("\x1b[2K"))
		Expect(l).To(Equal(0))
		Expect(e).To(Equal(ansi.Effect{}))
	})

	DescribeTable("string effects",
		func(s string, set, reset ansi.Covered) {
			Expect(ansi.ScanEffect(s)).To(Equal(ansi.Effect{Set: set, Reset: reset}))
		},
		Entry("plain", "text", ansi.S_NONE, ansi.S_NONE),
		Entry("set and reset", "\x1b[1mbold\x1b[22m", ansi.S_NONE, ansi.S_HIGHLIGHT),
		Entry("left set", "\x1b[31mred\x1b[1m", ansi.S_FG_COLOR|ansi.S_HIGHLIGHT, ansi.S_NONE),
		Entry("reset all in between", "\x1b[31ma\x1b[0mb\x1b[3mc", ansi.S_ITALIC, ansi.S_ALL&^ansi.S_ITALIC),
		Entry("other control sequences", "\x1b[2Ka\x1b[1m", ansi.S_HIGHLIGHT, ansi.S_NONE),
	)
})
