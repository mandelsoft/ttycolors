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
	Context("control sequences", func() {
		erase := "\x1b[2K"
		title := "\x1b]0;title\a"

		It("are ignored for runes", func() {
			Expect(ansi.CharLen(erase + "ab" + title)).To(Equal(2))
			Expect(ansi.EscapeLength([]byte(erase))).To(Equal(4))
			Expect(ansi.SubString(erase+"abc"+title+"d", 1, 4)).To(Equal("bc" + title + "d"))
			Expect(ansi.ByteIndex([]byte(erase+"ab"), 'b')).To(Equal(5))
		})

		It("are kept when stripping output formats", func() {
			Expect(ansi.Strip(erase + ansi.Mode(ansi.MODE_BOLD) + "ab" + title)).To(Equal(erase + "ab" + title))
		})

		It("are removed when stripping all", func() {
			Expect(ansi.StripAll(erase + ansi.Mode(ansi.MODE_BOLD) + "äb" + title + ansi.Mode())).To(Equal("äb"))
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
		Expect(ansi.EscapeLength([]byte("\x1b[2Kx"))).To(Equal(4))
		Expect(ansi.EscapeLength([]byte("\x1b[2"))).To(Equal(0))
	})
})
