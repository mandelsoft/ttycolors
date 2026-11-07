package plain

// Sequence provides a formatted string using the default context
// (see ttycolors.Sequence).
func Sequence(seq ...any) string {
	return defaultContext.Sequence(seq...)
}

// Sequence provides a formatted string using the context
// (see ttycolors.TTYContext.Sequence).
func (c *_context) Sequence(seq ...any) string {
	return c.ctx.Sequence(seq...).String()
}
