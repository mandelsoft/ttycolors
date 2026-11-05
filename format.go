package ttycolors

import (
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/renderer"
	"github.com/modern-go/reflect2"
)

type FormatInfo = *_FormatInfo

// groupCode describes the code for a group.
type groupCode struct {
	code ansi.CodeSequence
	// feature is the feature required for code.
	feature Features
	// fallback is used if the feature is not enabled.
	fallback ansi.CodeSequence
}

// effective provides the code to use for the given features.
func (c groupCode) effective(f Features) ansi.CodeSequence {
	if f.Has(c.feature) {
		return c.code
	}
	return c.fallback
}

// _FormatInfo describes a composed output format.
// It keeps a code for every configured group.
type _FormatInfo struct {
	renderer.Settings
	codes map[*renderer.Group]groupCode
}

var (
	_ Format      = (FormatInfo)(nil)
	_ ContextInfo = (FormatInfo)(nil)
)

func newGroupFormat(g *renderer.Group, c groupCode) FormatInfo {
	return &_FormatInfo{
		Settings: defaultContext.current(),
		codes:    map[*renderer.Group]groupCode{g: c},
	}
}

// effectiveCodes provides the codes to use according to the
// enabled features.
func (f *_FormatInfo) effectiveCodes() map[*renderer.Group]ansi.CodeSequence {
	r := map[*renderer.Group]ansi.CodeSequence{}
	for g, c := range f.codes {
		if e := c.effective(f.Features()); e != nil {
			r[g] = e
		}
	}
	return r
}

// Start provides the escape sequence establishing the format.
func (f *_FormatInfo) Start() []byte {
	if !f.IsEnabled() {
		return nil
	}
	codes := f.effectiveCodes()
	var seq []ansi.CodeSequence
	for _, g := range renderer.Groups() {
		seq = append(seq, codes[g])
	}
	return ansi.AppendModes(nil, seq...)
}

// End provides the escape sequence resetting all groups
// configured by the format to the terminal default.
func (f *_FormatInfo) End() []byte {
	if !f.IsEnabled() {
		return nil
	}
	codes := f.effectiveCodes()
	var seq []ansi.CodeSequence
	for _, g := range renderer.Groups() {
		if codes[g] != nil {
			seq = append(seq, g.Reset())
		}
	}
	return ansi.AppendModes(nil, seq...)
}

// Add adds the settings of another format.
// Settings for the same group replace the actual ones.
// A nil format is ignored.
func (f *_FormatInfo) Add(a FormatInfo) {
	if a == nil || len(a.codes) == 0 {
		return
	}
	if f.codes == nil {
		f.codes = map[*renderer.Group]groupCode{}
	}
	for g, c := range a.codes {
		f.codes[g] = c
	}
}

func (f *_FormatInfo) Format() FormatInfo {
	return f
}

func (f *_FormatInfo) String(args ...any) renderer.String {
	var r renderer.String

	for _, g := range renderer.Groups() {
		if c, ok := f.codes[g]; ok {
			r = renderer.NewFeatureFormat(f.Settings, g, c.feature, c.code, c.fallback, args)
			args = []any{r}
		}
	}
	if r == nil {
		return renderer.Sequence(f.Settings, args)
	}
	return r
}

// New provides a new composed output format using the
// settings of the default context.
func New(fmts ...FormatProvider) FormatInfo {
	return newFormat(defaultContext.current(), fmts...)
}

// newFormat provides a new composed output format with the
// given settings.
func newFormat(settings renderer.Settings, fmts ...FormatProvider) FormatInfo {
	r := &_FormatInfo{Settings: settings}
	for _, f := range fmts {
		if !reflect2.IsNil(f) {
			r.Add(f.Format())
		}
	}
	return r
}

func Composer(fmts ...FormatProvider) renderer.ComposerFunc {
	return New(fmts...).String
}

////////////////////////////////////////////////////////////////////////////////

// fmtGroup is a native output format providing
// a code for a dedicated group.
type fmtGroup struct {
	group *renderer.Group
	groupCode
}

func newFmtGroup(g *renderer.Group, c ansi.CodeSequence) fmtGroup {
	return fmtGroup{g, groupCode{code: c}}
}

// newFeatureFmtGroup provides a format whose code requires
// a feature. Without the feature, the fallback is used.
func newFeatureFmtGroup(g *renderer.Group, f Features, c, fallback ansi.CodeSequence) fmtGroup {
	return fmtGroup{g, groupCode{code: c, feature: f, fallback: fallback}}
}

