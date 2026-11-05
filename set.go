package ttycolors

import (
	"io"
)

// Set applies the given formats to the standard output
// using the settings of the default context (see TTYContext.Set).
func Set(f ...FormatProvider) error {
	return defaultContext.Set(f...)
}

// SetFor applies the given formats to the given writer
// using the settings of the default context (see TTYContext.SetFor).
func SetFor(w io.Writer, f ...FormatProvider) error {
	return defaultContext.SetFor(w, f...)
}

// UnSet reverts the settings of the given formats for the standard output
// using the settings of the default context (see TTYContext.UnSet).
func UnSet(f ...FormatProvider) error {
	return defaultContext.UnSet(f...)
}

// UnSetFor reverts the settings of the given formats for the given writer
// using the settings of the default context (see TTYContext.UnSetFor).
func UnSetFor(w io.Writer, f ...FormatProvider) error {
	return defaultContext.UnSetFor(w, f...)
}
