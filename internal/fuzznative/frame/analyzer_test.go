package frame_test

import (
	"encoding/json"
	"fmt"
	"go/token"
	"go/types"
	"os"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative/frame"
)

type imports map[string]*types.Package

func (i imports) Import(path string) (*types.Package, error) {
	if p := i[path]; p != nil {
		return p, nil
	}
	return nil, fmt.Errorf("unexpected import %s", path)
}

func signature(receiver types.Type, params, results []types.Type, variadic bool) *types.Signature {
	variables := func(values []types.Type) *types.Tuple {
		var vars []*types.Var
		for n, value := range values {
			vars = append(vars, types.NewVar(token.NoPos, nil, fmt.Sprintf("p%d", n), value))
		}
		return types.NewTuple(vars...)
	}
	var recv *types.Var
	if receiver != nil {
		recv = types.NewVar(token.NoPos, nil, "", receiver)
	}
	return types.NewSignatureType(recv, nil, nil, variables(params), variables(results), variadic)
}

func testImports(identity frame.Identity) imports {
	result := imports{}
	addPackage := func(path, name string) *types.Package { p := types.NewPackage(path, name); result[path] = p; return p }
	fn := func(p *types.Package, name string, params, returns []types.Type) {
		p.Scope().Insert(types.NewFunc(token.NoPos, p, name, signature(nil, params, returns, false)))
	}
	bytesType := types.NewSlice(types.Typ[types.Uint8])
	boolean := types.Typ[types.Bool]
	anyType := types.Universe.Lookup("any").Type()
	args := types.NewSlice(anyType)
	testingPackage := addPackage("testing", "testing")
	named := func(name string) *types.Named {
		obj := types.NewTypeName(token.NoPos, testingPackage, name, nil)
		value := types.NewNamed(obj, types.NewStruct(nil, nil), nil)
		testingPackage.Scope().Insert(obj)
		return value
	}
	typeT, typeF := named("T"), named("F")
	method := func(typ *types.Named, name string, params, returns []types.Type, variadic bool) {
		typ.AddMethod(types.NewFunc(token.NoPos, testingPackage, name, signature(types.NewPointer(typ), params, returns, variadic)))
	}
	method(typeF, "Add", []types.Type{args}, nil, true)
	method(typeF, "Fuzz", []types.Type{anyType}, nil, false)
	method(typeT, "Helper", nil, nil, false)
	method(typeT, "Name", nil, []types.Type{types.Typ[types.String]}, false)
	for _, name := range []string{"Error", "Fatal", "Skip", "Log"} {
		method(typeT, name, []types.Type{args}, nil, true)
	}
	for _, name := range []string{"Errorf", "Fatalf"} {
		method(typeT, name, []types.Type{types.Typ[types.String], args}, nil, true)
	}
	bytesPackage := addPackage("bytes", "bytes")
	fn(bytesPackage, "Equal", []types.Type{bytesType, bytesType}, []types.Type{boolean})
	fn(bytesPackage, "Clone", []types.Type{bytesType}, []types.Type{bytesType})
	hexPackage := addPackage("encoding/hex", "hex")
	fn(hexPackage, "AppendEncode", []types.Type{bytesType, bytesType}, []types.Type{bytesType})
	fn(hexPackage, "AppendDecode", []types.Type{bytesType, bytesType}, []types.Type{bytesType, types.Universe.Lookup("error").Type()})
	fn(hexPackage, "EncodeToString", []types.Type{bytesType}, []types.Type{types.Typ[types.String]})
	subject := addPackage(identity.SubjectPath(), "subject")
	fn(subject, "Encode", []types.Type{bytesType}, []types.Type{bytesType})
	fn(subject, "Decode", []types.Type{bytesType}, []types.Type{bytesType, boolean})
	for _, p := range result {
		p.MarkComplete()
	}
	return result
}

// The finite examples were authored for this repository's private mechanism
// proof and reviewed against ADR 0032. They contain no upstream task content.
// Pure tests use a typed import fixture; tagged integration also exercises the
// accepted source forms against real compiler exports in a disposable module.
func TestReviewedFiniteForms(t *testing.T) {
	data, err := os.ReadFile("testdata/finite-forms.json")
	if err != nil {
		t.Fatal(err)
	}
	var cases []struct {
		Name       string
		Accepted   bool
		Diagnostic string
		Sources    map[string]string
	}
	if err := json.Unmarshal(data, &cases); err != nil {
		t.Fatal(err)
	}
	if len(cases) != 61 {
		t.Fatalf("incomplete finite source matrix: %d", len(cases))
	}
	identity := frame.Identity{ModulePath: "task.local/typed-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			imports := testImports(identity)
			if strings.HasPrefix(test.Name, "subject-named-") {
				subject := types.NewPackage(identity.SubjectPath(), "subject")
				byteSlice := types.NewSlice(types.Typ[types.Uint8])
				named := types.NewNamed(types.NewTypeName(token.NoPos, subject, "Payload", nil), byteSlice, nil)
				var encodeResult, decodeResult types.Type = byteSlice, byteSlice
				if test.Name == "subject-named-encode-result" {
					encodeResult = named
				} else {
					decodeResult = named
				}
				subject.Scope().Insert(types.NewFunc(token.NoPos, subject, "Encode", signature(nil, []types.Type{byteSlice}, []types.Type{encodeResult}, false)))
				subject.Scope().Insert(types.NewFunc(token.NoPos, subject, "Decode", signature(nil, []types.Type{byteSlice}, []types.Type{decodeResult, types.Typ[types.Bool]}, false)))
				subject.MarkComplete()
				imports[identity.SubjectPath()] = subject
			}
			inspection, err := frame.Analyze(token.NewFileSet(), imports, test.Sources, identity)
			if test.Accepted {
				if err != nil {
					t.Fatal(err)
				}
				if _, err := inspection.Rewrite(identity.ObserverPath()); err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), test.Diagnostic) {
				t.Fatalf("want useful %s rejection, got %v", test.Diagnostic, err)
			}
		})
	}
}

