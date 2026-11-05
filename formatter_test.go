package ttycolors_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Formatter", func() {
	It("uses the settings of the default context", func() {
		f := NewFormatter(FmtRed)
		Expect(f.IsEnabled()).To(Equal(!NoColors))
		Expect(f.Features()).To(Equal(EnabledFeatures()))
	})

	It("uses the settings of a context", func() {
		ctx := NewContext(true)
		ctx.EnableFeatures(FeatureUnderlineStyles)
		f := ctx.NewFormatter(FmtCurlyUnderline)
		Expect(f.IsEnabled()).To(BeTrue())
		Expect(f.Features()).To(Equal(FeatureFormatting | FeatureUnderlineStyles))
		Expect(f.Sprint("x")).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})

	It("omits formatting without text", func() {
		f := NewContext(true).NewFormatter(FmtRed, FmtBold)
		Expect(f.Sprint()).To(Equal(""))
		Expect(f.Sprint("")).To(Equal(""))
		Expect(f.Sprintf("")).To(Equal(""))
	})

	It("sets features", func() {
		f := NewContext(true).NewFormatter(FmtCurlyUnderline)
		Expect(f.Sprint("x")).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
		f.EnableFeatures(FeatureUnderlineStyles)
		Expect(f.Sprint("x")).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})
})
