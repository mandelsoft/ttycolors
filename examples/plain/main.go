// This example shows the usage of package plain to directly
// create formatted Go strings.
//
// The output is colored, if the standard output is a terminal.
package main

import (
	"fmt"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/plain"
)

func main() {
	// the functions directly provide Go strings
	var s string = plain.Blue("a blue text with a ", plain.Bold("bold"), " word")
	fmt.Println(s)

	// composed formats are provided as functions of the same kind
	redItalic := plain.New(ttycolors.FmtRed, ttycolors.FmtItalic)
	fmt.Println(redItalic("a red italic text with an ", plain.NonItalic("upright"), " word"))

	// predefined and composed format functions can be used interchangeably
	highlights := map[string]plain.Func{
		"info":    plain.Green,
		"warning": plain.New(ttycolors.FmtYellow, ttycolors.FmtBold),
		"error":   plain.New(ttycolors.FmtWhite, ttycolors.FmtBgRed, ttycolors.FmtBold),
	}
	for _, level := range []string{"info", "warning", "error"} {
		fmt.Printf("%-8s %s\n", level+":", highlights[level]("message with ", plain.Underline("underlined"), " text"))
	}

	// the strings can be used with all string functions
	fmt.Println(plain.Sequence("length of the formatted string: ", fmt.Sprint(len(s))))
}
