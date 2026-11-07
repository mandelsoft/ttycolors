package ttycolors

import (
	"fmt"

	"github.com/mandelsoft/ttycolors/renderer"
)

// SplitAt splits a string containing escape sequences after n
// characters (runes). Like String.SplitAt, it keeps the output format:
// both parts are self-contained and can be embedded or formatted
// independently. The first part resets the output formats still set at
// the split position, the second part sets them again and resets the
// formats explicitly reset before. Output formats untouched up to the
// split position are not affected, so that the text still inherits them
// from an outer format. Output format sequences directly at the split
// position are replaced by these codes, other escape sequences there
// belong to the first part.
//
// It panics for a negative n, an n beyond the end provides an empty
// second part. Codes not affecting an output format (for example fonts)
// are not tracked. For splitting strings without considering output
// formats, see ansi.SplitAt.
func SplitAt(s string, n int) (string, string) {
	return renderer.SplitText(s, n)
}

// StripIfDisabled removes the output format sequences from a text,
// if output formatting is disabled for the default context
// (see TTYContext.StripIfDisabled).
func StripIfDisabled(s string) string {
	return defaultContext.StripIfDisabled(s)
}

// SubString extracts the characters (runes) from index start to end-1
// of a string containing escape sequences. Like the parts provided by
// SplitAt, the result is self-contained. If start equals end, the result
// is empty. It panics for negative indices and if end is smaller than start.
func SubString(s string, start, end int) string {
	if start < 0 || end < start {
		panic(fmt.Sprintf("ttycolors.SubString: invalid range %d:%d", start, end))
	}
	if start == end {
		return ""
	}
	_, r := SplitAt(s, start)
	l, _ := SplitAt(r, end-start)
	return l
}
