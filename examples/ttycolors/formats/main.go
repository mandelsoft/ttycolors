// This example shows the composition of output formats, their usage
// to format strings, the Formatter and the direct control of the
// output format of a writer.
//
// The output is colored, if the standard output is a terminal.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors"
)

func main() {
	// composed formats can be used like the predefined ones
	redOnWhite := ttycolors.New(ttycolors.FmtBgWhite, ttycolors.FmtRed)
	warning := ttycolors.New(ttycolors.FmtYellow, ttycolors.FmtBold)

	fmt.Printf("%s\n", redOnWhite.String("red on white with a ", ttycolors.Bold("bold"), " word"))
	fmt.Printf("%s\n", warning.String("warning: ", ttycolors.NonHighlighted("details in standard font")))

	// later formats replace earlier ones for the same mode or color
	fmt.Printf("%s\n", ttycolors.New(ttycolors.FmtBold, ttycolors.FmtDim, ttycolors.FmtBlue).String("dim blue, not bold"))

	// a formatter offers the print functions for a format
	f := ttycolors.NewFormatter(ttycolors.FmtGreen, ttycolors.FmtItalic)
	f.Printf("green italic formatted output: %d", 42)
	fmt.Println()
	fmt.Println(f.Sprint("with nested ", ttycolors.Red("red"), " text"))

	// formats can directly be applied to the standard output
	ttycolors.Set(redOnWhite)
	fmt.Printf("red on white")
	ttycolors.Set(ttycolors.FmtBlue)
	fmt.Printf(" blue on white")
	ttycolors.UnSet(ttycolors.FmtBgWhite)
	fmt.Printf(" blue")
	ttycolors.UnSet(ttycolors.FmtBlue)
	fmt.Printf(" back to normal\n")
}
