package plain

// NonHighlighted provides a formatted string using the default context
// (see ttycolors.NonHighlighted).
func NonHighlighted(s ...any) string {
	return defaultContext.NonHighlighted(s...)
}

// NonHighlighted provides a formatted string using the context
// (see ttycolors.TTYContext.NonHighlighted).
func (c *_context) NonHighlighted(s ...any) string {
	return c.ctx.NonHighlighted(s...).String()
}

// Bold provides a formatted string using the default context
// (see ttycolors.Bold).
func Bold(s ...any) string {
	return defaultContext.Bold(s...)
}

// Bold provides a formatted string using the context
// (see ttycolors.TTYContext.Bold).
func (c *_context) Bold(s ...any) string {
	return c.ctx.Bold(s...).String()
}

// Dim provides a formatted string using the default context
// (see ttycolors.Dim).
func Dim(s ...any) string {
	return defaultContext.Dim(s...)
}

// Dim provides a formatted string using the context
// (see ttycolors.TTYContext.Dim).
func (c *_context) Dim(s ...any) string {
	return c.ctx.Dim(s...).String()
}

// Italic provides a formatted string using the default context
// (see ttycolors.Italic).
func Italic(s ...any) string {
	return defaultContext.Italic(s...)
}

// Italic provides a formatted string using the context
// (see ttycolors.TTYContext.Italic).
func (c *_context) Italic(s ...any) string {
	return c.ctx.Italic(s...).String()
}

// Blinking provides a formatted string using the default context
// (see ttycolors.Blinking).
func Blinking(s ...any) string {
	return defaultContext.Blinking(s...)
}

// Blinking provides a formatted string using the context
// (see ttycolors.TTYContext.Blinking).
func (c *_context) Blinking(s ...any) string {
	return c.ctx.Blinking(s...).String()
}

// Underline provides a formatted string using the default context
// (see ttycolors.Underline).
func Underline(s ...any) string {
	return defaultContext.Underline(s...)
}

// Underline provides a formatted string using the context
// (see ttycolors.TTYContext.Underline).
func (c *_context) Underline(s ...any) string {
	return c.ctx.Underline(s...).String()
}

// Reverse provides a formatted string using the default context
// (see ttycolors.Reverse).
func Reverse(s ...any) string {
	return defaultContext.Reverse(s...)
}

// Reverse provides a formatted string using the context
// (see ttycolors.TTYContext.Reverse).
func (c *_context) Reverse(s ...any) string {
	return c.ctx.Reverse(s...).String()
}

// StrikeThrough provides a formatted string using the default context
// (see ttycolors.StrikeThrough).
func StrikeThrough(s ...any) string {
	return defaultContext.StrikeThrough(s...)
}

// StrikeThrough provides a formatted string using the context
// (see ttycolors.TTYContext.StrikeThrough).
func (c *_context) StrikeThrough(s ...any) string {
	return c.ctx.StrikeThrough(s...).String()
}

// Overline provides a formatted string using the default context
// (see ttycolors.Overline).
func Overline(s ...any) string {
	return defaultContext.Overline(s...)
}

// Overline provides a formatted string using the context
// (see ttycolors.TTYContext.Overline).
func (c *_context) Overline(s ...any) string {
	return c.ctx.Overline(s...).String()
}

// Invisible provides a formatted string using the default context
// (see ttycolors.Invisible).
func Invisible(s ...any) string {
	return defaultContext.Invisible(s...)
}

// Invisible provides a formatted string using the context
// (see ttycolors.TTYContext.Invisible).
func (c *_context) Invisible(s ...any) string {
	return c.ctx.Invisible(s...).String()
}

// NonItalic provides a formatted string using the default context
// (see ttycolors.NonItalic).
func NonItalic(s ...any) string {
	return defaultContext.NonItalic(s...)
}

// NonItalic provides a formatted string using the context
// (see ttycolors.TTYContext.NonItalic).
func (c *_context) NonItalic(s ...any) string {
	return c.ctx.NonItalic(s...).String()
}

