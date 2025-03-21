package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type _Mode struct {
	_format
	end ansi.Code
}

var _ String = (*_Mode)(nil)
var _ Renderer = (*_Mode)(nil)

func NewMode(start, end ansi.Code, nested []any) String {
	m := &_Mode{end: end}
	m._format = newFormat(m, start, nested)
	return m
}

func (m *_Mode) Render(enabled bool, state *state) (*state, string) {
	if enabled {
		state, sub := m._Sequence.Render(enabled, state)
		return state, ansi.Mode(m.start) + sub + ansi.Mode(m.end)
	}
	return m._Sequence.Render(enabled, state)
}
