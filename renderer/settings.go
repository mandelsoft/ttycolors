package renderer

import (
	"github.com/mandelsoft/goutils/optionutils"
)

// Settings keeps the settings for rendering. It implements
// ContextInfo and can be embedded into types providing
// a rendering context. Enabled output formatting is
// described by the feature FeatureFormatting.
type Settings struct {
	features Features
}

var _ ContextInfo = (*Settings)(nil)

// NewSettings provides settings with the given features.
func NewSettings(f Features) Settings {
	return Settings{features: f}
}

// IsEnabled checks whether output formatting
// (FeatureFormatting) is enabled.
func (s Settings) IsEnabled() bool {
	return s.features.Has(FeatureFormatting)
}

// Enable enables (default) or disables output formatting
// (FeatureFormatting).
func (s *Settings) Enable(b ...bool) {
	if optionutils.BoolOption(b...) {
		s.EnableFeatures(FeatureFormatting)
	} else {
		s.DisableFeatures(FeatureFormatting)
	}
}

// Features provides the enabled features.
func (s Settings) Features() Features {
	return s.features
}

// SetFeatures sets the enabled features.
func (s *Settings) SetFeatures(f Features) {
	s.features = f
}

// EnableFeatures enables the given features.
func (s *Settings) EnableFeatures(f Features) {
	s.features |= f
}

// DisableFeatures disables the given features.
func (s *Settings) DisableFeatures(f Features) {
	s.features &^= f
}
