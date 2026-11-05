package renderer

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

// set provides a state with the given codes.
func set(codes ...any) state {
	var s state
	for i := 0; i < len(codes); i += 2 {
		var c ansi.CodeSequence
		if codes[i+1] != nil {
			c = codes[i+1].(ansi.CodeSequence)
		}
		s = s.Set(codes[i].(*Group), c)
	}
	return s
}

var _ = Describe("State", func() {
	Context("groups", func() {
		It("are ordered", func() {
			Expect(groups).To(Equal([]*Group{FG, BG, Highlight, Italic, Underline, Blinking, Reverse, Invisible, StrikeThrough, Overline, UnderlineColor}))
		})

		It("are provided as copy", func() {
			g := Groups()
			Expect(g).To(Equal(groups))
			g[0] = nil
			Expect(groups[0]).To(BeIdenticalTo(FG))
		})

		It("provide reset codes", func() {
			Expect(FG.Reset()).To(Equal(ansi.COLOR_FG_DEFAULT))
			Expect(Highlight.Reset()).To(Equal(ansi.MODE_RESET_BOLD))
			Expect(Italic.Reset()).To(Equal(ansi.MODE_RESET_ITALIC))
			Expect(StrikeThrough.Reset()).To(Equal(ansi.MODE_RESET_STRIKE))
		})

		It("cover all distinct areas", func() {
			var all ansi.Covered
			for _, g := range groups {
				Expect(g.covered).NotTo(Equal(ansi.S_NONE))
				Expect(all & g.covered).To(Equal(ansi.S_NONE))
				all |= g.covered
			}
			Expect(all).To(Equal(ansi.S_ALL))
		})

		It("map to all distinct fields of the state", func() {
			var s state
			v := reflect.ValueOf(&s).Elem()
			fields := map[uintptr]string{}
			for i := 0; i < v.NumField(); i++ {
				fields[v.Field(i).Addr().Pointer()] = v.Type().Field(i).Name
			}
			used := map[uintptr]bool{}
			for _, g := range groups {
				addr := reflect.ValueOf(g.field(&s)).Pointer()
				Expect(fields).To(HaveKey(addr))
				Expect(used).NotTo(HaveKey(addr), "field %s used twice", fields[addr])
				used[addr] = true
			}
			Expect(used).To(HaveLen(len(fields)))
		})
	})

	Context("zero state", func() {
		It("provides defaults", func() {
			var s state
			Expect(s.Get(FG)).To(BeNil())
			Expect(s.Repeat()).To(Equal(""))
			Expect(s.EmitEnds()).To(Equal(""))
		})
	})

	Context("Set", func() {
		It("does not modify the original state", func() {
			s := set(FG, ansi.COLOR_FG_RED)
			n := s.Set(FG, ansi.COLOR_FG_BLUE)
			Expect(s.Get(FG)).To(Equal(ansi.COLOR_FG_RED))
			Expect(n.Get(FG)).To(Equal(ansi.COLOR_FG_BLUE))
		})

		It("copies the terminal state", func() {
			s := set(FG, ansi.COLOR_FG_RED)
			n := s.Set(BG, ansi.COLOR_BG_RED)
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_RED)))
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		})

		It("keeps the terminal state of all groups", func() {
			s := set(BG, ansi.COLOR_BG_RED)
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_BG_RED)))
			n := s.Set(Highlight, ansi.MODE_BOLD)
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.MODE_BOLD)))
			Expect(n.EmitEnds()).To(Equal(ansi.Mode(ansi.COLOR_BG_DEFAULT, ansi.MODE_RESET_BOLD)))
		})
	})

	Context("Repeat", func() {
		It("renders only changes", func() {
			s := set(FG, ansi.COLOR_FG_RED)
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
			Expect(s.Repeat()).To(Equal(""))
			n := s.Set(FG, ansi.COLOR_FG_RED)
			Expect(n.Repeat()).To(Equal(""))
			n = s.Set(FG, ansi.NewCodeSequence(ansi.COLOR_FG_RED))
			Expect(n.Repeat()).To(Equal(""))
			n = s.Set(FG, ansi.COLOR_FG_BLUE)
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE)))
		})

		It("renders in group order", func() {
			s := set(Highlight, ansi.MODE_BOLD, FG, ansi.COLOR_FG_RED)
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD)))
		})

		It("renders sub parameter codes separately", func() {
			s := set(Underline, ansi.UNDERLINE_CURLY, FG, ansi.COLOR_FG_RED, Highlight, ansi.MODE_BOLD)
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD) + ansi.Mode(ansi.UNDERLINE_CURLY)))
		})

		It("renders a reset for a removed code", func() {
			s := set(BG, ansi.COLOR_BG_RED)
			s.Repeat()
			n := s.Set(BG, nil)
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_BG_DEFAULT)))
			Expect(n.EmitEnds()).To(Equal(""))
		})

		It("replaces a reset by a new code", func() {
			s := set(BG, ansi.COLOR_BG_RED)
			s.Repeat()
			n := s.Set(BG, nil).Set(BG, ansi.COLOR_BG_BLUE)
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_BG_BLUE)))
		})

		It("renders a code equivalent to the default, if the terminal state differs", func() {
			s := set(Italic, ansi.MODE_RESET_ITALIC)
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.MODE_RESET_ITALIC)))
			n := s.Set(Italic, nil)
			Expect(n.Repeat()).To(Equal(""))
		})
	})

	Context("EmitEnds", func() {
		It("resets groups not in default state", func() {
			s := set(FG, ansi.COLOR_FG_RED, Highlight, ansi.MODE_BOLD)
			s.Repeat()
			Expect(s.EmitEnds()).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT, ansi.MODE_RESET_BOLD)))
			Expect(s.EmitEnds()).To(Equal(""))
		})

		It("ignores groups not yet rendered", func() {
			s := set(FG, ansi.COLOR_FG_RED)
			Expect(s.EmitEnds()).To(Equal(""))
		})

		It("ignores groups with code equivalent to the default", func() {
			s := set(Italic, ansi.MODE_RESET_ITALIC, Highlight, ansi.MODE_RESET_BOLD)
			s.Repeat()
			Expect(s.EmitEnds()).To(Equal(""))
		})
	})

	Context("Embedded", func() {
		It("keeps the state for plain text", func() {
			Expect(state{}.Embedded("text")).To(Equal(state{}))
			s := set(FG, ansi.COLOR_FG_RED)
			s.Repeat()
			n := s.Embedded("text")
			Expect(n).To(Equal(s))
			Expect(n.Repeat()).To(Equal(""))
		})

		It("does not modify the original state", func() {
			var s state
			n := s.Embedded("\x1b[1mbold")
			Expect(s).To(Equal(state{}))
			Expect(n.Repeat()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD)))
		})

		It("reestablishes the desired state for set groups", func() {
			s := set(FG, ansi.COLOR_FG_BLUE)
			s.Repeat()
			s = s.Embedded("\x1b[31mred")
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE)))
		})

		It("knows reset groups", func() {
			s := set(FG, ansi.COLOR_FG_BLUE, Highlight, ansi.MODE_BOLD)
			s.Repeat()
			s = s.Embedded("\x1b[0m")
			Expect(s.Repeat()).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE, ansi.MODE_BOLD)))
			s = s.Set(FG, nil).Set(Highlight, nil).Embedded("\x1b[0m")
			Expect(s.Repeat()).To(Equal(""))
			Expect(s.EmitEnds()).To(Equal(""))
		})

		It("resets groups left set at the end", func() {
			s := state{}.Embedded("\x1b[1;4:3mx")
			Expect(s.EmitEnds()).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_RESET_UNDERLINE)))
		})
	})

	Context("helpers", func() {
		It("compares codes", func() {
			Expect(sameCode(nil, nil)).To(BeTrue())
			Expect(sameCode(nil, ansi.MODE_BOLD)).To(BeFalse())
			Expect(sameCode(ansi.MODE_BOLD, ansi.NewCodeSequence(ansi.MODE_BOLD))).To(BeTrue())
			Expect(sameCode(ansi.MODE_UNDERLINE, ansi.UNDERLINE_SINGLE)).To(BeFalse())
		})

		It("detects default codes", func() {
			Expect(Italic.isDefault(nil)).To(BeTrue())
			Expect(Italic.isDefault(ansi.MODE_RESET_ITALIC)).To(BeTrue())
			Expect(Italic.isDefault(ansi.MODE_ITALIC)).To(BeFalse())
		})
	})
})
