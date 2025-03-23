package ttycolors

func Sequence(seq ...any) String {
	return defaultContext.Sequence(seq...)
}
