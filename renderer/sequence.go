package renderer

import (
	"fmt"

	"github.com/mandelsoft/ttycolors/ansi"
)

type _Sequence struct {
	Settings
	self Renderer
	seq  []any
}

var (
	_ String      = (*_Sequence)(nil)
	_ Renderer    = (*_Sequence)(nil)
	_ ContextInfo = (*_Sequence)(nil)
)

// newSequence creates a sequence with the given settings.
func newSequence(self Renderer, settings Settings, nested []any) _Sequence {
	return _Sequence{
		Settings: settings,
		self:     self,
		seq:      nested,
	}
}

// Sequence provides a String rendering a sequence of elements
// with the given settings.
func Sequence(settings Settings, nested []any) String {
	s := newSequence(nil, settings, nested)
	s.self = &s
	return &s
}

func (s *_Sequence) Append(a ...any) String {
	return Sequence(s.Settings, append([]any{s.self}, a...))
}

// String renders the sequence. The settings of the
// sequence are used for all nested elements.
func (s *_Sequence) String() string {
	st, r := s.self.Render(s.Settings, state{})
	return r + st.EmitEnds()
}

func (s *_Sequence) Render(cfg Settings, st state) (state, string) {
	sub := ""
	r := ""
	for _, e := range s.seq {
		st, sub = render(cfg, st, e)
		r += sub
	}
	return st, r
}

func render(cfg Settings, st state, s any) (state, string) {
	var r string

	switch e := s.(type) {
	case Renderer:
		return e.Render(cfg, st)
	case string:
		r = e
	case []byte:
		r = string(e)
	default:
		r = fmt.Sprintf("%v", e)
	}
	if r == "" {
		// empty text does not require any formatting
		return st, ""
	}
	if cfg.IsEnabled() {
		prefix := st.Repeat()
		return st.Embedded(r), prefix + r
	}
	// output format sequences embedded in texts are removed,
	// if output formats are disabled.
	return st, ansi.Strip(r)
}
