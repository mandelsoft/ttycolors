package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// Sequence provides a formatted string (see ttycolors.Sequence).
func Sequence(seq ...any) string {
	return ttycolors.Sequence(seq...).String()
}
