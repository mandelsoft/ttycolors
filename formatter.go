package ttycolors

import (
	"fmt"
	"io"
	"os"

	"github.com/mandelsoft/goutils/optionutils"
)

type formatter struct {
	fmt FormatInfo
}

var (
	_ Formatter      = (*formatter)(nil)
	_ FormatProvider = (*formatter)(nil)
)

func NewFormatter(f ...FormatProvider) Formatter {
	return newFormatter(NoColors, f...)
}

func newFormatter(enabled bool, f ...FormatProvider) Formatter {
	return &formatter{
		newFormat(enabled, f...),
	}
}

func (f formatter) IsEnabled() bool {
	return f.fmt.enabled
}

func (f formatter) Enable(b ...bool) {
	f.fmt.enabled = optionutils.BoolOption(b...)
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
