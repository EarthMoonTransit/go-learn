package subject

import (
	"go/ast"
	"go/importer"
	"go/parser"
	"go/token"
	"go/types"
	"strconv"
	"testing"
)

func p18Package(t *testing.T) (*token.FileSet, []*ast.File, *types.Package, *types.Info) {
	t.Helper()
	set := token.NewFileSet()
	files := []*ast.File{}
	for _, name := range []string{"api.go", "implementation.go"} {
		file, err := parser.ParseFile(set, name, nil, 0)
		if err != nil {
			t.Fatal(err)
		}
		files = append(files, file)
	}
	info := &types.Info{Types: make(map[ast.Expr]types.TypeAndValue)}
	config := types.Config{Importer: importer.ForCompiler(set, "source", nil), GoVersion: "go1.27"}
	pkg, err := config.Check("task.local/mapped-records-001/subject", set, files, info)
	if err != nil {
		t.Fatalf("implementation type check: %v", err)
	}
	return set, files, pkg, info
}

func TestGenericDeclarationAndCallMatrix(t *testing.T) {
	set, files, pkg, _ := p18Package(t)
	object := pkg.Scope().Lookup("mapEntries")
	fn, ok := object.(*types.Func)
	if !ok {
		t.Fatalf("mapEntries declaration = %T, want a function declaration", object)
	}
	sig := fn.Type().(*types.Signature)
	if got := sig.TypeParams().Len(); got != 2 {
		t.Fatalf("mapEntries type parameters = %d, want 2 independent type parameters", got)
	}
	a, b := sig.TypeParams().At(0), sig.TypeParams().At(1)
	unrestricted := types.NewInterfaceType(nil, nil).Complete()
	for i, parameter := range []*types.TypeParam{a, b} {
		if !types.Identical(parameter.Constraint().Underlying(), unrestricted) {
			t.Fatalf("mapEntries type parameter %d constraint = %s, want unrestricted any", i+1, parameter.Constraint())
		}
	}
	callback := types.NewSignatureType(nil, nil, nil,
		types.NewTuple(types.NewVar(token.NoPos, nil, "", a)),
		types.NewTuple(types.NewVar(token.NoPos, nil, "", b)), false)
	entry, err := types.Instantiate(nil, pkg.Scope().Lookup("Entry").Type(), []types.Type{a, b}, true)
	if err != nil {
		t.Fatal(err)
	}
	if sig.Variadic() || sig.Params().Len() != 2 || sig.Results().Len() != 1 {
		t.Fatalf("mapEntries signature = %s, want []A plus func(A) B returning []Entry[A, B]", sig)
	}
	wants := []struct {
		label     string
		got, want types.Type
	}{
		{"input", sig.Params().At(0).Type(), types.NewSlice(a)},
		{"callback", sig.Params().At(1).Type(), callback},
		{"result", sig.Results().At(0).Type(), types.NewSlice(entry)},
	}
	for _, c := range wants {
		if !types.Identical(c.got, c.want) {
			t.Fatalf("mapEntries %s type = %s, want %s; preserve both type parameters", c.label, c.got, c.want)
		}
	}
	// This source is type-checked in memory. It is never added to the Go test
	// package, so the deliberately non-generic starter remains compilable.
	const matrix = `package subject
 type p18MatrixInput struct { Data []byte }
 type p18MatrixInputs []p18MatrixInput
 type p18MatrixOutput []byte
 var _ func([]p18MatrixInput, func(p18MatrixInput) []byte) []Entry[p18MatrixInput, []byte] = mapEntries[p18MatrixInput, []byte]
 var _ func([]int, func(int) map[string]int) []Entry[int, map[string]int] = mapEntries[int, map[string]int]
 func p18MatrixCalls() {
  var input p18MatrixInputs
  mapper := func(v p18MatrixInput) p18MatrixOutput { return v.Data }
  var inferred []Entry[p18MatrixInput, p18MatrixOutput] = mapEntries(input, mapper)
  var explicit []Entry[p18MatrixInput, p18MatrixOutput] = mapEntries[p18MatrixInput, p18MatrixOutput](input, mapper)
  var numeric []Entry[string, int] = mapEntries([]string{"x"}, func(v string) int { return len(v) })
  _, _, _ = inferred, explicit, numeric
 }
 `
	calls, err := parser.ParseFile(set, "call_matrix.go", matrix, 0)
	if err != nil {
		t.Fatal(err)
	}
	config := types.Config{Importer: importer.ForCompiler(set, "source", nil), GoVersion: "go1.27"}
	if _, err := config.Check("task.local/mapped-records-001/subject", set, append(files, calls), nil); err != nil {
		t.Fatalf("explicit/inferred call matrix rejected: %v; want named input and unrestricted independent result types", err)
	}
}

func TestNoValueErasureOrReflection(t *testing.T) {
	set, files, _, info := p18Package(t)
	for _, file := range files {
		for _, imported := range file.Imports {
			name, err := strconv.Unquote(imported.Path.Value)
			if err != nil {
				t.Fatal(err)
			}
			if name == "reflect" {
				t.Error("implementation imports reflect; want direct typed mapping without reflection")
			}
		}
		for _, declaration := range file.Decls {
			fn, ok := declaration.(*ast.FuncDecl)
			if !ok || fn.Body == nil {
				continue
			}
			for expression, value := range info.Types {
				if expression.Pos() < fn.Body.Pos() || expression.End() > fn.Body.End() || value.Type == nil {
					continue
				}
				typ := types.Unalias(value.Type)
				if _, parameter := typ.(*types.TypeParam); parameter {
					continue
				}
				if _, erased := typ.Underlying().(*types.Interface); erased {
					t.Errorf("value type at %s = %s; want type parameters without interface-value erasure", set.Position(expression.Pos()), typ)
				}
			}
		}
	}
}

func TestNoGoroutineLaunch(t *testing.T) {
	set, files, _, _ := p18Package(t)
	for _, file := range files {
		ast.Inspect(file, func(node ast.Node) bool {
			if statement, ok := node.(*ast.GoStmt); ok {
				t.Errorf("goroutine launch at %s; want synchronous mapping without go statements", set.Position(statement.Pos()))
			}
			return true
		})
	}
}
