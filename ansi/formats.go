package ansi

func Code_FG_RGB(r, g, b byte) Code {
	return COLOR_FG_RGB.AppendBytes(r, g, b)
}

func Code_BG_RGB(r, g, b byte) Code {
	return COLOR_BG_RGB.AppendBytes(r, g, b)
}

func Code_FG_ID(id byte) Code {
	return COLOR_FG_ID.AppendBytes(id)
}

func Code_BG_ID(id byte) Code {
	return COLOR_BG_ID.AppendBytes(id)
}
