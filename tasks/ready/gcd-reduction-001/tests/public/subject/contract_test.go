package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"task.local/gcd-reduction-001/subject"
	"testing"
)

func TestGCDValues(t *testing.T) {
	for _, c := range []struct{ a, b, want int }{
		{0, 0, 0}, {0, 9, 9}, {9, 0, 9}, {12, 12, 12},
		{7, 4, 1}, {4, 7, 1}, {18, 30, 6}, {30, 18, 6},
		{1, 999, 1}, {100, 25, 25},
	} {
		if got := subject.GCD(c.a, c.b); got != c.want {
			t.Errorf("GCD(%d, %d) = %d, want %d", c.a, c.b, got, c.want)
		}
	}
}

func TestConditionOnlyLoop(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "implementation.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	loops := 0
	remainders := 0
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok {
			continue
		}
		ast.Inspect(fn.Body, func(node ast.Node) bool {
			switch n := node.(type) {
			case *ast.ForStmt:
				loops++
				if n.Init != nil || n.Post != nil || n.Cond == nil {
					t.Error("GCD loop form = non-condition-only, want for condition { ... } without init/post clauses")
				}
				ast.Inspect(n.Body, func(bodyNode ast.Node) bool {
					if expression, ok := bodyNode.(*ast.BinaryExpr); ok && expression.Op == token.REM {
						remainders++
					}
					if assignment, ok := bodyNode.(*ast.AssignStmt); ok && assignment.Tok == token.REM_ASSIGN {
						remainders++
					}
					return true
				})
			case *ast.RangeStmt:
				t.Error("GCD loop form = range, want condition-only for")
			case *ast.CallExpr:
				if name, ok := n.Fun.(*ast.Ident); ok && name.Name == fn.Name.Name {
					t.Error("GCD recursion present = true, want false")
				}
			}
			return true
		})
	}
	if loops == 0 {
		t.Error("GCD condition-only for loops = 0, want at least 1")
	}
	if remainders == 0 {
		t.Error("GCD remainder reduction in loop = false, want true: use % or %= in the loop body")
	}
}