const completeBody = `wire:=subject.Encode(input);decoded,ok:=subject.Decode(wire);if !ok||!bytes.Equal(decoded,input){t.Error("payload/status")}`

func source(identity frame.Identity, body, extra string) string {
	return `package subject_test
import("testing";"bytes";"` + identity.SubjectPath() + `")
func FuzzHex(f *testing.F){f.Add([]byte(nil));f.Add([]byte{0,128});f.Fuzz(func(t *testing.T,input []byte){` + body + `})}
func equal(a,b []byte)bool{return bytes.Equal(a,b)}
` + extra
}

func TestFiniteSourceFrame(t *testing.T) {
	identity := frame.Identity{ModulePath: "task.local/typed-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	tests := []struct{ name, body, extra, want string }{
		{"direct", completeBody, "", ""},
		{"aliases", `encode:=subject.Encode;decode:=subject.Decode;wire:=encode(input);decoded,ok:=decode(wire);if !ok||!bytes.Equal(decoded,input){t.Fatal("mismatch")}`, "", ""},
		{"clones", `wire:=subject.Encode(bytes.Clone(input));decoded,ok:=subject.Decode(bytes.Clone(wire));if !ok||!bytes.Equal(bytes.Clone(decoded),input){t.Error("mismatch")}`, "", ""},
		{"tuple-helper", `wire:=subject.Encode(input);decoded,ok:=decode(wire);if !ok||!bytes.Equal(decoded,input){t.Error("mismatch")}`, `func decode(wire []byte)([]byte,bool){return subject.Decode(wire)}`, ""},
		{"unused-decode", `wire:=subject.Encode(input);_ = wire`, `func unused(wire []byte){subject.Decode(wire)}`, ""},
		{"other-bytes-encode", `wire:=subject.Encode([]byte{0});subject.Decode(wire)`, "", "origin.argument"},
		{"decoded-reuse", `wire:=subject.Encode(input);decoded,_:=subject.Decode(wire);subject.Encode(decoded)`, "", "origin.argument"},
		{"missing-decode", `wire:=subject.Encode(input);_ = wire`, "", "frame.participation"},
		{"element-write", `input[0]=0;` + completeBody, "", "frame.write"},
		{"rebound-bytes", `input=[]byte{0};` + completeBody, "", "frame.write"},
		{"native-name", `t.Name();` + completeBody, "", "frame.call"},
		{"extra-test", completeBody, `func TestExtra(t *testing.T){}`, "frame.top-level"},
		{"cyclic-helper", `loop(input);` + completeBody, `func loop(input []byte){loop(input)}`, "frame.recursion"},
		{"tuple-var", `wire:=subject.Encode(input);var decoded,ok=subject.Decode(wire);_ = decoded;_ = ok`, "", "frame.declaration"},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			inspection, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{"fuzz_test.go": source(identity, test.body, test.extra)}, identity)
			if test.want != "" {
				if err == nil || !strings.Contains(err.Error(), test.want) {
					t.Fatalf("want %s diagnostic, got %v", test.want, err)
				}
				return
			}
			if err != nil {
				t.Fatal(err)
			}
			rewritten, err := inspection.Rewrite(identity.ObserverPath())
			if err != nil {
				t.Fatal(err)
			}
			if !strings.Contains(rewritten.Sources["fuzz_test.go"], identity.ObserverPath()) {
				t.Fatal("rewrite lost generated observer identity")
			}
		})
	}
}

func TestSourceSelectionIsExplicit(t *testing.T) {
	identity := frame.Identity{ModulePath: "task.local/typed-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	for _, name := range []string{"fuzz_linux_test.go", "fuzz_arm64_test.go", "fuzz_windows_amd64_test.go", "_hidden_test.go", ".hidden_test.go"} {
		t.Run(name, func(t *testing.T) {
			_, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{name: source(identity, completeBody, "")}, identity)
			if err == nil || !strings.Contains(err.Error(), "frame.files") {
				t.Fatalf("want file selection diagnostic, got %v", err)
			}
		})
	}
	for _, directive := range []string{"//go:build linux", "// +build linux", "//   +build !ignore", "//line redirected.go:10", "//go:generate arbitrary"} {
		t.Run(directive, func(t *testing.T) {
			_, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{"fuzz_test.go": directive + "\n\n" + source(identity, completeBody, "")}, identity)
			if err == nil || !strings.Contains(err.Error(), "frame.directives") {
				t.Fatalf("want directive diagnostic, got %v", err)
			}
		})
	}
}
