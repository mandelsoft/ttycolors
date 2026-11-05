// This example shows the underline styles and underline colors.
// They require optional terminal features, which must be enabled
// explicitly. Without them, the styles are rendered as single
// underline and the underline colors are ignored.
//
// The output is colored, if the standard output is a terminal.
package main

import (
	"fmt"
	"os"

	"github.com/mandelsoft/ttycolors"
)

func main() {
	ctx := ttycolors.NewContext(ttycolors.IsTerminal(os.Stdout))
	ctx.EnableFeatures(ttycolors.ExtendedFeatures)

	// underline styles (some terminals, for example GNOME Terminal with
	// VTE 0.60, render dotted and dashed like curly underlines)
	fmt.Printf("%s %s %s %s %s\n",
		ctx.Underline("single"),
		ctx.DoubleUnderline("double"),
		ctx.CurlyUnderline("curly"),
		ctx.DottedUnderline("dotted"),
		ctx.DashedUnderline("dashed"))

	// underline colors given by id or RGB values
	fmt.Printf("%s %s %s\n",
		ctx.Underline(ctx.UnderlineIdColor(196, "red id")),
		ctx.DoubleUnderline(ctx.UnderlineIdColor(46, "green id")),
		ctx.DashedUnderline(ctx.UnderlineRGBColor(0, 128, 255, "blue RGB")))

	// styles and colors are independent of the text color
	fmt.Printf("%s\n", ctx.Blue("a blue text with a ",
		ctx.CurlyUnderline(ctx.UnderlineIdColor(196, "spelling")), " error and a ",
		ctx.DottedUnderline(ctx.UnderlineIdColor(208, "grammar")), " hint"))

	// nested styles restore the outer style
	fmt.Printf("%s\n", ctx.DashedUnderline("dashed with ", ctx.CurlyUnderline("curly"), " and ", ctx.NonUnderline("no"), " underline"))

	// formats can be composed with styles and colors
	spelling := ctx.New(ttycolors.FmtCurlyUnderline, ttycolors.FmtUnderlineIdColor(196))
	fmt.Printf("%s\n", ctx.Sequence("a composed format for ", spelling.String("spelling"), " errors"))

	// without features, styles are rendered as single underline
	// and underline colors are ignored
	plain := ttycolors.NewContext(ttycolors.IsTerminal(os.Stdout))
	fmt.Printf("%s %s\n",
		plain.CurlyUnderline("curly without feature"),
		plain.Underline(plain.UnderlineIdColor(196, "color without feature")))
}
