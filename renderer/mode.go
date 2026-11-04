package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type _Mode struct {
	_base
	end ansi.Code
}

var _ String = (*_Mode)(nil)
var _ Renderer = (*_Mode)(nil)

func NewMode(enabled bool, start, end ansi.Code, nested []any) String {
	m := &_Mode{end: end}
	m._base = newFormat(m, enabled, start, nested)
	return m
}

func (m *_Mode) Render(enabled bool, st state) (state, string) {
	if enabled {
		st, sub := m._Sequence.Render(enabled, st)
		return st, ansi.Mode(m.start) + sub + ansi.Mode(m.end)
	}
	return m._Sequence.Render(enabled, st)
}

func NewHighlight(enabled bool, c ansi.Code, nested []any) String {
	return newGroupFormat(highlightGroup, enabled, c, nested)
}
