package ansi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Constants", func() {
	It("provides the palette ids of the standard and bright colors", func() {
		ids := []byte{
			ansi.COLOR_ID_BLACK, ansi.COLOR_ID_RED, ansi.COLOR_ID_GREEN, ansi.COLOR_ID_YELLOW,
			ansi.COLOR_ID_BLUE, ansi.COLOR_ID_MAGENTA, ansi.COLOR_ID_CYAN, ansi.COLOR_ID_WHITE,
			ansi.COLOR_ID_BRIGHT_BLACK, ansi.COLOR_ID_BRIGHT_RED, ansi.COLOR_ID_BRIGHT_GREEN, ansi.COLOR_ID_BRIGHT_YELLOW,
			ansi.COLOR_ID_BRIGHT_BLUE, ansi.COLOR_ID_BRIGHT_MAGENTA, ansi.COLOR_ID_BRIGHT_CYAN, ansi.COLOR_ID_BRIGHT_WHITE,
		}
		for i, id := range ids {
			Expect(id).To(Equal(byte(i)))
		}
	})

	It("are usable for underline colors", func() {
		Expect(ansi.Mode(ansi.Code_UL_ID(ansi.COLOR_ID_RED))).To(Equal("\x1b[58:5:1m"))
		Expect(ansi.Mode(ansi.Code_FG_ID(ansi.COLOR_ID_BRIGHT_BLUE))).To(Equal("\x1b[38;5;12m"))
	})
})
