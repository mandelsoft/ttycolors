package ttycolors

import (
	"io"

	"github.com/mandelsoft/ttycolors/colorstring/renderer"
)

type FormatProvider interface {
	Format() *FormatInfo
}

type Format interface {
	FormatProvider
	renderer.Composer
}

type Formatter interface {
	Enabled() bool
	Enable(...bool)

	Print(i ...interface{}) (n int, err error)
	Printf(string, ...interface{}) (n int, err error)
	Sprint(i ...interface{}) string
	Sprintf(string, ...interface{}) string
	Fprint(io.Writer, ...interface{}) (n int, err error)
	Fprintf(io.Writer, string, ...interface{}) (n int, err error)
}
