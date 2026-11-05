// This example shows the usage of output contexts to explicitly
// enable or disable colors, to enable optional terminal features
// and to control the output format of other writers.
package main

import (
	"fmt"
	"os"

	"github.com/mandelsoft/ttycolors"
)

func main() {
	on := ttycolors.NewContext(true)   // always colored
	off := ttycolors.NewContext(false) // never colored

	fmt.Printf("%s\n", on.Blue("always colored, even if the output is redirected"))
	fmt.Printf("%s\n", off.Blue("never colored"))

	// the outermost string decides about the coloring
	fmt.Printf("%s\n", on.Sequence("colored: ", off.Red("red")))
	fmt.Printf("%s\n", off.Sequence("not colored: ", on.Red("red")))

	// optional terminal features must be enabled explicitly
	styled := ttycolors.NewContext(true)
	styled.EnableFeatures(ttycolors.ExtendedFeatures)
	fmt.Printf("%s\n", styled.CurlyUnderline("a ", styled.UnderlineIdColor(196, "spelling"), " error"))
	fmt.Printf("%s\n", on.CurlyUnderline("without features: a single underline"))

	// a context for another writer, colored only if it is a terminal
	stderr := ttycolors.NewContext(ttycolors.IsTerminal(os.Stderr))
	stderr.SetFor(os.Stderr, ttycolors.FmtRed)
	fmt.Fprintf(os.Stderr, "an error message on stderr")
	stderr.UnSetFor(os.Stderr, ttycolors.FmtRed)
	fmt.Fprintln(os.Stderr)
}
