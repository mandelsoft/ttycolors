package renderer

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("State", func() {
	It("orders groups by creation", func() {
		Expect(groups).To(Equal([]*group{fgGroup, bgGroup, highlightGroup}))
	})

	It("indexes groups by creation order", func() {
		Expect(groupIndex).To(HaveLen(len(groups)))
		for i, g := range groups {
			Expect(groupIndex[g]).To(Equal(i))
		}
	})

	Context("nil state", func() {
		var s state

		It("provides defaults", func() {
			Expect(s.Get(fgGroup)).To(BeNil())
			Expect(s.Repeat(nil)).To(Equal(""))
			Expect(s.EmitEnds()).To(Equal(""))
			s.RequestEnd(fgGroup)
		})

		It("creates a state for all groups", func() {
			n := s.Set(bgGroup, ansi.COLOR_BG_RED)
			Expect(n).To(HaveLen(len(groups)))
			Expect(n.Get(bgGroup)).To(Equal(ansi.COLOR_BG_RED))
			Expect(n.Get(fgGroup)).To(BeNil())
		})
	})

	Context("Set", func() {
		It("does not share group states with the original state", func() {
			s := state(nil).Set(fgGroup, ansi.COLOR_FG_RED)
			n := s.Set(bgGroup, ansi.COLOR_BG_RED)
			Expect(n.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.COLOR_BG_RED)))
			Expect(s.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		})

		It("does not modify the original state", func() {
			s := state(nil).Set(fgGroup, ansi.COLOR_FG_RED)
			n := s.Set(fgGroup, ansi.COLOR_FG_BLUE)
			Expect(s.Get(fgGroup)).To(Equal(ansi.COLOR_FG_RED))
			Expect(n.Get(fgGroup)).To(Equal(ansi.COLOR_FG_BLUE))
		})

		It("keeps pending resets of all groups", func() {
			s := state(nil).Set(fgGroup, nil)
			s.RequestEnd(bgGroup)
			n := s.Set(highlightGroup, ansi.MODE_BOLD)
			Expect(n.EmitEnds()).To(Equal(ansi.Mode(ansi.COLOR_BG_DEFAULT)))
		})

		It("requires rendering of the changed group", func() {
			s := state(nil).Set(fgGroup, ansi.COLOR_FG_RED)
			Expect(s.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
			Expect(s.Repeat(nil)).To(Equal(""))
			n := s.Set(fgGroup, ansi.COLOR_FG_RED)
			Expect(n.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		})
	})

	Context("Repeat", func() {
		It("renders in group order", func() {
			s := state(nil).Set(highlightGroup, ansi.MODE_BOLD).Set(fgGroup, ansi.COLOR_FG_RED)
			Expect(s.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED, ansi.MODE_BOLD)))
		})

		It("excludes a group", func() {
			s := state(nil).Set(highlightGroup, ansi.MODE_BOLD).Set(fgGroup, ansi.COLOR_FG_RED)
			Expect(s.Repeat(fgGroup)).To(Equal(ansi.Mode(ansi.MODE_BOLD)))
			Expect(s.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_FG_RED)))
		})

		It("renders a pending reset", func() {
			s := state(nil).Set(bgGroup, nil)
			s.RequestEnd(bgGroup)
			Expect(s.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_BG_DEFAULT)))
			Expect(s.EmitEnds()).To(Equal(""))
		})

		It("replaces a pending reset by a code", func() {
			s := state(nil).Set(bgGroup, nil)
			s.RequestEnd(bgGroup)
			n := s.Set(bgGroup, ansi.COLOR_BG_RED)
			Expect(n.Repeat(nil)).To(Equal(ansi.Mode(ansi.COLOR_BG_RED)))
			Expect(n.EmitEnds()).To(Equal(""))
		})
	})

	Context("RequestEnd", func() {
		It("ignores groups with active code", func() {
			s := state(nil).Set(fgGroup, ansi.COLOR_FG_RED)
			s.RequestEnd(fgGroup)
			Expect(s.EmitEnds()).To(Equal(""))
		})

		It("emits resets in group order once", func() {
			s := state(nil).Set(fgGroup, nil)
			s.RequestEnd(highlightGroup)
			s.RequestEnd(fgGroup)
			Expect(s.EmitEnds()).To(Equal(ansi.Mode(ansi.COLOR_FG_DEFAULT) + ansi.Mode(ansi.MODE_RESET_BOLD)))
			Expect(s.EmitEnds()).To(Equal(""))
		})
	})
})
