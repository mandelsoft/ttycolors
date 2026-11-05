// This example shows the handling of strings containing escape
// sequences with package ansi. It prints the code together with
// its results.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/examples/internal/show"
)

func main() {
	p := show.New()
	s := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"

	p.Title("runes instead of bytes")
	p.Code(`s := ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED) + "Hällo" + ansi.Mode(ansi.MODE_RESET_ALL) + " world"`)
	p.Call(`len(s)`, len(s))
	p.Call(`ansi.CharLen(s)`, ansi.CharLen(s))
	p.Call(`ansi.Strip(s)`, ansi.Strip(s))

	p.Title("indices and substrings")
	p.Call(`ansi.CharIndex(s, 1)`, ansi.CharIndex(s, 1))
	p.Call(`ansi.CharIndexBefore(s, 5)`, ansi.CharIndexBefore(s, 5))
	p.Call(`ansi.ByteIndex([]byte(s), 'w')`, ansi.ByteIndex([]byte(s), 'w'))
	p.Call(`ansi.SubString(s, 1, 4)`, ansi.SubString(s, 1, 4))
	p.Call(`ansi.SubString(s, 2, 2)`, ansi.SubString(s, 2, 2))
	p.Call(`ansi.SplitAt(s, 5)`, show.R2(ansi.SplitAt(s, 5))...)

	p.Title("limitations: other control sequences are text")
	p.Call(`ansi.CharLen("\x1b[2Kab")`, ansi.CharLen("\x1b[2Kab"))
	p.Call(`ansi.Strip("\x1b[2K\x1b[1mab")`, ansi.Strip("\x1b[2K\x1b[1mab"))
	p.Flush()

	fmt.Printf("\nrendered: %s (%d runes)\n", s, ansi.CharLen(s))
}
