package ansi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Formats", func() {
	It("underline color id", func() {
		c := ansi.Code_UL_ID(9)
		Expect(ansi.Mode(c)).To(Equal("\x1b[58:5:9m"))
		Expect(c).To(Equal(ansi.Code{58, 5, 9}))
	})

	It("rgb and id colors", func() {
		Expect(ansi.Code_FG_RGB(1, 2, 3)).To(Equal(ansi.NewCodeSequence(38, 2, 1, 2, 3)))
		Expect(ansi.Code_BG_RGB(1, 2, 3)).To(Equal(ansi.NewCodeSequence(48, 2, 1, 2, 3)))
		Expect(ansi.Code_FG_ID(7)).To(Equal(ansi.NewCodeSequence(38, 5, 7)))
		Expect(ansi.Code_BG_ID(7)).To(Equal(ansi.NewCodeSequence(48, 5, 7)))
	})

	It("underline rgb color", func() {
		c := ansi.Code_UL_RGB(1, 2, 3)
		Expect(ansi.Mode(c)).To(Equal("\x1b[58:2::1:2:3m"))
		Expect(c).To(Equal(ansi.Code{58, 2, ansi.Omitted, 1, 2, 3}))
	})

	It("underline colors are detected by the scanner", func() {
		Expect(ansi.Scan(ansi.Mode(ansi.Code_UL_ID(9)) + "x" + ansi.Mode(ansi.Code_UL_RGB(1, 2, 3)))).To(Equal(ansi.S_UL_COLOR))
	})

	It("underline styles are detected by the scanner", func() {
		Expect(ansi.Scan(ansi.Mode(ansi.UNDERLINE_DASHED))).To(Equal(ansi.S_UNDERLINE))
	})
})
