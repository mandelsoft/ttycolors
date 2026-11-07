package renderer

// Features is a set of terminal features used for rendering.
// Formats requiring a feature use a fallback code,
// if the feature is not enabled for rendering.
type Features uint

const (
	// FeatureFormatting enables output formatting at all.
	// Without this feature, no escape sequences are rendered.
	// Texts are never modified, output format sequences
	// embedded in texts are kept.
	FeatureFormatting Features = 1 << iota
	// FeatureUnderlineStyles enables underline styles (double, curly, dotted, dashed).
	FeatureUnderlineStyles
	// FeatureUnderlineColor enables underline colors.
	FeatureUnderlineColor
)

const (
	// NoFeatures is the empty feature set.
	NoFeatures Features = 0
	// ExtendedFeatures are the features, which cannot be detected
	// and are not supported by all terminals.
	ExtendedFeatures = FeatureUnderlineStyles | FeatureUnderlineColor
	// AllFeatures are all features.
	AllFeatures = FeatureFormatting | ExtendedFeatures
)

// With provides the features f with the features o enabled
// or disabled.
func (f Features) With(o Features, enabled bool) Features {
	if enabled {
		return f | o
	}
	return f &^ o
}

// Has checks whether all features of o are contained in f.
func (f Features) Has(o Features) bool {
	return f&o == o
}