// NonBlinking provides a formatted string using the default context
// (see ttycolors.NonBlinking).
func NonBlinking(s ...any) string {
	return defaultContext.NonBlinking(s...)
}

// NonBlinking provides a formatted string using the context
// (see ttycolors.TTYContext.NonBlinking).
func (c *_context) NonBlinking(s ...any) string {
	return c.ctx.NonBlinking(s...).String()
}

// NonUnderline provides a formatted string using the default context
// (see ttycolors.NonUnderline).
func NonUnderline(s ...any) string {
	return defaultContext.NonUnderline(s...)
}

// NonUnderline provides a formatted string using the context
// (see ttycolors.TTYContext.NonUnderline).
func (c *_context) NonUnderline(s ...any) string {
	return c.ctx.NonUnderline(s...).String()
}

// NonReverse provides a formatted string using the default context
// (see ttycolors.NonReverse).
func NonReverse(s ...any) string {
	return defaultContext.NonReverse(s...)
}

// NonReverse provides a formatted string using the context
// (see ttycolors.TTYContext.NonReverse).
func (c *_context) NonReverse(s ...any) string {
	return c.ctx.NonReverse(s...).String()
}

// NonStrikeThrough provides a formatted string using the default context
// (see ttycolors.NonStrikeThrough).
func NonStrikeThrough(s ...any) string {
	return defaultContext.NonStrikeThrough(s...)
}

// NonStrikeThrough provides a formatted string using the context
// (see ttycolors.TTYContext.NonStrikeThrough).
func (c *_context) NonStrikeThrough(s ...any) string {
	return c.ctx.NonStrikeThrough(s...).String()
}

// NonOverline provides a formatted string using the default context
// (see ttycolors.NonOverline).
func NonOverline(s ...any) string {
	return defaultContext.NonOverline(s...)
}

// NonOverline provides a formatted string using the context
// (see ttycolors.TTYContext.NonOverline).
func (c *_context) NonOverline(s ...any) string {
	return c.ctx.NonOverline(s...).String()
}

// NonInvisible provides a formatted string using the default context
// (see ttycolors.NonInvisible).
func NonInvisible(s ...any) string {
	return defaultContext.NonInvisible(s...)
}

// NonInvisible provides a formatted string using the context
// (see ttycolors.TTYContext.NonInvisible).
func (c *_context) NonInvisible(s ...any) string {
	return c.ctx.NonInvisible(s...).String()
}

// DoubleUnderline provides a formatted string using the default context
// (see ttycolors.DoubleUnderline).
func DoubleUnderline(s ...any) string {
	return defaultContext.DoubleUnderline(s...)
}

// DoubleUnderline provides a formatted string using the context
// (see ttycolors.TTYContext.DoubleUnderline).
func (c *_context) DoubleUnderline(s ...any) string {
	return c.ctx.DoubleUnderline(s...).String()
}

// CurlyUnderline provides a formatted string using the default context
// (see ttycolors.CurlyUnderline).
func CurlyUnderline(s ...any) string {
	return defaultContext.CurlyUnderline(s...)
}

// CurlyUnderline provides a formatted string using the context
// (see ttycolors.TTYContext.CurlyUnderline).
func (c *_context) CurlyUnderline(s ...any) string {
	return c.ctx.CurlyUnderline(s...).String()
}

// DottedUnderline provides a formatted string using the default context
// (see ttycolors.DottedUnderline).
func DottedUnderline(s ...any) string {
	return defaultContext.DottedUnderline(s...)
}

// DottedUnderline provides a formatted string using the context
// (see ttycolors.TTYContext.DottedUnderline).
func (c *_context) DottedUnderline(s ...any) string {
	return c.ctx.DottedUnderline(s...).String()
}

// DashedUnderline provides a formatted string using the default context
// (see ttycolors.DashedUnderline).
func DashedUnderline(s ...any) string {
	return defaultContext.DashedUnderline(s...)
}

// DashedUnderline provides a formatted string using the context
// (see ttycolors.TTYContext.DashedUnderline).
func (c *_context) DashedUnderline(s ...any) string {
	return c.ctx.DashedUnderline(s...).String()
}
