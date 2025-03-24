package ttycolors

import (
	"github.com/mandelsoft/goutils/general"
	"github.com/mandelsoft/goutils/generics"
	"github.com/mandelsoft/goutils/optionutils"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/renderer"
)

var NoColors = false

type TTYContext = *_context

var defaultContext = NewContext()

type _context struct {
	disabled *bool
}

var _ ContextInfo = (*_context)(nil)

// NewContext provides a new output context
// with enabled or disabled colors.
// The default setting is taken from NoColors.
func NewContext(b ...bool) TTYContext {
	if len(b) == 0 {
		return &_context{disabled: &NoColors}
	}
	return &_context{disabled: generics.PointerTo(!general.Optional(b...))}
}

func (c *_context) IsEnabled() bool {
	return !*c.disabled
}

func (c *_context) Enable(b ...bool) {
	c.disabled = generics.PointerTo(!optionutils.BoolOption(b...))
}

////////////////////////////////////////////////////////////////////////////////

// New provides a new composed output format.
func (c *_context) New(f ...FormatProvider) FormatInfo {
	return newFormat(!*c.disabled, f...)
}

// NewFormatter provides a new Formatter, which
// offers the usable Print functions based on
// the context state.
func (c *_context) NewFormatter(f ...FormatProvider) Formatter {
	return newFormatter(!*c.disabled, f...)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) Sequence(seq ...any) String {
	return renderer.Sequence(!*c.disabled, seq)
}

func (c *_context) String(seq ...any) String {
	return renderer.Sequence(!*c.disabled, seq)
}

func (c *_context) StringWith(f FormatProvider, seq ...any) String {
	return c.Sequence(f.Format().String(seq...))
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) Bold(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_BOLD, ansi.MODE_RESET_BOLD, s)
}
func (c *_context) Italic(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_ITALIC, ansi.MODE_RESET_ITALIC, s)
}
func (c *_context) Blinking(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_BLINK, ansi.MODE_RESET_BLINK, s)
}
func (c *_context) Underline(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_UNDERLINE, ansi.MODE_RESET_UNDERLINE, s)
}
func (c *_context) Reverse(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_REVERSE, ansi.MODE_RESET_REVERSE, s)
}
func (c *_context) StrikeThrough(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_STRIKE, ansi.MODE_RESET_STRIKE, s)
}
func (c *_context) Invisible(s ...any) String {
	return renderer.NewMode(!*c.disabled, ansi.MODE_INVIS, ansi.MODE_RESET_INVIS, s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) Black(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BLACK, s)
}
func (c *_context) Red(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_RED, s)
}
func (c *_context) Green(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_GREEN, s)
}
func (c *_context) Yellow(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_YELLOW, s)
}
func (c *_context) Blue(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BLUE, s)
}
func (c *_context) Magenta(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_MAGENTA, s)
}
func (c *_context) Cyan(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_CYAN, s)
}
func (c *_context) White(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_WHITE, s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BgBlack(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BLACK, s)
}
func (c *_context) BgRed(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_RED, s)
}
func (c *_context) BgGreen(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_GREEN, s)
}
func (c *_context) BgYellow(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_YELLOW, s)
}
func (c *_context) BgBlue(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BLUE, s)
}
func (c *_context) BgMagenta(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_MAGENTA, s)
}
func (c *_context) BgCyan(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_CYAN, s)
}
func (c *_context) BgWhite(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_WHITE, s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BrightBlack(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_BLACK, s)
}
func (c *_context) BrightRed(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_RED, s)
}
func (c *_context) BrightGreen(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_GREEN, s)
}
func (c *_context) BrightYellow(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_YELLOW, s)
}
func (c *_context) BrightBlue(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_BLUE, s)
}
func (c *_context) BrightMagenta(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_MAGENTA, s)
}
func (c *_context) BrightCyan(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_CYAN, s)
}
func (c *_context) BrightWhite(s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_BRIGHT_WHITE, s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BgBrightBlack(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_BLACK, s)
}
func (c *_context) BgBrightRed(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_RED, s)
}
func (c *_context) BgBrightGreen(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_GREEN, s)
}
func (c *_context) BgBrightYellow(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_YELLOW, s)
}
func (c *_context) BgBrightBlue(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_BLUE, s)
}
func (c *_context) BgBrightMagenta(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_MAGENTA, s)
}
func (c *_context) BgBrightCyan(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_CYAN, s)
}
func (c *_context) BgBrightWhite(s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_BRIGHT_WHITE, s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) RGBColor(r, g, b byte, s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_RGB.AppendBytes(r, g, b), s)
}
func (c *_context) BgRGBColor(r, g, b byte, s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_RGB.AppendBytes(r, g, b), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) IdColor(id byte, s ...any) String {
	return renderer.NewFGColor(!*c.disabled, ansi.COLOR_FG_ID.AppendBytes(id), s)
}
func (c *_context) BgIdColor(id byte, s ...any) String {
	return renderer.NewBGColor(!*c.disabled, ansi.COLOR_BG_ID.AppendBytes(id), s)
}
