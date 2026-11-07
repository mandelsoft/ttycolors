package ansi

import (
	"bytes"
	"fmt"
	"iter"
	"math/bits"
	"strconv"
	"strings"
)

// Covered is a bit set describing the output format areas
// affected by an ANSI code sequence. The topmost bit (S_CONTROL)
// indicates control sequences, which are no output format sequences.
type Covered uint

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

	// S_CONTROL indicates escape sequences, which are no output format
	// sequences, for example cursor movements. It is not an output format
	// area and therefore not part of S_ALL. It uses the topmost bit, so that
	// new output format areas can be added without changing it.
	S_CONTROL Covered = 1 << (bits.UintSize - 1)
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
	{S_CONTROL, "S_CONTROL"},
}

// String provides the names of the covered areas separated by |,
// for example S_HIGHLIGHT|S_FG_COLOR. S_NONE and S_ALL are
// provided by their names, for example S_ALL|S_CONTROL.
func (c Covered) String() string {
	if c == S_NONE {
		return "S_NONE"
	}
	var names []string
	if c&S_ALL == S_ALL {
		names = append(names, "S_ALL")
		c &^= S_ALL
	}
	for _, e := range coveredNames {
		if c&e.area != 0 {
			names = append(names, e.name)
			c &^= e.area
		}
	}
	if c != 0 {
		names = append(names, fmt.Sprintf("0x%x", uint(c)))
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

// ScanSequence detects a leading escape sequence in data.
// It returns the byte length of the complete sequence and the output
// format areas affected by it. All escape sequences in their 7-bit
// form according to ECMA-48 are recognized:
//   - control sequences (ESC [ ... final byte),
//   - control strings (OSC, DCS, SOS, PM and APC, for example
//     ESC ] ... terminated by ST (ESC \) or BEL),
//   - other escape sequences (ESC intermediate bytes final byte,
//     for example ESC 7 or ESC ( B).
//
// Only output format sequences (Select Graphic Rendition, ESC [ ... m)
// affect output format areas, for all other sequences S_CONTROL is
// returned. The 8-bit forms of the control functions are not recognized,
// because those bytes are used by UTF-8 encoded characters.
// If data does not start with a complete escape sequence, the length is 0.
//
// Code 21 (MODE_DOUBLE_UNDERLINE) is mapped to S_UNDERLINE (doubly underlined
// according to ECMA-48), although some terminals interpret it as resetting bold.
// Parameters following an extended color with an unknown
// color kind cannot be interpreted and are ignored.
func ScanSequence(data []byte) (int, Covered) {
	l, e := ScanSequenceEffect(data)
	return l, e.Covered()
}

// controlEffect is the effect of a sequence, which is no
// output format sequence.
var controlEffect = Effect{Set: S_CONTROL}

// ScanSequenceEffect works like ScanSequence, but provides the effect
// of the sequence: the areas set to a value and the areas reset to the
// terminal default. The codes of the sequence are applied in order,
// for example, 1;0 resets all areas, 0;1 resets all areas but sets
// the highlight area. For sequences, which are no output format
// sequences, S_CONTROL is set.
func ScanSequenceEffect(data []byte) (int, Effect) {
	if len(data) < 2 || data[0] != '\x1b' {
		return 0, Effect{}
	}
	switch c := data[1]; {
	case c == '[':
		return scanControlSequence(data)
	case c == ']' || c == 'P' || c == 'X' || c == '^' || c == '_':
		return scanControlString(data)
	case isIntermediateByte(c):
		for i := 2; i < len(data); i++ {
			switch b := data[i]; {
			case isIntermediateByte(b):
			case isEscapeFinalByte(b):
				return i + 1, controlEffect
			default:
				return 0, Effect{}
			}
		}
	case isEscapeFinalByte(c):
		return 2, controlEffect
	}
	return 0, Effect{}
}

// scanControlSequence scans a control sequence (CSI parameter bytes,
// intermediate bytes, final byte).
func scanControlSequence(data []byte) (int, Effect) {
	i := 2
	for i < len(data) && isParameterByte(data[i]) {
		i++
	}
	params := data[2:i]
	f := i
	for f < len(data) && isIntermediateByte(data[f]) {
		f++
	}
	if f >= len(data) || !isFinalByte(data[f]) {
		return 0, Effect{}
	}
	if data[f] == 'm' && f == i && isSGRParameters(params) {
		return f + 1, sgrEffect(params)
	}
	return f + 1, controlEffect
}

// scanControlString scans a control string terminated by
// ST (ESC \) or BEL.
func scanControlString(data []byte) (int, Effect) {
	for i := 2; i < len(data); i++ {
		switch data[i] {
		case '\a':
			return i + 1, controlEffect
		case '\x1b':
			if i+1 < len(data) && data[i+1] == '\\' {
				return i + 2, controlEffect
			}
			return 0, Effect{}
		}
	}
	return 0, Effect{}
}

// isParameterByte checks for a parameter byte of a control sequence [ECMA-48].
func isParameterByte(c byte) bool {
	return c >= '0' && c <= '?' // 0x30-0x3F
}

// isIntermediateByte checks for an intermediate byte of an escape
// or control sequence [ECMA-48].
func isIntermediateByte(c byte) bool {
	return c >= ' ' && c <= '/' // 0x20-0x2F
}

// isFinalByte checks for the final byte of a control sequence [ECMA-48].
func isFinalByte(c byte) bool {
	return c >= '@' && c <= '~' // 0x40-0x7E
}

// isEscapeFinalByte checks for the final byte of an escape sequence
// (other than a control sequence or control string) [ECMA-48].
func isEscapeFinalByte(c byte) bool {
	return c >= '0' && c <= '~' // 0x30-0x7E
}

func isSGRParameters(params []byte) bool {
	for _, c := range params {
		if !isSGRParameterByte(c) {
			return false
		}
	}
	return true
}

// Scan scans a complete string for escape sequences
// (see ScanSequence) and returns the union of the output format
// areas affected by all found sequences. S_CONTROL is included,
// if other control sequences are found.
// A sequence resetting all modes (ESC[0m or ESC[m) affects all
// areas (S_ALL), therefore the result contains S_ALL for every
// string containing such a sequence.
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

// ScanEffect scans a complete string for escape sequences
// and provides their combined effect (see ScanSequenceEffect).
// S_CONTROL is set, if other control sequences are found.
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

// ScanCodes provides the codes of all output format sequences of s
// in order, together with the output format area affected by them.
// A reset of all modes (ESC[0m or ESC[m) affects S_ALL, all other codes
// a single area. The arguments of extended colors given as separate
// parameters are part of the code of the color, for example 38;5;208.
// Codes not affecting an output format area (for example fonts) and
// other escape sequences are omitted.
func ScanCodes(s string) iter.Seq2[Covered, CodeSequence] {
	return func(yield func(Covered, CodeSequence) bool) {
		data := []byte(s)
		for {
			i := bytes.IndexByte(data, '\x1b')
			if i < 0 {
				return
			}
			l, e := ScanSequenceEffect(data[i:])
			if l == 0 {
				l = 1
			}
			if e.Covered() != S_NONE && e.Set&S_CONTROL == 0 {
				cont := true
				sgrCodes(data[i+2:i+l-1], func(code int, text []byte) {
					if area := sgrArea(code); cont && area != S_NONE {
						cont = yield(area, parseCodeSequence(text))
					}
				})
				if !cont {
					return
				}
			}
			data = data[i+l:]
		}
	}
}

// parseCodeSequence parses the parameter text of an SGR code. A single
// parameter is provided as Param, a single code with sub parameters
// as Code, empty parameters as Omitted.
func parseCodeSequence(text []byte) CodeSequence {
	var seq codeSequence
	for _, p := range bytes.Split(text, []byte{PARAM_SEPARATOR}) {
		var code Code
		for _, sub := range bytes.Split(p, []byte{SUBPARAM_SEPARATOR}) {
			v, err := strconv.Atoi(string(sub))
			if err != nil || v < 0 {
				v = int(Omitted)
			}
			code = append(code, Param(v))
		}
		seq = append(seq, code)
	}
	switch {
	case len(seq) > 1:
		return seq
	case len(seq[0]) > 1:
		return seq[0]
	default:
		return seq[0][0]
	}
}

func isSGRParameterByte(c byte) bool {
	return (c >= '0' && c <= '9') || c == PARAM_SEPARATOR || c == SUBPARAM_SEPARATOR
}

// sgrEffect determines the effect of the
// parameters of a Select Graphic Rendition sequence.
func sgrEffect(params []byte) Effect {
	var effect Effect

	sgrCodes(params, func(code int, _ []byte) {
		if isResetCode(code) {
			effect = effect.Then(Effect{Reset: sgrArea(code)})
		} else {
			effect = effect.Then(Effect{Set: sgrArea(code)})
		}
	})
	return effect
}

// sgrCodes calls f for all codes of the parameters of a Select
// Graphic Rendition sequence in order, with the code and its
// parameter text. The arguments of an extended color given as
// separate parameters are part of the text of the color code.
func sgrCodes(params []byte, f func(code int, text []byte)) {
	args := bytes.Split(params, []byte{PARAM_SEPARATOR})
	for i := 0; i < len(args); i++ {
		code, sub, ok := sgrCode(args[i])
		if !ok {
			continue
		}
		n := 1
		if isExtendedColor(code) && !sub && i+1 < len(args) {
			switch kind, _, _ := sgrCode(args[i+1]); kind {
			case COLOR_ID:
				n = 3
			case COLOR_RGB:
				n = 5
			default:
				// the number of arguments of an unknown
				// color kind is unknown, so the remaining
				// parameters cannot be interpreted.
				f(code, args[i])
				return
			}
		}
		end := min(i+n, len(args))
		f(code, bytes.Join(args[i:end], []byte{PARAM_SEPARATOR}))
		i = end - 1
	}
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
