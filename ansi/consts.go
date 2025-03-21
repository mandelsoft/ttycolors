package ansi

const CSI = "\x1b["

const (
	MODE_RESET_ALL       Simple = 0
	MODE_RESET_BOLD      Simple = 22
	MODE_RESET_ITALIC    Simple = 23
	MODE_RESET_UNDERLINE Simple = 24
	MODE_RESET_BLINK     Simple = 25
	MODE_RESET_REVERSE   Simple = 27
	MODE_RESET_INVIS     Simple = 28
	MODE_RESET_STRIKE    Simple = 29

	MODE_BOLD      Simple = 1
	MODE_ITALIC    Simple = 3
	MODE_UNDERLINE Simple = 4
	MODE_BLINK     Simple = 5
	MODE_REVERSE   Simple = 7
	MODE_INVIS     Simple = 8
	MODE_STRIKE    Simple = 9
)

const (
	COLOR_FG_BLACK Simple = 30 + iota
	COLOR_FG_RED
	COLOR_FG_GREEN
	COLOR_FG_YELLOW
	COLOR_FG_BLUE
	COLOR_FG_MAGENTA
	COLOR_FG_CYAN
	COLOR_FG_WHITE
	COLOR_FG
	COLOR_FG_DEFAULT
)

const (
	COLOR_BG_BLACK Simple = 40 + iota
	COLOR_BG_RED
	COLOR_BG_GREEN
	COLOR_BG_YELLOW
	COLOR_BG_BLUE
	COLOR_BG_MAGENTA
	COLOR_BG_CYAN
	COLOR_BG_WHITE
	COLOR_BG
	COLOR_BG_DEFAULT
)

const (
	COLOR_FG_BRIGHT_BLACK Simple = 90 + iota
	COLOR_FG_BRIGHT_RED
	COLOR_FG_BRIGHT_GREEN
	COLOR_FG_BRIGHT_YELLOW
	COLOR_FG_BRIGHT_BLUE
	COLOR_FG_BRIGHT_MAGENTA
	COLOR_FG_BRIGHT_CYAN
	COLOR_FG_BRIGHT_WHITE
)

const (
	COLOR_BG_BRIGHT_BLACK Simple = 100 + iota
	COLOR_BG_BRIGHT_RED
	COLOR_BG_BRIGHT_GREEN
	COLOR_BG_BRIGHT_YELLOW
	COLOR_BG_BRIGHT_BLUE
	COLOR_BG_BRIGHT_MAGENTA
	COLOR_BG_BRIGHT_CYAN
	COLOR_BG_BRIGHT_WHITE
)

const (
	COLOR_ID  = 5
	COLOR_RGB = 2
)

var (
	COLOR_FG_ID  = Sequence{byte(COLOR_FG), COLOR_ID}
	COLOR_BG_ID  = Sequence{byte(COLOR_BG), COLOR_ID}
	COLOR_FG_RGB = Sequence{byte(COLOR_FG), COLOR_RGB}
	COLOR_BG_RGB = Sequence{byte(COLOR_BG), COLOR_RGB}
)
