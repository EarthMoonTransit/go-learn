package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"slices"
	"task.local/defer-trace-001/subject"
	"testing"
)

func TestRegistrationAndLIFO(t *testing.T) {
	for _, seed := range []int{0, 2, -11, 1000, -1000} {
		want := []int{seed + 9, subject.Checkpoint, seed + 4, seed}
		if got := subject.Trace(seed); !slices.Equal(got, want) {
			t.Errorf("Trace(%d) = %v, want %v; checkpoint precedes deferred LIFO calls with saved arguments", seed, got, want)
		}
	}
}

func TestDeferredCallsRemainInFunction(t *testing.T) {
	file, err := parser.ParseFile(token.NewFileSet(), "implementation.go", nil, 0)
	if err != nil {
		t.Fatal(err)
	}
	calls := 0
	for _, decl := range file.Decls {
		fn, ok := decl.(*ast.FuncDecl)
		if !ok || fn.Name.Name != "runTrace" {
			continue
		}
		ast.Inspect(fn.Body, func(node ast.Node) bool {
			if _, ok := node.(*ast.DeferStmt); ok {
				calls++
			}
			return true
		})
	}
	if calls < 2 {
		t.Errorf("runTrace deferred calls = %d, want at least 2 in the supplied function", calls)
	}
}
