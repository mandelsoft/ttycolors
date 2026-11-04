package ansi

import (
	"bytes"
	"strconv"
)

// Covered is a bit set describing the output format areas
// affected by an ANSI code sequence.
type Covered int

const (
	S_HIGHLIGHT Covered = 1 << iota
	S_ITALIC
	S_UNDERLINE
	S_BLINK
	S_REVERSE
	S_INVIS
	S_STRIKE
	S_FG_COLOR
	S_BG_COLOR
	S_UL_COLOR
)

const (
	// S_NONE is used for sequences not affecting any output format area.
	S_NONE Covered = 0
	// S_COLOR covers all color areas (foreground, background and underline color).
	S_COLOR = S_FG_COLOR | S_BG_COLOR | S_UL_COLOR
	// S_ALL covers all output format areas (used for a reset of all modes).
	S_ALL = S_HIGHLIGHT | S_ITALIC | S_UNDERLINE | S_BLINK | S_REVERSE | S_INVIS | S_STRIKE | S_COLOR
)

// Has checks whether all areas of o are covered by c.
func (c Covered) Has(o Covered) bool {
	return c&o == o
}

// ScanSequence detects a leading ANSI output format sequence
// (Select Graphic Rendition, ESC [ ... m) in data.
// It returns the byte length of the complete sequence
// and the output format areas affected by it.
// If data does not start with a complete output format sequence,
// the length is 0.
// Parameters following an extended color with an unknown
// color kind cannot be interpreted and are ignored.
func ScanSequence(data []byte) (int, Covered) {
	if len(data) < 3 || data[0] != '\x1b' || data[1] != '[' {
		return 0, S_NONE
	}
	for i := 2; i < len(data); i++ {
		c := data[i]
		if c == 'm' {
			return i + 1, sgrCovered(data[2:i])
		}
		if !isSGRParameterByte(c) {
			return 0, S_NONE
		}
	}
	return 0, S_NONE
}

// Scan scans a complete string for ANSI output format sequences
// (see ScanSequence) and returns the union of the output format
// areas affected by all found sequences.
func Scan(s string) Covered {
	var covered Covered

	data := []byte(s)
	for {
		i := bytes.IndexByte(data, '\x1b')
		if i < 0 {
			return covered
		}
		l, c := ScanSequence(data[i:])
		if l == 0 {
			l = 1
		}
		covered |= c
		data = data[i+l:]
	}
}

func isSGRParameterByte(c byte) bool {
	return (c >= '0' && c <= '9') || c == ';' || c == ':'
}

// sgrCovered determines the areas covered by the
// parameters of a Select Graphic Rendition sequence.
func sgrCovered(params []byte) Covered {
	var covered Covered

	args := bytes.Split(params, []byte{';'})
	for i := 0; i < len(args); i++ {
		code, sub, ok := sgrCode(args[i])
		if !ok {
			continue
		}
		covered |= sgrArea(code)
		if isExtendedColor(code) && !sub && i+1 < len(args) {
			// skip the arguments of an extended color
			// given as separate parameters.
			switch kind, _, _ := sgrCode(args[i+1]); kind {
			case COLOR_ID:
				i += 2
			case COLOR_RGB:
				i += 4
			default:
				// the number of arguments of an unknown
				// color kind is unknown, so the remaining
				// parameters cannot be interpreted.
				return covered
			}
		}
	}
	return covered
}

// sgrCode parses a single SGR parameter. An empty parameter
// is interpreted as 0. sub reports whether the parameter
// contains colon separated sub parameters.
func sgrCode(arg []byte) (code int, sub bool, ok bool) {
	if idx := bytes.IndexByte(arg, ':'); idx >= 0 {
		arg = arg[:idx]
		sub = true
	}
	if len(arg) == 0 {
		return 0, sub, true
	}
	code, err := strconv.Atoi(string(arg))
	if err != nil || code < 0 {
		return 0, sub, false
	}
	return code, sub, true
}

func isExtendedColor(code int) bool {
	return code == int(COLOR_FG) || code == int(COLOR_BG) || code == int(COLOR_UL)
}

// sgrArea maps an SGR code to the output format area it affects.
func sgrArea(code int) Covered {
	switch {
	case code == int(MODE_RESET_ALL):
		return S_ALL
	case code == int(MODE_BOLD) || code == int(MODE_DIM) || code == int(MODE_RESET_BOLD):
		return S_HIGHLIGHT
	case code == int(MODE_ITALIC) || code == int(MODE_RESET_ITALIC):
		return S_ITALIC
	case code == int(MODE_UNDERLINE) || code == 21 || code == int(MODE_RESET_UNDERLINE):
		return S_UNDERLINE
	case code == int(MODE_BLINK) || code == 6 || code == int(MODE_RESET_BLINK):
		return S_BLINK
	case code == int(MODE_REVERSE) || code == int(MODE_RESET_REVERSE):
		return S_REVERSE
	case code == int(MODE_INVIS) || code == int(MODE_RESET_INVIS):
		return S_INVIS
	case code == int(MODE_STRIKE) || code == int(MODE_RESET_STRIKE):
		return S_STRIKE
	case code >= int(COLOR_FG_BLACK) && code <= int(COLOR_FG_DEFAULT),
		code >= int(COLOR_FG_BRIGHT_BLACK) && code <= int(COLOR_FG_BRIGHT_WHITE):
		return S_FG_COLOR
	case code >= int(COLOR_BG_BLACK) && code <= int(COLOR_BG_DEFAULT),
		code >= int(COLOR_BG_BRIGHT_BLACK) && code <= int(COLOR_BG_BRIGHT_WHITE):
		return S_BG_COLOR
	case code == int(COLOR_UL) || code == int(COLOR_UL_DEFAULT):
		return S_UL_COLOR
	default:
		return S_NONE
	}
}
