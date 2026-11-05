package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// NonHighlighted provides a formatted string (see ttycolors.NonHighlighted).
func NonHighlighted(s ...any) string {
	return ttycolors.NonHighlighted(s...).String()
}

// Bold provides a formatted string (see ttycolors.Bold).
func Bold(s ...any) string {
	return ttycolors.Bold(s...).String()
}

// Dim provides a formatted string (see ttycolors.Dim).
func Dim(s ...any) string {
	return ttycolors.Dim(s...).String()
}

// Italic provides a formatted string (see ttycolors.Italic).
func Italic(s ...any) string {
	return ttycolors.Italic(s...).String()
}

// Blinking provides a formatted string (see ttycolors.Blinking).
func Blinking(s ...any) string {
	return ttycolors.Blinking(s...).String()
}

// Underline provides a formatted string (see ttycolors.Underline).
func Underline(s ...any) string {
	return ttycolors.Underline(s...).String()
}

// Reverse provides a formatted string (see ttycolors.Reverse).
func Reverse(s ...any) string {
	return ttycolors.Reverse(s...).String()
}

// StrikeThrough provides a formatted string (see ttycolors.StrikeThrough).
func StrikeThrough(s ...any) string {
	return ttycolors.StrikeThrough(s...).String()
}

// Invisible provides a formatted string (see ttycolors.Invisible).
func Invisible(s ...any) string {
	return ttycolors.Invisible(s...).String()
}

// NonItalic provides a formatted string (see ttycolors.NonItalic).
func NonItalic(s ...any) string {
	return ttycolors.NonItalic(s...).String()
}

// NonBlinking provides a formatted string (see ttycolors.NonBlinking).
func NonBlinking(s ...any) string {
	return ttycolors.NonBlinking(s...).String()
}

// NonUnderline provides a formatted string (see ttycolors.NonUnderline).
func NonUnderline(s ...any) string {
	return ttycolors.NonUnderline(s...).String()
}

// NonReverse provides a formatted string (see ttycolors.NonReverse).
func NonReverse(s ...any) string {
	return ttycolors.NonReverse(s...).String()
}

// NonStrikeThrough provides a formatted string (see ttycolors.NonStrikeThrough).
func NonStrikeThrough(s ...any) string {
	return ttycolors.NonStrikeThrough(s...).String()
}

// NonInvisible provides a formatted string (see ttycolors.NonInvisible).
func NonInvisible(s ...any) string {
	return ttycolors.NonInvisible(s...).String()
}

// DoubleUnderline provides a formatted string (see ttycolors.DoubleUnderline).
func DoubleUnderline(s ...any) string {
	return ttycolors.DoubleUnderline(s...).String()
}

// CurlyUnderline provides a formatted string (see ttycolors.CurlyUnderline).
func CurlyUnderline(s ...any) string {
	return ttycolors.CurlyUnderline(s...).String()
}

// DottedUnderline provides a formatted string (see ttycolors.DottedUnderline).
func DottedUnderline(s ...any) string {
	return ttycolors.DottedUnderline(s...).String()
}

// DashedUnderline provides a formatted string (see ttycolors.DashedUnderline).
func DashedUnderline(s ...any) string {
	return ttycolors.DashedUnderline(s...).String()
}
