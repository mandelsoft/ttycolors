package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

// contexts used by the tests.
var (
	on  = NewSettings(FeatureFormatting)
	off = NewSettings(NoFeatures)
	all = NewSettings(AllFeatures)
)

var _ = Describe("Format", func() {
	It("renders a group format", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"red"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "red" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("restores the outer code of the same group", func() {
		s := NewFormat(on, BG, ansi.COLOR_BG_BLUE, []any{NewFormat(on, BG, ansi.COLOR_BG_RED, []any{"red"}), "blue"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RED) + "red" +
			ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
	})

	It("keeps other groups", func() {
		s := NewFormat(on, Highlight, ansi.MODE_BOLD, []any{NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"red"}), "bold"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD) + "red" +
			ansi.Mode(ansi.COLOR_FG_DEFAULT) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
	})

	It("combines the codes of nested formats", func() {
		s := NewFormat(on, Italic, ansi.MODE_ITALIC, []any{NewFormat(on, BG, ansi.COLOR_BG_RED, []any{NewFormat(on, FG, ansi.COLOR_FG_BLUE, []any{"x"})})})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE, ansi.COLOR_BG_RED, ansi.MODE_ITALIC) + "x" +
			ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.COLOR_BG_DEFAULT, ansi.MODE_RESET_ITALIC)))
	})

	It("combines restores and resets", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_BLUE, []any{NewFormat(on, Italic, ansi.MODE_ITALIC, []any{NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"a"})}), "b"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_ITALIC) + "a" +
			ansi.Mode(ansi.COLOR_FG_BLUE, ansi.MODE_RESET_ITALIC) + "b" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("switches off a mode in nested elements", func() {
		s := NewFormat(on, Italic, ansi.MODE_ITALIC, []any{"a", NewFormat(on, Italic, ansi.MODE_RESET_ITALIC, []any{"b"}), "c"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "a" +
			ansi.Mode(ansi.MODE_RESET_ITALIC) + "b" +
			ansi.Mode(ansi.MODE_ITALIC) + "c" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
	})

	It("is a no-op without text", func() {
		Expect(NewFormat(on, Italic, ansi.MODE_ITALIC, nil).String()).To(Equal(""))
		Expect(NewFormat(on, Italic, ansi.MODE_ITALIC, []any{"", NewFormat(on, FG, ansi.COLOR_FG_RED, nil)}).String()).To(Equal(""))
	})

	It("keeps the state for empty nested formats", func() {
		s := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"a", NewFormat(on, Italic, ansi.MODE_ITALIC, []any{""}), "b"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "ab" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("requires a group", func() {
		Expect(func() { NewFormat(on, nil, ansi.MODE_ITALIC, nil) }).To(PanicWith("renderer.NewFormat: group required"))
	})

	It("renders nothing if disabled", func() {
		s := NewFormat(off, FG, ansi.COLOR_FG_RED, []any{NewFormat(off, Highlight, ansi.MODE_BOLD, []any{"text"})})
		Expect(s.String()).To(Equal("text"))
	})

	Context("features", func() {
		curly := func(ctx Settings, nested ...any) String {
			return NewFeatureFormat(ctx, Underline, FeatureUnderlineStyles, ansi.UNDERLINE_CURLY, ansi.MODE_UNDERLINE, nested)
		}
		ulColor := func(ctx Settings, nested ...any) String {
			return NewFeatureFormat(ctx, UnderlineColor, FeatureUnderlineColor, ansi.Code_UL_ID(1), nil, nested)
		}

		It("uses the feature code if enabled", func() {
			Expect(curly(all, "x").String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("uses the fallback if not enabled", func() {
			Expect(curly(on, "x").String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("ignores the group without fallback if not enabled", func() {
			Expect(ulColor(on, "x").String()).To(Equal("x"))
			Expect(NewFormat(on, FG, ansi.COLOR_FG_RED, []any{ulColor(on, "x")}).String()).
				To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("uses the features of the outermost string", func() {
			s := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{curly(all, "x")})
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_UNDERLINE) + "x" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_UNDERLINE)))
		})

		It("renders sub parameter codes in separate sequences", func() {
			s := NewFormat(all, FG, ansi.COLOR_FG_RED, []any{curly(all, ulColor(all, "x"))})
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) +
				ansi.Mode(ansi.UNDERLINE_CURLY) + ansi.Mode(ansi.Code_UL_ID(1)) + "x" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_UNDERLINE, ansi.COLOR_UL_DEFAULT)))
		})

		It("restores an underline style after a nested style", func() {
			s := NewFeatureFormat(all, Underline, FeatureUnderlineStyles, ansi.UNDERLINE_DOTTED, ansi.MODE_UNDERLINE,
				[]any{"a", curly(all, "b"), "c"})
			Expect(s.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_DOTTED) + "a" +
				ansi.Mode(ansi.UNDERLINE_CURLY) + "b" +
				ansi.Mode(ansi.UNDERLINE_DOTTED) + "c" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})
	})

	Context("splitting", func() {
		It("counts the characters of the content", func() {
			s := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"ab", NewFormat(on, Highlight, ansi.MODE_BOLD, []any{"cd"})})
			Expect(s.CharLen()).To(Equal(4))
		})

		It("uses the format for both parts", func() {
			l, r := NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"abcd"}).SplitAt(2)
			Expect(l.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "ab" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
			Expect(r.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "cd" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("splits nested formats", func() {
			s := NewFormat(on, Italic, ansi.MODE_ITALIC, []any{"aa", NewFormat(on, Highlight, ansi.MODE_BOLD, []any{"bbbb"}), "cc"})
			l, r := s.SplitAt(4)
			Expect(l.String()).To(Equal(ansi.Mode(ansi.MODE_ITALIC) + "aa" + ansi.Mode(ansi.MODE_BOLD) + "bb" +
				ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_RESET_ITALIC)))
			Expect(r.String()).To(Equal(ansi.Mode(ansi.MODE_BOLD, ansi.MODE_ITALIC) + "bb" +
				ansi.Mode(ansi.MODE_RESET_BOLD) + "cc" + ansi.Mode(ansi.MODE_RESET_ITALIC)))
		})

		It("renders the original when the parts are joined", func() {
			s := NewFormat(on, Italic, ansi.MODE_ITALIC, []any{"aa", NewFormat(on, Highlight, ansi.MODE_BOLD, []any{"bbbb"}), "cc"})
			for n := 0; n <= s.CharLen()+1; n++ {
				l, r := s.SplitAt(n)
				Expect(Sequence(on, []any{l, r}).String()).To(Equal(s.String()), "split at %d", n)
			}
		})

		It("keeps feature and fallback", func() {
			f := NewFeatureFormat(on, Underline, FeatureUnderlineStyles, ansi.UNDERLINE_CURLY, ansi.MODE_UNDERLINE, []any{"ab"})
			l, r := f.SplitAt(1)
			Expect(l.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "a" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
			r.(ContextInfo).EnableFeatures(FeatureUnderlineStyles)
			Expect(r.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "b" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		})

		It("panics for a negative index", func() {
			Expect(func() { NewFormat(on, FG, ansi.COLOR_FG_RED, []any{"ab"}).SplitAt(-1) }).To(Panic())
		})
	})
})
