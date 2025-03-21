package ttycolors

import (
	"io"
	"os"
)

func Set(f ...FormatProvider) error {
	return set(os.Stdout, New(f...))
}

func SetFor(w io.Writer, f ...FormatProvider) error {
	return set(w, New(f...))
}

func set(w io.Writer, f FormatProvider) error {
	_, err := w.Write(f.Format().Start())
	return err
}

func UnSet(f ...FormatProvider) error {
	return unset(os.Stdout, New(f...))
}

func UnSetFor(w io.Writer, f ...FormatProvider) error {
	return unset(w, New(f...))
}

func unset(w io.Writer, f FormatProvider) error {
	_, err := w.Write(f.Format().End())
	return err
}
