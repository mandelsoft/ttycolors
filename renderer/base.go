package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type _base struct {
	start ansi.Code
	_Sequence
}

func newFormat(self Renderer, enabled bool, start ansi.Code, nested []any) _base {
	return _base{start: start, _Sequence: newSequence(self, enabled, nested)}
}
