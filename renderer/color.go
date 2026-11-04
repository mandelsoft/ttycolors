package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

func NewFGColor(enabled bool, c ansi.Code, nested []any) String {
	return newGroupFormat(fgGroup, enabled, c, nested)
}

func NewBGColor(enabled bool, c ansi.Code, nested []any) String {
	return newGroupFormat(bgGroup, enabled, c, nested)
}
