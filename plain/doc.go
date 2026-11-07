// Package plain provides the output formats of package ttycolors
// as functions directly providing formatted Go strings instead of
// ttycolors.String objects. Composed formats are provided by New
// as functions of the same kind.
//
// The functions and formats use the settings (enabled colors and features)
// of the default context of package ttycolors at the time they are called.
// Because its settings are initialized according to the standard output,
// the package functions provide formatted strings only, if the standard
// output is a terminal (or the global settings are changed). For is used
// to get the same functions as methods of a Context bound to another context
// of package ttycolors, for example to provide formatted strings for
// another output independently of the global settings:
//
//	p := plain.For(ttycolors.NewContext(true))
//	s := p.Blue("a ", p.Bold("bold"), " blue")
//
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
//
// SplitAt and SubString split formatted strings into self-contained
// parts keeping their output format, Normalize rewrites the output format
// sequences into minimal codes, and ScanFormat provides a format
// function for the output format established at the end of a formatted
// string (see the functions of the same names in package ttycolors):
//
//	first, second := plain.SplitAt(s, 3)
//	s = plain.Reverse(first) + second
//
//	s = s + plain.ScanFormat(s)(" continued")
package plain
