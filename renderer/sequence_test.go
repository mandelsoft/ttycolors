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

	It("keeps embedded sequences if disabled", func() {
		s := Sequence(off, []any{"a", "\x1b[1mb\x1b[22m", NewFormat(off, FG, ansi.COLOR_FG_RED, []any{"\x1b[3mc"})})
		Expect(s.String()).To(Equal("a\x1b[1mb\x1b[22m\x1b[3mc"))
	})

	It("tracks embedded sequences if enabled", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_BLUE, []any{"a", "\x1b[31mb\x1b[39m", "c"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "a\x1b[31mb\x1b[39m" + ansi.Mode(ansi.COLOR_FG_BLUE) + "c" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("keeps control sequences if disabled", func() {
		s := Sequence(off, []any{"a\x1b[2Kb\x1b[1mc\x1b]0;title\a"})
		Expect(s.String()).To(Equal("a\x1b[2Kb\x1b[1mc\x1b]0;title\a"))
	})

	It("does not track control sequences if enabled", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"a\x1b[2Kb", "c"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "a\x1b[2Kbc" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	Context("splitting", func() {
		It("counts characters", func() {
			s := Sequence(on, []any{"äb", "\x1b[1mc\x1b[22m", 42, []byte("x"), NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"de"})})
			Expect(s.CharLen()).To(Equal(8))
			Expect(Sequence(on, nil).CharLen()).To(Equal(0))
		})

		DescribeTable("splits at element boundaries and inside elements",
			func(n int, left, right string) {
				s := Sequence(on, []any{"ab", 42, []byte("cd")})
				l, r := s.SplitAt(n)
				Expect(l.String()).To(Equal(left))
				Expect(r.String()).To(Equal(right))
			},
			Entry("start", 0, "", "ab42cd"),
			Entry("inside string", 1, "a", "b42cd"),
			Entry("element boundary", 2, "ab", "42cd"),
			Entry("inside other value", 3, "ab4", "2cd"),
			Entry("inside bytes", 5, "ab42c", "d"),
			Entry("end", 6, "ab42cd", ""),
			Entry("beyond end", 7, "ab42cd", ""),
		)

		It("splits multi-byte characters", func() {
			l, r := Sequence(on, []any{"äöü"}).SplitAt(1)
			Expect(l.String()).To(Equal("ä"))
			Expect(r.String()).To(Equal("öü"))
		})

		It("keeps the settings", func() {
			l, r := Sequence(all, []any{"ab"}).SplitAt(1)
			Expect(l.(ContextInfo).Features()).To(Equal(all.features))
			Expect(r.(ContextInfo).Features()).To(Equal(all.features))
		})

		It("panics for a negative index", func() {
			Expect(func() { Sequence(on, []any{"ab"}).SplitAt(-1) }).To(Panic())
		})
	})
})
