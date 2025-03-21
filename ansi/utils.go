package ansi

import (
	"slices"
	"strconv"
	"strings"

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
