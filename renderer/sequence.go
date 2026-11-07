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

// element is implemented by all Strings of this package.
type element interface {
	Renderer
	String
}

// text provides the text of an element, which is no Renderer.
func text(e any) string {
	switch e := e.(type) {
	case string:
		return e
	case []byte:
		return string(e)
	default:
		return fmt.Sprintf("%v", e)
	}
}

// charLen provides the number of characters of an element.
func charLen(e any) int {
	if s, ok := e.(element); ok {
		return s.CharLen()
	}
	return ansi.CharLen(text(e))
}

// splitElement splits an element after n characters.
func splitElement(e any, n int) (any, any) {
	if s, ok := e.(element); ok {
		l, r := s.SplitAt(n)
		return l, r
	}
	return SplitText(text(e), n)
}

func checkSplitIndex(n int) {
	if n < 0 {
		panic(fmt.Sprintf("renderer: negative split index %d", n))
	}
}

// CharLen provides the number of characters of all elements.
func (s *_Sequence) CharLen() int {
	n := 0
	for _, e := range s.seq {
		n += charLen(e)
	}
	return n
}

// splitElements splits the elements after n characters. Elements
// completely before the split position are part of the left part,
// elements completely after it of the right part, only the element
// containing the split position is split.
func (s *_Sequence) splitElements(n int) (left, right []any) {
	for i, e := range s.seq {
		if n <= 0 {
			return left, append(right, s.seq[i:]...)
		}
		l := charLen(e)
		if n < l {
			a, b := splitElement(e, n)
			return append(left, a), append(append(right, b), s.seq[i+1:]...)
		}
		left = append(left, e)
		n -= l
	}
	return left, right
}

// SplitAt splits the sequence after n characters (see String).
func (s *_Sequence) SplitAt(n int) (String, String) {
	checkSplitIndex(n)
	l, r := s.splitElements(n)
	return Sequence(s.Settings, l), Sequence(s.Settings, r)
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
	if e, ok := s.(Renderer); ok {
		return e.Render(cfg, st)
	}
	r := text(s)
	if r == "" {
		// empty text does not require any formatting
		return st, ""
	}
	if cfg.IsEnabled() {
		prefix := st.Repeat()
		return st.Embedded(r), prefix + r
	}
	// texts are never modified, output format sequences
	// embedded in texts are kept, if output formats are disabled.
	return st, r
}
