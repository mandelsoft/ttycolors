// Package plain provides the output formats of package ttycolors
// as functions directly providing formatted Go strings instead of
// ttycolors.String objects. Composed formats are provided by New
// as functions of the same kind.
//
// The functions and formats use the settings (enabled colors and features)
// of the default context of package ttycolors at the time they are called.
// Because every call is rendered immediately, nested calls are rendered
// with their own settings, and every string resets its output format at
// its end. The surrounding calls track the output format sequences embedded
// in such strings and establish their own output format again afterwards.
// Therefore, nested calls can be used like with package ttycolors:
//
//	s := plain.Blue("a ", plain.Bold("bold"), " blue")
//
//	redItalic := plain.New(ttycolors.FmtRed, ttycolors.FmtItalic)
//	s = redItalic("a ", plain.Bold("bold"), " red italic")
//
// The resulting strings may contain more escape sequences than strings
// composed with package ttycolors, which renders nested formats structurally.
package plain
