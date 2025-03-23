package renderer

import (
	"fmt"

	"github.com/mandelsoft/goutils/optionutils"
)

type _Sequence struct {
	self    Renderer
	enabled bool
	seq     []any
}

var _ String = (*_Sequence)(nil)
var _ Renderer = (*_Sequence)(nil)

func newSequence(self Renderer, enabled bool, nested []any) _Sequence {
	return _Sequence{
		self:    self,
		enabled: enabled,
		seq:     nested,
	}
}

func Sequence(enabled bool, nested []any) String {
	s := newSequence(nil, enabled, nested)
	s.self = &s
	return &s
}

func (s *_Sequence) Enabled() bool {
	return s.enabled
}

func (s *_Sequence) Enable(m ...bool) {
	s.enabled = optionutils.BoolOption(m...)
}

func (s *_Sequence) Append(a ...any) String {
	r := Sequence(s.enabled, append([]any{s.self}, a...))
	r.Enable(s.Enabled())
	return r
}

func (s *_Sequence) String() string {
	state, r := s.self.Render(s.enabled, nil)
	return r + state.EmitEndFG() + state.EmitEndBG()
}

func (s *_Sequence) Render(enabled bool, state *state) (*state, string) {
	sub := ""
	r := ""
	for _, e := range s.seq {
		state, sub = render(enabled, state, e)
		r += sub
	}
	return state, r
}

func render(enabled bool, state *state, s any) (*state, string) {
	var r string

	switch e := s.(type) {
	case Renderer:
		return e.Render(enabled, state)
	case string:
		r = e
	default:
		r = fmt.Sprintf("%v", e)
	}
	if enabled {
		return state, state.RepeatAll() + r
	}
	return state, r
}
