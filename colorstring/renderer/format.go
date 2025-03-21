package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type _format struct {
	start ansi.Code
	_Sequence
}

func newFormat(self Renderer, start ansi.Code, nested []any) _format {
	return _format{start: start, _Sequence: newSequence(self, nested)}
}
