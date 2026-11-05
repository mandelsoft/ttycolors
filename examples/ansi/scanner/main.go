// This example shows the scanning of escape sequences with
// package ansi. It prints the code together with its results.
package main

import (
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/examples/internal/show"
)

func main() {
	p := show.New()
	s := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"

	p.Title("leading sequences")
	p.Code(`s := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"`)
	p.Call(`ansi.ScanSequence([]byte(s))`, show.R2(ansi.ScanSequence([]byte(s)))...)
	p.Call(`ansi.ScanSequence([]byte("\x1b[38;5;208;4:3m"))`, show.R2(ansi.ScanSequence([]byte("\x1b[38;5;208;4:3m")))...)
	p.Call(`ansi.ScanSequence([]byte("\x1b[2K"))`, show.R2(ansi.ScanSequence([]byte("\x1b[2K")))...)
	p.Call(`ansi.ScanSequence([]byte("text"))`, show.R2(ansi.ScanSequence([]byte("text")))...)

	p.Title("complete strings")
	p.Call(`ansi.Scan(s)`, ansi.Scan(s))
	p.Call(`ansi.Scan("\x1b[1mbold\x1b[22m and \x1b[3mitalic\x1b[23m")`, ansi.Scan("\x1b[1mbold\x1b[22m and \x1b[3mitalic\x1b[23m"))

	p.Title("effects: areas set and reset")
	p.Call(`ansi.ScanSequenceEffect([]byte("\x1b[0;1m"))`, show.R2(ansi.ScanSequenceEffect([]byte("\x1b[0;1m")))...)
	p.Call(`ansi.ScanEffect("\x1b[1mbold\x1b[22m")`, ansi.ScanEffect("\x1b[1mbold\x1b[22m"))
	p.Call(`ansi.ScanEffect("\x1b[31mred\x1b[1m")`, ansi.ScanEffect("\x1b[31mred\x1b[1m"))
	p.Flush()
}
