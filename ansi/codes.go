package ansi

import (
	"slices"
	"strconv"
)

// Sequence is an encoded escape sequence or a part of it.
type Sequence []byte

// CodeSequence is a sequence of Codes, the parameter list of a
// Select Graphic Rendition sequence. It is implemented by Param,
// Code and the sequences provided by NewCodeSequence and Join.
type CodeSequence interface {
	// Codes provides the codes of the sequence.
	Codes() []Code
	// AppendTo appends the encoded sequence to seq.
	AppendTo(seq Sequence) Sequence
	// Bytes provides the encoded sequence.
	Bytes() []byte
}

////////////////////////////////////////////////////////////////////////////////

// Param is a single numeric parameter value.
// It is encoded as decimal digits. Negative values, like Omitted,
// are encoded as empty value.
// As CodeSequence, it consists of a single Code with this parameter.
type Param int

// Omitted denotes an omitted (empty) parameter value.
const Omitted Param = -1

var _ CodeSequence = Param(0)

func (p Param) Codes() []Code {
	return []Code{{p}}
}

func (p Param) AppendTo(seq Sequence) Sequence {
	if p < 0 {
		return seq
	}
	return strconv.AppendInt(seq, int64(p), 10)
}

func (p Param) Bytes() []byte {
	return p.AppendTo(nil)
}

////////////////////////////////////////////////////////////////////////////////

// Code is a sequence of parameters separated by SUBPARAM_SEPARATOR,
// for example 1 or 4:3. As CodeSequence, it consists of this single Code.
// Codes with more than one parameter (sub parameters) are not supported
// by all terminals. They should only be used, if a function cannot be
// expressed by separate parameters (see package documentation).
// Functions with arguments, like colors, should be expressed by a
// CodeSequence with a Code per parameter (see NewCodeSequence).
type Code []Param

var _ CodeSequence = Code(nil)

func (c Code) Codes() []Code {
	return []Code{c}
}

func (c Code) AppendTo(seq Sequence) Sequence {
	for i, p := range c {
		if i > 0 {
			seq = append(seq, SUBPARAM_SEPARATOR)
		}
		seq = p.AppendTo(seq)
	}
	return seq
}

func (c Code) Bytes() []byte {
	return c.AppendTo(nil)
}

// HasSubParams checks whether the code consists of more than one parameter.
// Such codes are not supported by all terminals and are put into
// separate sequences by Modes.
func (c Code) HasSubParams() bool {
	return len(c) > 1
}

////////////////////////////////////////////////////////////////////////////////

// codeSequence is a sequence of Codes separated by PARAM_SEPARATOR,
// for example 1;38;5;3;4:3.
type codeSequence []Code

var _ CodeSequence = codeSequence(nil)

// NewCodeSequence provides a CodeSequence consisting of
// a Code for every given parameter, for example 38;5.
func NewCodeSequence(params ...Param) CodeSequence {
	r := make(codeSequence, len(params))
	for i, p := range params {
		r[i] = Code{p}
	}
	return r
}

// Join provides a CodeSequence consisting of the codes
// of the given sequences. Nil sequences are ignored.
func Join(seqs ...CodeSequence) CodeSequence {
	var r codeSequence
	for _, s := range seqs {
		if s != nil {
			r = append(r, s.Codes()...)
		}
	}
	return r
}

func (s codeSequence) Codes() []Code {
	return slices.Clone(s)
}

func (s codeSequence) AppendTo(seq Sequence) Sequence {
	for i, c := range s {
		if i > 0 {
			seq = append(seq, PARAM_SEPARATOR)
		}
		seq = c.AppendTo(seq)
	}
	return seq
}

func (s codeSequence) Bytes() []byte {
	return s.AppendTo(nil)
}

////////////////////////////////////////////////////////////////////////////////

// codes provides the non-empty codes of the given sequences.
func codes(seqs []CodeSequence) codeSequence {
	var r codeSequence
	for _, e := range seqs {
		if e == nil {
			continue
		}
		for _, c := range e.Codes() {
			if len(c) > 0 {
				r = append(r, c)
			}
		}
	}
	return r
}

func appendMode(seq Sequence, c codeSequence) Sequence {
	seq = append(seq, CSI...)
	seq = c.AppendTo(seq)
	return append(seq, 'm')
}

// AppendMode appends the escape sequence for a sequence of output
// format code sequences to seq.
// Nil sequences and empty codes are ignored. If only such sequences
// are given, nothing is appended. Without any sequence the sequence
// resetting all modes is appended.
func AppendMode(seq Sequence, seqs ...CodeSequence) Sequence {
	c := codes(seqs)
	if len(c) == 0 && len(seqs) > 0 {
		return seq
	}
	return appendMode(seq, c)
}

// Mode provides the escape sequence for a sequence of output
// format code sequences (see AppendMode).
func Mode(seqs ...CodeSequence) string {
	return string(AppendMode(nil, seqs...))
}

// AppendModes appends the escape sequences for a sequence of output
// format code sequences to seq. Contrary to AppendMode, codes with
// sub parameters are put into separate sequences following the
// sequence for the other codes. This way, terminals not supporting
// sub parameters only drop those codes, if they drop complete
// sequences with unknown parameter formats.
// Nil sequences and empty codes are ignored. Without any remaining
// code, nothing is appended.
func AppendModes(seq Sequence, seqs ...CodeSequence) Sequence {
	var plain, sub codeSequence
	for _, c := range codes(seqs) {
		if c.HasSubParams() {
			sub = append(sub, c)
		} else {
			plain = append(plain, c)
		}
	}
	if len(plain) > 0 {
		seq = appendMode(seq, plain)
	}
	for _, c := range sub {
		seq = appendMode(seq, codeSequence{c})
	}
	return seq
}

// Modes provides the escape sequences for a sequence of output
// format code sequences (see AppendModes).
func Modes(seqs ...CodeSequence) string {
	return string(AppendModes(nil, seqs...))
}
