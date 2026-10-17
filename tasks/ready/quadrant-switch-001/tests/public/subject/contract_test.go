package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"task.local/quadrant-switch-001/subject"
	"testing"
)

func TestNineRegions(t *testing.T) {
	for _, c := range []struct {
		x, y int
		want string
	}{
		{0, 0, "origin"}, {3, 0, "positive-x"}, {-3, 0, "negative-x"},
		{0, 4, "positive-y"}, {0, -4, "negative-y"},
		{2, 5, "quadrant-i"}, {-2, 5, "quadrant-ii"}, {-2, -5, "quadrant-iii"}, {2, -5, "quadrant-iv"},
	} {
		if got := subject.Locate(c.x, c.y); got != c.want {
			t.Errorf("Locate(%d, %d) = %q, want %q; origin/axes take priority", c.x, c.y, got, c.want)
		}
	}
}

func TestExpressionlessSwitch(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "implementation.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	count := 0
	ast.Inspect(file, func(node ast.Node) bool {
		switch n := node.(type) {
		case *ast.SwitchStmt:
			count++
			if n.Tag != nil {
				t.Error("Locate switch tag present = true, want false: use an expressionless switch")
			}
		case *ast.BranchStmt:
			if n.Tok == token.FALLTHROUGH {
				t.Error("Locate fallthrough present = true, want false")
			}
		}
		return true
	})
	if count == 0 {
		t.Error("Locate expressionless switches = 0, want at least 1")
	}
}
