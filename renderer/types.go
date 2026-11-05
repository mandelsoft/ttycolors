package renderer

type String interface {
	Enable(...bool)
	Append(a ...any) String
	String() string
}

// Renderer renders an element according to the
// settings of the outermost rendered String.
type Renderer interface {
	Render(cfg Settings, s state) (state, string)
}

type Composer interface {
	String(...any) String
}

type ComposerFunc = func(...any) String

// ContextInfo is implemented by all objects providing settings,
// like contexts, Strings and formats. Enabled output formatting is
// described by the feature FeatureFormatting.
// For rendering, the settings are copied into a Settings value.
type ContextInfo interface {
	// IsEnabled checks whether output formatting is enabled.
	IsEnabled() bool
	// Enable enables (default) or disables output formatting.
	Enable(b ...bool)

	// Features provides the enabled features.
	Features() Features
	// SetFeatures sets the enabled features.
	SetFeatures(f Features)
	// EnableFeatures enables the given features.
	EnableFeatures(f Features)
	// DisableFeatures disables the given features.
	DisableFeatures(f Features)
}
