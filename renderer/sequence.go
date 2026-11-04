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

var (
	_ String      = (*_Sequence)(nil)
	_ Renderer    = (*_Sequence)(nil)
	_ ContextInfo = (*_Sequence)(nil)
)

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

func (s *_Sequence) IsEnabled() bool {
	return s.enabled
}

func (s *_Sequence) Enable(m ...bool) {
	s.enabled = optionutils.BoolOption(m...)
}

func (s *_Sequence) Append(a ...any) String {
	r := Sequence(s.enabled, append([]any{s.self}, a...))
	r.Enable(s.IsEnabled())
	return r
}

func (s *_Sequence) String() string {
	st, r := s.self.Render(s.enabled, nil)
	return r + st.EmitEnds()
}

func (s *_Sequence) Render(enabled bool, st state) (state, string) {
	sub := ""
	r := ""
	for _, e := range s.seq {
		st, sub = render(enabled, st, e)
		r += sub
	}
	return st, r
}

func render(enabled bool, st state, s any) (state, string) {
	var r string

	switch e := s.(type) {
	case Renderer:
		return e.Render(enabled, st)
	case string:
		r = e
	case []byte:
		r = string(e)
	default:
		r = fmt.Sprintf("%v", e)
	}
	if enabled {
		return st, st.Repeat(nil) + r
	}
	return st, r
}
