package plain

import (
	"github.com/mandelsoft/ttycolors"
)

// Context provides the functions of this package for a context of
// package ttycolors. They use the settings (enabled output formatting
// and features) of the context at the time they are called, like the
// methods of the context. The package functions use the default context
// of package ttycolors, which depends on the global settings.
// A Context can be used to provide formatted strings independently of
// the global settings, for example for another output than the standard
// output:
//
//	p := plain.For(ttycolors.NewContext(true))
//	s := p.Blue("a ", p.Bold("bold"), " blue")
type Context = *_context

// _context implements Context for a context of package ttycolors.
type _context struct {
	ctx ttycolors.TTYContext
}

// defaultContext is used by the package functions.
var defaultContext = For(nil)

// For provides the functions of this package for a context of
// package ttycolors. A nil context means the default context.
func For(ctx ttycolors.TTYContext) Context {
	if ctx == nil {
		ctx = ttycolors.DefaultContext()
	}
	return &_context{ctx}
}

// TTYContext provides the context of package ttycolors used
// by the Context.
func (c *_context) TTYContext() ttycolors.TTYContext {
	return c.ctx
}
