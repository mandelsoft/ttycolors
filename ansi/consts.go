package ansi

// The origin of the constants is given in brackets
// (see section Standards of the package documentation):
//   - [ECMA-48]: ECMA-48 (ISO/IEC 6429, successor of ANSI X3.64)
//   - [T.416]: ITU-T T.416 (ISO/IEC 8613-6)
//   - [xterm], [aixterm], [kitty]: de-facto standards introduced by
//     these terminals and adopted by most other terminal emulators.

// CSI is the 7-bit representation of the Control Sequence Introducer [ECMA-48].
const CSI = "\x1b["

const (
	// PARAM_SEPARATOR separates the parameters of a control sequence [ECMA-48].
	PARAM_SEPARATOR = ';'
	// SUBPARAM_SEPARATOR separates the sub parameters of a parameter [ECMA-48].
	SUBPARAM_SEPARATOR = ':'
)

// Graphic rendition modes (parameters of SGR) [ECMA-48].
const (
	// MODE_RESET_ALL resets all modes to the default rendition [ECMA-48].
	MODE_RESET_ALL Param = 0
	// MODE_RESET_BOLD resets to normal intensity, neither bold nor faint [ECMA-48].
	MODE_RESET_BOLD Param = 22
	// MODE_RESET_ITALIC resets italic (and Fraktur) [ECMA-48].
	MODE_RESET_ITALIC Param = 23
	// MODE_RESET_UNDERLINE resets single and double underline [ECMA-48]
	// and all underline styles [kitty].
	MODE_RESET_UNDERLINE Param = 24
	// MODE_RESET_BLINK resets slow and rapid blinking (steady) [ECMA-48].
	MODE_RESET_BLINK Param = 25
	// MODE_RESET_REVERSE resets reverse video (positive image) [ECMA-48].
	MODE_RESET_REVERSE Param = 27
	// MODE_RESET_INVIS resets invisibility (revealed characters) [ECMA-48].
	MODE_RESET_INVIS Param = 28
	// MODE_RESET_STRIKE resets strike through (not crossed-out) [ECMA-48].
	MODE_RESET_STRIKE Param = 29

	// MODE_BOLD sets bold or increased intensity [ECMA-48].
	MODE_BOLD Param = 1
	// MODE_DIM sets faint or decreased intensity [ECMA-48].
	MODE_DIM Param = 2
	// MODE_ITALIC sets italic [ECMA-48].
	MODE_ITALIC Param = 3
	// MODE_UNDERLINE sets single underline [ECMA-48].
	MODE_UNDERLINE Param = 4
	// MODE_BLINK sets slow blinking, less than 150 per minute [ECMA-48].
	MODE_BLINK Param = 5
	// MODE_REVERSE sets reverse video (negative image) [ECMA-48].
	MODE_REVERSE Param = 7
	// MODE_INVIS sets invisible (concealed) characters [ECMA-48].
	MODE_INVIS Param = 8
	// MODE_STRIKE sets strike through (crossed-out characters) [ECMA-48].
	MODE_STRIKE Param = 9
	// MODE_OVERLINE sets overline [ECMA-48]. It is supported by most
	// modern terminals, but not, for example, by xterm and the Linux console.
	MODE_OVERLINE Param = 53
	// MODE_RESET_OVERLINE resets overline (not overlined) [ECMA-48].
	MODE_RESET_OVERLINE Param = 55

	// MODE_RAPID_BLINK sets rapid blinking, 150 per minute or more [ECMA-48].
	// It is rarely supported, most terminals ignore it or render it
	// like MODE_BLINK. It is reset by MODE_RESET_BLINK.
	MODE_RAPID_BLINK Param = 6
	// MODE_DOUBLE_UNDERLINE sets double underline [ECMA-48].
	// Some terminals, for example the Linux console, interpret it as
	// resetting bold, instead. It is reset by MODE_RESET_UNDERLINE.
	// UNDERLINE_DOUBLE should be preferred.
	MODE_DOUBLE_UNDERLINE Param = 21
)

// Foreground colors [ECMA-48].
const (
	COLOR_FG_BLACK Param = 30 + iota
	COLOR_FG_RED
	COLOR_FG_GREEN
	COLOR_FG_YELLOW
	COLOR_FG_BLUE
	COLOR_FG_MAGENTA
	COLOR_FG_CYAN
	COLOR_FG_WHITE
	// COLOR_FG introduces an extended foreground color, see COLOR_ID
	// and COLOR_RGB. It is reserved by [ECMA-48] for the color format
	// defined by [T.416].
	COLOR_FG
	// COLOR_FG_DEFAULT resets the foreground color [ECMA-48].
	COLOR_FG_DEFAULT
)

