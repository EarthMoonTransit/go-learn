package subject_test

import (
	"go/ast"
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
	config := types.Config{}
	pkg, e := config.Check("subject", set, files, nil)
	if e != nil {
		t.Fatal(e)
	}
	obj := pkg.Scope().Lookup("numeric")
	if obj == nil {
		t.Fatal("missing numeric constraint")
	}
	iface, ok := obj.Type().Underlying().(*types.Interface)
	if !ok || iface.NumEmbeddeds() != 1 || iface.NumMethods() != 0 {
		t.Fatal("numeric must be a three-term union")
	}
	union, ok := iface.EmbeddedType(0).(*types.Union)
	if !ok || union.Len() != 3 {
		t.Fatal("numeric must be a three-term union")
	}
	want := map[string]bool{"int": true, "int64": true, "float64": true}
	for i := 0; i < union.Len(); i++ {
		term := union.Term(i)
		name := types.TypeString(types.Unalias(term.Type()), nil)
		if !term.Tilde() || !want[name] {
			t.Errorf("unexpected numeric term %s, require ~int | ~int64 | ~float64", term)
		}
		delete(want, name)
	}
	if len(want) != 0 {
		t.Errorf("missing numeric terms %v", want)
	}
}