var _ Format = fmtGroup{}

func (f fmtGroup) Format() FormatInfo {
	return newGroupFormat(f.group, f.groupCode)
}

func (f fmtGroup) String(args ...any) renderer.String {
	return f.Format().String(args...)
}

// render provides a String for the format using the given settings.
func (f fmtGroup) render(settings renderer.Settings, args []any) String {
	return renderer.NewFeatureFormat(settings, f.group, f.feature, f.code, f.fallback, args)
}

// Bold and Dim are handled as alternative values of the
// highlight group (together with NonHighlighted). Because bold and dim
// are additive on many terminals, they always reset the group first.
var (
	// FmtNonHighlighted is the standard font without bold or dim.
	FmtNonHighlighted = newFmtGroup(renderer.Highlight, ansi.MODE_RESET_BOLD)
	FmtBold           = newFmtGroup(renderer.Highlight, ansi.NewCodeSequence(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD))
	FmtDim            = newFmtGroup(renderer.Highlight, ansi.NewCodeSequence(ansi.MODE_RESET_BOLD, ansi.MODE_DIM))

	FmtItalic        = newFmtGroup(renderer.Italic, ansi.MODE_ITALIC)
	FmtUnderline     = newFmtGroup(renderer.Underline, ansi.MODE_UNDERLINE)
	FmtBlinking      = newFmtGroup(renderer.Blinking, ansi.MODE_BLINK)
	FmtReverse       = newFmtGroup(renderer.Reverse, ansi.MODE_REVERSE)
	FmtInvisible     = newFmtGroup(renderer.Invisible, ansi.MODE_INVIS)
	FmtStrikeThrough = newFmtGroup(renderer.StrikeThrough, ansi.MODE_STRIKE)
	FmtOverline      = newFmtGroup(renderer.Overline, ansi.MODE_OVERLINE)

	FmtNonItalic        = newFmtGroup(renderer.Italic, ansi.MODE_RESET_ITALIC)
	FmtNonUnderline     = newFmtGroup(renderer.Underline, ansi.MODE_RESET_UNDERLINE)
	FmtNonBlinking      = newFmtGroup(renderer.Blinking, ansi.MODE_RESET_BLINK)
	FmtNonReverse       = newFmtGroup(renderer.Reverse, ansi.MODE_RESET_REVERSE)
	FmtNonInvisible     = newFmtGroup(renderer.Invisible, ansi.MODE_RESET_INVIS)
	FmtNonStrikeThrough = newFmtGroup(renderer.StrikeThrough, ansi.MODE_RESET_STRIKE)
	FmtNonOverline      = newFmtGroup(renderer.Overline, ansi.MODE_RESET_OVERLINE)
)

// Underline styles require the feature FeatureUnderlineStyles.
// Without it, a single underline is used.
// Terminals supporting underline styles do not necessarily support all
// of them. For example, GNOME Terminal (VTE 0.60) renders dotted and
// dashed underlines like curly underlines.
var (
	FmtDoubleUnderline = newFeatureFmtGroup(renderer.Underline, FeatureUnderlineStyles, ansi.UNDERLINE_DOUBLE, ansi.MODE_UNDERLINE)
	FmtCurlyUnderline  = newFeatureFmtGroup(renderer.Underline, FeatureUnderlineStyles, ansi.UNDERLINE_CURLY, ansi.MODE_UNDERLINE)
	FmtDottedUnderline = newFeatureFmtGroup(renderer.Underline, FeatureUnderlineStyles, ansi.UNDERLINE_DOTTED, ansi.MODE_UNDERLINE)
	FmtDashedUnderline = newFeatureFmtGroup(renderer.Underline, FeatureUnderlineStyles, ansi.UNDERLINE_DASHED, ansi.MODE_UNDERLINE)
)

////////////////////////////////////////////////////////////////////////////////

// Standard foreground colors.
var (
	FmtBlack   = newFmtGroup(renderer.FG, ansi.COLOR_FG_BLACK)
	FmtRed     = newFmtGroup(renderer.FG, ansi.COLOR_FG_RED)
	FmtGreen   = newFmtGroup(renderer.FG, ansi.COLOR_FG_GREEN)
	FmtYellow  = newFmtGroup(renderer.FG, ansi.COLOR_FG_YELLOW)
	FmtBlue    = newFmtGroup(renderer.FG, ansi.COLOR_FG_BLUE)
	FmtMagenta = newFmtGroup(renderer.FG, ansi.COLOR_FG_MAGENTA)
	FmtCyan    = newFmtGroup(renderer.FG, ansi.COLOR_FG_CYAN)
	FmtWhite   = newFmtGroup(renderer.FG, ansi.COLOR_FG_WHITE)
)

