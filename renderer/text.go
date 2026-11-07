package renderer

import (
	"strings"
	"unicode/utf8"

	"github.com/mandelsoft/ttycolors/ansi"
)

// SplitText splits a text containing escape sequences after n
// characters (runes). It panics for a negative n, an n beyond the end
// provides an empty second part.
//
// The output format sequences of the text are interpreted as follows:
// a code sets the format of its group, a reset (of the group or of all
// modes) ends a format set before, and the group is untouched again.
// A reset code for a group not set before explicitly switches it off.
//
// Both parts are self-contained: the first part resets the groups
// still set at the split position, the second part starts with the
// output format established at the split position. It sets the set
// groups again and resets the groups explicitly reset before, so that
// its text keeps its output format, if it is embedded into an outer
// format. Groups untouched up to the split position are not affected,
// the text still inherits their outer format.
//
// Output format sequences directly at the split position are replaced
// by these codes, other escape sequences there belong to the first part.
// No codes are added at the start or the end of the text. Codes not
// affecting an output format area (for example fonts) are not tracked,
// sequences consisting only of such codes at the split position belong
// to the second part.
func SplitText(s string, n int) (string, string) {
	checkSplitIndex(n)

	var st, before state
	var ctrl, untracked []byte
	data := []byte(s)
	cnt, i, p := 0, 0, 0
	for {
		// escape sequences before the next rune
		p, before, ctrl, untracked = i, st, nil, nil
		for i < len(data) {
			l, e := ansi.ScanSequenceEffect(data[i:])
			if l == 0 {
				break
			}
			seq := data[i : i+l]
			switch {
			case e.Set&ansi.S_CONTROL != 0:
				ctrl = append(ctrl, seq...)
			case e.Covered() == ansi.S_NONE:
				untracked = append(untracked, seq...)
			default:
				st.scan(string(seq))
			}
			i += l
		}
		if cnt == n || i >= len(data) {
			break
		}
		_, l := utf8.DecodeRune(data[i:])
		i += l
		cnt++
	}

	switch {
	case i >= len(data):
		return s + st.EmitEnds(), ""
	case cnt == 0:
		return "", s
	default:
		return s[:p] + string(ctrl) + before.EmitEnds(), st.reopen() + string(untracked) + s[i:]
	}
}

// NormalizeText rewrites the output format sequences of a text into the
// minimal codes required for the changes of the output format (see
// SplitText for the interpretation of resets). A reset of all modes ending
// some formats is replaced by the resets of these groups, redundant codes
// are omitted, and consecutive output format sequences are combined.
// A reset of all modes is only used, if all groups are set and have to
// be reset. Texts and other escape sequences are kept. Codes not
// affecting an output format area (for example fonts) are not tracked,
// sequences consisting only of such codes are kept, but such codes in
// sequences with other codes are omitted.
func NormalizeText(s string) string {
	var scanned, out state
	var b strings.Builder

	flush := func() {
		for _, g := range groups {
			g.field(&out).code = scanned.Get(g)
		}
		b.WriteString(out.Repeat())
	}

	data := []byte(s)
	for i := 0; i < len(data); {
		l, e := ansi.ScanSequenceEffect(data[i:])
		switch {
		case l == 0:
			flush()
			_, l = utf8.DecodeRune(data[i:])
			b.Write(data[i : i+l])
		case e.Set&ansi.S_CONTROL != 0 || e.Covered() == ansi.S_NONE:
			flush()
			b.Write(data[i : i+l])
		default:
			scanned.scan(string(data[i : i+l]))
		}
		i += l
	}
	flush()
	return b.String()
}

// ScanFormat provides the codes of the groups established by the output
// format sequences of a text at its end (see SplitText for the
// interpretation of resets). Groups explicitly switched off get their
// reset code, untouched groups are omitted.
func ScanFormat(s string) map[*Group]ansi.CodeSequence {
	var st state
	st.scan(s)
	r := map[*Group]ansi.CodeSequence{}
	for _, g := range groups {
		if c := st.Get(g); c != nil {
			r[g] = c
		}
	}
	return r
}
