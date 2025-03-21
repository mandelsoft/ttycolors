package renderer

import (
	"fmt"

	"github.com/mandelsoft/goutils/general"
)

type _Sequence struct {
	self    Renderer
	enabled bool
	seq     []any
}

var _ String = (*_Sequence)(nil)
var _ Renderer = (*_Sequence)(nil)

func newSequence(self Renderer, nested []any) _Sequence {
	return _Sequence{
		self:    self,
		enabled: !NoColors,
		seq:     nested,
	}
}

func Sequence(seq ...any) String {
	s := newSequence(nil, seq)
	s.self = &s
	return &s
}

func (s *_Sequence) Enabled() bool {
	return s.enabled
}

func (s *_Sequence) Enable(m ...bool) {
	s.enabled = general.OptionalDefaultedBool(true, m...)
}

func (s *_Sequence) Append(a ...any) String {
	r := Sequence(append([]any{s.self}, a...)...)
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
