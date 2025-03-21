package renderer

var NoColors = false

type String interface {
	Enable(...bool)
	Append(a ...any) String
	String() string
}

type Renderer interface {
	Render(enabled bool, s *state) (*state, string)
}

type Composer interface {
	String(...any) String
}

type ComposerFunc = func(...any) String
