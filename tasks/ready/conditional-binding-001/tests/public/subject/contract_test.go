package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"task.local/conditional-binding-001/subject"
	"testing"
)

func TestConditionalValues(t *testing.T) {
	for _, c := range []struct {
		initial, next int
		enabled       bool
		want          int
	}{
		{0, 13, true, 13}, {0, 13, false, 0},
		{7, -4, true, -4}, {7, -4, false, 7},
		{-8, 0, true, 0}, {-8, 0, false, -8}, {5, 5, true, 5},
	} {
		if got := subject.SelectValue(c.initial, c.next, c.enabled); got != c.want {
			t.Errorf("SelectValue(%d, %d, %t) = %d, want %d; check which block owns result", c.initial, c.next, c.enabled, got, c.want)
		}
	}
}

func TestRetainedConditionalBlock(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "implementation.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	found := false
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "SelectValue" || fn.Body == nil {
			continue
		}
		for _, stmt := range fn.Body.List {
			branch, ok := stmt.(*ast.IfStmt)
			if !ok {
				continue
			}
			name, ok := branch.Cond.(*ast.Ident)
			if ok && name.Name == "useNext" && len(branch.Body.List) > 0 {
				found = true
			}
		}
	}
	if !found {
		t.Error("SelectValue conditional block present = false, want true: retain if useNext { ... }")
	}
}
