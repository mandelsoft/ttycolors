# Color Output for ANSI terminals

This Go package provides support for colorized output
in terms of [ANSI Escape Code](http://en.wikipedia.org/wiki/ANSI_escape_code#Colors).

The library consists of three basic APIs:
- [composition of output formats](#output-format-management)
- [enable colored output on `io.Writer` and `os.StdOut`](#managing-colorized-output).
- [compose colored strings](#managing-colorized-strings)

It supports the following ANSI output formats:
- output modes:
  - Bold
  - Italic
  - Underline
  - Blinking
  - Reverse
  - StrikeThrough
- ANSI standard colors like Blue and Green, etc.
- ANSI color ids
- ANSI RGB color settings
- Foreground and Background Colors
- Bright Colors


The APIs support cascaded output format management.
Foreground and background colors as well as the output modes can be enabled and disabled independently.

The main package `github.com/mandelsoft/ttycolors`
contains the management of native and composed 
output formats, the operation to colorize
an output stream according to those formats and
the composition of colored strings supporting cascaded format handling.

The package `github.com/mandelsoft/ttycolors/ansi`
defines constants for the various ANSI escape codes.

## Output Format Management

The package `github.com/mandelsoft/ttycolors`
provides constants for all native ANSI output formats:
- output modes, for example `FmtBold`.
- standard colors for foreground, background and bright colors, for example `FmtBlue`, `FmtBgBlue`, `FmtBrightBlue` or `FmtBgBrightBlue`.
- color ids using `FmtIdColor(id)` and `FmtBgIdColor(id)`
- RGB colors using `RGBColorId(r,g,b)` and `FmtRGBColorId(r,g,b)`

All those constants provide the composition method `String(...)` usable to compose colored strings (type String).

With the function `New(fmt...)` new
formats can be composed based on any other format, for example `New(FmtBGWhite, FmtRed, FmtBold)`.

### Examples

Predefined standard ANSI formats can be composed to new formats:

```golang
RedOnWhite := ttycolors.New(ttycolors.FmtBgWhite, ttycolors.FmtRed)
```

Those composed formats as well as the standard formats can be used to
create [colorized strings](#managing-colorized-strings)

```golang
str := RedOnWhite.String("red on white")
```
or to directly control an [output format](#managing-colorized-output):

```golang
ttycolors.Set(RedOnWhite, ttycolors.FmtBold)
```

## Managing colorized Output

The package `github.com/mandelsoft/ttycolors`
provides operations to apply output formats 
to `io.Writer` streams.
- `Set(fmt...) error` applies the format set to the standard output (`os.StdOut`)
- `SetFor(w, fmt...)` applies the format set to a dedicated writer (which should write to an ANSI terminal)
- UnSet(fmt...) error reverts the described settings for the standard output again. Output modes, foreground and background colors are handled separately.
- UnSetFor(w, fmt...) error reverts the described settings again for the given writer.

With `NewFormatter(fmt...)` a formatter for the given format is created.
It provides the usual methods for printing:
- Print
- Printf
- Sprint
- Sprint
- Fprint
- Fprintf


### Examples

```golang
ttycolors.Set(ttycolors.FmtBgWhite, ttycolors.FmtRed)
fmt.Printf("red on white")
ttycolors.Set(ttycolors.FmtBlue)
fmt.Printf("blue on white")
ttycolors.UnSet(ttycolors.FmtBgWhite)
fmt.Printf("blue")
ttycolors.UnSet(ttycolors.FmtBlue)
fmt.Printf("\n")
fmt.Printf("back to normal\n")
```

The formatter can be used to apply a format.

```golang
f := ttycolors.NewFormatter(ttycolors.FmtBlue, ttycolors.FmtBold)
str := f.Sprint("an ", ttycolors.Italic("italic"), " string.")
# This is basically identical to
str = ttycolors.New(ttycolors.FmtBlue, ttycolors.FmtBold).String("an ", ttycolors.Italic("italic"), " string.").String()
```

As can be seen, the two APIs for colorized output are combinable

## Managing colorized Strings

The package `github.com/mandelsoft/ttycolors`
let you compose colored strings.
The base type for such a string is `String`.
It can be rendered to a Go string using the `String() string` method.

The package provides functions for all
ANSI output formats providing an appropriate formatted `String` object.

- output modes, for example `Bold`.
- standard colors for foreground, background and bright colors, for example `Blue`, `BgBlue`, `BrightBlue` or `BgBrightBlue`.
- color ids using `IdColor(id)` and `BgIdColor(id)`
- RGB colors using `RGBColor(r,g,b)` and `BgRGBColor(r,g,b)`

Foreground and Background color settings as wells as the output modes are handled separately. This way the Strings added
to a formatting functions can locally use own
modes or colors independently of the surrounding settings.

With the method `Sequence(text...)` a sequence of Strings can be created.
All composition functions access an arbitrary list of Golang strings or other
`String` objects. All other types will be converted to a string using `fmt.Sprintf("%v")`.

### Examples

```golang
str:= ttycolors.Blue("this blue text contains a ", ttycolors.Italic(ttycolors.Red("red"), " italic"), " word")
```

This expression provides a `ttycolors.String` object. In between, there is an italic part with a red word. All the rest is still blue.

A formal colorized string can be used as argument to all composition functions.
For example, with
```golang
str=str.Append(".")
```
it can be extended with a sequence of one or more strings.

To render such an object to a Golang string, the `String()` method can be called:

```golang
s:=str.String()
```

If colors are enabled, it contains the required ANSI escape sequences,
otherwise, the native text is returned.

Because of this method `String` objects can be directly passed to a
function of the `fmt.Printf` family.

For example,

```golang
fmt.Printf("%s\n", str)
```

outputs the color string to the standard output.

A `String` object can be used to enable or disable coloring
by using the `Enable(b...)` method. The default coloring
is defined by the `NoColors` variable, it is defaulted to `true`.

The state of top level `String` object determines the color enablement, regardless what sublevel objects define.

## The Output Context

The package `github.com/mandelsoft/ttycolors` offers plain functions
or format constants for the string composition. They are based
on the variable `NoColors`, which can be used to globally enable or
disable the color handling. It is defaulted to `false` if
`os.StdOut` represents a terminal device.

Additionally, there is a `TTYContext` object offering the same functions as methods.
It can be enabled or disabled, also. The offered methods provide
objects reflecting this colorizing state.
If nothing selected during the context creation, it directly reflects
the state of the `NoColors` variable.

### Examples

```golang

ctx := ttycolors.NewContext(true) // provide a context always enabled

s := ctx.Blue("a blue text with some ", ctx.Italic("italic"), " word")
```

The crucial element here is the first call. It decided, whether coloring
is enabled ot not. Subsequent elements derive this information.
Therefore, it could have been written as follows, also:

```golang
s := ctx.Blue("a blue text with some ", ttycolors.Italic("italic"), " word")
```




