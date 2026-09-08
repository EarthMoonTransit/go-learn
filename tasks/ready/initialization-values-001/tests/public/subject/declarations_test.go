package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestDeclaredDependency(t *testing.T) {
	set := token.NewFileSet()
	f, e := parser.ParseFile(set, "values.go", nil, 0)
	if e != nil {
		t.Fatal(e)
	}
	var a, b token.Pos
	usesBase := false
	for _, d := range f.Decls {
		g, ok := d.(*ast.GenDecl)
		if !ok || g.Tok != token.VAR {
			continue
		}
		for _, s := range g.Specs {
			v := s.(*ast.ValueSpec)
			for _, n := range v.Names {
				if n.Name == "base" {
					b = n.Pos()
				}
				if n.Name == "doubled" {
					a = n.Pos()
					for _, x := range v.Values {
						ast.Inspect(x, func(node ast.Node) bool {
							if id, ok := node.(*ast.Ident); ok && id.Name == "base" {
								usesBase = true
							}
							return true
						})
					}
				}
			}
		}
	}
	if a == 0 || b == 0 || a >= b || !usesBase {
		t.Error("declare doubled before base with an explicit base dependency")
	}
}
