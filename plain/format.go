package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// Func is a function providing a formatted string for the given
// elements, like the format functions of this package.
type Func func(args ...any) string

// New provides a format function for a composed output format based on
// the given formats using the settings of the default context of
// package ttycolors (see ttycolors.New). The settings are taken when
// the function is created.
func New(fmts ...ttycolors.FormatProvider) Func {
	f := ttycolors.New(fmts...)
	return func(args ...any) string {
		return f.String(args...).String()
	}
}
