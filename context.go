package ttycolors

import (
	"io"
	"os"

	"github.com/mandelsoft/goutils/optionutils"

	"github.com/mandelsoft/ttycolors/renderer"
)

// NoColors is the global switch for output formatting. It is part of
// the global settings (see SetFeatures) and corresponds to the absence
// of the feature FeatureFormatting: setting NoColors to true disables
// output formatting for the default context, and disabling the feature
// FeatureFormatting for the global settings sets NoColors to true.
// It is initialized to true, if the standard output is not a terminal.
var NoColors = !IsTerminal(os.Stdout)

// globalSettings are the global settings used by the default context.
// The feature FeatureFormatting is kept by the variable NoColors,
// all other features by the embedded settings.
type globalSettings struct {
	renderer.Settings
}

var _ ContextInfo = (*globalSettings)(nil)

func (s *globalSettings) IsEnabled() bool {
	return !NoColors
}

func (s *globalSettings) Enable(b ...bool) {
	s.SetFeatures(s.Features().With(FeatureFormatting, optionutils.BoolOption(b...)))
}

func (s *globalSettings) Features() Features {
	return s.Settings.Features().With(FeatureFormatting, !NoColors)
}

func (s *globalSettings) SetFeatures(f Features) {
	NoColors = !f.Has(FeatureFormatting)
	s.Settings.SetFeatures(f &^ FeatureFormatting)
}

func (s *globalSettings) EnableFeatures(f Features) {
	s.SetFeatures(s.Features() | f)
}

func (s *globalSettings) DisableFeatures(f Features) {
	s.SetFeatures(s.Features() &^ f)
}

// sharedSettings use the settings of a base context until they
// are changed. With the first change, a local copy is created.
type sharedSettings struct {
	base ContextInfo
	// local is nil, as long as the settings are shared.
	local *renderer.Settings
}

var _ ContextInfo = (*sharedSettings)(nil)

// settings provides the actual settings.
func (s *sharedSettings) settings() ContextInfo {
	if s.local == nil {
		return s.base
	}
	return s.local
}

// copy provides the local settings, they are created on first use.
func (s *sharedSettings) copy() *renderer.Settings {
	if s.local == nil {
		c := renderer.NewSettings(s.base.Features())
		s.local = &c
	}
	return s.local
}

func (s *sharedSettings) IsEnabled() bool {
	return s.settings().IsEnabled()
}

func (s *sharedSettings) Enable(b ...bool) {
	s.copy().Enable(b...)
}

func (s *sharedSettings) Features() Features {
	return s.settings().Features()
}

func (s *sharedSettings) SetFeatures(f Features) {
	s.copy().SetFeatures(f)
}

func (s *sharedSettings) EnableFeatures(f Features) {
	s.copy().EnableFeatures(f)
}

func (s *sharedSettings) DisableFeatures(f Features) {
	s.copy().DisableFeatures(f)
}

type TTYContext = *_context

// defaultContext is the default context used by the package functions
// and formats. It uses the global settings: output formatting
// (FeatureFormatting) is enabled, if the standard output is a terminal
// (see NoColors). The extended features cannot be detected and are
// disabled by default.
var defaultContext = &_context{&globalSettings{}}

// SetFeatures sets the global settings (the features of the default
// context used by the package functions and formats).
// Already created contexts, strings and formats are not affected.
func SetFeatures(f Features) {
	defaultContext.SetFeatures(f)
}

// EnableFeatures enables features for the global settings (see SetFeatures).
// For example, EnableFeatures(ExtendedFeatures) enables all features,
// which cannot be detected.
func EnableFeatures(f Features) {
	defaultContext.EnableFeatures(f)
}

// DisableFeatures disables features for the global settings (see SetFeatures).
// For example, DisableFeatures(FeatureFormatting) disables output formatting,
// like setting NoColors to true.
func DisableFeatures(f Features) {
	defaultContext.DisableFeatures(f)
}

