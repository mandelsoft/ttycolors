package ansi

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
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
	S_OVERLINE
)

const (
	// S_NONE is used for sequences not affecting any output format area.
	S_NONE Covered = 0
	// S_COLOR covers all color areas (foreground, background and underline color).
	S_COLOR = S_FG_COLOR | S_BG_COLOR | S_UL_COLOR
	// S_ALL covers all output format areas (used for a reset of all modes).
	S_ALL = S_HIGHLIGHT | S_ITALIC | S_UNDERLINE | S_BLINK | S_REVERSE | S_INVIS | S_STRIKE | S_OVERLINE | S_COLOR
)

var coveredNames = []struct {
	area Covered
	name string
}{
	{S_HIGHLIGHT, "S_HIGHLIGHT"},
	{S_ITALIC, "S_ITALIC"},
	{S_UNDERLINE, "S_UNDERLINE"},
	{S_BLINK, "S_BLINK"},
	{S_REVERSE, "S_REVERSE"},
	{S_INVIS, "S_INVIS"},
	{S_STRIKE, "S_STRIKE"},
	{S_FG_COLOR, "S_FG_COLOR"},
	{S_BG_COLOR, "S_BG_COLOR"},
	{S_UL_COLOR, "S_UL_COLOR"},
	{S_OVERLINE, "S_OVERLINE"},
}

// String provides the names of the covered areas separated by |,
// for example S_HIGHLIGHT|S_FG_COLOR. S_NONE and S_ALL are
// provided by their names.
func (c Covered) String() string {
	switch c {
	case S_NONE:
		return "S_NONE"
	case S_ALL:
		return "S_ALL"
	}
	var names []string
	for _, e := range coveredNames {
		if c&e.area != 0 {
			names = append(names, e.name)
			c &^= e.area
		}
	}
	if c != 0 {
		names = append(names, fmt.Sprintf("0x%x", int(c)))
	}
	return strings.Join(names, "|")
}

// Has checks whether all areas of o are covered by c.
func (c Covered) Has(o Covered) bool {
	return c&o == o
}

// Effect describes the effect of output format sequences
// on the output format areas.
type Effect struct {
	// Set are the areas set to a value other than the terminal default.
	Set Covered
	// Reset are the areas reset to the terminal default.
	Reset Covered
}

// String provides the set and reset areas.
func (e Effect) String() string {
	return fmt.Sprintf("{Set: %s, Reset: %s}", e.Set, e.Reset)
}

// Covered provides all areas affected by the effect.
func (e Effect) Covered() Covered {
	return e.Set | e.Reset
}

// Then provides the combined effect of e followed by o.
func (e Effect) Then(o Effect) Effect {
	return Effect{
		Set:   (e.Set &^ o.Reset) | o.Set,
		Reset: (e.Reset &^ o.Set) | o.Reset,
	}
}

// ScanSequence detects a leading ANSI output format sequence
// (Select Graphic Rendition, ESC [ ... m) in data.
// Other control sequences are not recognized.
// Code 21 (MODE_DOUBLE_UNDERLINE) is mapped to S_UNDERLINE (doubly underlined according to
// ECMA-48), although some terminals interpret it as resetting bold.
// It returns the byte length of the complete sequence
// and the output format areas affected by it.
// If data does not start with a complete output format sequence,
// the length is 0.
// Parameters following an extended color with an unknown
// color kind cannot be interpreted and are ignored.
func ScanSequence(data []byte) (int, Covered) {
	l, e := ScanSequenceEffect(data)
	return l, e.Covered()
}

