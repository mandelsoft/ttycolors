package ttycolors

import (
	"fmt"
	"io"
	"os"

	"github.com/mandelsoft/ttycolors/renderer"
)

type formatter struct {
	fmt FormatInfo
}

var (
	_ Formatter      = (*formatter)(nil)
	_ FormatProvider = (*formatter)(nil)
)

// NewFormatter provides a Formatter using the settings
// of the default context.
func NewFormatter(f ...FormatProvider) Formatter {
	return newFormatter(defaultContext.current(), f...)
}

func newFormatter(settings renderer.Settings, f ...FormatProvider) Formatter {
	return &formatter{
		newFormat(settings, f...),
	}
}

func (f formatter) IsEnabled() bool {
	return f.fmt.IsEnabled()
}

func (f formatter) Enable(b ...bool) {
	f.fmt.Enable(b...)
}

func (f formatter) Features() Features {
	return f.fmt.Features()
}

func (f formatter) SetFeatures(features Features) {
	f.fmt.SetFeatures(features)
}

func (f formatter) EnableFeatures(features Features) {
	f.fmt.EnableFeatures(features)
}

func (f formatter) DisableFeatures(features Features) {
	f.fmt.DisableFeatures(features)
}

func (f formatter) Format() FormatInfo {
	return f.fmt
}

func (f formatter) String(seq ...any) String {
	return f.fmt.String(seq...)
}

func (f formatter) Print(i ...interface{}) (n int, err error) {
	return os.Stdout.Write([]byte(f.fmt.String(Sequence(i...)).String()))
}

func (f formatter) Printf(s string, i ...interface{}) (n int, err error) {
	return os.Stdout.Write([]byte(f.fmt.String(fmt.Sprintf(s, i...)).String()))
}

func (f formatter) Sprint(i ...interface{}) string {
	return f.fmt.String(Sequence(i...)).String()
}

func (f formatter) Sprintf(s string, i ...interface{}) string {
	return f.fmt.String(fmt.Sprintf(s, i...)).String()
}

func (f formatter) Fprint(writer io.Writer, i ...interface{}) (n int, err error) {
	return writer.Write([]byte(f.fmt.String(Sequence(i...)).String()))
}

func (f formatter) Fprintf(writer io.Writer, s string, i ...interface{}) (n int, err error) {
	return writer.Write([]byte(f.fmt.String(fmt.Sprintf(s, i...)).String()))
}