// EnabledFeatures provides the global settings (the features
// of the default context).
func EnabledFeatures() Features {
	return defaultContext.Features()
}

// _context is an output context. Output formatting is described
// by the feature FeatureFormatting.
type _context struct {
	settings ContextInfo
}

var _ ContextInfo = (*_context)(nil)

// NewContext provides a new output context with a copy of the global
// settings (see SetFeatures). Later changes of the global settings
// do not affect the context.
// Optionally, output formatting can be enabled or disabled for the
// new context (see Enable).
func NewContext(b ...bool) TTYContext {
	settings := defaultContext.current()
	c := &_context{&settings}
	if len(b) > 0 {
		c.Enable(b...)
	}
	return c
}

// NewSharedContext provides a new output context sharing the settings of
// the given base context (default: the default context, which uses the
// global settings). As long as its settings are not changed, it follows
// the settings of the base context, also later changes of them.
// With the first change, it gets a local copy of the actual settings of
// the base context and is independent of the base context afterwards.
func NewSharedContext(base ...TTYContext) TTYContext {
	b := defaultContext
	if len(base) > 0 && base[0] != nil {
		b = base[0]
	}
	return &_context{&sharedSettings{base: b}}
}

// Share provides a new output context sharing the settings of this
// context (see NewSharedContext).
func (c *_context) Share() TTYContext {
	return NewSharedContext(c)
}

// current provides a copy of the actual settings of the context.
// For the default context, the feature FeatureFormatting is
// derived from NoColors.
func (c *_context) current() renderer.Settings {
	return renderer.NewSettings(c.settings.Features())
}

// IsEnabled checks whether output formatting is enabled for the context.
func (c *_context) IsEnabled() bool {
	return c.settings.IsEnabled()
}

// Enable enables (default) or disables output formatting for the context.
func (c *_context) Enable(b ...bool) {
	c.settings.Enable(b...)
}

// Features provides the features enabled for the context.
func (c *_context) Features() Features {
	return c.settings.Features()
}

// SetFeatures sets the features enabled for the context.
// Strings and formats created by the context copy the features
// when they are created.
func (c *_context) SetFeatures(f Features) {
	c.settings.SetFeatures(f)
}

// EnableFeatures enables features for the context.
func (c *_context) EnableFeatures(f Features) {
	c.settings.EnableFeatures(f)
}

// DisableFeatures disables features for the context.
func (c *_context) DisableFeatures(f Features) {
	c.settings.DisableFeatures(f)
}

////////////////////////////////////////////////////////////////////////////////

// New provides a new composed output format.
func (c *_context) New(f ...FormatProvider) FormatInfo {
	return newFormat(c.current(), f...)
}

// NewFormatter provides a new Formatter, which
// offers the usable Print functions based on
// the context state.
func (c *_context) NewFormatter(f ...FormatProvider) Formatter {
	return newFormatter(c.current(), f...)
}

////////////////////////////////////////////////////////////////////////////////

// Set applies the given formats to the standard output.
// The settings (enabled colors and features) of the context are used,
// the settings of the given formats are ignored.
func (c *_context) Set(f ...FormatProvider) error {
	return c.SetFor(os.Stdout, f...)
}

// SetFor applies the given formats to the given writer.
// The settings (enabled colors and features) of the context are used,
// the settings of the given formats are ignored.
func (c *_context) SetFor(w io.Writer, f ...FormatProvider) error {
	_, err := w.Write(newFormat(c.current(), f...).Start())
	return err
}

// UnSet reverts the settings of the given formats for the standard output.
// The settings (enabled colors and features) of the context are used,
// the settings of the given formats are ignored.
func (c *_context) UnSet(f ...FormatProvider) error {
	return c.UnSetFor(os.Stdout, f...)
}