// ScanSequenceEffect works like ScanSequence, but provides the effect
// of the sequence: the areas set to a value and the areas reset to the
// terminal default. The codes of the sequence are applied in order,
// for example, 1;0 resets all areas, 0;1 resets all areas but sets
// the highlight area.
func ScanSequenceEffect(data []byte) (int, Effect) {
	if len(data) < 3 || data[0] != '\x1b' || data[1] != '[' {
		return 0, Effect{}
	}
	for i := 2; i < len(data); i++ {
		c := data[i]
		if c == 'm' {
			return i + 1, sgrEffect(data[2:i])
		}
		if !isSGRParameterByte(c) {
			return 0, Effect{}
		}
	}
	return 0, Effect{}
}

// Scan scans a complete string for ANSI output format sequences
// (see ScanSequence) and returns the union of the output format
// areas affected by all found sequences.
// A sequence resetting all modes (ESC[0m or ESC[m) affects all
// areas (S_ALL), therefore the result is S_ALL for every string
// containing such a sequence.
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

// ScanEffect scans a complete string for ANSI output format sequences
// and provides their combined effect (see ScanSequenceEffect).
// The sequences are applied in order, so that, for example,
// ESC[1mbold ESC[22m has the effect of resetting the highlight area.
func ScanEffect(s string) Effect {
	var effect Effect

	data := []byte(s)
	for {
		i := bytes.IndexByte(data, '\x1b')
		if i < 0 {
			return effect
		}
		l, e := ScanSequenceEffect(data[i:])
		if l == 0 {
			l = 1
		}
		effect = effect.Then(e)
		data = data[i+l:]
	}
}

func isSGRParameterByte(c byte) bool {
	return (c >= '0' && c <= '9') || c == PARAM_SEPARATOR || c == SUBPARAM_SEPARATOR
}

// sgrEffect determines the effect of the
// parameters of a Select Graphic Rendition sequence.
func sgrEffect(params []byte) Effect {
	var effect Effect

	args := bytes.Split(params, []byte{PARAM_SEPARATOR})
	for i := 0; i < len(args); i++ {
		code, sub, ok := sgrCode(args[i])
		if !ok {
			continue
		}
		if isResetCode(code) {
			effect = effect.Then(Effect{Reset: sgrArea(code)})
		} else {
			effect = effect.Then(Effect{Set: sgrArea(code)})
		}
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
				return effect
			}
		}
	}
	return effect
}

// isResetCode checks whether a code resets its area
// to the terminal default.
func isResetCode(code int) bool {
	switch code {
	case int(MODE_RESET_ALL), int(MODE_RESET_BOLD), int(MODE_RESET_ITALIC),
		int(MODE_RESET_UNDERLINE), int(MODE_RESET_BLINK), int(MODE_RESET_REVERSE),
		int(MODE_RESET_INVIS), int(MODE_RESET_STRIKE), int(MODE_RESET_OVERLINE),
		int(COLOR_FG_DEFAULT), int(COLOR_BG_DEFAULT), int(COLOR_UL_DEFAULT):
		return true
	}
	return false
}

// sgrCode parses a single SGR parameter. An empty parameter
// is interpreted as 0. sub reports whether the parameter
// contains colon separated sub parameters.
func sgrCode(arg []byte) (code int, sub bool, ok bool) {
	if idx := bytes.IndexByte(arg, SUBPARAM_SEPARATOR); idx >= 0 {
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
	case code == int(MODE_UNDERLINE) || code == int(MODE_DOUBLE_UNDERLINE) || code == int(MODE_RESET_UNDERLINE):
		return S_UNDERLINE
	case code == int(MODE_BLINK) || code == int(MODE_RAPID_BLINK) || code == int(MODE_RESET_BLINK):
		return S_BLINK
	case code == int(MODE_REVERSE) || code == int(MODE_RESET_REVERSE):
		return S_REVERSE
	case code == int(MODE_INVIS) || code == int(MODE_RESET_INVIS):
		return S_INVIS
	case code == int(MODE_STRIKE) || code == int(MODE_RESET_STRIKE):
		return S_STRIKE
	case code == int(MODE_OVERLINE) || code == int(MODE_RESET_OVERLINE):
		return S_OVERLINE
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
