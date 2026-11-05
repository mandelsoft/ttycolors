package ansi

import (
	"fmt"
	"unicode/utf8"
)

// EscapeLength returns the byte length for an ANSI output format escape sequence
// if the data starts with such a prefix.
// Other control sequences are not recognized (see package documentation).
func EscapeLength(data []byte) int {
	l, _ := ScanSequence(data)
	return l
}

// checkIndex panics for negative rune indices.
func checkIndex(idx int) {
	if idx < 0 {
		panic(fmt.Sprintf("ansi: negative index %d", idx))
	}
}

// SplitAt splits a string containing ANSI code sequences
// after the n-th rune. idx should be between
// 0 and CharLen(s) (including). It panics for
// a negative index.
func SplitAt(s string, idx int) (string, string) {
	offset := CharIndex(s, idx)
	return string([]byte(s)[:offset]), string([]byte(s)[offset:])
}

// SubString extracts runes from index start to end-1, excluding
// leading and trailing ANSI code sequences.
// If start equals end, the result is empty. Like for
// slices, end must not be smaller than start.
// It panics for negative indices.
func SubString(s string, start, end int) string {
	checkIndex(start)
	checkIndex(end)
	if start == end {
		return ""
	}
	bstart := CharIndex(s, start)
	bend := CharIndexBefore(s, end)
	return string([]byte(s)[bstart:bend])
}

// CharIndex returns the byte index of the n-th
// rune in a string containing ANSI code sequences.
// It panics for a negative index.
func CharIndex(s string, idx int) int {
	checkIndex(idx)
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

// CharIndexBefore returns the byte index of the n-th
// rune in a string containing ANSI code sequences
// including a leading ANSI code sequence.
// It panics for a negative index.
func CharIndexBefore(s string, idx int) int {
	checkIndex(idx)
	data := []byte(s)
	cnt := 0
	i := 0
	for i < len(data) {
		if cnt == idx {
			break
		}
		l := EscapeLength(data[i:])
		if l == 0 {
			_, l = utf8.DecodeRune(data[i:])
			cnt++
		}
		i += l
	}
	return i
}

// CharLen determines the number of runes
// in a string containing ANSI code sequences.
// Only output format sequences are ignored, other control
// sequences are counted as text (see package documentation).
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

// ByteIndex returns the byte index of the first occurrence
// of the given rune in data containing ANSI code sequences,
// or -1 if the rune is not present.
func ByteIndex(data []byte, c rune) int {
	var r rune

	i := 0
	for len(data) > 0 {
		l := EscapeLength(data)
		if l == 0 {
			r, l = utf8.DecodeRune(data)
			if r == c {
				return i
			}
		}
		i += l
		data = data[l:]
	}
	return -1
}

// Strip removes all ANSI output format sequences from a string.
// Other control sequences are kept (see package documentation).
func Strip(s string) string {
	data := []byte(s)
	i := 0
	for len(data[i:]) > 0 {
		l := EscapeLength(data[i:])
		if l == 0 {
			_, l = utf8.DecodeRune(data[i:])
			i += l
		} else {
			data = append(data[:i], data[i+l:]...)
		}
	}
	return string(data)
}
