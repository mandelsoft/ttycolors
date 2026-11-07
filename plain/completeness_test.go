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
	return declarations(dir, result, "")
}

// stringMethods provides the parameter types of all exported methods
// of the given receiver type of a package directory returning the given
// result type.
func stringMethods(dir string, receiver string, result string) map[string]string {
	return declarations(dir, result, receiver)
}

// declarations provides the parameter types of all exported functions
// (no receiver type given) or methods of the given receiver type of a
// package directory returning the given result type.
func declarations(dir string, result string, receiver string) map[string]string {
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
				if !ok || receiverType(fd) != receiver || !fd.Name.IsExported() || fd.Type.Results == nil || len(fd.Type.Results.List) != 1 {
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

// receiverType provides the name of the receiver type of a
// method or the empty string for a function.
func receiverType(fd *ast.FuncDecl) string {
	if fd.Recv == nil {
		return ""
	}
	t := fd.Recv.List[0].Type
	if s, ok := t.(*ast.StarExpr); ok {
		t = s.X
	}
	if id, ok := t.(*ast.Ident); ok {
		return id.Name
	}
	return "?"
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
		// functions already providing Go strings are provided, also.
		for n, p := range stringFunctions("..", "string") {
			expected[n] = p
		}
		Expect(expected).To(HaveKey("SubString"))
		Expect(stringFunctions(".", "string")).To(Equal(expected))
	})

	It("provides a method for every string method of a context of package ttycolors", func() {
		expected := stringMethods("..", "_context", "String")
		Expect(expected).To(HaveKey("Blue"))
		// String and StringWith are covered by Sequence and New.
		delete(expected, "String")
		delete(expected, "StringWith")
		// methods already providing Go strings are provided, also.
		for n, p := range stringMethods("..", "_context", "string") {
			expected[n] = p
		}
		Expect(expected).To(HaveKey("StripIfDisabled"))
		Expect(stringMethods(".", "_context", "string")).To(Equal(expected))
	})
})
