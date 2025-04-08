package ansi

import (
	"slices"
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/mandelsoft/goutils/sliceutils"
)

func CodeToString(i Code) string {
	return strings.Join(sliceutils.Transform(i.Bytes(), ByteToString), ";")
}

func ByteToString(i byte) string {
	return strconv.Itoa(int(i))
}

// Mode provides an espace sequence for a sequence of output
// format codes.
func Mode(args ...Code) string {
	return CSI + strings.Join(sliceutils.Transform(args, CodeToString), ";") + "m"
}

// Code represents a sequence of Control Sequence Introducer arguments.
// for Select Graphic Rendition parameters.
type Code interface {
	Bytes() []byte
}

// Simple is a Code consisting of a single argument.
type Simple byte

func (s Simple) Bytes() []byte {
	return []byte{byte(s)}
}

// Sequence is a Code consisting of a sequence of arguments.
type Sequence []byte

func (s Sequence) Bytes() []byte {
	return []byte(s)
}
func (s Sequence) Append(codes ...Code) Sequence {
	seq := slices.Clone(s)
	for _, c := range codes {
		if c != nil {
			seq = append(seq, c.Bytes()...)
		}
	}
	return Sequence(seq)
}

func (s Sequence) AppendBytes(codes ...byte) Sequence {
	return append(slices.Clone(s), codes...)
}

// EscapeLength returns the byte length for an ANSI output format escape sequence
// if the data starts with such a prefix.
func EscapeLength(data []byte) int {
	if len(data) <= 3 || data[0] != '\x1b' || data[1] != '[' {
		return 0
	}

	for i, c := range data[2:] {
		if (c < '0' || c > '9') && c != ';' {
			if c != 'm' {
				return 0
			}
			return i + 3
		}
	}
	return 0
}

// SplitAt splits a string containing ANSI code sequences
// after the n-th rune. idx should be between
// 0 and CharLen(s) (including).
func SplitAt(s string, idx int) (string, string) {
	offset := CharIndex(s, idx)
	return string([]byte(s)[:offset]), string([]byte(s)[offset:])
}

// CharIndex returns the byte index of the n-th
// rune in a string containing ANI code sequences.
func CharIndex(s string, idx int) int {
	data := []byte(s)
	cnt := 0
	i := 0
	for i < len(data) {
		l := EscapeLength(data[i:])
		if l == 0 {
			_, l = utf8.DecodeRune(data[i:])
			if cnt == idx {
				break
			}
			cnt++
		}
		i += l
	}
	return i
}

// CharLen determines the number of runes
// in a string containing ANSI code sequences.
func CharLen(s string) int {
	data := []byte(s)
	cnt := 0
	for len(data) > 0 {
		l := EscapeLength(data)
		if l == 0 {
			_, l = utf8.DecodeRune(data)
			cnt++
		}
		data = data[l:]
	}
	return cnt
}
