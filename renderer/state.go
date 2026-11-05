package renderer

import (
	"bytes"
	"slices"

	"github.com/mandelsoft/ttycolors/ansi"
)

// Group describes a kind of output format, whose settings
// replace each other, for example the foreground color.
// It is shared by all states, the state of a group
// is kept by a field of type groupState in a state.
// Groups are only provided by the predefined variables.
//
// The set of groups is static. To introduce a new group:
//
//  1. Add a field of type groupState to the type state.
//  2. Define the group as a package variable with its reset code,
//     the covered output format area (an area of package ansi, which
//     must not be covered by another group, so a new area has to be
//     added to ansi.Covered, if required) and the accessor for the
//     new field of the state.
//  3. Add the group to the list groups at the position, which
//     defines its render order.
//  4. Provide formats for the group (see NewFormat), for example
//     format constants and functions in package ttycolors and the
//     corresponding functions in package plain.
//
// The tests verify, that every field of the state is used by exactly
// one group and that the groups cover all output format areas.
type Group struct {
	reset ansi.CodeSequence
	// covered is the output format area described by the group.
	// It is used to track output format sequences embedded in texts.
	covered ansi.Covered
	// field provides the state of the group in a state.
	field func(s *state) *groupState
}

// Reset provides the code resetting the group to the
// terminal default.
func (g *Group) Reset() ansi.CodeSequence {
	return g.reset
}

// The groups (see Group for the steps to introduce a new group).
var (
	FG = &Group{ansi.COLOR_FG_DEFAULT, ansi.S_FG_COLOR,
		func(s *state) *groupState { return &s.fg }}
	BG = &Group{ansi.COLOR_BG_DEFAULT, ansi.S_BG_COLOR,
		func(s *state) *groupState { return &s.bg }}
	Highlight = &Group{ansi.MODE_RESET_BOLD, ansi.S_HIGHLIGHT,
		func(s *state) *groupState { return &s.highlight }}
	Italic = &Group{ansi.MODE_RESET_ITALIC, ansi.S_ITALIC,
		func(s *state) *groupState { return &s.italic }}
	Underline = &Group{ansi.MODE_RESET_UNDERLINE, ansi.S_UNDERLINE,
		func(s *state) *groupState { return &s.underline }}
	Blinking = &Group{ansi.MODE_RESET_BLINK, ansi.S_BLINK,
		func(s *state) *groupState { return &s.blinking }}
	Reverse = &Group{ansi.MODE_RESET_REVERSE, ansi.S_REVERSE,
		func(s *state) *groupState { return &s.reverse }}
	Invisible = &Group{ansi.MODE_RESET_INVIS, ansi.S_INVIS,
		func(s *state) *groupState { return &s.invisible }}
	StrikeThrough = &Group{ansi.MODE_RESET_STRIKE, ansi.S_STRIKE,
		func(s *state) *groupState { return &s.strike }}
	Overline = &Group{ansi.MODE_RESET_OVERLINE, ansi.S_OVERLINE,
		func(s *state) *groupState { return &s.overline }}
	// UnderlineColor requires the feature FeatureUnderlineColor.
	UnderlineColor = &Group{ansi.COLOR_UL_DEFAULT, ansi.S_UL_COLOR,
		func(s *state) *groupState { return &s.ulColor }}
)

// groups is the list of all groups. It defines the order
// the groups are rendered.
var groups = []*Group{FG, BG, Highlight, Italic, Underline, Blinking, Reverse, Invisible, StrikeThrough, Overline, UnderlineColor}

// Groups provides the list of all groups in render order.
func Groups() []*Group {
	return slices.Clone(groups)
}

////////////////////////////////////////////////////////////////////////////////

// groupState is the rendering state of a group.
type groupState struct {
	// code is the desired code, nil means terminal default.
	code ansi.CodeSequence
	// current is the code last sent to the terminal,
	// nil means terminal default.
	current ansi.CodeSequence
	// unknown indicates that the terminal state has been changed
	// by output format sequences embedded in a text.
	unknown bool
}

// sameCode checks whether two codes have the same encoding.
func sameCode(a, b ansi.CodeSequence) bool {
	if a == nil || b == nil {
		return a == nil && b == nil
	}
	return bytes.Equal(a.Bytes(), b.Bytes())
}

// isDefault checks whether the code c describes the terminal
// default for the group g.
func (g *Group) isDefault(c ansi.CodeSequence) bool {
	return c == nil || sameCode(c, g.reset)
}

// repeat appends the code required to establish the desired
// state of the group g on the terminal, if it differs from the
// current terminal state.
func (e *groupState) repeat(g *Group, codes []ansi.CodeSequence) []ansi.CodeSequence {
	if e.code == nil {
		if e.unknown || !g.isDefault(e.current) {
			codes = append(codes, g.reset)
		}
	} else if e.unknown || !sameCode(e.code, e.current) {
		codes = append(codes, e.code)
	}
	e.current = e.code
	e.unknown = false
	return codes
}

// emitEnd provides the reset code for the group g, if
// the terminal state is not the terminal default.
func (e *groupState) emitEnd(g *Group) ansi.CodeSequence {
	c, unknown := e.current, e.unknown
	e.current = nil
	e.unknown = false
	if !unknown && g.isDefault(c) {
		return nil
	}
	return g.reset
}

////////////////////////////////////////////////////////////////////////////////

// state is the rendering state for all groups. It keeps the desired
// code and the code last sent to the terminal for every group.
// It is a plain value, copying a state is just an assignment.
// The zero value describes the terminal default for all groups.
type state struct {
	fg, bg, highlight, italic, underline,
	blinking, reverse, invisible, strike, overline, ulColor groupState
}

// Get returns the desired code for a group.
func (s state) Get(g *Group) ansi.CodeSequence {
	return g.field(&s).code
}

// Set provides a new state with the given desired code for a group.
// The actual state is not modified.
func (s state) Set(g *Group, c ansi.CodeSequence) state {
	g.field(&s).code = c
	return s
}

// Embedded provides a state reflecting output format sequences
// embedded in a text. Groups reset by the sequences are in terminal
// default state afterwards. The terminal state of all groups set by
// the sequences becomes unknown, so that the desired state is
// established again for the next text, and they are reset at the end.
// The actual state is not modified.
func (s state) Embedded(text string) state {
	effect := ansi.ScanEffect(text)
	if effect.Covered() == ansi.S_NONE {
		return s
	}
	for _, g := range groups {
		e := g.field(&s)
		switch {
		case g.covered&effect.Set != 0:
			e.unknown = true
		case g.covered&effect.Reset != 0:
			e.current = nil
			e.unknown = false
		}
	}
	return s
}

// Repeat provides the escape sequences required to establish
// the desired state of all groups on the terminal and updates
// the terminal state. Only groups whose desired code differs from
// the terminal state are considered. All changes are combined into
// one sequence (see ansi.Modes).
func (s *state) Repeat() string {
	var codes []ansi.CodeSequence
	for _, g := range groups {
		codes = g.field(s).repeat(g, codes)
	}
	return ansi.Modes(codes...)
}

// EmitEnds provides the escape sequence resetting all groups,
// whose terminal state is not the terminal default, and updates
// the terminal state. All resets are combined into one sequence.
func (s *state) EmitEnds() string {
	var codes []ansi.CodeSequence
	for _, g := range groups {
		if c := g.field(s).emitEnd(g); c != nil {
			codes = append(codes, c)
		}
	}
	return ansi.Modes(codes...)
}
