// This example shows the creation of escape sequences with
// package ansi. It prints the code together with its results.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/examples/internal/show"
)

func main() {
	p := show.New()

	p.Title("parameters, codes and code sequences")
	p.Call(`ansi.MODE_BOLD.Bytes()`, ansi.MODE_BOLD.Bytes())
	p.Call(`ansi.UNDERLINE_CURLY.Bytes()`, ansi.UNDERLINE_CURLY.Bytes())
	p.Call(`ansi.NewCodeSequence(38, 5, 208).Bytes()`, ansi.NewCodeSequence(38, 5, 208).Bytes())
	p.Call(`ansi.Join(ansi.MODE_ITALIC, ansi.Code_FG_ID(4)).Bytes()`, ansi.Join(ansi.MODE_ITALIC, ansi.Code_FG_ID(4)).Bytes())
	p.Call(`ansi.Join(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY).Codes()`, ansi.Join(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY).Codes())

	p.Title("escape sequences")
	p.Call(`ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED)`, ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED))
	p.Call(`ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(10, 20, 30))`, ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(10, 20, 30)))
	p.Call(`ansi.Mode()`, ansi.Mode())
	p.Call(`ansi.Mode(nil)`, ansi.Mode(nil))

	p.Title("sub parameters are put into separate sequences by Modes")
	p.Call(`ansi.Mode(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY)`, ansi.Mode(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY))
	p.Call(`ansi.Modes(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY, ansi.Code_UL_ID(1))`, ansi.Modes(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY, ansi.Code_UL_ID(1)))

	p.Title("appending to a byte sequence")
	p.Call(`ansi.AppendMode(ansi.Sequence("x"), ansi.MODE_ITALIC)`, ansi.AppendMode(ansi.Sequence("x"), ansi.MODE_ITALIC))
	p.Flush()

	fmt.Printf("\nrendered: %sbold red%s %sorange on dark blue%s\n",
		ansi.Mode(ansi.MODE_BOLD, ansi.COLOR_FG_RED), ansi.Mode(),
		ansi.Mode(ansi.Code_FG_ID(208), ansi.Code_BG_RGB(10, 20, 30)), ansi.Mode())
}
