package ansi_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors/ansi"
)

var _ = Describe("Codes", func() {
	It("separators", func() {
		Expect(string(ansi.PARAM_SEPARATOR)).To(Equal(";"))
		Expect(string(ansi.SUBPARAM_SEPARATOR)).To(Equal(":"))
	})

	Context("Param", func() {
		It("encodes", func() {
			Expect(string(ansi.MODE_RESET_ALL.Bytes())).To(Equal("0"))
			Expect(string(ansi.COLOR_BG_BRIGHT_WHITE.Bytes())).To(Equal("107"))
			Expect(string(ansi.Param(300).Bytes())).To(Equal("300"))
			Expect(ansi.Omitted.Bytes()).To(BeEmpty())
		})

		It("appends", func() {
			Expect(string(ansi.MODE_BOLD.AppendTo(ansi.Sequence("x")))).To(Equal("x1"))
		})

		It("provides codes", func() {
			Expect(ansi.MODE_BOLD.Codes()).To(Equal([]ansi.Code{{1}}))
		})
	})

	Context("Code", func() {
		It("encodes", func() {
			Expect(ansi.Code{}.Bytes()).To(BeEmpty())
			Expect(string(ansi.Code{1}.Bytes())).To(Equal("1"))
			Expect(string(ansi.UNDERLINE_CURLY.Bytes())).To(Equal("4:3"))
			Expect(string(ansi.Code{58, 2, ansi.Omitted, 1, 2, 3}.Bytes())).To(Equal("58:2::1:2:3"))
		})

		It("appends", func() {
			Expect(string(ansi.UNDERLINE_CURLY.AppendTo(ansi.Sequence("x")))).To(Equal("x4:3"))
		})

		It("provides codes", func() {
			Expect(ansi.UNDERLINE_CURLY.Codes()).To(Equal([]ansi.Code{{4, 3}}))
		})

		It("detects sub parameters", func() {
			Expect(ansi.Code{1}.HasSubParams()).To(BeFalse())
			Expect(ansi.UNDERLINE_CURLY.HasSubParams()).To(BeTrue())
		})
	})

	Context("CodeSequence", func() {
		It("is created from params", func() {
			Expect(ansi.NewCodeSequence(38, 5).Codes()).To(Equal([]ansi.Code{{38}, {5}}))
			Expect(ansi.NewCodeSequence().Codes()).To(BeEmpty())
		})

		It("is joined from sequences", func() {
			s := ansi.Join(ansi.MODE_BOLD, nil, ansi.UNDERLINE_CURLY, ansi.Code_FG_ID(4))
			Expect(s.Codes()).To(Equal([]ansi.Code{{1}, {4, 3}, {38}, {5}, {4}}))
			Expect(ansi.Join().Codes()).To(BeEmpty())
		})

		It("encodes", func() {
			Expect(ansi.Join().Bytes()).To(BeEmpty())
			Expect(string(ansi.Code_FG_RGB(1, 22, 255).Bytes())).To(Equal("38;2;1;22;255"))
			Expect(string(ansi.Join(ansi.Param(1), ansi.UNDERLINE_CURLY, ansi.Omitted, ansi.Param(2)).Bytes())).To(Equal("1;4:3;;2"))
		})

		It("appends", func() {
			Expect(string(ansi.COLOR_FG_ID.AppendTo(ansi.Sequence("x")))).To(Equal("x38;5"))
		})

		It("does not share codes", func() {
			s := ansi.NewCodeSequence(1, 2)
			c := s.Codes()
			c[0] = ansi.Code{9}
			Expect(s.Codes()).To(Equal([]ansi.Code{{1}, {2}}))
		})
	})

	Context("Mode", func() {
		It("renders elements", func() {
			Expect(ansi.Mode()).To(Equal("\x1b[m"))
			Expect(ansi.Mode(ansi.MODE_BOLD, ansi.Code_FG_ID(5))).To(Equal("\x1b[1;38;5;5m"))
			Expect(ansi.Mode(ansi.MODE_BOLD, ansi.UNDERLINE_CURLY)).To(Equal("\x1b[1;4:3m"))
			Expect(ansi.Mode(ansi.Join(ansi.Omitted, ansi.Param(1)))).To(Equal("\x1b[;1m"))
		})

		It("ignores nil elements and empty codes", func() {
			Expect(ansi.Mode(nil)).To(Equal(""))
			Expect(ansi.Mode(ansi.Join())).To(Equal(""))
			Expect(ansi.Mode(ansi.NewCodeSequence(), nil, ansi.Code{})).To(Equal(""))
			Expect(ansi.Mode(nil, ansi.MODE_BOLD, ansi.Join(), ansi.MODE_ITALIC)).To(Equal("\x1b[1;3m"))
			Expect(ansi.Mode(ansi.Code{}, ansi.MODE_BOLD)).To(Equal("\x1b[1m"))
		})

		It("appends to a sequence", func() {
			Expect(string(ansi.AppendMode(ansi.Sequence("x"), ansi.MODE_BOLD))).To(Equal("x\x1b[1m"))
			Expect(string(ansi.AppendMode(ansi.Sequence("x"), nil))).To(Equal("x"))
		})
	})

	Context("Modes", func() {
		It("renders elements", func() {
			Expect(ansi.Modes()).To(Equal(""))
			Expect(ansi.Modes(nil, ansi.Join())).To(Equal(""))
			Expect(ansi.Modes(ansi.MODE_BOLD, ansi.MODE_ITALIC)).To(Equal("\x1b[1;3m"))
			Expect(ansi.Modes(ansi.UNDERLINE_CURLY)).To(Equal("\x1b[4:3m"))
		})

		It("separates codes with sub parameters", func() {
			Expect(ansi.Modes(ansi.MODE_BOLD, ansi.UNDERLINE_DOTTED, ansi.Code_UL_ID(3), ansi.MODE_ITALIC)).
				To(Equal("\x1b[1;3m\x1b[4:4m\x1b[58:5:3m"))
			Expect(ansi.Modes(ansi.Join(ansi.Param(1), ansi.UNDERLINE_CURLY, ansi.Param(2)))).To(Equal("\x1b[1;2m\x1b[4:3m"))
		})

		It("appends to a sequence", func() {
			Expect(string(ansi.AppendModes(ansi.Sequence("x"), ansi.MODE_BOLD, ansi.UNDERLINE_CURLY))).To(Equal("x\x1b[1m\x1b[4:3m"))
			Expect(ansi.AppendModes(nil)).To(BeNil())
		})
	})
})
