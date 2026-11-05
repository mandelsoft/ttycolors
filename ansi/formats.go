package ansi

// Code_FG_RGB provides a foreground RGB color.
// It uses separate parameters (38;2;r;g;b) for maximum compatibility
// (see package documentation).
func Code_FG_RGB(r, g, b byte) CodeSequence {
	return Join(COLOR_FG_RGB, Param(r), Param(g), Param(b))
}

// Code_BG_RGB provides a background RGB color.
// It uses separate parameters (48;2;r;g;b) for maximum compatibility
// (see package documentation).
func Code_BG_RGB(r, g, b byte) CodeSequence {
	return Join(COLOR_BG_RGB, Param(r), Param(g), Param(b))
}

// Code_FG_ID provides a foreground color id.
// It uses separate parameters (38;5;id) for maximum compatibility
// (see package documentation).
func Code_FG_ID(id byte) CodeSequence {
	return Join(COLOR_FG_ID, Param(id))
}

// Code_BG_ID provides a background color id.
// It uses separate parameters (48;5;id) for maximum compatibility
// (see package documentation).
func Code_BG_ID(id byte) CodeSequence {
	return Join(COLOR_BG_ID, Param(id))
}

// Code_UL_ID provides an underline color id.
// It uses the sub parameter format (58:5:id), because
// terminals not supporting underline colors would misinterpret
// the semicolon separated arguments as other output modes.
// Package ttycolors uses it only with the feature FeatureUnderlineColor.
func Code_UL_ID(id byte) CodeSequence {
	return Code{COLOR_UL, COLOR_ID, Param(id)}
}

// Code_UL_RGB provides an underline RGB color.
// It uses the sub parameter format (58:2::r:g:b), because
// terminals not supporting underline colors would misinterpret
// the semicolon separated arguments as other output modes.
// Package ttycolors uses it only with the feature FeatureUnderlineColor.
func Code_UL_RGB(r, g, b byte) CodeSequence {
	return Code{COLOR_UL, COLOR_RGB, Omitted, Param(r), Param(g), Param(b)}
}
