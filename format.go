package ttycolors

import (
	"bytes"

	"github.com/mandelsoft/goutils/optionutils"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/renderer"
)

func containsCode(list []ansi.Code, code ansi.Code) bool {
	for _, e := range list {
		if bytes.Equal(e.Bytes(), code.Bytes()) {
			return true
		}
	}
	return false
}

type FormatInfo = *_FormatInfo

type _FormatInfo struct {
	enabled   bool
	modeStart []ansi.Code
	modeEnd   []ansi.Code
	fg        ansi.Code
	bg        ansi.Code
}

var (
	_ Format      = (FormatInfo)(nil)
	_ ContextInfo = (FormatInfo)(nil)
)

func (f *_FormatInfo) IsEnabled() bool {
	return f.enabled
}

func (f *_FormatInfo) Enable(b ...bool) {
	f.enabled = optionutils.BoolOption(b...)
}

func (f *_FormatInfo) Start() []byte {
	var s []byte

	if !f.enabled {
		return nil
	}
	for _, c := range f.modeStart {
		s = append(s, c.Bytes()...)
	}
	if f.fg != nil {
		s = append(s, f.fg.Bytes()...)
	}
	if f.bg != nil {
		s = append(s, f.bg.Bytes()...)
	}
	return []byte(ansi.Mode(ansi.Sequence(s)))
}

func (f *_FormatInfo) End() []byte {
	var s []byte

	if !f.enabled {
		return nil
	}
	for _, c := range f.modeEnd {
		s = append(s, c.Bytes()...)
	}
	if f.fg != nil {
		s = append(s, ansi.COLOR_FG_DEFAULT.Bytes()...)
	}
	if f.bg != nil {
		s = append(s, ansi.COLOR_BG_DEFAULT.Bytes()...)
	}
	return []byte(ansi.Mode(ansi.Sequence(s)))
}

func (f *_FormatInfo) Add(a FormatInfo) {
	if a.bg != nil {
		f.bg = a.bg
	}
	if a.fg != nil {
		f.fg = a.fg
	}
	if len(a.modeStart) > 0 {
		for i, c := range a.modeStart {
			if !containsCode(f.modeStart, c) {
				f.modeStart = append(f.modeStart, c)
				f.modeEnd = append(f.modeEnd, a.modeEnd[i])
			}
		}
	}
}

func (f *_FormatInfo) Format() FormatInfo {
	return f
}

func (f *_FormatInfo) String(args ...any) renderer.String {
	var r renderer.String

	if len(f.modeStart) > 0 {
		r = renderer.NewMode(f.enabled, ansi.Sequence{}.Append(f.modeStart...), ansi.Sequence{}.Append(f.modeEnd...), args)
		args = []any{r}
	}
	if f.fg != nil {
		r = renderer.NewFGColor(f.enabled, f.fg, args)
		args = []any{r}
	}
	if f.bg != nil {
		r = renderer.NewBGColor(f.enabled, f.bg, args)
	}
	if r == nil {
		return renderer.Sequence(f.enabled, args)
	}
	return r
}

func New(fmts ...FormatProvider) FormatInfo {
	return newFormat(!NoColors, fmts...)
}

func newFormat(enabled bool, fmts ...FormatProvider) FormatInfo {
	r := &_FormatInfo{enabled: enabled}
	for _, f := range fmts {
		if f != nil {
			r.Add(f.Format())
		}
	}
	return r
}

func Composer(fmts ...FormatProvider) renderer.ComposerFunc {
	return New(fmts...).String
}

////////////////////////////////////////////////////////////////////////////////

type fmtMode byte

var _ Format = (fmtMode)(0)

func (c fmtMode) Format() FormatInfo {
	f := &_FormatInfo{
		enabled: !NoColors,
		modeEnd: []ansi.Code{ansi.MODE_RESET_BOLD + ansi.Simple(c)},
	}
	if c > 0 {
		c++
	}
	f.modeStart = []ansi.Code{ansi.MODE_BOLD + ansi.Simple(c)}
	return f
}

func (c fmtMode) String(args ...any) renderer.String {
	return c.Format().String(args...)
}

const (
	FmtBold fmtMode = iota
	FmtItalic
	FmtUnderline
	FmtBlinking
	fmtModeUnused6
	FmtReverse
	FmtInvisible
	FmtStrikeThrough
)

////////////////////////////////////////////////////////////////////////////////

type fmtFgColor byte

var _ Format = (fmtFgColor)(0)

func (c fmtFgColor) Format() FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		fg:      ansi.COLOR_FG_BLACK + ansi.Simple(c),
	}
}

func (c fmtFgColor) String(args ...any) renderer.String {
	return c.Format().String(args...)
}

const (
	FmtBlack fmtFgColor = iota
	FmtRed
	FmtGreen
	FmtYellow
	FmtBlue
	FmtMagenta
	FmtCyan
	FmtWhite
)

////////////////////////////////////////////////////////////////////////////////

type fmtBgColor byte

var _ Format = (fmtBgColor)(0)

func (c fmtBgColor) Format() FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		bg:      ansi.COLOR_BG_BLACK + ansi.Simple(c),
	}
}

func (c fmtBgColor) String(args ...any) renderer.String {
	return c.Format().String(args...)
}

const (
	FmtBgBlack fmtBgColor = iota
	FmtBgRed
	FmtBgGreen
	FmtBgYellow
	FmtBgBlue
	FmtBgMagenta
	FmtBgCyan
	FmtBgWhite
)

////////////////////////////////////////////////////////////////////////////////

type fmtFgBrightColor byte

var _ Format = (fmtFgBrightColor)(0)

func (c fmtFgBrightColor) Format() FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		fg:      ansi.COLOR_FG_BRIGHT_BLACK + ansi.Simple(c),
	}
}

func (c fmtFgBrightColor) String(args ...any) renderer.String {
	return c.Format().String(args...)
}

const (
	FmtBrightBlack fmtFgBrightColor = iota
	FmtBrightRed
	FmtBrightGreen
	FmtBrightYellow
	FmtBrightBlue
	FmtBrightMagenta
	FmtBrightCyan
	FmtBrightWhite
)

////////////////////////////////////////////////////////////////////////////////

type fmtBgBrightColor byte

var _ Format = (fmtBgBrightColor)(0)

func (c fmtBgBrightColor) Format() FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		bg:      ansi.COLOR_BG_BRIGHT_BLACK + ansi.Simple(c),
	}
}

func (c fmtBgBrightColor) String(args ...any) renderer.String {
	return c.Format().String(args...)
}

const (
	FmtBgBrightBlack fmtBgBrightColor = iota
	FmtBgBrightRed
	FmtBgBrightGreen
	FmtBgBrightYellow
	FmtBgBrightBlue
	FmtBgBrightMagenta
	FmtBgBrightCyan
	FmtBgBrightWhite
)

////////////////////////////////////////////////////////////////////////////////

func FmtRGBColor(r, g, b byte) FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		fg:      ansi.COLOR_FG_RGB.AppendBytes(r, g, b),
	}
}

func FmtBgRGBColor(r, g, b byte) FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		bg:      ansi.COLOR_BG_RGB.AppendBytes(r, g, b),
	}
}

func FmtIdColor(id byte) FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		fg:      ansi.COLOR_FG_ID.AppendBytes(id),
	}
}

func FmtBgIdColor(id byte) FormatInfo {
	return &_FormatInfo{
		enabled: !NoColors,
		bg:      ansi.COLOR_BG_ID.AppendBytes(id),
	}
}
