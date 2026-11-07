package ttycolors_test

import (
	"os"

	"github.com/mandelsoft/ttycolors/ansi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
)

var _ = Describe("Test Environment", func() {
	It("disabled", func() {
		ctx := ttycolors.NewContext(false) // provide a context always disabled

		s := ctx.Blue("a blue text with some ", ctx.Italic("italic"), " word")
		Expect(s.String()).To(Equal("a blue text with some italic word"))
	})

	It("enabled", func() {
		ctx := ttycolors.NewContext(true) // provide a context always enabled

		s := ctx.Blue("a blue text with some ", ctx.Italic("italic"), " word")
		Expect(s.String()).To(Equal(
			ansi.Mode(ansi.COLOR_FG_BLUE) + "a blue text with some " +
				ansi.Mode(ansi.MODE_ITALIC) + "italic" +
				ansi.Mode(ansi.MODE_RESET_ITALIC) + " word" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("omits formatting without text", func() {
		ctx := ttycolors.NewContext(true)
		Expect(ctx.Bold().String()).To(Equal(""))
		Expect(ctx.Blue(ctx.Italic(""), ctx.BgRed()).String()).To(Equal(""))
		Expect(ctx.Blue("a", ctx.Italic(""), "b").String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "ab" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	Context("features", func() {
		var ctx ttycolors.TTYContext

		BeforeEach(func() {
			ctx = ttycolors.NewContext(true)
		})

		It("has no extended features by default", func() {
			Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting))
		})

		It("uses fallbacks without features", func() {
			s := ctx.CurlyUnderline("a", ctx.UnderlineIdColor(1, "b"))
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "a" + "b" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("uses features", func() {
			ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles | ttycolors.FeatureUnderlineColor)
			s := ctx.CurlyUnderline("a", ctx.UnderlineIdColor(1, "b"), ctx.UnderlineRGBColor(1, 2, 3, "c"))
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "a" +
				ansi.Mode(ansi.Code_UL_ID(1)) + "b" +
				ansi.Mode(ansi.Code_UL_RGB(1, 2, 3)) + "c" +
				ansi.Mode(ansi.MODE_RESET_UNDERLINE, ansi.COLOR_UL_DEFAULT)))
		})

		DescribeTable("underline styles",
			func(f func(ttycolors.TTYContext, ...any) ttycolors.String, style ansi.CodeSequence) {
				ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
				Expect(f(ctx, "x").String()).To(Equal(ansi.Mode(style) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
			},
			Entry("double", ttycolors.TTYContext.DoubleUnderline, ansi.UNDERLINE_DOUBLE),
			Entry("curly", ttycolors.TTYContext.CurlyUnderline, ansi.UNDERLINE_CURLY),
			Entry("dotted", ttycolors.TTYContext.DottedUnderline, ansi.UNDERLINE_DOTTED),
			Entry("dashed", ttycolors.TTYContext.DashedUnderline, ansi.UNDERLINE_DASHED),
		)

		It("copies the features when creating a string", func() {
			s := ctx.CurlyUnderline("x")
			ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("uses the features of the outermost string", func() {
			styled := ttycolors.NewContext(true)
			styled.EnableFeatures(ttycolors.FeatureUnderlineStyles)

			s := ctx.Sequence(styled.CurlyUnderline("x"))
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))

			s = styled.Sequence(ctx.CurlyUnderline("x"))
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("restores an underline style after a single underline", func() {
			ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
			s := ctx.DottedUnderline("a", ctx.Underline("b"), "c")
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_DOTTED) + "a" +
				ansi.Mode(ansi.MODE_UNDERLINE) + "b" +
				ansi.Mode(ansi.UNDERLINE_DOTTED) + "c" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("sets the features of the default context", func() {
			old := ttycolors.EnabledFeatures()
			defer ttycolors.SetFeatures(old)

			ttycolors.EnableFeatures(ttycolors.FeatureUnderlineStyles)
			Expect(ttycolors.EnabledFeatures()).To(Equal(ttycolors.FeatureUnderlineStyles))
			s := ttycolors.CurlyUnderline("x")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})
	})

	Context("embedded formatted strings", func() {
		var ctx ttycolors.TTYContext

		BeforeEach(func() {
			ctx = ttycolors.NewContext(true)
		})

		DescribeTable("restores the outer format",
			func(f func() ttycolors.String, expected string) {
				Expect(f().String()).To(Equal(expected))
			},
			Entry("pre-rendered string", func() ttycolors.String { return ctx.Blue("a", ctx.Red("red").String(), "c") },
				"\x1b[34ma\x1b[31mred\x1b[39m\x1b[34mc\x1b[39m"),
			Entry("reset all", func() ttycolors.String { return ctx.Blue("a", "x\x1b[0my", "c") },
				"\x1b[34max\x1b[0my\x1b[34mc\x1b[39m"),
			Entry("foreign mode left set", func() ttycolors.String { return ctx.Blue("a", "\x1b[1mbold", "c") },
				"\x1b[34ma\x1b[1mbold\x1b[22mc\x1b[39m"),
			Entry("foreign mode left set at the end", func() ttycolors.String { return ctx.Sequence("\x1b[1mbold") },
				"\x1b[1mbold\x1b[22m"),
			Entry("embedded reset", func() ttycolors.String { return ctx.Italic("a", "\x1b[4:3mu\x1b[24m", "c") },
				"\x1b[3ma\x1b[4:3mu\x1b[24mc\x1b[23m"),
			Entry("switched off outer mode", func() ttycolors.String { return ctx.Italic("a", "\x1b[23mx", "c") },
				"\x1b[3ma\x1b[23mx\x1b[3mc\x1b[23m"),
		)

		It("keeps embedded sequences if disabled", func() {
			red := ctx.Red("red").String()
			Expect(ttycolors.NewContext(false).Blue("a", red, "\x1b[1mc").String()).To(Equal("a" + red + "\x1b[1mc"))
		})

		It("strips embedded sequences on request if disabled", func() {
			red := ctx.Red("red").String()
			off := ttycolors.NewContext(false)
			Expect(off.Blue("a", off.StripIfDisabled(red), "c").String()).To(Equal("aredc"))
		})
	})

	Context("StripIfDisabled", func() {
		s := "\x1b[2K\x1b[1mbold\x1b[22m"

		It("keeps texts if enabled", func() {
			Expect(ttycolors.NewContext(true).StripIfDisabled(s)).To(Equal(s))
		})

		It("removes output format sequences if disabled", func() {
			Expect(ttycolors.NewContext(false).StripIfDisabled(s)).To(Equal("\x1b[2Kbold"))
		})

		It("uses the settings at the time of the call", func() {
			ctx := ttycolors.NewContext(true)
			Expect(ctx.StripIfDisabled(s)).To(Equal(s))
			ctx.Enable(false)
			Expect(ctx.StripIfDisabled(s)).To(Equal("\x1b[2Kbold"))
		})
	})
})

var _ = Describe("Global settings", func() {
	var noColors bool
	var features ttycolors.Features

	BeforeEach(func() {
		noColors = ttycolors.NoColors
		features = ttycolors.EnabledFeatures()
	})

	AfterEach(func() {
		ttycolors.SetFeatures(features)
		Expect(ttycolors.NoColors).To(Equal(noColors))
	})

	It("are initialized according to the standard output", func() {
		Expect(ttycolors.NoColors).To(Equal(!ttycolors.IsTerminal(os.Stdout)))
		Expect(ttycolors.EnabledFeatures().Has(ttycolors.FeatureFormatting)).To(Equal(!ttycolors.NoColors))
	})

	It("disable extended features by default", func() {
		Expect(ttycolors.EnabledFeatures() & ttycolors.ExtendedFeatures).To(Equal(ttycolors.NoFeatures))
	})

	It("keep NoColors and FeatureFormatting in sync", func() {
		ttycolors.NoColors = false
		Expect(ttycolors.EnabledFeatures().Has(ttycolors.FeatureFormatting)).To(BeTrue())
		ttycolors.NoColors = true
		Expect(ttycolors.EnabledFeatures().Has(ttycolors.FeatureFormatting)).To(BeFalse())

		ttycolors.EnableFeatures(ttycolors.FeatureFormatting)
		Expect(ttycolors.NoColors).To(BeFalse())
		ttycolors.DisableFeatures(ttycolors.FeatureFormatting)
		Expect(ttycolors.NoColors).To(BeTrue())
		ttycolors.SetFeatures(ttycolors.AllFeatures)
		Expect(ttycolors.NoColors).To(BeFalse())
	})

	It("keep the extended features independent of NoColors", func() {
		ttycolors.EnableFeatures(ttycolors.ExtendedFeatures)
		ttycolors.NoColors = true
		Expect(ttycolors.EnabledFeatures()).To(Equal(ttycolors.ExtendedFeatures))
		ttycolors.NoColors = false
		Expect(ttycolors.EnabledFeatures()).To(Equal(ttycolors.AllFeatures))
	})

	It("are used by the package functions", func() {
		ttycolors.SetFeatures(ttycolors.FeatureFormatting)
		s := ttycolors.CurlyUnderline("x")
		Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))

		ttycolors.EnableFeatures(ttycolors.ExtendedFeatures)
		s = ttycolors.CurlyUnderline("x")
		Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))

		ttycolors.DisableFeatures(ttycolors.FeatureFormatting)
		Expect(ttycolors.Red("x").String()).To(Equal("x"))
	})

	It("are copied by new contexts", func() {
		ttycolors.SetFeatures(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor)
		ctx := ttycolors.NewContext()
		Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor))

		ttycolors.SetFeatures(ttycolors.NoFeatures)
		Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor))
		Expect(ctx.Red("x").String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("copy the output formatting state given by NoColors to new contexts", func() {
		ttycolors.NoColors = false
		on := ttycolors.NewContext()
		ttycolors.NoColors = true
		off := ttycolors.NewContext()

		Expect(on.IsEnabled()).To(BeTrue())
		Expect(on.Red("x").String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(off.IsEnabled()).To(BeFalse())
		Expect(off.Red("x").String()).To(Equal("x"))
	})

	It("can be changed for new contexts", func() {
		ttycolors.SetFeatures(ttycolors.FeatureUnderlineStyles)
		on := ttycolors.NewContext(true)
		Expect(on.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineStyles))
		on.DisableFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(on.Features()).To(Equal(ttycolors.FeatureFormatting))
		Expect(ttycolors.EnabledFeatures()).To(Equal(ttycolors.FeatureUnderlineStyles))

		ttycolors.SetFeatures(ttycolors.AllFeatures)
		off := ttycolors.NewContext(false)
		Expect(off.Features()).To(Equal(ttycolors.ExtendedFeatures))
		Expect(off.IsEnabled()).To(BeFalse())
	})
})

var _ = Describe("Shared contexts", func() {
	var features ttycolors.Features

	BeforeEach(func() {
		features = ttycolors.EnabledFeatures()
	})

	AfterEach(func() {
		ttycolors.SetFeatures(features)
	})

	It("follow the global settings until changed", func() {
		ttycolors.SetFeatures(ttycolors.FeatureFormatting)
		ctx := ttycolors.NewSharedContext()
		Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting))

		ttycolors.EnableFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineStyles))
		ttycolors.NoColors = true
		Expect(ctx.IsEnabled()).To(BeFalse())
		ttycolors.NoColors = false
		Expect(ctx.IsEnabled()).To(BeTrue())
		Expect(ctx.CurlyUnderline("x").String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})

	It("get a local copy with the first change", func() {
		ttycolors.SetFeatures(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor)
		ctx := ttycolors.NewSharedContext()
		ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(ctx.Features()).To(Equal(ttycolors.AllFeatures))
		Expect(ttycolors.EnabledFeatures()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor))

		ttycolors.SetFeatures(ttycolors.NoFeatures)
		Expect(ctx.Features()).To(Equal(ttycolors.AllFeatures))
		Expect(ctx.Red("x").String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	DescribeTable("copy on every kind of change",
		func(change func(ttycolors.TTYContext)) {
			ttycolors.SetFeatures(ttycolors.FeatureFormatting)
			ctx := ttycolors.NewSharedContext()
			change(ctx)
			before := ctx.Features()
			ttycolors.SetFeatures(ttycolors.AllFeatures)
			Expect(ctx.Features()).To(Equal(before))
		},
		Entry("Enable", func(c ttycolors.TTYContext) { c.Enable(false) }),
		Entry("SetFeatures", func(c ttycolors.TTYContext) { c.SetFeatures(ttycolors.FeatureUnderlineColor) }),
		Entry("EnableFeatures", func(c ttycolors.TTYContext) { c.EnableFeatures(ttycolors.FeatureUnderlineColor) }),
		Entry("DisableFeatures", func(c ttycolors.TTYContext) { c.DisableFeatures(ttycolors.FeatureFormatting) }),
	)

	It("share the settings of another context", func() {
		base := ttycolors.NewContext(true)
		ctx := ttycolors.NewSharedContext(base)
		shared := ctx.Share()

		base.EnableFeatures(ttycolors.FeatureUnderlineColor)
		Expect(ctx.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor))
		Expect(shared.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineColor))

		ctx.Enable(false)
		Expect(base.IsEnabled()).To(BeTrue())
		Expect(shared.IsEnabled()).To(BeFalse())

		shared.Enable()
		Expect(ctx.IsEnabled()).To(BeFalse())
		Expect(shared.IsEnabled()).To(BeTrue())
	})

	It("use the default context for a nil base", func() {
		ttycolors.SetFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(ttycolors.NewSharedContext(nil).Features()).To(Equal(ttycolors.FeatureUnderlineStyles))
	})

	It("copy their settings to new strings", func() {
		ttycolors.SetFeatures(ttycolors.FeatureFormatting)
		ctx := ttycolors.NewSharedContext()
		s := ctx.Red("x")
		ttycolors.SetFeatures(ttycolors.NoFeatures)
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(ctx.Red("x").String()).To(Equal("x"))
	})

	It("scans formats with the settings of the context", func() {
		ctx := ttycolors.NewContext(true)
		ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
		f := ctx.ScanFormat(ansi.Mode(ansi.UNDERLINE_CURLY))
		Expect(f.IsEnabled()).To(BeTrue())
		Expect(f.Features()).To(Equal(ttycolors.FeatureFormatting | ttycolors.FeatureUnderlineStyles))
		Expect(f.String("x").String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})

	It("provides the default context", func() {
		ctx := ttycolors.DefaultContext()
		Expect(ctx).NotTo(BeNil())
		Expect(ctx.IsEnabled()).To(Equal(!ttycolors.NoColors))
		Expect(ctx.Features()).To(Equal(ttycolors.EnabledFeatures()))
	})
})
