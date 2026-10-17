package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"strings"
	"task.local/named-result-adjustment-001/subject"
	"testing"
)

func TestAdjustedResultAndCaller(t *testing.T) {
	for _, c := range []struct{ base, delta, initial int }{{4, 7, 99}, {0, -3, 11}, {-8, 0, 17}, {-1000, 1000, -5}, {1000, -1000, 6}} {
		value, calls := c.initial, 0
		got := subject.Adjust(c.base, c.delta, &value, func() { calls++ })
		if want := c.base + 3 + c.delta; got != want {
			t.Errorf("Adjust(%d,%d) = %d, want %d after deferred result adjustment", c.base, c.delta, got, want)
		}
		if value != c.initial {
			t.Errorf("caller value = %d, want unchanged %d", value, c.initial)
		}
		if calls != 1 {
			t.Errorf("work calls = %d, want 1", calls)
		}
	}
}

func capturePanic(call func()) (value any) {
	defer func() { value = recover() }()
	call()
	return nil
}

func TestPendingPanicIsPreserved(t *testing.T) {
	token := &struct{ code int }{23}
	caller, calls := 41, 0
	got := capturePanic(func() { subject.Adjust(4, 7, &caller, func() { calls++; panic(token) }) })
	if got != token {
		t.Errorf("panic identity = %v, want original %v", got, token)
	}
	if caller != 41 {
		t.Errorf("caller after panic = %d, want 41", caller)
	}
	if calls != 1 {
		t.Errorf("panicking work calls = %d, want 1", calls)
	}
}

func TestNamedDeferredResultForm(t *testing.T) {
	paths, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	named, defers := false, 0
	for _, path := range paths {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, decl := range file.Decls {
			fn, ok := decl.(*ast.FuncDecl)
			if !ok || fn.Name.Name != "Adjust" {
				continue
			}
			if fn.Type.Results != nil {
				for _, field := range fn.Type.Results.List {
					if len(field.Names) > 0 {
						named = true
					}
				}
			}
			ast.Inspect(fn.Body, func(node ast.Node) bool {
				if _, ok := node.(*ast.DeferStmt); ok {
					defers++
				}
				return true
			})
		}
	}
	if !named {
		t.Error("Adjust named result present = false, want true")
	}
	if defers == 0 {
		t.Error("Adjust deferred calls = 0, want at least 1")
	}
}
