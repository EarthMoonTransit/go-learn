package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestMonthConstBlock(t *testing.T) {
	f, e := parser.ParseFile(token.NewFileSet(), "months.go", nil, 0)
	if e != nil {
		t.Fatal(e)
	}
	for _, d := range f.Decls {
		g, ok := d.(*ast.GenDecl)
		if !ok || g.Tok != token.CONST {
			continue
		}
		names := map[string]bool{}
		hasIota := false
		for _, s := range g.Specs {
			v := s.(*ast.ValueSpec)
			for _, n := range v.Names {
				names[n.Name] = true
			}
		}
		ast.Inspect(g, func(n ast.Node) bool {
			if id, ok := n.(*ast.Ident); ok && id.Name == "iota" {
				hasIota = true
			}
			return true
		})
		all := true
		for _, n := range []string{"jan", "feb", "mar", "apr", "may", "jun", "jul", "aug", "sep", "oct", "nov", "dec"} {
			all = all && names[n]
		}
		if all && hasIota {
			return
		}
	}
	t.Fatal("month constants must share one const block using iota")
}
