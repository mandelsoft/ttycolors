package ttycolors_test

import (
	. "github.com/mandelsoft/ttycolors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("ColorString Test Environment", func() {
	Context("disabled", func() {
		It("mode", func() {
			s := Bold("bold")
			s.Enable(false)
			Expect(s.String()).To(Equal("bold"))
		})

		It("standard color", func() {
			s := Blue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg color", func() {
			s := BgBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bright color", func() {
			s := BrightBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg bright color", func() {
			s := BgBrightBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb color", func() {
			s := RGBColor(1, 255, 2, "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb bg color", func() {
			s := BgRGBColor(1, 255, 2, "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("id color", func() {
			s := IdColor(1, "id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("id bg color", func() {
			s := BgIdColor(1, "id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("complex color", func() {
			s := Blue("blue", Green("green"), Red("red"), "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue" +
				"green" +
				"red" +
				"blue"))
		})

		It("mixed", func() {
			s := Blue("blue", Bold(Green("bold green"), "bold blue"), Red("red"), "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue" +
				"bold green" +
				"bold blue" +
				"red" +
				"blue"))
		})
	})

	Context("enabled", func() {
		It("mode", func() {
			s := Bold("bold")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("standard color", func() {
			s := Blue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg color", func() {
			s := BgBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("standard bright color", func() {
			s := BrightBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg bright color", func() {
			s := BgBrightBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("rgb color", func() {
			s := RGBColor(1, 255, 2, "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RGB, ansi.Param(1), ansi.Param(255), ansi.Param(2)) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("rgb bg color", func() {
			s := BgRGBColor(1, 255, 2, "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RGB, ansi.Param(1), ansi.Param(255), ansi.Param(2)) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("id color", func() {
			s := IdColor(1, "id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_ID, ansi.Param(1)) + "id" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("id bg color", func() {
			s := BgIdColor(1, "id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_ID, ansi.Param(1)) + "id" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("complex color", func() {
			s := Blue("blue", Green("green"), Red("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("complex bg color", func() {
			s := BgBlue("blue", BgGreen("green"), BgRed("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_BG_RED) + "red" +
				ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("mixed color", func() {
			s := Blue("blue", BgWhite("white back"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("mixed bg", func() {
			s := Blue("blue", Bold(Green("bold green"), "bold blue", BgWhite("white back"), BgYellow("yellow back")), Red("red"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_YELLOW) + "yellow back" +
				ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_DEFAULT, ansi.MODE_RESET_BOLD) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("mixed bg", func() {
			s := Blue("blue", Bold(Green("bold green", Yellow("bold yellow"), "bold green"), "bold blue", BgWhite("white back"), BgYellow("yellow back")), Red("red"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_YELLOW) + "bold yellow" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_YELLOW) + "yellow back" +
				ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_DEFAULT, ansi.MODE_RESET_BOLD) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("resets bg nested in highlight", func() {
			s := Bold(BgYellow("yellow back"), "bold")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_YELLOW, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "yellow back" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT) + "bold" +
				ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("renders outer highlight once", func() {
			s := Bold(Green("bold green"))
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_GREEN, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})

		DescribeTable("non modes",
			func(mode, non func(...any) String, on, off ansi.CodeSequence) {
				s := mode("a", non("b"), "c")
				s.Enable()
				Expect(s.String()).To(Equal(ansi.Mode(on) + "a" + ansi.Mode(off) + "b" + ansi.Mode(on) + "c" + ansi.Mode(off)))
			},
			Entry("italic", Italic, NonItalic, ansi.MODE_ITALIC, ansi.MODE_RESET_ITALIC),
			Entry("underline", Underline, NonUnderline, ansi.MODE_UNDERLINE, ansi.MODE_RESET_UNDERLINE),
			Entry("blinking", Blinking, NonBlinking, ansi.MODE_BLINK, ansi.MODE_RESET_BLINK),
			Entry("reverse", Reverse, NonReverse, ansi.MODE_REVERSE, ansi.MODE_RESET_REVERSE),
			Entry("invisible", Invisible, NonInvisible, ansi.MODE_INVIS, ansi.MODE_RESET_INVIS),
			Entry("strike through", StrikeThrough, NonStrikeThrough, ansi.MODE_STRIKE, ansi.MODE_RESET_STRIKE),
			Entry("overline", Overline, NonOverline, ansi.MODE_OVERLINE, ansi.MODE_RESET_OVERLINE),
		)

		It("nested modes", func() {
			s := Italic("a", Underline("b", NonItalic("c")), "d")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "a" +
				ansi.Mode(ansi.MODE_UNDERLINE) + "b" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + "c" +
				ansi.Mode(ansi.MODE_ITALIC, ansi.MODE_RESET_UNDERLINE) + "d" +
				ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("non mode at top level", func() {
			s := NonItalic("a")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_ITALIC) + "a"))
		})

		It("non highlighted at top level", func() {
			s := NonHighlighted("a")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD) + "a"))
		})

		DescribeTable("omits unchanged codes",
			func(s String, expected string) {
				s.Enable()
				Expect(s.String()).To(Equal(expected))
			},
			Entry("same mode nested", Italic("a", Italic("b"), "c"),
				ansi.Mode(ansi.MODE_ITALIC)+"abc"+ansi.Mode(ansi.MODE_RESET_ITALIC)),
			Entry("same color nested", Blue("a", Blue("b"), "c"),
				ansi.Mode(ansi.COLOR_FG_BLUE)+"abc"+ansi.Mode(ansi.COLOR_FG_DEFAULT)),
			Entry("same mode in other format", Bold("a", Italic(Bold("b"))),
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD)+"a"+ansi.Mode(ansi.MODE_ITALIC)+"b"+ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_RESET_ITALIC)),
			Entry("same color in composed format", New(FmtRed, FmtBold).String("a", Red("b")),
				ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD)+"ab"+ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)),
			Entry("non mode in non mode", NonItalic("a", NonItalic("b")),
				ansi.Mode(ansi.MODE_RESET_ITALIC)+"ab"),
			Entry("changed color", Red("a", Blue("b"), "c"),
				ansi.Mode(ansi.COLOR_FG_RED)+"a"+ansi.Mode(ansi.COLOR_FG_BLUE)+"b"+ansi.Mode(ansi.COLOR_FG_RED)+"c"+ansi.Mode(ansi.COLOR_FG_DEFAULT)),
		)

		DescribeTable("empty formats are no-ops",
			func(s String, expected string) {
				s.Enable()
				Expect(s.String()).To(Equal(expected))
			},
			Entry("without content", Bold(), ""),
			Entry("with empty string", Italic(""), ""),
			Entry("nested", Blue(Bold(Italic(), "")), ""),
			Entry("inside text", Blue("a", Bold(), "b"), ansi.Mode(ansi.COLOR_FG_BLUE)+"ab"+ansi.Mode(ansi.COLOR_FG_DEFAULT)),
			Entry("composed", New(FmtRed, FmtItalic).String(), ""),
		)

		DescribeTable("sequences with formats nested in the same format",
			nestedSequenceTable(func(outer, first, second func(...any) String, prerendered bool, expected string) {
				ctx := NewContext(true)
				var b any = "bbb"
				if second != nil {
					b = second("bbb")
				}
				var i any = ctx.Sequence(first("aaaa"), b)
				if prerendered {
					i = i.(String).String()
				}
				s := outer(i, "ccc")
				s.Enable()
				Expect(s.String()).To(Equal(expected))
			})...,
		)

		DescribeTable("uncolored",
			func(s String, expected string) {
				s.Enable()
				Expect(s.String()).To(Equal(expected))
			},
			Entry("foreground", Blue("a", UncoloredFG("b"), "c"),
				ansi.Mode(ansi.COLOR_FG_BLUE)+"a"+ansi.Mode(ansi.COLOR_FG_DEFAULT)+"b"+ansi.Mode(ansi.COLOR_FG_BLUE)+"c"+ansi.Mode(ansi.COLOR_FG_DEFAULT)),
			Entry("bright and id foreground", BrightRed("a", UncoloredFG("b", IdColor(5, "c"))),
				ansi.Mode(ansi.COLOR_FG_BRIGHT_RED)+"a"+ansi.Mode(ansi.COLOR_FG_DEFAULT)+"b"+ansi.Mode(ansi.Code_FG_ID(5))+"c"+ansi.Mode(ansi.COLOR_FG_DEFAULT)),
			Entry("background", BgYellow("a", UncoloredBG("b"), "c"),
				ansi.Mode(ansi.COLOR_BG_YELLOW)+"a"+ansi.Mode(ansi.COLOR_BG_DEFAULT)+"b"+ansi.Mode(ansi.COLOR_BG_YELLOW)+"c"+ansi.Mode(ansi.COLOR_BG_DEFAULT)),
			Entry("background keeps foreground", Blue(BgYellow("a", UncoloredBG("b"), "c")),
				ansi.Mode(ansi.COLOR_FG_BLUE, ansi.COLOR_BG_YELLOW)+"a"+ansi.Mode(ansi.COLOR_BG_DEFAULT)+"b"+ansi.Mode(ansi.COLOR_BG_YELLOW)+"c"+ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.COLOR_BG_DEFAULT)),
			Entry("rgb background", BgRGBColor(1, 2, 3, "a", UncoloredBG("b")),
				ansi.Mode(ansi.Code_BG_RGB(1, 2, 3))+"a"+ansi.Mode(ansi.COLOR_BG_DEFAULT)+"b"),
			Entry("top level", UncoloredFG("x"), ansi.Mode(ansi.COLOR_FG_DEFAULT)+"x"),
		)

		It("alternating", func() {
			s := Blue("blue", BgGreen("green back", Red("red", BgBlack("black back"), "red"), "green back"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "green back" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_BG_BLACK) + "black back" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "green back" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})
	})

	Context("splitting", func() {
		It("keeps inheriting plain texts", func() {
			ctx := NewContext(true)
			l, r := ctx.Sequence(ctx.Italic("aa"), "bb").SplitAt(3)
			Expect(ctx.Italic(l, "|", r).String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "aab|b" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("keeps switched off formats", func() {
			ctx := NewContext(true)
			l, r := ctx.Sequence(ctx.NonItalic("aabb")).SplitAt(2)
			Expect(ctx.Italic(l, "|", r).String()).To(Equal(ansi.Mode(ansi.MODE_RESET_ITALIC) + "aa" +
				ansi.Mode(ansi.MODE_ITALIC) + "|" + ansi.Mode(ansi.MODE_RESET_ITALIC) + "bb"))
		})

		It("splits embedded formatted texts", func() {
			ctx := NewContext(true)
			l, r := ctx.Sequence("x"+ansi.Mode(ansi.MODE_BOLD)+"yz", 42).SplitAt(2)
			Expect(l.String()).To(Equal("x" + ansi.Mode(ansi.MODE_BOLD) + "y" + ansi.Mode(ansi.MODE_RESET_BOLD)))
			Expect(r.String()).To(Equal(ansi.Mode(ansi.MODE_BOLD) + "z" + ansi.Mode(ansi.MODE_RESET_BOLD) + "42"))
		})

		It("keeps explicitly switched off formats of embedded formatted texts", func() {
			ctx := NewContext(true)
			pre := ctx.Sequence(ctx.Italic("aa"), ctx.NonItalic("bb")).String()
			l, r := ctx.Sequence(pre).SplitAt(3)
			Expect(r.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_ITALIC) + "b"))
			// the second "b" stays upright in an outer italic format
			Expect(ctx.Italic(l, r, "c").String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + ansi.Mode(ansi.MODE_ITALIC) + "aa" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" + ansi.Mode(ansi.MODE_ITALIC) + ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" +
				ansi.Mode(ansi.MODE_ITALIC) + "c" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("keeps formats of embedded formatted texts for an outer format", func() {
			ctx := NewContext(true)
			l, r := ctx.Sequence(ctx.Blue("ab", "cd").String()).SplitAt(1)
			Expect(ctx.Sequence(ctx.Reverse(l), r).String()).To(Equal(ansi.Mode(ansi.MODE_REVERSE) +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "a" + ansi.Mode(ansi.COLOR_FG_DEFAULT) + ansi.Mode(ansi.MODE_RESET_REVERSE) +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bcd" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("keeps formats of the second part for an outer format", func() {
			ctx := NewContext(true)
			l, r := ctx.Blue("ab", "cd").SplitAt(1)
			Expect(ctx.Sequence(ctx.Reverse(l), r).String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE, ansi.MODE_REVERSE) + "a" +
				ansi.Mode(ansi.MODE_RESET_REVERSE) + "bcd" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})
	})
})

// nestedSequenceTable provides the arguments for DescribeTable (the
// test function followed by the table entries) for sequences with
// formats nested in the same format for all pairs of a format and
// its counterpart: outer(Sequence(first("aaaa"), second("bbb")), "ccc").
// A nil second element means a plain string.
//
// A plain string inherits the format of the outer format, while the
// counterpart switches it off. This difference is lost for pre-rendered
// strings: a rendered string cannot express inheriting a format, it
// always describes the terminal state for the default rendition.
func nestedSequenceTable(body any) []any {
	type pair struct {
		name    string
		mode    func(...any) String
		non     func(...any) String
		on, off ansi.CodeSequence
	}
	pairs := []pair{
		{"bold", Bold, NonHighlighted, ansi.NewCodeSequence(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD), ansi.MODE_RESET_BOLD},
		{"dim", Dim, NonHighlighted, ansi.NewCodeSequence(ansi.MODE_RESET_BOLD, ansi.MODE_DIM), ansi.MODE_RESET_BOLD},
		{"italic", Italic, NonItalic, ansi.MODE_ITALIC, ansi.MODE_RESET_ITALIC},
		{"underline", Underline, NonUnderline, ansi.MODE_UNDERLINE, ansi.MODE_RESET_UNDERLINE},
		{"blinking", Blinking, NonBlinking, ansi.MODE_BLINK, ansi.MODE_RESET_BLINK},
		{"reverse", Reverse, NonReverse, ansi.MODE_REVERSE, ansi.MODE_RESET_REVERSE},
		{"invisible", Invisible, NonInvisible, ansi.MODE_INVIS, ansi.MODE_RESET_INVIS},
		{"strike through", StrikeThrough, NonStrikeThrough, ansi.MODE_STRIKE, ansi.MODE_RESET_STRIKE},
		{"overline", Overline, NonOverline, ansi.MODE_OVERLINE, ansi.MODE_RESET_OVERLINE},
		{"foreground color", Blue, UncoloredFG, ansi.COLOR_FG_BLUE, ansi.COLOR_FG_DEFAULT},
		{"background color", BgBlue, UncoloredBG, ansi.COLOR_BG_BLUE, ansi.COLOR_BG_DEFAULT},
	}
	args := []any{body}
	for _, p := range pairs {
		on, off := ansi.Mode(p.on), ansi.Mode(p.off)
		args = append(args,
			// the plain string inherits the outer format
			Entry(p.name+": plain string inherits", p.mode, p.mode, nil, false,
				on+"aaaabbbccc"+off),
			// the counterpart switches the outer format off
			Entry(p.name+": counterpart switches off", p.mode, p.mode, p.non, false,
				on+"aaaa"+off+"bbb"+on+"ccc"+off),
			// pre-rendered, both are identical: the plain string does not inherit
			Entry(p.name+": plain string pre-rendered", p.mode, p.mode, nil, true,
				on+on+"aaaa"+off+"bbb"+on+"ccc"+off),
			Entry(p.name+": counterpart pre-rendered", p.mode, p.mode, p.non, true,
				on+on+"aaaa"+off+"bbb"+on+"ccc"+off),
			// counterpart as first element
			Entry(p.name+": counterpart first", p.mode, p.non, nil, false,
				off+"aaaa"+on+"bbbccc"+off),
			Entry(p.name+": counterpart first pre-rendered", p.mode, p.non, nil, true,
				on+off+"aaaabbb"+on+"ccc"+off),
		)
	}
	return args
}
