package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Format", func() {
	It("renders a group format", func() {
		s := NewFGColor(true, ansi.COLOR_FG_RED, []any{"red"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED) + "red" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
	})

	It("restores the outer code of the same group", func() {
		s := NewBGColor(true, ansi.COLOR_BG_BLUE, []any{NewBGColor(true, ansi.COLOR_BG_RED, []any{"red"}), "blue"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RED) + "red" +
			ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
	})

	It("keeps other groups", func() {
		s := NewHighlight(true, ansi.MODE_BOLD, []any{NewFGColor(true, ansi.COLOR_FG_RED, []any{"red"}), "bold"})
		Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_BOLD) + ansi.Mode(ansi.COLOR_FG_RED) + "red" +
			ansi.Mode(ansi.COLOR_FG_DEFAULT) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
	})

	It("renders nothing if disabled", func() {
		s := NewFGColor(false, ansi.COLOR_FG_RED, []any{NewHighlight(false, ansi.MODE_BOLD, []any{"text"})})
		Expect(s.String()).To(Equal("text"))
	})
})
