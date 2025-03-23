package ttycolors_test

import (
	. "github.com/mandelsoft/ttycolors"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("ColorString Test Environment", func() {
	Context("disabled", func() {
		It("mode", func() {
			s := Bold("bold")
			s.Enable(false)
			Expect(s.String()).To(Equal("bold"))
		})

		It("standard color", func() {
			s := Blue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg color", func() {
			s := BgBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bright color", func() {
			s := BrightBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg bright color", func() {
			s := BgBrightBlue("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb color", func() {
			s := RGBColor(1, 255, 2, "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb bg color", func() {
			s := BgRGBColor(1, 255, 2, "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("id color", func() {
			s := IdColor(1, "id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("id bg color", func() {
			s := BgIdColor(1, "id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("complex color", func() {
			s := Blue("blue", Green("green"), Red("red"), "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue" +
				"green" +
				"red" +
				"blue"))
		})

		It("mixed", func() {
			s := Blue("blue", Bold(Green("bold green"), "bold blue"), Red("red"), "blue")
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
			s := Bold("bold")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("standard color", func() {
			s := Blue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg color", func() {
			s := BgBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("standard bright color", func() {
			s := BrightBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg bright color", func() {
			s := BgBrightBlue("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("rgb color", func() {
			s := RGBColor(1, 255, 2, "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RGB, ansi.Simple(1), ansi.Simple(255), ansi.Simple(2)) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("rgb bg color", func() {
			s := BgRGBColor(1, 255, 2, "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RGB, ansi.Simple(1), ansi.Simple(255), ansi.Simple(2)) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("id color", func() {
			s := IdColor(1, "id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_ID, ansi.Simple(1)) + "id" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("id bg color", func() {
			s := BgIdColor(1, "id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_ID, ansi.Simple(1)) + "id" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("complex color", func() {
			s := Blue("blue", Green("green"), Red("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("complex bg color", func() {
			s := BgBlue("blue", BgGreen("green"), BgRed("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_BG_RED) + "red" +
				ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("mixed color", func() {
			s := Blue("blue", BgWhite("white back"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("mixed bg", func() {
			s := Blue("blue", Bold(Green("bold green"), "bold blue", BgWhite("white back"), BgYellow("yellow back")), Red("red"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.MODE_BOLD) + ansi.Mode(ansi.COLOR_FG_GREEN) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_YELLOW) + "yellow back" +
				ansi.Mode(ansi.MODE_RESET_BOLD) + ansi.Mode(ansi.COLOR_BG_DEFAULT) + ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("mixed bg", func() {
			s := Blue("blue", Bold(Green("bold green", Yellow("bold yellow"), "bold green"), "bold blue", BgWhite("white back"), BgYellow("yellow back")), Red("red"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.MODE_BOLD) + ansi.Mode(ansi.COLOR_FG_GREEN) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_YELLOW) + "bold yellow" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.COLOR_BG_WHITE) + "white back" +
				ansi.Mode(ansi.COLOR_BG_YELLOW) + "yellow back" +
				ansi.Mode(ansi.MODE_RESET_BOLD) + ansi.Mode(ansi.COLOR_BG_DEFAULT) + ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})

		It("alternating", func() {
			s := Blue("blue", BgGreen("green back", Red("red", BgBlack("black back"), "red"), "green back"), "blue")
			s.Enable()
			x := s.String()
			e := ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "green back" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_BG_BLACK) + "black back" +
				ansi.Mode(ansi.COLOR_BG_GREEN) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "green back" +
				ansi.Mode(ansi.COLOR_BG_DEFAULT) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)
			Expect(x).To(Equal(e))
		})
	})

})
