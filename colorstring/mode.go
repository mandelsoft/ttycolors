package colorstring

import (
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/colorstring/renderer"
)

func Bold(s ...any) String {
	return renderer.NewMode(ansi.MODE_BOLD, ansi.MODE_RESET_BOLD, s)
}

func Italic(s ...any) String {
	return renderer.NewMode(ansi.MODE_ITALIC, ansi.MODE_RESET_ITALIC, s)
}

func Blinking(s ...any) String {
	return renderer.NewMode(ansi.MODE_BLINK, ansi.MODE_RESET_BLINK, s)
}

func Underline(s ...any) String {
	return renderer.NewMode(ansi.MODE_UNDERLINE, ansi.MODE_RESET_UNDERLINE, s)
}

func Reverse(s ...any) String {
	return renderer.NewMode(ansi.MODE_REVERSE, ansi.MODE_RESET_REVERSE, s)
}

func StrikeThrough(s ...any) String {
	return renderer.NewMode(ansi.MODE_STRIKE, ansi.MODE_RESET_STRIKE, s)
}

func Invisible(s ...any) String {
	return renderer.NewMode(ansi.MODE_INVIS, ansi.MODE_RESET_INVIS, s)
}
