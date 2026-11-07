package plain

// UncoloredFG provides a formatted string using the default context
// (see ttycolors.UncoloredFG).
func UncoloredFG(s ...any) string {
	return defaultContext.UncoloredFG(s...)
}

// UncoloredFG provides a formatted string using the context
// (see ttycolors.TTYContext.UncoloredFG).
func (c *_context) UncoloredFG(s ...any) string {
	return c.ctx.UncoloredFG(s...).String()
}

// UncoloredBG provides a formatted string using the default context
// (see ttycolors.UncoloredBG).
func UncoloredBG(s ...any) string {
	return defaultContext.UncoloredBG(s...)
}

// UncoloredBG provides a formatted string using the context
// (see ttycolors.TTYContext.UncoloredBG).
func (c *_context) UncoloredBG(s ...any) string {
	return c.ctx.UncoloredBG(s...).String()
}

// Black provides a formatted string using the default context
// (see ttycolors.Black).
func Black(s ...any) string {
	return defaultContext.Black(s...)
}

// Black provides a formatted string using the context
// (see ttycolors.TTYContext.Black).
func (c *_context) Black(s ...any) string {
	return c.ctx.Black(s...).String()
}

// Red provides a formatted string using the default context
// (see ttycolors.Red).
func Red(s ...any) string {
	return defaultContext.Red(s...)
}

// Red provides a formatted string using the context
// (see ttycolors.TTYContext.Red).
func (c *_context) Red(s ...any) string {
	return c.ctx.Red(s...).String()
}

// Green provides a formatted string using the default context
// (see ttycolors.Green).
func Green(s ...any) string {
	return defaultContext.Green(s...)
}

// Green provides a formatted string using the context
// (see ttycolors.TTYContext.Green).
func (c *_context) Green(s ...any) string {
	return c.ctx.Green(s...).String()
}

// Yellow provides a formatted string using the default context
// (see ttycolors.Yellow).
func Yellow(s ...any) string {
	return defaultContext.Yellow(s...)
}

// Yellow provides a formatted string using the context
// (see ttycolors.TTYContext.Yellow).
func (c *_context) Yellow(s ...any) string {
	return c.ctx.Yellow(s...).String()
}

// Blue provides a formatted string using the default context
// (see ttycolors.Blue).
func Blue(s ...any) string {
	return defaultContext.Blue(s...)
}

// Blue provides a formatted string using the context
// (see ttycolors.TTYContext.Blue).
func (c *_context) Blue(s ...any) string {
	return c.ctx.Blue(s...).String()
}

// Magenta provides a formatted string using the default context
// (see ttycolors.Magenta).
func Magenta(s ...any) string {
	return defaultContext.Magenta(s...)
}

// Magenta provides a formatted string using the context
// (see ttycolors.TTYContext.Magenta).
func (c *_context) Magenta(s ...any) string {
	return c.ctx.Magenta(s...).String()
}

// Cyan provides a formatted string using the default context
// (see ttycolors.Cyan).
func Cyan(s ...any) string {
	return defaultContext.Cyan(s...)
}

// Cyan provides a formatted string using the context
// (see ttycolors.TTYContext.Cyan).
func (c *_context) Cyan(s ...any) string {
	return c.ctx.Cyan(s...).String()
}

// White provides a formatted string using the default context
// (see ttycolors.White).
func White(s ...any) string {
	return defaultContext.White(s...)
}

// White provides a formatted string using the context
// (see ttycolors.TTYContext.White).
func (c *_context) White(s ...any) string {
	return c.ctx.White(s...).String()
}

// BgBlack provides a formatted string using the default context
// (see ttycolors.BgBlack).
func BgBlack(s ...any) string {
	return defaultContext.BgBlack(s...)
}

// BgBlack provides a formatted string using the context
// (see ttycolors.TTYContext.BgBlack).
func (c *_context) BgBlack(s ...any) string {
	return c.ctx.BgBlack(s...).String()
}