// Background colors [ECMA-48].
const (
	COLOR_BG_BLACK Param = 40 + iota
	COLOR_BG_RED
	COLOR_BG_GREEN
	COLOR_BG_YELLOW
	COLOR_BG_BLUE
	COLOR_BG_MAGENTA
	COLOR_BG_CYAN
	COLOR_BG_WHITE
	// COLOR_BG introduces an extended background color, see COLOR_ID
	// and COLOR_RGB. It is reserved by [ECMA-48] for the color format
	// defined by [T.416].
	COLOR_BG
	// COLOR_BG_DEFAULT resets the background color [ECMA-48].
	COLOR_BG_DEFAULT
)

// Bright foreground colors [aixterm]. They are not part of ECMA-48.
const (
	COLOR_FG_BRIGHT_BLACK Param = 90 + iota
	COLOR_FG_BRIGHT_RED
	COLOR_FG_BRIGHT_GREEN
	COLOR_FG_BRIGHT_YELLOW
	COLOR_FG_BRIGHT_BLUE
	COLOR_FG_BRIGHT_MAGENTA
	COLOR_FG_BRIGHT_CYAN
	COLOR_FG_BRIGHT_WHITE
)

// Bright background colors [aixterm]. They are not part of ECMA-48.
const (
	COLOR_BG_BRIGHT_BLACK Param = 100 + iota
	COLOR_BG_BRIGHT_RED
	COLOR_BG_BRIGHT_GREEN
	COLOR_BG_BRIGHT_YELLOW
	COLOR_BG_BRIGHT_BLUE
	COLOR_BG_BRIGHT_MAGENTA
	COLOR_BG_BRIGHT_CYAN
	COLOR_BG_BRIGHT_WHITE
)

// Underline colors [kitty]. The codes are reserved for future
// standardization by ECMA-48. The color is given in the
// format defined by [T.416] with sub parameters.
const (
	// COLOR_UL introduces an underline color, see Code_UL_ID and Code_UL_RGB.
	COLOR_UL Param = 58
	// COLOR_UL_DEFAULT resets the underline color.
	COLOR_UL_DEFAULT Param = 59
)

// Underline styles [kitty]. They use the sub parameter syntax of
// [ECMA-48] and require sub parameters, because 4;3 would
// mean underline and italic. Sub parameters are not supported by all
// terminals (see package documentation). Package ttycolors uses them
// only with the feature FeatureUnderlineStyles. Use Modes to put
// them into separate sequences.
// Terminals supporting underline styles do not necessarily support all
// of them. For example, GNOME Terminal (VTE 0.60) renders dotted and
// dashed underlines like curly underlines.
var (
	UNDERLINE_SINGLE = Code{MODE_UNDERLINE, 1}
	UNDERLINE_DOUBLE = Code{MODE_UNDERLINE, 2}
	UNDERLINE_CURLY  = Code{MODE_UNDERLINE, 3}
	UNDERLINE_DOTTED = Code{MODE_UNDERLINE, 4}
	UNDERLINE_DASHED = Code{MODE_UNDERLINE, 5}
)

// Color kinds of extended colors (COLOR_FG, COLOR_BG and COLOR_UL) [T.416].
const (
	// COLOR_ID is an indexed color. With the 256 color palette of
	// [xterm], 0-7 are the standard colors, 8-15 the bright colors,
	// 16-231 a 6x6x6 color cube and 232-255 a gray scale.
	COLOR_ID = 5
	// COLOR_RGB is a direct color given by red, green and blue values.
	// [T.416] specifies a color space id as first argument, which is
	// omitted by the semicolon form of [xterm].
	COLOR_RGB = 2
)

// Extended colors in the semicolon form [xterm] (38;5;n and 38;2;r;g;b).
// [T.416] specifies the form with sub parameters (38:5:n and
// 38:2::r:g:b), which is not supported by all terminals.
var (
	COLOR_FG_ID  = NewCodeSequence(COLOR_FG, COLOR_ID)
	COLOR_BG_ID  = NewCodeSequence(COLOR_BG, COLOR_ID)
	COLOR_FG_RGB = NewCodeSequence(COLOR_FG, COLOR_RGB)
	COLOR_BG_RGB = NewCodeSequence(COLOR_BG, COLOR_RGB)
)
