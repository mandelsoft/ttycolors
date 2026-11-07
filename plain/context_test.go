package plain_test

import (
	"reflect"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/mandelsoft/ttycolors"
	"github.com/mandelsoft/ttycolors/ansi"
	"github.com/mandelsoft/ttycolors/plain"
)

var _ = Describe("Context", func() {
	var old bool
	BeforeEach(func() {
		old = ttycolors.NoColors
		ttycolors.NoColors = true
	})
	AfterEach(func() {
		ttycolors.NoColors = old
	})

	It("uses the given context", func() {
		ctx := ttycolors.NewContext(true)
		p := plain.For(ctx)
		Expect(p.TTYContext()).To(BeIdenticalTo(ctx))
		Expect(p.Blue("x")).To(Equal(ansi.Mode(ansi.COLOR_FG_BLUE) + "x" + ansi.Mode(ansi.COLOR_FG_DEFAULT)))
		Expect(plain.Blue("x")).To(Equal("x"))
	})

	It("uses the default context for nil", func() {
		Expect(plain.For(nil).TTYContext()).To(BeIdenticalTo(ttycolors.DefaultContext()))
	})

	It("uses the settings of the context at the time of the call", func() {
		ctx := ttycolors.NewContext(false)
		p := plain.For(ctx)
		Expect(p.Bold("x")).To(Equal("x"))
		ctx.Enable()
		Expect(p.Bold("x")).To(Equal(ansi.Mode(ansi.MODE_RESET_BOLD, ansi.MODE_BOLD) + "x" + ansi.Mode(ansi.MODE_RESET_BOLD)))
	})

	It("uses the features of the context", func() {
		ctx := ttycolors.NewContext(true)
		ctx.EnableFeatures(ttycolors.FeatureUnderlineStyles)
		Expect(plain.For(ctx).CurlyUnderline("x")).To(Equal(ansi.Mode(ansi.UNDERLINE_CURLY) + "x" + ansi.Mode(ansi.MODE_RESET_UNDERLINE)))
	})

	It("provides the strings of the ttycolors context for all methods", func() {
		ctx := ttycolors.NewContext(true)
		ctx.EnableFeatures(ttycolors.AllFeatures)
		p := reflect.ValueOf(plain.For(ctx))
		c := reflect.ValueOf(ctx)
		n := 0
		for i := 0; i < p.NumMethod(); i++ {
			m := p.Type().Method(i)
			t := m.Type
			if t.NumOut() != 1 || t.Out(0).Kind() != reflect.String || !t.IsVariadic() {
				continue
			}
			var args []reflect.Value
			for j := 1; j < t.NumIn()-1; j++ {
				args = append(args, reflect.ValueOf(byte(j)))
			}
			args = append(args, reflect.ValueOf("a"), reflect.ValueOf(1))
			expected := c.MethodByName(m.Name).Call(args)[0].Interface().(ttycolors.String).String()
			Expect(p.Method(i).Call(args)[0].String()).To(Equal(expected), m.Name)
			n++
		}
		Expect(n).To(BeNumerically(">", 60))
	})
})
