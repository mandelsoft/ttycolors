package renderer

// String is a formatted string. It keeps the structure of nested
// formats and texts and is rendered with String.
type String interface {
	Enable(...bool)
	Append(a ...any) String
	String() string

	// CharLen provides the number of characters (runes) of the
	// text, escape sequences embedded in texts are not counted.
	CharLen() int
	// SplitAt splits the String after n characters. Both parts keep
	// the structure of nested formats: a format containing the split
	// position is used for both parts, nested texts without own format
	// still inherit the format of a surrounding format. Rendered one
	// after the other, both parts provide the same output as the
	// original String. Already formatted texts embedded in the String
	// are split with SplitText into self-contained parts. A negative n
	// causes a panic, an n beyond the end provides an empty second part.
	SplitAt(n int) (String, String)
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
