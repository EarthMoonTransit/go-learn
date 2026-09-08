package subject_test

import (
	"go/ast"
	"go/constant"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"strings"
	"testing"
)

func TestDeclaredConstantKinds(t *testing.T) {
	set := token.NewFileSet()

	paths, e := filepath.Glob("*.go")
	if e != nil {
		t.Fatal(e)
	}
	files := []*ast.File{}
	for _, name := range paths {
		if strings.HasSuffix(name, "_test.go") {
			continue
		}
		f, e := parser.ParseFile(set, name, nil, 0)
		if e != nil {
			t.Fatal(e)
		}
		files = append(files, f)
	}
	config := types.Config{Importer: importer.ForCompiler(set, "source", nil)}
	pkg, e := config.Check("subject", set, files, nil)
	if e != nil {
		t.Fatal(e)
	}
	for _, c := range []struct {
		name string
		kind types.BasicKind
		want constant.Value
	}{{"server", types.UntypedString, constant.MakeString("localhost")}, {"port", types.Int, constant.MakeInt64(8080)}} {
		obj, ok := pkg.Scope().Lookup(c.name).(*types.Const)
		if !ok {
			t.Errorf("%s must be a constant", c.name)
			continue
		}
		if !constant.Compare(obj.Val(), token.EQL, c.want) {
			t.Errorf("%s constant value got %s want %s", c.name, obj.Val().ExactString(), c.want.ExactString())
		}
		typ, ok := types.Unalias(obj.Type()).(*types.Basic)
		if !ok || typ.Kind() != c.kind {
			t.Errorf("%s constant type got %s want %s", c.name, obj.Type(), types.Typ[c.kind])
		}
	}
}