// BgRed provides a formatted string using the default context
// (see ttycolors.BgRed).
func BgRed(s ...any) string {
	return defaultContext.BgRed(s...)
}

// BgRed provides a formatted string using the context
// (see ttycolors.TTYContext.BgRed).
func (c *_context) BgRed(s ...any) string {
	return c.ctx.BgRed(s...).String()
}

// BgGreen provides a formatted string using the default context
// (see ttycolors.BgGreen).
func BgGreen(s ...any) string {
	return defaultContext.BgGreen(s...)
}

// BgGreen provides a formatted string using the context
// (see ttycolors.TTYContext.BgGreen).
func (c *_context) BgGreen(s ...any) string {
	return c.ctx.BgGreen(s...).String()
}

// BgYellow provides a formatted string using the default context
// (see ttycolors.BgYellow).
func BgYellow(s ...any) string {
	return defaultContext.BgYellow(s...)
}

// BgYellow provides a formatted string using the context
// (see ttycolors.TTYContext.BgYellow).
func (c *_context) BgYellow(s ...any) string {
	return c.ctx.BgYellow(s...).String()
}

// BgBlue provides a formatted string using the default context
// (see ttycolors.BgBlue).
func BgBlue(s ...any) string {
	return defaultContext.BgBlue(s...)
}

// BgBlue provides a formatted string using the context
// (see ttycolors.TTYContext.BgBlue).
func (c *_context) BgBlue(s ...any) string {
	return c.ctx.BgBlue(s...).String()
}

// BgMagenta provides a formatted string using the default context
// (see ttycolors.BgMagenta).
func BgMagenta(s ...any) string {
	return defaultContext.BgMagenta(s...)
}

// BgMagenta provides a formatted string using the context
// (see ttycolors.TTYContext.BgMagenta).
func (c *_context) BgMagenta(s ...any) string {
	return c.ctx.BgMagenta(s...).String()
}

// BgCyan provides a formatted string using the default context
// (see ttycolors.BgCyan).
func BgCyan(s ...any) string {
	return defaultContext.BgCyan(s...)
}

// BgCyan provides a formatted string using the context
// (see ttycolors.TTYContext.BgCyan).
func (c *_context) BgCyan(s ...any) string {
	return c.ctx.BgCyan(s...).String()
}

// BgWhite provides a formatted string using the default context
// (see ttycolors.BgWhite).
func BgWhite(s ...any) string {
	return defaultContext.BgWhite(s...)
}

// BgWhite provides a formatted string using the context
// (see ttycolors.TTYContext.BgWhite).
func (c *_context) BgWhite(s ...any) string {
	return c.ctx.BgWhite(s...).String()
}

// BrightBlack provides a formatted string using the default context
// (see ttycolors.BrightBlack).
func BrightBlack(s ...any) string {
	return defaultContext.BrightBlack(s...)
}

// BrightBlack provides a formatted string using the context
// (see ttycolors.TTYContext.BrightBlack).
func (c *_context) BrightBlack(s ...any) string {
	return c.ctx.BrightBlack(s...).String()
}

// BrightRed provides a formatted string using the default context
// (see ttycolors.BrightRed).
func BrightRed(s ...any) string {
	return defaultContext.BrightRed(s...)
}

// BrightRed provides a formatted string using the context
// (see ttycolors.TTYContext.BrightRed).
func (c *_context) BrightRed(s ...any) string {
	return c.ctx.BrightRed(s...).String()
}

// BrightGreen provides a formatted string using the default context
// (see ttycolors.BrightGreen).
func BrightGreen(s ...any) string {
	return defaultContext.BrightGreen(s...)
}

// BrightGreen provides a formatted string using the context
// (see ttycolors.TTYContext.BrightGreen).
func (c *_context) BrightGreen(s ...any) string {
	return c.ctx.BrightGreen(s...).String()
}

