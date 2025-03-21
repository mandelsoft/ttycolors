package ttycolors_test

import (
	. "github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/colorstring"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Format Test Environment", func() {
	Context("disabled", func() {
		It("mode", func() {
			s := FmtBold.String("bold")
			s.Enable(false)
			Expect(s.String()).To(Equal("bold"))
		})

		It("standard color", func() {
			s := FmtBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg color", func() {
			s := FmtBgBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bright color", func() {
			s := FmtBrightBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("standard bg bright color", func() {
			s := FmtBgBrightBlue.String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb color", func() {
			s := FmtRGBColor(1, 255, 2).String("blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue"))
		})

		It("rgb bg color", func() {
			s := FmtBgRGBColor(1, 255, 2).String("green")
			s.Enable(false)
			Expect(s.String()).To(Equal("green"))
		})

		It("id color", func() {
			s := FmtIdColor(1).String("id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("id bg color", func() {
			s := FmtBgIdColor(1).String("id")
			s.Enable(false)
			Expect(s.String()).To(Equal("id"))
		})

		It("complex color", func() {
			s := FmtBlue.String("blue", colorstring.Green("green"), colorstring.Red("red"), "blue")
			s.Enable(false)
			Expect(s.String()).To(Equal("blue" +
				"green" +
				"red" +
				"blue"))
		})

		It("mixed", func() {
			s := FmtBlue.String("blue", colorstring.Bold(colorstring.Green("bold green"), "bold blue"), colorstring.Red("red"), "blue")
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
			s := FmtBold.String("bold")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.MODE_BOLD) + "bold" + ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("standard color", func() {
			s := FmtBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg color", func() {
			s := FmtBgBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("standard bright color", func() {
			s := FmtBrightBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("standard bg bright color", func() {
			s := FmtBgBrightBlue.String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_BRIGHT_BLUE) + "blue" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("rgb color", func() {
			s := FmtRGBColor(1, 255, 2).String("blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_RGB, ansi.Simple(1), ansi.Simple(255), ansi.Simple(2)) + "blue" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("rgb bg color", func() {
			s := FmtBgRGBColor(1, 255, 2).String("green")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_RGB, ansi.Simple(1), ansi.Simple(255), ansi.Simple(2)) + "green" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("id color", func() {
			s := FmtIdColor(1).String("id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_ID, ansi.Simple(1)) + "id" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("id gb color", func() {
			s := FmtBgIdColor(1).String("id")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_BG_ID, ansi.Simple(1)) + "id" + ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("complex color", func() {
			s := FmtBlue.String("blue", colorstring.Green("green"), colorstring.Red("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_GREEN) + "green" +
				ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})

		It("mixed", func() {
			s := FmtBlue.String("blue", colorstring.Bold(colorstring.Green("bold green"), "bold blue"), colorstring.Red("red"), "blue")
			s.Enable()
			Expect(s.String()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.MODE_BOLD) + ansi.Mode(ansi.COLOR_FG_GREEN) + "bold green" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "bold blue" +
				ansi.Mode(ansi.MODE_RESET_BOLD) + ansi.Mode(ansi.COLOR_FG_RED) + "red" +
				ansi.Mode(ansi.COLOR_FG_BLUE) + "blue" +
				ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		})
	})

})
