// This example shows the composition of colored strings with
// cascaded output formats. Nested strings can locally use their
// own output formats, the surrounding formats are restored afterwards.
//
// The output is colored, if the standard output is a terminal.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors"
)

func main() {
	fmt.Printf("%s\n", ttycolors.Blue("this blue text contains a ", ttycolors.Italic(ttycolors.Red("red"), " italic"), " word"))
	fmt.Printf("%s\n", ttycolors.Blue(ttycolors.Bold("different"), "highlight", ttycolors.Dim("levels")))

	// modes can be switched off for nested strings
	fmt.Printf("%s\n", ttycolors.Italic("an italic text with an ", ttycolors.NonItalic("upright"), " word"))
	fmt.Printf("%s\n", ttycolors.Bold("bold text with ", ttycolors.NonHighlighted("standard"), " text"))

	// foreground and background colors are handled independently
	fmt.Printf("%s\n", ttycolors.BgYellow("yellow background with ", ttycolors.Blue("blue ", ttycolors.Bold("bold")), " text"))

	// bright, id and RGB colors
	fmt.Printf("%s %s %s\n", ttycolors.BrightGreen("bright green"), ttycolors.IdColor(208, "orange id"), ttycolors.RGBColor(128, 0, 255, "purple RGB"))

	// strings can be composed incrementally
	s := ttycolors.Sequence("normal, ", ttycolors.Underline("underlined"))
	s = s.Append(", ", ttycolors.Reverse("reverse"))
	fmt.Printf("%s\n", s)

	// modes not supported by all terminals
	fmt.Printf("%s\n", ttycolors.Sequence(ttycolors.StrikeThrough("strike through"), ", ",
		ttycolors.Overline("overlined"), ", ", ttycolors.Blinking("blinking")))
}
