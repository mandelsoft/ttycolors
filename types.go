package ttycolors

import (
	"io"

	"github.com/mandelsoft/ttycolors/renderer"
)

type String = renderer.String
type Renderer renderer.Renderer
type StringComposer = renderer.Composer
type ComposerFunc = renderer.ComposerFunc
type ContextInfo = renderer.ContextInfo

// Features is a set of optional terminal features.
// Formats requiring a feature, which is not enabled,
// use a fallback.
type Features = renderer.Features

const (
	// FeatureFormatting enables output formatting at all.
	// Without it, no escape sequences are generated. Texts are never
	// modified, output format sequences embedded in texts are kept
	// (see StripIfDisabled).
	FeatureFormatting = renderer.FeatureFormatting
	// FeatureUnderlineStyles enables underline styles (double, curly, dotted, dashed).
	// Without it, they are rendered as single underline.
	FeatureUnderlineStyles = renderer.FeatureUnderlineStyles
	// FeatureUnderlineColor enables underline colors.
	// Without it, they are ignored.
	FeatureUnderlineColor = renderer.FeatureUnderlineColor

	// NoFeatures is the empty feature set.
	NoFeatures = renderer.NoFeatures
	// ExtendedFeatures are the features, which cannot be detected
	// and are not supported by all terminals. They are disabled
	// by default and can be enabled with EnableFeatures(ExtendedFeatures).
	ExtendedFeatures = renderer.ExtendedFeatures
	// AllFeatures are all features.
	AllFeatures = renderer.AllFeatures
)

type FormatProvider interface {
	Format() FormatInfo
}

type Format interface {
	FormatProvider
	renderer.Composer
}

type Formatter interface {
	FormatProvider
	StringComposer
	ContextInfo

	Print(i ...interface{}) (n int, err error)
	Printf(string, ...interface{}) (n int, err error)
	Sprint(i ...interface{}) string
	Sprintf(string, ...interface{}) string
	Fprint(io.Writer, ...interface{}) (n int, err error)
	Fprintf(io.Writer, string, ...interface{}) (n int, err error)
}
