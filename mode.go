package ttycolors

// NonHighlighted provides a string in the standard font
// without bold or dim.
func NonHighlighted(s ...any) String {
	return defaultContext.NonHighlighted(s...)
}

func Bold(s ...any) String {
	return defaultContext.Bold(s...)
}
func Dim(s ...any) String {
	return defaultContext.Dim(s...)
}
func Italic(s ...any) String {
	return defaultContext.Italic(s...)
}
func Blinking(s ...any) String {
	return defaultContext.Blinking(s...)
}
func Underline(s ...any) String {
	return defaultContext.Underline(s...)
}
func Reverse(s ...any) String {
	return defaultContext.Reverse(s...)
}
func StrikeThrough(s ...any) String {
	return defaultContext.StrikeThrough(s...)
}
func Invisible(s ...any) String {
	return defaultContext.Invisible(s...)
}
func NonItalic(s ...any) String {
	return defaultContext.NonItalic(s...)
}
func NonBlinking(s ...any) String {
	return defaultContext.NonBlinking(s...)
}
func NonUnderline(s ...any) String {
	return defaultContext.NonUnderline(s...)
}
func NonReverse(s ...any) String {
	return defaultContext.NonReverse(s...)
}
func NonStrikeThrough(s ...any) String {
	return defaultContext.NonStrikeThrough(s...)
}
func NonInvisible(s ...any) String {
	return defaultContext.NonInvisible(s...)
}
func DoubleUnderline(s ...any) String {
	return defaultContext.DoubleUnderline(s...)
}
func CurlyUnderline(s ...any) String {
	return defaultContext.CurlyUnderline(s...)
}

// DottedUnderline provides a string with a dotted underline.
// It requires the feature FeatureUnderlineStyles, otherwise a single
// underline is used. Some terminals render it like a curly underline.
func DottedUnderline(s ...any) String {
	return defaultContext.DottedUnderline(s...)
}

// DashedUnderline provides a string with a dashed underline.
// It requires the feature FeatureUnderlineStyles, otherwise a single
// underline is used. Some terminals render it like a curly underline.
func DashedUnderline(s ...any) String {
	return defaultContext.DashedUnderline(s...)
}
