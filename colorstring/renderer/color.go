package renderer

import (
	"github.com/mandelsoft/ttycolors/ansi"
)

type _FGColor struct {
	_format
}

type _BGColor struct {
	_format
}

func NewFGColor(c ansi.Code, nested []any) String {
	col := &_FGColor{}
	col._format = newFormat(col, c, nested)
	return col
}

func (c *_FGColor) Render(enabled bool, state *state) (*state, string) {
	if enabled {
		next, sub := c._Sequence.Render(enabled, state.SetFG(c.start))
		next = next.SetFG(state.GetFG())
		next.RequestEndFG()
		return next, state.RepeatNonFG() + sub
	}
	return c._Sequence.Render(enabled, nil)
}

func NewBGColor(c ansi.Code, nested []any) String {
	col := &_BGColor{}
	col._format = newFormat(col, c, nested)
	return col
}

func (c *_BGColor) Render(enabled bool, state *state) (*state, string) {
	if enabled {
		next, sub := c._Sequence.Render(enabled, state.SetBG(c.start))
		next = next.SetBG(state.GetBG())
		next.RequestEndBG()
		return next, state.RepeatNonBG() + sub
	}
	return c._Sequence.Render(enabled, state.SetBG(c.start))
}
