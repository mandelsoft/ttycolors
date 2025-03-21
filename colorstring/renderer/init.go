package renderer

import (
	"os"

	"golang.org/x/term"
)

func IsTerminal(f *os.File) bool {
	return term.IsTerminal(int(f.Fd()))
}

func init() {
	if !IsTerminal(os.Stdout) {
		NoColors = true
	}
}