// BrightYellow provides a formatted string using the default context
// (see ttycolors.BrightYellow).
func BrightYellow(s ...any) string {
	return defaultContext.BrightYellow(s...)
}

// BrightYellow provides a formatted string using the context
// (see ttycolors.TTYContext.BrightYellow).
func (c *_context) BrightYellow(s ...any) string {
	return c.ctx.BrightYellow(s...).String()
}

// BrightBlue provides a formatted string using the default context
// (see ttycolors.BrightBlue).
func BrightBlue(s ...any) string {
	return defaultContext.BrightBlue(s...)
}

// BrightBlue provides a formatted string using the context
// (see ttycolors.TTYContext.BrightBlue).
func (c *_context) BrightBlue(s ...any) string {
	return c.ctx.BrightBlue(s...).String()
}

// BrightMagenta provides a formatted string using the default context
// (see ttycolors.BrightMagenta).
func BrightMagenta(s ...any) string {
	return defaultContext.BrightMagenta(s...)
}

// BrightMagenta provides a formatted string using the context
// (see ttycolors.TTYContext.BrightMagenta).
func (c *_context) BrightMagenta(s ...any) string {
	return c.ctx.BrightMagenta(s...).String()
}

// BrightCyan provides a formatted string using the default context
// (see ttycolors.BrightCyan).
func BrightCyan(s ...any) string {
	return defaultContext.BrightCyan(s...)
}

// BrightCyan provides a formatted string using the context
// (see ttycolors.TTYContext.BrightCyan).
func (c *_context) BrightCyan(s ...any) string {
	return c.ctx.BrightCyan(s...).String()
}

// BrightWhite provides a formatted string using the default context
// (see ttycolors.BrightWhite).
func BrightWhite(s ...any) string {
	return defaultContext.BrightWhite(s...)
}

// BrightWhite provides a formatted string using the context
// (see ttycolors.TTYContext.BrightWhite).
func (c *_context) BrightWhite(s ...any) string {
	return c.ctx.BrightWhite(s...).String()
}

// BgBrightBlack provides a formatted string using the default context
// (see ttycolors.BgBrightBlack).
func BgBrightBlack(s ...any) string {
	return defaultContext.BgBrightBlack(s...)
}

// BgBrightBlack provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightBlack).
func (c *_context) BgBrightBlack(s ...any) string {
	return c.ctx.BgBrightBlack(s...).String()
}

// BgBrightRed provides a formatted string using the default context
// (see ttycolors.BgBrightRed).
func BgBrightRed(s ...any) string {
	return defaultContext.BgBrightRed(s...)
}

// BgBrightRed provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightRed).
func (c *_context) BgBrightRed(s ...any) string {
	return c.ctx.BgBrightRed(s...).String()
}

// BgBrightGreen provides a formatted string using the default context
// (see ttycolors.BgBrightGreen).
func BgBrightGreen(s ...any) string {
	return defaultContext.BgBrightGreen(s...)
}

// BgBrightGreen provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightGreen).
func (c *_context) BgBrightGreen(s ...any) string {
	return c.ctx.BgBrightGreen(s...).String()
}

// BgBrightYellow provides a formatted string using the default context
// (see ttycolors.BgBrightYellow).
func BgBrightYellow(s ...any) string {
	return defaultContext.BgBrightYellow(s...)
}

// BgBrightYellow provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightYellow).
func (c *_context) BgBrightYellow(s ...any) string {
	return c.ctx.BgBrightYellow(s...).String()
}

// BgBrightBlue provides a formatted string using the default context
// (see ttycolors.BgBrightBlue).
func BgBrightBlue(s ...any) string {
	return defaultContext.BgBrightBlue(s...)
}

// BgBrightBlue provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightBlue).
func (c *_context) BgBrightBlue(s ...any) string {
	return c.ctx.BgBrightBlue(s...).String()
}

