package main

import (
	"fmt"

	. "github.com/mandelsoft/ttycolors"
)

func demo() {
	s := Sequence("normal")
	fmt.Printf("%s\n", s)
	fmt.Printf("%s\n", Bold("bold"))
	fmt.Printf("%s\n", StrikeThrough("strike"))
	fmt.Printf("%s\n", Sequence(Italic(Bold("italic bold"), "italic"), "normal"))
	fmt.Printf("%s\n", Sequence(Blue("blue", Green("green", Bold("green bold")), Bold("blue bold"), "blue")))
	fmt.Printf("%s\n", Sequence(Red("red"), "normal"))
	fmt.Printf("%s\n", Sequence(RGBColor(255, 0, 0, "red"), "normal"))
	fmt.Printf("%s\n", Bold(IdColor(100, "bold100"), "normal"))
	fmt.Printf("%s\n", Sequence(BgYellow("yellow", Blue("blue"), "yellow"), "normal"))
	fmt.Printf("%s\n", Sequence("normal", BgYellow("yellow", Blue("blue", Bold("blue bold", Italic("blue bold italic"), "blue bold")), "yellow"), "normal"))
	fmt.Printf("%s\n", Blue("blue", Green("green"), Red("red"), "blue"))

	F := New(FmtBlue, FmtBgYellow, FmtBold, FmtStrikeThrough)

	fmt.Printf("%s\n", Italic("a", F.String("test"), " for all"))

	fmt.Printf("%s\n", BgBlack(White("black")))

	RedOnWhite := New(FmtBgWhite, FmtRed)
	Set(RedOnWhite)
	fmt.Printf("red on white")
	Set(FmtBlue)
	fmt.Printf("blue on white")
	UnSet(FmtBgWhite)
	fmt.Printf("blue")
	UnSet(FmtBlue)
	fmt.Printf("\n")

	fmt.Printf("back to normal\n")

	str := F.String("test")
	// str = Bold("test")
	str.Enable(false)
	if str.String() != "test" {
		fmt.Printf("disable failed\n")
	}
}

func main() {
	demo()
	x := Blue("blue", Red("red"), "blue").String()
	_ = x
	str := Blue("this blue text contains a ", Italic(Red("red"), " italic"), " word")
	x = str.String()
	fmt.Printf("%s\n", str)
	str = str.Append(".")
	fmt.Printf("%s\n", str)

	f := NewFormatter(FmtGreen, FmtBold)
	f.Printf("green bold\n")

	fmt.Printf("%s\n", f.Sprint("an ", Italic("italic"), " string."))
	fmt.Printf("%s\n", New(FmtGreen, FmtBold).String("an ", Italic("italic"), " string.").String())

	fmt.Printf("%s\n", BgGreen("   test   "))
	str = FmtBgRGBColor(0, 255, 0).String("   test   ")
	x = str.String()
	fmt.Printf("%s\n", x)
	str = Blue("blue", BgGreen("green back", Red("red", BgYellow("black back"), "red"), "green back"), "blue")
	x = str.String()
	fmt.Printf("%s\n", x)
}
