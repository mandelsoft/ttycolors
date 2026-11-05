// Package show provides a printer for the examples showing
// code lines together with their results as comment.
package show

import (
	"fmt"
	"io"
	"os"
	"reflect"
	"strings"
	"text/tabwriter"
)

// Printer prints code lines together with their results,
// aligned in columns.
type Printer struct {
	w *tabwriter.Writer
}

// New provides a Printer writing to the standard output.
func New() *Printer {
	return NewFor(os.Stdout)
}

// NewFor provides a Printer writing to the given writer.
func NewFor(w io.Writer) *Printer {
	return &Printer{w: tabwriter.NewWriter(w, 0, 0, 1, ' ', 0)}
}

// Title prints a section title. Pending lines are flushed before.
func (p *Printer) Title(title string) {
	p.Flush()
	fmt.Fprintf(p.w, "\n// %s\n", title)
}

// Code prints a code line without result.
func (p *Printer) Code(code string) {
	fmt.Fprintf(p.w, "%s\n", code)
}

// Call prints a code line together with its results as comment.
// Strings and byte slices are quoted, so that escape sequences are visible.
func (p *Printer) Call(code string, results ...any) {
	var r []string
	for _, e := range results {
		r = append(r, Format(e))
	}
	fmt.Fprintf(p.w, "%s\t// %s\n", code, strings.Join(r, ", "))
}

// Flush prints all pending lines.
func (p *Printer) Flush() {
	p.w.Flush()
}

// Format formats a result value. Strings and byte slices are quoted.
func Format(v any) string {
	switch e := v.(type) {
	case string:
		return fmt.Sprintf("%q", e)
	case []byte:
		return fmt.Sprintf("%q", string(e))
	default:
		if r := reflect.ValueOf(v); r.Kind() == reflect.Slice && r.Type().Elem().Kind() == reflect.Uint8 {
			return fmt.Sprintf("%q", string(r.Bytes()))
		}
		return fmt.Sprintf("%v", e)
	}
}

// R2 provides the two results of a function call as list,
// to be passed to Call.
func R2[A, B any](a A, b B) []any {
	return []any{a, b}
}