// BgBrightMagenta provides a formatted string using the default context
// (see ttycolors.BgBrightMagenta).
func BgBrightMagenta(s ...any) string {
	return defaultContext.BgBrightMagenta(s...)
}

// BgBrightMagenta provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightMagenta).
func (c *_context) BgBrightMagenta(s ...any) string {
	return c.ctx.BgBrightMagenta(s...).String()
}

// BgBrightCyan provides a formatted string using the default context
// (see ttycolors.BgBrightCyan).
func BgBrightCyan(s ...any) string {
	return defaultContext.BgBrightCyan(s...)
}

// BgBrightCyan provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightCyan).
func (c *_context) BgBrightCyan(s ...any) string {
	return c.ctx.BgBrightCyan(s...).String()
}

// BgBrightWhite provides a formatted string using the default context
// (see ttycolors.BgBrightWhite).
func BgBrightWhite(s ...any) string {
	return defaultContext.BgBrightWhite(s...)
}

// BgBrightWhite provides a formatted string using the context
// (see ttycolors.TTYContext.BgBrightWhite).
func (c *_context) BgBrightWhite(s ...any) string {
	return c.ctx.BgBrightWhite(s...).String()
}

// RGBColor provides a formatted string using the default context
// (see ttycolors.RGBColor).
func RGBColor(r, g, b byte, s ...any) string {
	return defaultContext.RGBColor(r, g, b, s...)
}

// RGBColor provides a formatted string using the context
// (see ttycolors.TTYContext.RGBColor).
func (c *_context) RGBColor(r, g, b byte, s ...any) string {
	return c.ctx.RGBColor(r, g, b, s...).String()
}

// BgRGBColor provides a formatted string using the default context
// (see ttycolors.BgRGBColor).
func BgRGBColor(r, g, b byte, s ...any) string {
	return defaultContext.BgRGBColor(r, g, b, s...)
}

// BgRGBColor provides a formatted string using the context
// (see ttycolors.TTYContext.BgRGBColor).
func (c *_context) BgRGBColor(r, g, b byte, s ...any) string {
	return c.ctx.BgRGBColor(r, g, b, s...).String()
}

// IdColor provides a formatted string using the default context
// (see ttycolors.IdColor).
func IdColor(id byte, s ...any) string {
	return defaultContext.IdColor(id, s...)
}

// IdColor provides a formatted string using the context
// (see ttycolors.TTYContext.IdColor).
func (c *_context) IdColor(id byte, s ...any) string {
	return c.ctx.IdColor(id, s...).String()
}

// BgIdColor provides a formatted string using the default context
// (see ttycolors.BgIdColor).
func BgIdColor(id byte, s ...any) string {
	return defaultContext.BgIdColor(id, s...)
}

// BgIdColor provides a formatted string using the context
// (see ttycolors.TTYContext.BgIdColor).
func (c *_context) BgIdColor(id byte, s ...any) string {
	return c.ctx.BgIdColor(id, s...).String()
}

// UnderlineIdColor provides a formatted string using the default context
// (see ttycolors.UnderlineIdColor).
func UnderlineIdColor(id byte, s ...any) string {
	return defaultContext.UnderlineIdColor(id, s...)
}

// UnderlineIdColor provides a formatted string using the context
// (see ttycolors.TTYContext.UnderlineIdColor).
func (c *_context) UnderlineIdColor(id byte, s ...any) string {
	return c.ctx.UnderlineIdColor(id, s...).String()
}

// UnderlineRGBColor provides a formatted string using the default context
// (see ttycolors.UnderlineRGBColor).
func UnderlineRGBColor(r, g, b byte, s ...any) string {
	return defaultContext.UnderlineRGBColor(r, g, b, s...)
}

// UnderlineRGBColor provides a formatted string using the context
// (see ttycolors.TTYContext.UnderlineRGBColor).
func (c *_context) UnderlineRGBColor(r, g, b byte, s ...any) string {
	return c.ctx.UnderlineRGBColor(r, g, b, s...).String()
}
