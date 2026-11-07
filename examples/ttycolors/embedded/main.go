// This example shows the embedding of already formatted strings,
// for example the output of other tools, into colored strings.
// The output format sequences embedded in such strings are tracked,
// so that the surrounding output format is established again afterwards.
//
// The example prints the code together with the generated escape
// sequences, followed by the rendered result.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/examples/internal/show"
)

func main() {
	ctx := ttycolors.NewContext(true)
	p := show.New()

	p.Title("already formatted strings")
	red := ctx.Red("red").String()
	p.Call(`red := ctx.Red("red").String()`, red)
	s1 := ctx.Blue("blue ", red, " blue again")
	p.Call(`ctx.Blue("blue ", red, " blue again")`, s1.String())

	p.Title("formats left set by embedded strings are reset")
	foreign := "\x1b[1mbold from elsewhere"
	s2 := ctx.Blue("blue ", foreign, " blue, but not bold")
	p.Call(`ctx.Blue("blue ", "\x1b[1mbold from elsewhere", " blue, but not bold")`, s2.String())

	p.Title("embedded resets are recognized")
	s3 := ctx.Blue("blue ", "\x1b[0mreset", " blue again")
	p.Call(`ctx.Blue("blue ", "\x1b[0mreset", " blue again")`, s3.String())

	p.Title("embedded sequences are kept without colors")
	off := ttycolors.NewContext(false)
	s4 := off.Blue("blue ", red, " blue again")
	p.Code(`off := ttycolors.NewContext(false)`)
	p.Call(`off.Blue("blue ", red, " blue again")`, s4.String())

	p.Title("embedded sequences are removed on request without colors")
	s5 := off.Blue("blue ", off.StripIfDisabled(red), " blue again")
	p.Call(`off.Blue("blue ", off.StripIfDisabled(red), " blue again")`, s5.String())
	p.Flush()

	fmt.Printf("\nrendered:\n%s\n%s\n%s\n%s\n%s\n", s1, s2, s3, s4, s5)
}