// UnSetFor reverts the settings of the given formats for the given writer.
// The settings (enabled colors and features) of the context are used,
// the settings of the given formats are ignored.
func (c *_context) UnSetFor(w io.Writer, f ...FormatProvider) error {
	_, err := w.Write(newFormat(c.current(), f...).End())
	return err
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) Sequence(seq ...any) String {
	return renderer.Sequence(c.current(), seq)
}

func (c *_context) String(seq ...any) String {
	return renderer.Sequence(c.current(), seq)
}

func (c *_context) StringWith(f FormatProvider, seq ...any) String {
	return c.Sequence(f.Format().String(seq...))
}

////////////////////////////////////////////////////////////////////////////////

// NonHighlighted provides a string in the standard font
// without bold or dim.
func (c *_context) NonHighlighted(s ...any) String {
	return FmtNonHighlighted.render(c.current(), s)
}
func (c *_context) Bold(s ...any) String {
	// BOLD and DIM should be in the same dimension according to ANSI spec. But modern terminals
	// can dim bold text. Bold Text is thicker and brighter, Dimmed Bold text is dimmer but still thicker.
	return FmtBold.render(c.current(), s)
}
func (c *_context) Dim(s ...any) String {
	return FmtDim.render(c.current(), s)
}
func (c *_context) Italic(s ...any) String {
	return FmtItalic.render(c.current(), s)
}
func (c *_context) Blinking(s ...any) String {
	return FmtBlinking.render(c.current(), s)
}
func (c *_context) Underline(s ...any) String {
	return FmtUnderline.render(c.current(), s)
}
func (c *_context) Reverse(s ...any) String {
	return FmtReverse.render(c.current(), s)
}
func (c *_context) StrikeThrough(s ...any) String {
	return FmtStrikeThrough.render(c.current(), s)
}
func (c *_context) Overline(s ...any) String {
	return FmtOverline.render(c.current(), s)
}
func (c *_context) Invisible(s ...any) String {
	return FmtInvisible.render(c.current(), s)
}
func (c *_context) NonItalic(s ...any) String {
	return FmtNonItalic.render(c.current(), s)
}
func (c *_context) NonBlinking(s ...any) String {
	return FmtNonBlinking.render(c.current(), s)
}
func (c *_context) NonUnderline(s ...any) String {
	return FmtNonUnderline.render(c.current(), s)
}
func (c *_context) NonReverse(s ...any) String {
	return FmtNonReverse.render(c.current(), s)
}
func (c *_context) NonStrikeThrough(s ...any) String {
	return FmtNonStrikeThrough.render(c.current(), s)
}
func (c *_context) NonOverline(s ...any) String {
	return FmtNonOverline.render(c.current(), s)
}
func (c *_context) NonInvisible(s ...any) String {
	return FmtNonInvisible.render(c.current(), s)
}
func (c *_context) DoubleUnderline(s ...any) String {
	return FmtDoubleUnderline.render(c.current(), s)
}
func (c *_context) CurlyUnderline(s ...any) String {
	return FmtCurlyUnderline.render(c.current(), s)
}
func (c *_context) DottedUnderline(s ...any) String {
	return FmtDottedUnderline.render(c.current(), s)
}
func (c *_context) DashedUnderline(s ...any) String {
	return FmtDashedUnderline.render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) Black(s ...any) String {
	return FmtBlack.render(c.current(), s)
}
func (c *_context) Red(s ...any) String {
	return FmtRed.render(c.current(), s)
}
func (c *_context) Green(s ...any) String {
	return FmtGreen.render(c.current(), s)
}
func (c *_context) Yellow(s ...any) String {
	return FmtYellow.render(c.current(), s)
}
func (c *_context) Blue(s ...any) String {
	return FmtBlue.render(c.current(), s)
}
func (c *_context) Magenta(s ...any) String {
	return FmtMagenta.render(c.current(), s)
}
func (c *_context) Cyan(s ...any) String {
	return FmtCyan.render(c.current(), s)
}
func (c *_context) White(s ...any) String {
	return FmtWhite.render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BgBlack(s ...any) String {
	return FmtBgBlack.render(c.current(), s)
}
func (c *_context) BgRed(s ...any) String {
	return FmtBgRed.render(c.current(), s)
}
func (c *_context) BgGreen(s ...any) String {
	return FmtBgGreen.render(c.current(), s)
}
func (c *_context) BgYellow(s ...any) String {
	return FmtBgYellow.render(c.current(), s)
}
func (c *_context) BgBlue(s ...any) String {
	return FmtBgBlue.render(c.current(), s)
}
func (c *_context) BgMagenta(s ...any) String {
	return FmtBgMagenta.render(c.current(), s)
}
func (c *_context) BgCyan(s ...any) String {
	return FmtBgCyan.render(c.current(), s)
}
func (c *_context) BgWhite(s ...any) String {
	return FmtBgWhite.render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BrightBlack(s ...any) String {
	return FmtBrightBlack.render(c.current(), s)
}
func (c *_context) BrightRed(s ...any) String {
	return FmtBrightRed.render(c.current(), s)
}
func (c *_context) BrightGreen(s ...any) String {
	return FmtBrightGreen.render(c.current(), s)
}
func (c *_context) BrightYellow(s ...any) String {
	return FmtBrightYellow.render(c.current(), s)
}
func (c *_context) BrightBlue(s ...any) String {
	return FmtBrightBlue.render(c.current(), s)
}
func (c *_context) BrightMagenta(s ...any) String {
	return FmtBrightMagenta.render(c.current(), s)
}
func (c *_context) BrightCyan(s ...any) String {
	return FmtBrightCyan.render(c.current(), s)
}
func (c *_context) BrightWhite(s ...any) String {
	return FmtBrightWhite.render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) BgBrightBlack(s ...any) String {
	return FmtBgBrightBlack.render(c.current(), s)
}
func (c *_context) BgBrightRed(s ...any) String {
	return FmtBgBrightRed.render(c.current(), s)
}
func (c *_context) BgBrightGreen(s ...any) String {
	return FmtBgBrightGreen.render(c.current(), s)
}
func (c *_context) BgBrightYellow(s ...any) String {
	return FmtBgBrightYellow.render(c.current(), s)
}
func (c *_context) BgBrightBlue(s ...any) String {
	return FmtBgBrightBlue.render(c.current(), s)
}
func (c *_context) BgBrightMagenta(s ...any) String {
	return FmtBgBrightMagenta.render(c.current(), s)
}
func (c *_context) BgBrightCyan(s ...any) String {
	return FmtBgBrightCyan.render(c.current(), s)
}
func (c *_context) BgBrightWhite(s ...any) String {
	return FmtBgBrightWhite.render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) RGBColor(r, g, b byte, s ...any) String {
	return fmtRGBColor(r, g, b).render(c.current(), s)
}
func (c *_context) BgRGBColor(r, g, b byte, s ...any) String {
	return fmtBgRGBColor(r, g, b).render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

func (c *_context) IdColor(id byte, s ...any) String {
	return fmtIdColor(id).render(c.current(), s)
}
func (c *_context) BgIdColor(id byte, s ...any) String {
	return fmtBgIdColor(id).render(c.current(), s)
}

////////////////////////////////////////////////////////////////////////////////

// UnderlineIdColor requires the feature FeatureUnderlineColor,
// otherwise the color is ignored.
func (c *_context) UnderlineIdColor(id byte, s ...any) String {
	return fmtUnderlineIdColor(id).render(c.current(), s)
}

// UnderlineRGBColor requires the feature FeatureUnderlineColor,
// otherwise the color is ignored.
func (c *_context) UnderlineRGBColor(r, g, b byte, s ...any) String {
	return fmtUnderlineRGBColor(r, g, b).render(c.current(), s)
}
