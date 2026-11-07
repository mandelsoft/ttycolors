package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// SplitAt splits a string containing escape sequences after n
// characters (runes) into self-contained parts keeping their output
// format (see ttycolors.SplitAt).
func SplitAt(s string, n int) (string, string) {
	return ttycolors.SplitAt(s, n)
}

// StripIfDisabled removes the output format sequences from a string,
// if output formatting is disabled for the default context of package
// ttycolors. Otherwise, the string is returned unchanged
// (see ttycolors.StripIfDisabled). The functions of this package never
// modify their arguments.
func StripIfDisabled(s string) string {
	return defaultContext.StripIfDisabled(s)
}

// StripIfDisabled removes the output format sequences from a string,
// if output formatting is disabled for the context. Otherwise, the
// string is returned unchanged (see ttycolors.TTYContext.StripIfDisabled).
func (c *_context) StripIfDisabled(s string) string {
	return c.ctx.StripIfDisabled(s)
}

// SubString extracts the characters (runes) from index start to end-1
// of a string containing escape sequences as self-contained string
// (see ttycolors.SubString).
func SubString(s string, start, end int) string {
	return ttycolors.SubString(s, start, end)
}
