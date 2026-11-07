package ttycolors_test

import (
	"fmt"

	. "github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format Test Environment", func() {
	Context("disabled", func() {
		It("mode", func() {
			s := FmtBold.String("bold")
			s.Enable(false)
			Expect(s.String()).To(Equal("bold"))
		})

		It("standard color", func() {
			s := FmtBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg color", func() {
			s := FmtBgBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bright color", func() {
			s := FmtBrightBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg bright color", func() {
			s := FmtBgBrightBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb color", func() {
			s := FmtRGBColor(1, 255, 2).String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb bg color", func() {
			s := FmtBgRGBColor(1, 255, 2).String("green")
			s.Enable(false)
			Expect(s.String()).To(Equal("green"))
		})

		It("id color", func() {
			s := FmtIdColor(1).String("id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("id bg color", func() {
			s := FmtBgIdColor(1).String("id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("complex color", func() {
			s := FmtBlue.String("blue", Green("green"), Red("red"), "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue" +
				"green" +
				"red" +
				"blue"))
		})

		It("mixed", func() {
			s := FmtBlue.String("blue", Bold(Green("bold green"), "bold blue"), Red("red"), "blue")
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
			s := FmtBold.String("bold")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("standard color", func() {
			s := FmtBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg color", func() {
			s := FmtBgBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("standard bright color", func() {
			s := FmtBrightBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg bright color", func() {
			s := FmtBgBrightBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("rgb color", func() {
			s := FmtRGBColor(1, 255, 2).String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RGB, ansi.Param(1), ansi.Param(255), ansi.Param(2)) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("rgb bg color", func() {
			s := FmtBgRGBColor(1, 255, 2).String("green")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RGB, ansi.Param(1), ansi.Param(255), ansi.Param(2)) + "green" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("id color", func() {
			s := FmtIdColor(1).String("id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_ID, ansi.Param(1)) + "id" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("id gb color", func() {
			s := FmtBgIdColor(1).String("id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_ID, ansi.Param(1)) + "id" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("complex color", func() {
			s := FmtBlue.String("blue", Green("green"), Red("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("mixed", func() {
			s := FmtBlue.String("blue", Bold(Green("bold green"), "bold blue"), Red("red"), "blue")
			s.Enable()
			fmt.Printf("\n%v\n", []byte(s.String()))
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_RESET_BOLD) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		DescribeTable("modes",
			func(f, non Format, on, off ansi.CodeSequence) {
				s := f.String("a", non.String("b"), "c")
				s.Enable()
				Expect(s.String()).To(Equal(ansi.Mode(on) + "a" + ansi.Mode(off) + "b" + ansi.Mode(on) + "c" + ansi.Mode(off)))
			},
			Entry("italic", FmtItalic, FmtNonItalic, ansi.MODE_ITALIC, ansi.MODE_RESET_ITALIC),
			Entry("underline", FmtUnderline, FmtNonUnderline, ansi.MODE_UNDERLINE, ansi.MODE_RESET_UNDERLINE),
			Entry("blinking", FmtBlinking, FmtNonBlinking, ansi.MODE_BLINK, ansi.MODE_RESET_BLINK),
			Entry("reverse", FmtReverse, FmtNonReverse, ansi.MODE_REVERSE, ansi.MODE_RESET_REVERSE),
			Entry("invisible", FmtInvisible, FmtNonInvisible, ansi.MODE_INVIS, ansi.MODE_RESET_INVIS),
			Entry("strike through", FmtStrikeThrough, FmtNonStrikeThrough, ansi.MODE_STRIKE, ansi.MODE_RESET_STRIKE),
			Entry("overline", FmtOverline, FmtNonOverline, ansi.MODE_OVERLINE, ansi.MODE_RESET_OVERLINE),
		)

		It("uncolored formats", func() {
			s := New(FmtRed, FmtBgWhite).String("a", FmtUncoloredFG.String("b"), FmtUncoloredBG.String("c"))
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_WHITE) + "a" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT) + "b" +
				ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_DEFAULT) + "c" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("non highlighted in bold", func() {
			s := FmtBold.String("a", FmtNonHighlighted.String("b"), "c")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "a" +
				ansi.Mode(ansi.MODE_RESET_BOLD) + "b" +
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "c" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("dim in bold", func() {
			s := FmtBold.String("a", FmtDim.String("b"))
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "a" +
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_DIM) + "b" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("composed format with modes", func() {
			f := New(FmtRed, FmtItalic, FmtUnderline)
			s := f.String("a", NonItalic("b"), "c")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_ITALIC, ansi.MODE_UNDERLINE) + "a" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" +
				ansi.Mode(ansi.MODE_ITALIC) + "c" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_ITALIC, ansi.MODE_RESET_UNDERLINE)))
		})

		It("composed non format", func() {
			s := FmtItalic.String("a", New(FmtNonItalic, FmtBlue).String("b"), "c")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "a" +
				ansi.Mode(ansi.COLOR_FG_BLUE, ansi.MODE_RESET_ITALIC) + "b" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_ITALIC) + "c" +
				ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("underline style with features", func() {
			f := New(FmtDashedUnderline)
			f.EnableFeatures(FeatureUnderlineStyles)
			s := f.String("a", Underline("b"), "c")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_DASHED) + "a" +
				ansi.Mode(ansi.MODE_UNDERLINE) + "b" +
				ansi.Mode(ansi.UNDERLINE_DASHED) + "c" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("underline style without features", func() {
			s := FmtDoubleUnderline.String("x")
			s.Enable()
			s.(ContextInfo).DisableFeatures(ExtendedFeatures)
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		DescribeTable("omits formatting without text",
			func(s String) {
				s.Enable()
				Expect(s.String()).To(Equal(""))
			},
			Entry("mode format", FmtBold.String()),
			Entry("color format", FmtRed.String("")),
			Entry("composed format", New(FmtRed, FmtItalic, FmtBgWhite).String("", Bold())),
			Entry("feature format", FmtCurlyUnderline.String()),
		)

		It("omits formatting of empty nested formats", func() {
			s := FmtRed.String("a", FmtItalic.String(), New(FmtBold, FmtBgWhite).String(""), "b")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "ab" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("ignores nil formats when adding", func() {
			var none FormatInfo
			f := New(FmtRed)
			f.Add(none)
			s := f.String("x")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
			Expect(New(FmtRed, none).String("x").String()).To(Equal(New(FmtRed).String("x").String()))
		})

		It("replaces settings of the same group", func() {
			s := New(FmtBold, FmtRed, FmtDim, FmtBlue).String("x")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE, ansi.MODE_RESET_BOLD, ansi.MODE_DIM) + "x" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})
	})

	Context("start and end", func() {
		It("composed format", func() {
			f := New(FmtItalic, FmtBgWhite, FmtBold, FmtRed)
			f.Enable()
			Expect(string(f.Start())).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_WHITE, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD, ansi.MODE_ITALIC)))
			Expect(string(f.End())).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.COLOR_BG_DEFAULT, ansi.MODE_RESET_BOLD, ansi.MODE_RESET_ITALIC)))
		})

		It("format with features", func() {
			f := New(FmtRed, FmtCurlyUnderline, FmtUnderlineIdColor(4))
			f.Enable()
			f.EnableFeatures(FeatureUnderlineStyles | FeatureUnderlineColor)
			Expect(string(f.Start())).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + ansi.Mode(ansi.UNDERLINE_CURLY) + ansi.Mode(ansi.Code_UL_ID(4))))
			Expect(string(f.End())).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_UNDERLINE, ansi.COLOR_UL_DEFAULT)))
		})

		It("format without features", func() {
			f := New(FmtRed, FmtCurlyUnderline, FmtUnderlineRGBColor(1, 2, 3))
			f.Enable()
			f.DisableFeatures(ExtendedFeatures)
			Expect(string(f.Start())).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_UNDERLINE)))
			Expect(string(f.End())).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_UNDERLINE)))
		})

		It("empty format", func() {
			f := New()
			f.Enable()
			Expect(f.Start()).To(BeEmpty())
			Expect(f.End()).To(BeEmpty())
		})

		It("disabled format", func() {
			f := New(FmtItalic)
			f.Enable(false)
			Expect(f.Start()).To(BeNil())
			Expect(f.End()).To(BeNil())
		})
	})

	Context("scanned formats", func() {
		ctx := NewContext(true)

		It("continues the format at the end of a text", func() {
			pre := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "ab" + ansi.Mode(ansi.MODE_RESET_ITALIC) + "cd"
			Expect(ctx.ScanFormat(pre).String("x").String()).To(Equal(
				ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD, ansi.MODE_RESET_ITALIC) + "x" +
					ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})

		It("keeps explicit resets in an outer format", func() {
			f := ctx.ScanFormat(ansi.Mode(ansi.MODE_RESET_ITALIC) + "a")
			Expect(ctx.Italic("a", f.String("b"), "c").String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "a" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" + ansi.Mode(ansi.MODE_ITALIC) + "c" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("inherits untouched formats", func() {
			Expect(ctx.Blue(ctx.ScanFormat("plain").String("x")).String()).To(Equal(
				ansi.Mode(ansi.COLOR_FG_BLUE) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("handles a reset of all modes", func() {
			f := ctx.ScanFormat(ansi.Mode(ansi.MODE_BOLD) + "a" + ansi.Mode(ansi.MODE_RESET_ALL, ansi.COLOR_BG_BLUE))
			Expect(ctx.Italic(f.String("x")).String()).To(Equal(ansi.Mode(ansi.MODE_RESET_ALL, ansi.COLOR_BG_BLUE) + "x" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("maps dim like FmtDim", func() {
			Expect(ctx.ScanFormat(ansi.Mode(ansi.MODE_DIM)).String("x").String()).To(Equal(ctx.New(FmtDim).String("x").String()))
		})

		It("handles underline styles like the underline style formats", func() {
			c := NewContext(true)
			f := c.ScanFormat(ansi.Mode(ansi.UNDERLINE_CURLY))
			Expect(f.String("x").String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
			c.EnableFeatures(FeatureUnderlineStyles)
			f = c.ScanFormat(ansi.Mode(ansi.UNDERLINE_CURLY))
			Expect(f.String("x").String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("handles underline style 0 as no underline", func() {
			f := ctx.ScanFormat(ansi.Mode(ansi.Code{ansi.MODE_UNDERLINE, 0}))
			Expect(ctx.Underline(f.String("x")).String()).To(Equal(ansi.Mode(ansi.MODE_RESET_UNDERLINE) + "x"))
		})

		It("handles underline colors like the underline color formats", func() {
			c := NewContext(true)
			f := c.ScanFormat(ansi.Mode(ansi.Code_UL_ID(5)))
			Expect(f.String("x").String()).To(Equal("x"))
			c.EnableFeatures(FeatureUnderlineColor)
			f = c.ScanFormat(ansi.Mode(ansi.Code_UL_ID(5)))
			Expect(f.String("x").String()).To(Equal(ansi.Mode(ansi.Code_UL_ID(5)) + "x" + ansi.Mode(ansi.COLOR_UL_DEFAULT)))
		})

		It("is composable", func() {
			f := New(ctx.ScanFormat(ansi.Mode(ansi.COLOR_FG_RED)), FmtItalic)
			f.Enable()
			Expect(f.String("x").String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_ITALIC) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_ITALIC)))
		})

		It("uses the settings of the default context", func() {
			f := ScanFormat(ansi.Mode(ansi.COLOR_FG_RED))
			Expect(f.IsEnabled()).To(Equal(!NoColors))
		})
	})
})
