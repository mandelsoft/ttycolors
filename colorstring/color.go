package colorstring

import (
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/colorstring/renderer"
)

func Black(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BLACK, s)
}
func Red(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_RED, s)
}
func Green(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_GREEN, s)
}
func Yellow(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_YELLOW, s)
}
func Blue(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BLUE, s)
}
func Magenta(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_MAGENTA, s)
}
func Cyan(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_CYAN, s)
}
func White(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_WHITE, s)
}

func BgBlack(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BLACK, s)
}
func BgRed(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_RED, s)
}
func BgGreen(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_GREEN, s)
}
func BgYellow(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_YELLOW, s)
}
func BgBlue(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BLUE, s)
}
func BgMagenta(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_MAGENTA, s)
}
func BgCyan(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_CYAN, s)
}
func BgWhite(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_WHITE, s)
}

func BrightBlack(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_BLACK, s)
}
func BrightRed(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_RED, s)
}
func BrightGreen(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_GREEN, s)
}
func BrightYellow(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_YELLOW, s)
}
func BrightBlue(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_BLUE, s)
}
func BrightMagenta(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_MAGENTA, s)
}
func BrightCyan(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_CYAN, s)
}
func BrightWhite(s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_BRIGHT_WHITE, s)
}

func BgBrightBlack(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_BLACK, s)
}
func BgBrightRed(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_RED, s)
}
func BgBrightGreen(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_GREEN, s)
}
func BgBrightYellow(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_YELLOW, s)
}
func BgBrightBlue(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_BLUE, s)
}
func BgBrightMagenta(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_MAGENTA, s)
}
func BgBrightCyan(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_CYAN, s)
}
func BgBrightWhite(s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_BRIGHT_WHITE, s)
}

func RGBColor(r, g, b byte, s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_RGB.AppendBytes(r, g, b), s)
}
func BgRGBColor(r, g, b byte, s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_RGB.AppendBytes(r, g, b), s)
}

func IdColor(id byte, s ...any) String {
	return renderer.NewFGColor(ansi.COLOR_FG_ID.AppendBytes(id), s)
}
func BgIdColor(id byte, s ...any) String {
	return renderer.NewBGColor(ansi.COLOR_BG_ID.AppendBytes(id), s)
}
