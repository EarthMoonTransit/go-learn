package subject_test

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"os"
	"strings"
	"testing"
)

func TestExactTypeSet(t *testing.T) {
	set := token.NewFileSet()
	pkgs, e := parser.ParseDir(set, ".", func(info os.FileInfo) bool { return !strings.HasSuffix(info.Name(), "_test.go") }, 0)
	if e != nil {
		t.Fatal(e)
	}
	files := []*ast.File{}
	source, ok := pkgs["subject"]
	if !ok {
		t.Fatal("missing subject package")
	}
	for _, f := range source.Files {
		files = append(files, f)
	}
	config := types.Config{Importer: importer.ForCompiler(set, "source", nil)}
	pkg, e := config.Check("subject", set, files, nil)
	if e != nil {
		t.Fatal(e)
	}
	obj := pkg.Scope().Lookup("numeric")
	if obj == nil {
		t.Fatal("missing numeric constraint")
	}
	iface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok {
		t.Fatal("numeric must be an interface constraint")
	}
	expected := types.NewInterfaceType(nil, []types.Type{types.NewUnion([]*types.Term{types.NewTerm(true, types.Typ[types.Int]), types.NewTerm(true, types.Typ[types.Int64]), types.NewTerm(true, types.Typ[types.Float64])})}).Complete()
	if !types.Identical(iface, expected) {
		t.Errorf("numeric type set got %s, want ~int | ~int64 | ~float64", iface)
	}
}
