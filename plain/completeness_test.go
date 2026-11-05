package plain_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"io/fs"
	"strings"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

// stringFunctions provides the parameter types of all exported functions
// of a package directory returning the given result type.
func stringFunctions(dir string, result string) map[string]string {
	fset := token.NewFileSet()
	pkgs, err := parser.ParseDir(fset, dir, func(fi fs.FileInfo) bool {
		return !strings.HasSuffix(fi.Name(), "_test.go")
	}, 0)
	Expect(err).To(Succeed())
	funcs := map[string]string{}
	for _, p := range pkgs {
		for _, f := range p.Files {
			for _, d := range f.Decls {
				fd, ok := d.(*ast.FuncDecl)
				if !ok || fd.Recv != nil || !fd.Name.IsExported() || fd.Type.Results == nil || len(fd.Type.Results.List) != 1 {
					continue
				}
				if id, ok := fd.Type.Results.List[0].Type.(*ast.Ident); !ok || id.Name != result {
					continue
				}
				var params []string
				for _, p := range fd.Type.Params.List {
					for range p.Names {
						params = append(params, types(p.Type))
					}
				}
				funcs[fd.Name.Name] = strings.Join(params, ",")
			}
		}
	}
	return funcs
}

func types(e ast.Expr) string {
	switch t := e.(type) {
	case *ast.Ident:
		return t.Name
	case *ast.Ellipsis:
		return "..." + types(t.Elt)
	case *ast.InterfaceType:
		return "interface{}"
	default:
		return "?"
	}
}

var _ = Describe("Completeness", func() {
	It("provides a function for every string function of package ttycolors", func() {
		expected := stringFunctions("..", "String")
		Expect(expected).To(HaveKey("Blue"))
		Expect(stringFunctions(".", "string")).To(Equal(expected))
	})
})
