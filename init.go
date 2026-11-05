package ttycolors

import (
	"os"

	"golang.org/x/term"
)

// IsTerminal checks whether a file is a terminal.
func IsTerminal(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}