// Standard background colors.
var (
	FmtBgBlack   = newFmtGroup(renderer.BG, ansi.COLOR_BG_BLACK)
	FmtBgRed     = newFmtGroup(renderer.BG, ansi.COLOR_BG_RED)
	FmtBgGreen   = newFmtGroup(renderer.BG, ansi.COLOR_BG_GREEN)
	FmtBgYellow  = newFmtGroup(renderer.BG, ansi.COLOR_BG_YELLOW)
	FmtBgBlue    = newFmtGroup(renderer.BG, ansi.COLOR_BG_BLUE)
	FmtBgMagenta = newFmtGroup(renderer.BG, ansi.COLOR_BG_MAGENTA)
	FmtBgCyan    = newFmtGroup(renderer.BG, ansi.COLOR_BG_CYAN)
	FmtBgWhite   = newFmtGroup(renderer.BG, ansi.COLOR_BG_WHITE)
)

// Bright foreground colors.
var (
	FmtBrightBlack   = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_BLACK)
	FmtBrightRed     = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_RED)
	FmtBrightGreen   = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_GREEN)
	FmtBrightYellow  = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_YELLOW)
	FmtBrightBlue    = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_BLUE)
	FmtBrightMagenta = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_MAGENTA)
	FmtBrightCyan    = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_CYAN)
	FmtBrightWhite   = newFmtGroup(renderer.FG, ansi.COLOR_FG_BRIGHT_WHITE)
)

// Bright background colors.
var (
	FmtBgBrightBlack   = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_BLACK)
	FmtBgBrightRed     = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_RED)
	FmtBgBrightGreen   = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_GREEN)
	FmtBgBrightYellow  = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_YELLOW)
	FmtBgBrightBlue    = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_BLUE)
	FmtBgBrightMagenta = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_MAGENTA)
	FmtBgBrightCyan    = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_CYAN)
	FmtBgBrightWhite   = newFmtGroup(renderer.BG, ansi.COLOR_BG_BRIGHT_WHITE)
)

////////////////////////////////////////////////////////////////////////////////

func fmtRGBColor(r, g, b byte) fmtGroup {
	return newFmtGroup(renderer.FG, ansi.Code_FG_RGB(r, g, b))
}

func fmtBgRGBColor(r, g, b byte) fmtGroup {
	return newFmtGroup(renderer.BG, ansi.Code_BG_RGB(r, g, b))
}

func fmtIdColor(id byte) fmtGroup {
	return newFmtGroup(renderer.FG, ansi.Code_FG_ID(id))
}

func fmtBgIdColor(id byte) fmtGroup {
	return newFmtGroup(renderer.BG, ansi.Code_BG_ID(id))
}

func fmtUnderlineIdColor(id byte) fmtGroup {
	return newFeatureFmtGroup(renderer.UnderlineColor, FeatureUnderlineColor, ansi.Code_UL_ID(id), nil)
}

func fmtUnderlineRGBColor(r, g, b byte) fmtGroup {
	return newFeatureFmtGroup(renderer.UnderlineColor, FeatureUnderlineColor, ansi.Code_UL_RGB(r, g, b), nil)
}

// FmtRGBColor provides a foreground RGB color.
func FmtRGBColor(r, g, b byte) FormatInfo {
	return fmtRGBColor(r, g, b).Format()
}

// FmtBgRGBColor provides a background RGB color.
func FmtBgRGBColor(r, g, b byte) FormatInfo {
	return fmtBgRGBColor(r, g, b).Format()
}

// FmtIdColor provides a foreground color id.
func FmtIdColor(id byte) FormatInfo {
	return fmtIdColor(id).Format()
}

// FmtBgIdColor provides a background color id.
func FmtBgIdColor(id byte) FormatInfo {
	return fmtBgIdColor(id).Format()
}

// FmtUnderlineIdColor provides an underline color id.
// It requires the feature FeatureUnderlineColor, otherwise
// it is ignored.
func FmtUnderlineIdColor(id byte) FormatInfo {
	return fmtUnderlineIdColor(id).Format()
}

// FmtUnderlineRGBColor provides an underline RGB color.
// It requires the feature FeatureUnderlineColor, otherwise
// it is ignored.
func FmtUnderlineRGBColor(r, g, b byte) FormatInfo {
	return fmtUnderlineRGBColor(r, g, b).Format()
}
