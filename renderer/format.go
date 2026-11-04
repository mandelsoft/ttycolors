package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

// _Format renders nested elements with a code for a dedicated group.
type _Format struct {
	_base
	group *group
}

var (
	_ String   = (*_Format)(nil)
	_ Renderer = (*_Format)(nil)
)

func newGroupFormat(g *group, enabled bool, c ansi.Code, nested []any) String {
	f := &_Format{group: g}
	f._base = newFormat(f, enabled, c, nested)
	return f
}

func (f *_Format) Render(enabled bool, st state) (state, string) {
	if !enabled {
		return f._Sequence.Render(enabled, nil)
	}
	prefix := st.Repeat(f.group)
	next, sub := f._Sequence.Render(enabled, st.Set(f.group, f.start))
	next = next.Set(f.group, st.Get(f.group))
	next.RequestEnd(f.group)
	return next, prefix + sub
}
