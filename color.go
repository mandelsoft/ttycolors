package ttycolors

func Black(s ...any) String {
	return defaultContext.Black(s...)
}
func Red(s ...any) String {
	return defaultContext.Red(s...)
}
func Green(s ...any) String {
	return defaultContext.Green(s...)
}
func Yellow(s ...any) String {
	return defaultContext.Yellow(s...)
}
func Blue(s ...any) String {
	return defaultContext.Blue(s...)
}
func Magenta(s ...any) String {
	return defaultContext.Magenta(s...)
}
func Cyan(s ...any) String {
	return defaultContext.Cyan(s...)
}
func White(s ...any) String {
	return defaultContext.White(s...)
}

func BgBlack(s ...any) String {
	return defaultContext.BgBlack(s...)
}
func BgRed(s ...any) String {
	return defaultContext.BgRed(s...)
}
func BgGreen(s ...any) String {
	return defaultContext.BgGreen(s...)
}
func BgYellow(s ...any) String {
	return defaultContext.BgYellow(s...)
}
func BgBlue(s ...any) String {
	return defaultContext.BgBlue(s...)
}
func BgMagenta(s ...any) String {
	return defaultContext.BgMagenta(s...)
}
func BgCyan(s ...any) String {
	return defaultContext.BgCyan(s...)
}
func BgWhite(s ...any) String {
	return defaultContext.BgWhite(s...)
}

func BrightBlack(s ...any) String {
	return defaultContext.BrightBlack(s...)
}
func BrightRed(s ...any) String {
	return defaultContext.BrightRed(s...)
}
func BrightGreen(s ...any) String {
	return defaultContext.BrightGreen(s...)
}
func BrightYellow(s ...any) String {
	return defaultContext.BrightYellow(s...)
}
func BrightBlue(s ...any) String {
	return defaultContext.BrightBlue(s...)
}
func BrightMagenta(s ...any) String {
	return defaultContext.BrightMagenta(s...)
}
func BrightCyan(s ...any) String {
	return defaultContext.BrightCyan(s...)
}
func BrightWhite(s ...any) String {
	return defaultContext.BrightWhite(s...)
}

////////////////////////////////////////////////////////////////////////////////

func BgBrightBlack(s ...any) String {
	return defaultContext.BgBrightBlack(s...)
}
func BgBrightRed(s ...any) String {
	return defaultContext.BgBrightRed(s...)
}
func BgBrightGreen(s ...any) String {
	return defaultContext.BgBrightGreen(s...)
}
func BgBrightYellow(s ...any) String {
	return defaultContext.BgBrightYellow(s...)
}
func BgBrightBlue(s ...any) String {
	return defaultContext.BgBrightBlue(s...)
}
func BgBrightMagenta(s ...any) String {
	return defaultContext.BgBrightMagenta(s...)
}
func BgBrightCyan(s ...any) String {
	return defaultContext.BgBrightCyan(s...)
}
func BgBrightWhite(s ...any) String {
	return defaultContext.BgBrightWhite(s...)
}

func RGBColor(r, g, b byte, s ...any) String {
	return defaultContext.RGBColor(r, g, b, s...)
}
func BgRGBColor(r, g, b byte, s ...any) String {
	return defaultContext.BgRGBColor(r, g, b, s...)
}

func IdColor(id byte, s ...any) String {
	return defaultContext.IdColor(id, s...)
}
func BgIdColor(id byte, s ...any) String {
	return defaultContext.BgIdColor(id, s...)
}
