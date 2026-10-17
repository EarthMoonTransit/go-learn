package subject_test

import (
	"go/ast"
	"go/parser"
	"go/token"
	"path/filepath"
	"slices"
	"strings"
	"task.local/merge-two-streams-001/subject"
	"testing"
	"testing/synctest"
)

func TestSelectAndDisabling(t *testing.T) {
	files, err := filepath.Glob("*.go")
	if err != nil {
		t.Fatal(err)
	}
	functions := map[string][]ast.Node{}
	for _, path := range files {
		if strings.HasSuffix(path, "_test.go") {
			continue
		}
		file, err := parser.ParseFile(token.NewFileSet(), path, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		for _, decl := range file.Decls {
			if fn, ok := decl.(*ast.FuncDecl); ok {
				functions[fn.Name.Name] = append(functions[fn.Name.Name], fn.Body)
			}
			if group, ok := decl.(*ast.GenDecl); ok {
				for _, spec := range group.Specs {
					if value, ok := spec.(*ast.ValueSpec); ok {
						for i, expr := range value.Values {
							if literal, ok := expr.(*ast.FuncLit); ok && i < len(value.Names) {
								name := value.Names[i].Name
								functions[name] = append(functions[name], literal.Body)
							}
						}
					}
				}
			}
		}
	}
	var nilExpression func(ast.Expr) bool
	nilExpression = func(expr ast.Expr) bool {
		switch value := expr.(type) {
		case *ast.Ident:
			return value.Name == "nil"
		case *ast.ParenExpr:
			return nilExpression(value.X)
		case *ast.CallExpr:
			return len(value.Args) == 1 && nilExpression(value.Args[0])
		}
		return false
	}
	// Overapproximate referenced local functions/methods and function values.
	// Behavior checks remain the authority for progress/order/closure; this is
	// a bounded construction guard, not a general dataflow proof.
	seen := map[string]bool{}
	pending := []string{"MergeTwo"}
	selects, nilAssignments := 0, 0
	for len(pending) > 0 {
		name := pending[0]
		pending = pending[1:]
		if seen[name] {
			continue
		}
		seen[name] = true
		for _, fn := range functions[name] {
			ast.Inspect(fn, func(node ast.Node) bool {
				if ident, ok := node.(*ast.Ident); ok && len(functions[ident.Name]) > 0 && !seen[ident.Name] {
					pending = append(pending, ident.Name)
				}
				if _, ok := node.(*ast.SelectStmt); ok {
					selects++
				}
				if assign, ok := node.(*ast.AssignStmt); ok {
					for _, rhs := range assign.Rhs {
						if nilExpression(rhs) {
							nilAssignments++
						}
					}
				}
				return true
			})
		}
	}
	if selects == 0 {
		t.Error("MergeTwo select statements = 0, want at least 1")
	}
	if nilAssignments == 0 {
		t.Error("MergeTwo nil disabling assignments = 0, want closed-input removal with nil")
	}
}

func drainMerge(t *testing.T, out <-chan int, limit int) ([]int, bool) {
	t.Helper()
	var values []int
	if out == nil {
		t.Error("MergeTwo output = nil, want non-nil channel")
		return nil, false
	}
	for step := 0; step < limit; step++ {
		synctest.Wait()
		select {
		case value, open := <-out:
			if !open {
				return values, true
			}
			values = append(values, value)
		default:
			return values, false
		}
	}
	t.Errorf("merge produced more than %d bounded observations, want finite input occurrences", limit)
	return values, false
}
func expectMergeValue(t *testing.T, out <-chan int, want int) {
	t.Helper()
	synctest.Wait()
	select {
	case got, open := <-out:
		if !open {
			t.Errorf("available input %d output = closed, want value", want)
		} else if got != want {
			t.Errorf("available input output = %d, want %d", got, want)
		}
	default:
		t.Errorf("available input %d output = absent, want progress from either open input", want)
	}
}
func TestInterleavedInputs(t *testing.T) {
	for _, swapped := range []bool{false, true} {
		name := "left-closes-first"
		if swapped {
			name = "right-closes-first"
		}
		t.Run(name, func(t *testing.T) {
			synctest.Test(t, func(t *testing.T) {
				left, right := make(chan int, 4), make(chan int, 4)
				var out <-chan int
				if swapped {
					out = subject.MergeTwo(right, left)
				} else {
					out = subject.MergeTwo(left, right)
				}
				leftClosed, rightClosed := false, false
				defer func() {
					if !leftClosed {
						close(left)
					}
					if !rightClosed {
						close(right)
					}
					drainMerge(t, out, 32)
					synctest.Wait()
				}()
				left <- 11
				expectMergeValue(t, out, 11)
				right <- 21
				expectMergeValue(t, out, 21)
				left <- 12
				expectMergeValue(t, out, 12)
				right <- 22
				expectMergeValue(t, out, 22)
				close(left)
				leftClosed = true
				synctest.Wait()
				select {
				case got, open := <-out:
					t.Errorf("one input closed output = (%d,%t), want pending until remaining input", got, open)
				default:
				}
				right <- 23
				expectMergeValue(t, out, 23)
				close(right)
				rightClosed = true
				rest, closed := drainMerge(t, out, 32)
				if len(rest) != 0 || !closed {
					t.Errorf("both closed remaining = %v, closed=%t, want [], true", rest, closed)
				}
			})
		})
	}
}
func checkBufferedMerge(t *testing.T, a, b []int) {
	t.Helper()
	synctest.Test(t, func(t *testing.T) {
		left, right := make(chan int, len(a)), make(chan int, len(b))
		for _, v := range a {
			left <- v
		}
		for _, v := range b {
			right <- v
		}
		close(left)
		close(right)
		out := subject.MergeTwo(left, right)
		defer func() { drainMerge(t, out, 64); synctest.Wait() }()
		got, closed := drainMerge(t, out, 64)
		if !closed {
			t.Errorf("finite merge closed = %t, want true", closed)
		}
		want := append(append([]int(nil), a...), b...)
		sortedGot := append([]int(nil), got...)
		slices.Sort(sortedGot)
		slices.Sort(want)
		if !slices.Equal(sortedGot, want) {
			t.Errorf("occurrences = %v, want %v", sortedGot, want)
		}
		// These tagged cases use disjoint ranges; interleaving has no total-order rule.
		if len(a) > 0 && len(b) > 0 && a[0] >= 100 && b[0] >= 200 {
			var fromLeft, fromRight []int
			for _, v := range got {
				if v < 200 {
					fromLeft = append(fromLeft, v)
				} else {
					fromRight = append(fromRight, v)
				}
			}
			if !slices.Equal(fromLeft, a) || !slices.Equal(fromRight, b) {
				t.Errorf("stream order = %v / %v, want %v / %v", fromLeft, fromRight, a, b)
			}
		}
	})
}
func TestFiniteOccurrencesAndOrder(t *testing.T) {
	checkBufferedMerge(t, nil, nil)
	checkBufferedMerge(t, nil, []int{0, 3})
	checkBufferedMerge(t, []int{0, 2, 2}, []int{0, 3, 3})
	checkBufferedMerge(t, []int{103, 101, 103}, []int{202, 200, 201})
}
