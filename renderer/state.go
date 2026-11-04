package renderer

import (
	"slices"

	"github.com/mandelsoft/ttycolors/ansi"
)

// group describes a kind of output format, whose settings
// replace each other, for example the foreground color.
// It is shared by all states, the state of a group
// is kept by groupState.
type group struct {
	reset ansi.Code
}

var (
	// groups is the list of all groups. The order of creation
	// defines the order the groups are rendered.
	groups []*group
	// groupIndex maps a group to its index in a state.
	groupIndex = map[*group]int{}
)

func newGroup(reset ansi.Code) *group {
	g := &group{reset: reset}
	groupIndex[g] = len(groups)
	groups = append(groups, g)
	return g
}

// The groups must be declared in a single block to
// assure a well-defined creation order.
var (
	fgGroup        = newGroup(ansi.COLOR_FG_DEFAULT)
	bgGroup        = newGroup(ansi.COLOR_BG_DEFAULT)
	highlightGroup = newGroup(ansi.MODE_RESET_BOLD)
)

////////////////////////////////////////////////////////////////////////////////

// groupState is the rendering state of a group.
type groupState struct {
	// code is the active code, nil means terminal default.
	code     ansi.Code
	rendered bool
	// pending indicates a required reset to the terminal default.
	pending bool
}

// repeat appends the code required to establish the
// state of the group g, if not yet rendered.
func (e *groupState) repeat(g *group, seq ansi.Sequence) ansi.Sequence {
	switch {
	case e.code != nil && !e.rendered:
		e.rendered = true
		e.pending = false
		return seq.Append(e.code)
	case e.pending:
		e.pending = false
		return seq.Append(g.reset)
	}
	return seq
}

func (e *groupState) requestEnd() {
	if e.code == nil {
		e.pending = true
	}
}

func (e *groupState) emitEnd(g *group) string {
	if !e.pending {
		return ""
	}
	e.pending = false
	return ansi.Mode(g.reset)
}

////////////////////////////////////////////////////////////////////////////////

// state is the rendering state for all groups, indexed
// according to groupIndex.
// A nil state describes the terminal default for all groups.
// Copies of a state share the group states, use Set
// to provide a modified state.
type state []groupState

// Get returns the active code for a group.
func (s state) Get(g *group) ansi.Code {
	if s == nil {
		return nil
	}
	return s[groupIndex[g]].code
}

// Set provides a new state with the given code for a group.
// The actual state is not modified.
func (s state) Set(g *group, c ansi.Code) state {
	n := slices.Clone(s)
	if n == nil {
		n = make(state, len(groups))
	}
	e := &n[groupIndex[g]]
	e.code = c
	e.rendered = false
	return n
}

// RequestEnd requests a reset of a group to the
// terminal default, if no code is active for it.
func (s state) RequestEnd(g *group) {
	if s != nil {
		s[groupIndex[g]].requestEnd()
	}
}

// Repeat provides the escape sequence required to establish
// the state of all groups except the given one.
// A nil group means all groups.
func (s state) Repeat(except *group) string {
	if s == nil {
		return ""
	}
	var seq ansi.Sequence
	for i, g := range groups {
		if g != except {
			seq = s[i].repeat(g, seq)
		}
	}
	if seq == nil {
		return ""
	}
	return ansi.Mode(seq)
}

// EmitEnds provides the escape sequences for all pending resets.
func (s state) EmitEnds() string {
	if s == nil {
		return ""
	}
	r := ""
	for i, g := range groups {
		r += s[i].emitEnd(g)
	}
	return r
}
