package colorstring

import (
	"github.com/mandelsoft/ttycolors/colorstring/renderer"
)

func Sequence(seq ...any) String {
	return renderer.Sequence(seq...)
}
