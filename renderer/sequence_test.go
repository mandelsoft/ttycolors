package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Sequence", func() {
	It("copies the settings", func() {
		settings := NewSettings(FeatureFormatting | FeatureUnderlineColor)
		s := Sequence(settings, []any{"x"})
		settings.Enable(false)
		settings.SetFeatures(NoFeatures)

		info := s.(ContextInfo)
		Expect(info.IsEnabled()).To(BeTrue())
		Expect(info.Features()).To(Equal(FeatureFormatting | FeatureUnderlineColor))
	})

	It("sets features", func() {
		s := Sequence(on, []any{NewFeatureFormat(on, Underline, FeatureUnderlineStyles, ansi.UNDERLINE_CURLY, ansi.MODE_UNDERLINE, []any{"x"})})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		s.(ContextInfo).EnableFeatures(FeatureUnderlineStyles)
		Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})

	It("appends with the settings of the sequence", func() {
		s := Sequence(all, []any{"a"}).Append("b")
		info := s.(ContextInfo)
		Expect(info.IsEnabled()).To(BeTrue())
		Expect(info.Features()).To(Equal(all.features))
	})

	It("strips embedded sequences if disabled", func() {
		s := Sequence(off, []any{"a", "\x1b[1mb\x1b[22m", NewFormat(off, FG, ansi.COLOR_FG_RED, []any{"\x1b[3mc"})})
		Expect(s.String()).To(Equal("abc"))
	})

	It("tracks embedded sequences if enabled", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_BLUE, []any{"a", "\x1b[31mb\x1b[39m", "c"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "a\x1b[31mb\x1b[39m" + ansi.Mode(ansi.COLOR_FG_BLUE) + "c" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})
})
