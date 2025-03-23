package ttycolors

func Bold(s ...any) String {
	return defaultContext.Bold(s...)
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
