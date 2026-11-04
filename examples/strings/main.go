package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors"
)

func main() {
	fmt.Printf("%s\n", ttycolors.Blue("this blue text contains a ", ttycolors.Italic(ttycolors.Red("red"), " italic"), " word"))
	fmt.Printf("%s\n", ttycolors.Blue(ttycolors.Bold("different"), "highlight", ttycolors.Dim("levels")))
}
