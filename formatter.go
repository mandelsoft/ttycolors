package ttycolors

import (
	"fmt"
	"io"
	"os"

	"github.com/mandelsoft/goutils/general"
	"github.com/mandelsoft/ttycolors/colorstring"
)

type formatter struct {
	fmt *FormatInfo
}

var _ Formatter = (*formatter)(nil)

func NewFormatter(f ...FormatProvider) Formatter {
	return &formatter{
		New(f...),
	}
}

func (f formatter) Enabled() bool {
	return f.fmt.enabled
}

func (f formatter) Enable(b ...bool) {
	f.fmt.enabled = general.OptionalDefaultedBool(true, b...)
}

func (f formatter) Print(i ...interface{}) (n int, err error) {
	return os.Stdout.Write([]byte(f.fmt.String(colorstring.Sequence(i...)).String()))
}

func (f formatter) Printf(s string, i ...interface{}) (n int, err error) {
	return os.Stdout.Write([]byte(f.fmt.String(fmt.Sprintf(s, i...)).String()))
}

func (f formatter) Sprint(i ...interface{}) string {
	return f.fmt.String(colorstring.Sequence(i...)).String()
}

func (f formatter) Sprintf(s string, i ...interface{}) string {
	return f.fmt.String(fmt.Sprintf(s, i...)).String()
}

func (f formatter) Fprint(writer io.Writer, i ...interface{}) (n int, err error) {
	return writer.Write([]byte(f.fmt.String(colorstring.Sequence(i...)).String()))
}

func (f formatter) Fprintf(writer io.Writer, s string, i ...interface{}) (n int, err error) {
	return writer.Write([]byte(f.fmt.String(fmt.Sprintf(s, i...)).String()))
}
