package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// Func is a function providing a formatted string for the given
// elements, like the format functions of this package. The elements
// are concatenated without separators:
//   - Go strings (for example the results of other functions of this
//     package) and byte slices are used as text, output format sequences
//     embedded in them are tracked,
//   - ttycolors.String objects are nested structurally,
//   - all other values are converted with fmt.Sprintf("%v").
type Func func(args ...any) string

// New provides a format function for a composed output format based on
// the given formats using the settings of the default context of
// package ttycolors (see ttycolors.New). The settings are taken when
// the function is created.
func New(fmts ...ttycolors.FormatProvider) Func {
	return defaultContext.New(fmts...)
}

// New provides a format function for a composed output format based on
// the given formats using the settings of the context
// (see ttycolors.TTYContext.New). The settings are taken when
// the function is created.
func (c *_context) New(fmts ...ttycolors.FormatProvider) Func {
	return formatFunc(c.ctx.New(fmts...))
}

// ScanFormat provides a format function for the output format
// established at the end of a string containing output format
// sequences using the default context (see ttycolors.ScanFormat).
// It can be used to continue a string with its final output format:
//
//	s = s + plain.ScanFormat(s)(" continued")
//
// The settings are taken when the function is created.
func ScanFormat(s string) Func {
	return defaultContext.ScanFormat(s)
}

// ScanFormat provides a format function for the output format
// established at the end of a string containing output format
// sequences using the context (see ttycolors.TTYContext.ScanFormat).
// The settings are taken when the function is created.
func (c *_context) ScanFormat(s string) Func {
	return formatFunc(c.ctx.ScanFormat(s))
}

// formatFunc provides a format function for a composed format.
func formatFunc(f ttycolors.FormatInfo) Func {
	return func(args ...any) string {
		return f.String(args...).String()
	}
}
