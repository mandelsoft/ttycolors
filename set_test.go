package ttycolors_test

import (
	"bytes"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Set", func() {
	var buf bytes.Buffer

	BeforeEach(func() {
		buf.Reset()
	})

	It("uses the default context", func() {
		Expect(SetFor(&buf, FmtRed)).To(Succeed())
		if NoColors {
			Expect(buf.String()).To(Equal(""))
		} else {
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		}
	})

	Context("context", func() {
		It("sets and unsets formats", func() {
			ctx := NewContext(true)
			Expect(ctx.SetFor(&buf, FmtRed, FmtBold)).To(Succeed())
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_RESET_BOLD, ansi.MODE_BOLD)))
			buf.Reset()
			Expect(ctx.UnSetFor(&buf, FmtRed, FmtBold)).To(Succeed())
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})

		It("writes nothing if disabled", func() {
			ctx := NewContext(false)
			Expect(ctx.SetFor(&buf, FmtRed)).To(Succeed())
			Expect(ctx.UnSetFor(&buf, FmtRed)).To(Succeed())
			Expect(buf.String()).To(Equal(""))
		})

		It("ignores the settings of the given formats", func() {
			f := NewContext(false).New(FmtRed)
			Expect(NewContext(true).SetFor(&buf, f)).To(Succeed())
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		})

		It("uses the features of the context", func() {
			ctx := NewContext(true)
			Expect(ctx.SetFor(&buf, FmtCurlyUnderline)).To(Succeed())
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.MODE_UNDERLINE)))
			buf.Reset()
			ctx.EnableFeatures(FeatureUnderlineStyles)
			Expect(ctx.SetFor(&buf, FmtCurlyUnderline)).To(Succeed())
			Expect(buf.String()).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY)))
		})
	})
})
