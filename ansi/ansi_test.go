package ansi_test

import (
	"github.com/mandelsoft/ttycolors/ansi"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Ansi Test Environment", func() {
	s := ansi.Mode(ansi.Code_BG_RGB('\n', '\n', '\n')) + "1234" + ansi.Mode(ansi.MODE_RESET_ALL) + "end"

	It("Len", func() {
		Expect(ansi.CharLen(s)).To(Equal(7))
	})

	It("ByteIndex", func() {
		Expect(ansi.ByteIndex([]byte(s), '1')).To(Equal(16))
	})

	It("CharIndex", func() {
		Expect(ansi.CharIndex(s, 0)).To(Equal(16))
		Expect(ansi.CharIndex(s, 1)).To(Equal(17))
		Expect(ansi.CharIndex(s, 4)).To(Equal(24))
	})
	It("CharIndexBefore", func() {
		Expect(ansi.CharIndexBefore(s, 0)).To(Equal(0))
		Expect(ansi.CharIndexBefore(s, 1)).To(Equal(17))
		Expect(ansi.CharIndexBefore(s, 4)).To(Equal(20))
	})
	It("SubString", func() {
		Expect(ansi.SubString(s, 1, 4)).To(Equal("234"))
	})
	It("SubString empty", func() {
		for i := 0; i <= ansi.CharLen(s); i++ {
			Expect(ansi.SubString(s, i, i)).To(Equal(""))
		}
	})
	It("panics for negative indices", func() {
		Expect(func() { ansi.CharIndex(s, -1) }).To(PanicWith("ansi: negative index -1"))
		Expect(func() { ansi.CharIndexBefore(s, -2) }).To(PanicWith("ansi: negative index -2"))
		Expect(func() { ansi.SplitAt(s, -1) }).To(Panic())
		Expect(func() { ansi.SubString(s, -1, 2) }).To(Panic())
		Expect(func() { ansi.SubString(s, 1, -2) }).To(Panic())
		Expect(func() { ansi.SubString(s, -1, -1) }).To(Panic())
	})
	Context("limitations", func() {
		erase := "\x1b[2K"

		It("counts other control sequences as text", func() {
			Expect(ansi.CharLen(erase + "ab")).To(Equal(6))
			Expect(ansi.EscapeLength([]byte(erase))).To(Equal(0))
		})

		It("keeps other control sequences when stripping", func() {
			Expect(ansi.Strip(erase + ansi.Mode(ansi.MODE_BOLD) + "ab")).To(Equal(erase + "ab"))
		})
	})
	It("SubString invalid range", func() {
		Expect(func() { ansi.SubString(s, 3, 1) }).To(Panic())
	})
	It("Strip", func() {
		r := ansi.Strip(s)
		Expect(r).To(Equal("1234end"))
	})
	It("Strip multi byte runes", func() {
		r := ansi.Strip("ä" + ansi.Mode(ansi.MODE_BOLD) + "öü" + ansi.Mode() + "ß")
		Expect(r).To(Equal("äöüß"))
	})
	It("EscapeLength", func() {
		Expect(ansi.EscapeLength([]byte("\x1b[mx"))).To(Equal(3))
		Expect(ansi.EscapeLength([]byte("\x1b[38:5:1mx"))).To(Equal(9))
		Expect(ansi.EscapeLength([]byte("\x1b[2Kx"))).To(Equal(0))
	})
})
