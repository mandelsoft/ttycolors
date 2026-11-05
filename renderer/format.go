package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

// _Format renders nested elements with a code for a dedicated group.
type _Format struct {
	_Sequence
	start ansi.CodeSequence
	group *Group
	// feature is the feature required for the code.
	feature Features
	// fallback is the code used if the feature is not enabled.
	fallback ansi.CodeSequence
}

var (
	_ String   = (*_Format)(nil)
	_ Renderer = (*_Format)(nil)
)

// NewFormat provides a String rendering the nested elements
// with the code c for the group g.
// Nested elements may use other codes for the same group,
// the code c is restored after them.
// The String uses the given settings.
func NewFormat(settings Settings, g *Group, c ansi.CodeSequence, nested []any) String {
	return NewFeatureFormat(settings, g, NoFeatures, c, nil, nested)
}

// NewFeatureFormat provides a String like NewFormat, whose code c
// requires the given features. If the features are not enabled
// for rendering, the fallback code is used instead. A nil fallback
// renders the nested elements without setting the group.
func NewFeatureFormat(settings Settings, g *Group, features Features, c, fallback ansi.CodeSequence, nested []any) String {
	if g == nil {
		panic("renderer.NewFormat: group required")
	}
	f := &_Format{start: c, group: g, feature: features, fallback: fallback}
	f._Sequence = newSequence(f, settings, nested)
	return f
}

// code provides the code to use for the given rendering settings.
func (f *_Format) code(cfg Settings) ansi.CodeSequence {
	if cfg.features.Has(f.feature) {
		return f.start
	}
	return f.fallback
}

func (f *_Format) Render(cfg Settings, st state) (state, string) {
	c := f.code(cfg)
	if !cfg.IsEnabled() || c == nil {
		return f._Sequence.Render(cfg, st)
	}
	// the codes are rendered together with the first nested text,
	// so that all pending changes are combined into one sequence.
	next, sub := f._Sequence.Render(cfg, st.Set(f.group, c))
	if sub == "" {
		// no nested text, so nothing has been rendered and
		// the format is a no-op.
		return st, ""
	}
	// restore the outer code, it is rendered together
	// with the next text or reset at the end, if required.
	next = next.Set(f.group, st.Get(f.group))
	return next, sub
}
