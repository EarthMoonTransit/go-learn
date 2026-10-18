// Finite typed admission for the hex-codec-v1 source frame.
package frame

import (
	"fmt"
	"go/ast"
	"go/constant"
	"go/parser"
	"go/token"
	"go/types"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
	"unicode"
)

type frameError struct{ Code, Position, Detail string }

func (e *frameError) Error() string { return e.Code + " at " + e.Position + ": " + e.Detail }

type origin uint8

const (
	invalidOrigin origin = iota
	scalarOrigin
	seedOrigin
	cloneOrigin
	resultOrigin
	otherBytesOrigin
	testingTOrigin
	testingFOrigin
	functionOrigin
	seedTableOrigin
	encodedOrigin
	encodedCloneOrigin
	decodedOrigin
	decodedCloneOrigin
	tupleOrigin
)

type fact struct {
	kind       origin
	function   *types.Func
	bytes      []byte
	table      [][]byte
	knownBytes bool
	parts      []fact
}
type binding struct {
	value         fact
	mutableScalar bool
}
type scopeState map[types.Object]binding

func cloneState(s scopeState) scopeState {
	out := scopeState{}
	for k, v := range s {
		out[k] = v
	}
	return out
}

type flow struct {
	returns []fact
	falls   bool
}

type frame struct {
	identity     Identity
	fset         *token.FileSet
	info         *types.Info
	files        []*ast.File
	functions    map[*types.Func]*ast.FuncDecl
	parent       *ast.FuncDecl
	callback     ast.Node
	callbackSig  *types.Signature
	callbackBody *ast.BlockStmt
	seeds        [][]byte
	stack        map[*types.Func]bool
	apiCalls     map[*ast.CallExpr]*types.Func
	apiAliases   map[types.Object]*types.Func
	apiAliasInit map[ast.Expr]*types.Func
	structural   bool
	rewritten    bool
	steps        int
	exhausted    error
}

type Inspection struct {
	SeedCount     int      `json:"seed_count"`
	HelperCount   int      `json:"helper_count"`
	NamedCallback bool     `json:"named_callback"`
	SeedValues    [][]byte `json:"seed_values"`
	engine        *frame   `json:"-"`
	AnalysisSteps int      `json:"analysis_steps"`
}

func (a *frame) fail(n ast.Node, code, format string, args ...any) error {
	position := "package"
	if n != nil {
		position = a.fset.Position(n.Pos()).String()
	}
	return &frameError{code, position, fmt.Sprintf(format, args...)}
}
func scalarType(t types.Type) bool {
	b, ok := t.(*types.Basic)
	if !ok {
		return false
	}
	return b.Info()&(types.IsBoolean|types.IsInteger|types.IsString) != 0 && b.Kind() != types.Uintptr
}
func byteSlice(t types.Type) bool {
	s, ok := t.(*types.Slice)
	if !ok {
		return false
	}
	b, ok := s.Elem().(*types.Basic)
	return ok && b.Kind() == types.Uint8
}
func seedTable(t types.Type) bool { s, ok := t.(*types.Slice); return ok && byteSlice(s.Elem()) }
func testingPointer(t types.Type, name string) bool {
	p, ok := t.(*types.Pointer)
	if !ok {
		return false
	}
	n, ok := p.Elem().(*types.Named)
	return ok && n.Obj().Pkg() != nil && n.Obj().Pkg().Path() == "testing" && n.Obj().Name() == name
}
func allowedValueType(t types.Type) bool {
	return scalarType(t) || byteSlice(t) || testingPointer(t, "T")
}
func nativeCallback(sig *types.Signature) bool {
	return sig != nil && !sig.Variadic() && sig.Params().Len() == 2 && testingPointer(sig.Params().At(0).Type(), "T") && byteSlice(sig.Params().At(1).Type()) && sig.Results().Len() == 0
}
func isSeed(f fact) bool { return f.kind == seedOrigin || f.kind == cloneOrigin }

func Analyze(fset *token.FileSet, imp types.Importer, sources map[string]string, identity Identity) (Inspection, error) {
	if err := identity.Validate(); err != nil {
		return Inspection{}, err
	}
	if err := preflight(sources); err != nil {
		return Inspection{}, err
	}
	a := &frame{identity: identity, fset: fset, functions: map[*types.Func]*ast.FuncDecl{}, stack: map[*types.Func]bool{}, apiCalls: map[*ast.CallExpr]*types.Func{}, apiAliases: map[types.Object]*types.Func{}, apiAliasInit: map[ast.Expr]*types.Func{}}
	a.info = &types.Info{Types: map[ast.Expr]types.TypeAndValue{}, Defs: map[*ast.Ident]types.Object{}, Uses: map[*ast.Ident]types.Object{}, Selections: map[*ast.SelectorExpr]*types.Selection{}, Scopes: map[ast.Node]*types.Scope{}, Implicits: map[ast.Node]types.Object{}}
	var names []string
	for name := range sources {
		names = append(names, name)
	}
	sort.Strings(names)
	imports := map[string]bool{}
	for _, name := range names {
		if filepath.Base(name) != name || !strings.HasSuffix(name, "_test.go") {
			return Inspection{}, a.fail(nil, "frame.files", "expected basename _test.go files")
		}
		f, err := parser.ParseFile(fset, name, sources[name], parser.ParseComments|parser.AllErrors)
		if err != nil {
			return Inspection{}, fmt.Errorf("parse original source: %w", err)
		}
		if f.Name.Name != "subject_test" {
			return Inspection{}, a.fail(f, "frame.package", "expected external subject_test package")
		}
		if err := a.checkSelection(name, sources[name], f); err != nil {
			return Inspection{}, err
		}
		for _, spec := range f.Imports {
			path, err := strconv.Unquote(spec.Path.Value)
			if err != nil {
				return Inspection{}, err
			}
			if path != "testing" && path != "bytes" && path != "encoding/hex" && path != a.identity.SubjectPath() {
				return Inspection{}, a.fail(spec, "frame.import", "unsupported import %q", path)
			}
			if spec.Name != nil && (spec.Name.Name == "." || spec.Name.Name == "_") {
				return Inspection{}, a.fail(spec, "frame.import", "dot/blank imports unsupported")
			}
			imports[path] = true
		}
		a.files = append(a.files, f)
	}
	if !imports["testing"] || !imports[a.identity.SubjectPath()] {
		return Inspection{}, a.fail(nil, "frame.import", "package must import testing and subject")
	}
	conf := types.Config{Importer: imp, GoVersion: "go1.27", Sizes: types.SizesFor("gc", runtime.GOARCH)}
	if _, err := conf.Check(a.identity.SubjectPath()+"_test", fset, a.files, a.info); err != nil {
		return Inspection{}, fmt.Errorf("type-check original source: %w", err)
	}
	for _, f := range a.files {
		for _, decl := range f.Decls {
			if gen, ok := decl.(*ast.GenDecl); ok && gen.Tok == token.IMPORT {
				continue
			}
			fn, ok := decl.(*ast.FuncDecl)
			if !ok {
				return Inspection{}, a.fail(decl, "frame.top-level", "only functions/imports are supported")
			}
			if fn.Recv != nil || fn.Body == nil || (fn.Type.TypeParams != nil && len(fn.Type.TypeParams.List) > 0) {
				return Inspection{}, a.fail(fn, "frame.signature", "methods, generics or bodyless functions unsupported")
			}
			obj := a.info.Defs[fn.Name].(*types.Func)
			sig := obj.Type().(*types.Signature)
			if fn.Name.Name == a.identity.Target {
				if a.parent != nil || sig.Params().Len() != 1 || !testingPointer(sig.Params().At(0).Type(), "F") || sig.Results().Len() != 0 || sig.Variadic() {
					return Inspection{}, a.fail(fn, "frame.signature", "target must be func(*testing.F)")
				}
				a.parent = fn
			} else {
				if fn.Name.Name == "init" || nativeEntryName(fn.Name.Name, "Test") || nativeEntryName(fn.Name.Name, "Fuzz") || nativeEntryName(fn.Name.Name, "Benchmark") || nativeEntryName(fn.Name.Name, "Example") {
					return Inspection{}, a.fail(fn, "frame.top-level", "additional native entrypoints/init unsupported")
				}
				if sig.Variadic() || sig.Results().Len() > 2 || (sig.Results().Len() == 2 && !tupleSignature(sig)) || sig.Params().Len() > 8 {
					return Inspection{}, a.fail(fn, "frame.signature", "helpers require at most 8 parameters, zero/one result or an exact ([]byte, bool) pair")
				}
				byteParams := 0
				for i := 0; i < sig.Params().Len(); i++ {
					if byteSlice(sig.Params().At(i).Type()) {
						byteParams++
					}
					if !allowedValueType(sig.Params().At(i).Type()) {
						return Inspection{}, a.fail(fn, "frame.signature", "unsupported helper parameter")
					}
				}
				if byteParams > 4 {
					return Inspection{}, a.fail(fn, "frame.signature", "helpers require at most 4 byte parameters")
				}
				if sig.Results().Len() == 1 && !scalarType(sig.Results().At(0).Type()) && !byteSlice(sig.Results().At(0).Type()) {
					return Inspection{}, a.fail(fn, "frame.signature", "unsupported helper result")
				}
			}
			a.functions[obj] = fn
		}
	}
	if a.parent == nil {
		return Inspection{}, a.fail(nil, "frame.registration", "missing exact native target %s", a.identity.Target)
	}
	parentSig := a.info.Defs[a.parent.Name].Type().(*types.Signature)
	parentState := scopeState{parentSig.Params().At(0): {value: fact{kind: testingFOrigin}}}
	if err := a.parentBlock(a.parent.Body, parentState, true); err != nil {
		return Inspection{}, err
	}
	if a.callback == nil {
		return Inspection{}, a.fail(a.parent, "frame.registration", "missing final F.Fuzz registration")
	}
	if len(a.seeds) == 0 || len(a.seeds) > 64 {
		return Inspection{}, a.fail(a.parent, "frame.seeds", "require 1..64 learner seeds")
	}
	helperCount := len(a.functions) - 1
	if _, ok := a.callback.(*ast.FuncDecl); ok {
		helperCount--
	}
	if helperCount > 16 {
		return Inspection{}, a.fail(a.parent, "frame.helpers", "at most 16 helpers")
	}
	// Hypothetical seed/T bindings validate every helper's finite syntax and
	// acyclic graph, including unused helpers. Actual calls are re-evaluated with
	// their exact origins; these hypothetical bindings grant no call-site credit.
	a.structural = true
	var orderedFunctions []*types.Func
	for obj := range a.functions {
		orderedFunctions = append(orderedFunctions, obj)
	}
	sort.Slice(orderedFunctions, func(i, j int) bool {
		if orderedFunctions[i].Name() != orderedFunctions[j].Name() {
			return orderedFunctions[i].Name() < orderedFunctions[j].Name()
		}
		return orderedFunctions[i].Pos() < orderedFunctions[j].Pos()
	})
	for _, obj := range orderedFunctions {
		fn := a.functions[obj]
		if fn == a.parent {
			continue
		}
		if err := a.helperRoles(obj, fn); err != nil {
			return Inspection{}, err
		}
	}
	a.structural = false
	state := scopeState{a.callbackSig.Params().At(0): {value: fact{kind: testingTOrigin}}, a.callbackSig.Params().At(1): {value: fact{kind: seedOrigin}}}
	if _, err := a.block(a.callbackBody, state, false); err != nil {
		return Inspection{}, err
	}
	participation := map[string]bool{}
	for _, fn := range a.apiCalls {
		participation[a.subjectAPIName(fn)] = true
	}
	if !participation["Encode"] || !participation["Decode"] {
		return Inspection{}, a.fail(a.parent, "frame.participation", "include statically resolved Encode and Decode calls; executed participation is checked separately")
	}
	_, named := a.callback.(*ast.FuncDecl)
	return Inspection{len(a.seeds), helperCount, named, a.seeds, a, a.steps}, nil
}

func (a *frame) parentBlock(body *ast.BlockStmt, state scopeState, top bool) error {
	last := -1
	for i, stmt := range body.List {
		if _, empty := stmt.(*ast.EmptyStmt); !empty {
			last = i
		}
	}
	for index, stmt := range body.List {
		if err := a.charge(stmt, 1); err != nil {
			return err
		}
		switch s := stmt.(type) {
		case *ast.EmptyStmt:
			continue
		case *ast.AssignStmt, *ast.DeclStmt:
			if err := a.local(stmt, state, true); err != nil {
				return err
			}
		case *ast.RangeStmt:
			if s.Tok != token.DEFINE && s.Key != nil {
				return a.fail(s, "frame.write", "range bindings must be fresh")
			}
			v, err := a.expr(s.X, state, true)
			if err != nil {
				return err
			}
			if v.kind != seedTableOrigin {
				return a.fail(s, "frame.seeds", "parent range needs finite [][]byte table")
			}
			// Check even an empty range body once for forbidden syntax, without
			// letting that hypothetical seed contribute to actual F.Add entries.
			values := v.table
			hypothetical := len(values) == 0
			if hypothetical {
				values = [][]byte{nil}
			}
			for _, seed := range values {
				if err := a.charge(s, 1); err != nil {
					return err
				}
				sub := cloneState(state)
				if id, ok := s.Key.(*ast.Ident); ok && id.Name != "_" {
					sub[a.info.Defs[id]] = binding{value: fact{kind: scalarOrigin}}
				}
				if id, ok := s.Value.(*ast.Ident); ok && id.Name != "_" {
					sub[a.info.Defs[id]] = binding{value: fact{kind: otherBytesOrigin, bytes: seed, knownBytes: true}}
				}
				before := len(a.seeds)
				if err := a.parentBlock(s.Body, sub, false); err != nil {
					return err
				}
				if hypothetical {
					a.seeds = a.seeds[:before]
				}
			}
		case *ast.ExprStmt:
			call, ok := s.X.(*ast.CallExpr)
			if !ok {
				return a.fail(s, "frame.registration", "expected F.Add/F.Fuzz")
			}
			method, receiver, ok := a.testingMethod(call.Fun, state)
			if !ok || receiver.kind != testingFOrigin {
				return a.fail(call, "frame.registration", "only native F.Add/F.Fuzz calls allowed in parent")
			}
			switch method {
			case "Add":
				if len(call.Args) != 1 || call.Ellipsis != token.NoPos {
					return a.fail(call, "frame.seeds", "one []byte seed per F.Add")
				}
				if !byteSlice(a.info.TypeOf(call.Args[0])) {
					return a.fail(call.Args[0], "frame.seeds", "F.Add argument must have original []byte type")
				}
				seed, err := a.expr(call.Args[0], state, true)
				if err != nil {
					return err
				}
				if seed.kind != otherBytesOrigin || !seed.knownBytes || len(seed.bytes) > 4096 {
					return a.fail(call, "frame.seeds", "seed must be finite []byte value <=4096 bytes")
				}
				a.seeds = append(a.seeds, append([]byte(nil), seed.bytes...))
				if len(a.seeds) > 64 {
					return a.fail(call, "frame.seeds", "too many learner seeds")
				}
			case "Fuzz":
				if !top || index != last || a.callback != nil || len(call.Args) != 1 || call.Ellipsis != token.NoPos {
					return a.fail(call, "frame.registration", "one final top-level F.Fuzz registration required")
				}
				switch cb := unparen(call.Args[0]).(type) {
				case *ast.FuncLit:
					a.callback = cb
					a.callbackSig = a.info.TypeOf(cb).(*types.Signature)
					a.callbackBody = cb.Body
				case *ast.Ident:
					obj, ok := a.info.Uses[cb].(*types.Func)
					if !ok || a.functions[obj] == nil {
						return a.fail(cb, "frame.registration", "callback must be literal or named submitted function")
					}
					fn := a.functions[obj]
					a.callback = fn
					a.callbackSig = obj.Type().(*types.Signature)
					a.callbackBody = fn.Body
				default:
					return a.fail(call, "frame.registration", "callback must be literal or named submitted function")
				}
				if !nativeCallback(a.callbackSig) {
					return a.fail(call, "frame.signature", "callback must be func(*testing.T,[]byte)")
				}
			default:
				return a.fail(call, "frame.registration", "unsupported native F method %s", method)
			}
		default:
			return a.fail(stmt, "frame.statement", "unsupported parent statement %T", stmt)
		}
	}
	return nil
}

func unparen(x ast.Expr) ast.Expr {
	for {
		p, ok := x.(*ast.ParenExpr)
		if !ok {
			return x
		}
		x = p.X
	}
}

func nativeEntryName(name, prefix string) bool {
	suffix, matched := strings.CutPrefix(name, prefix)
	return matched && (suffix == "" || !unicode.IsLower([]rune(suffix)[0]))
}
func (a *frame) testingMethod(x ast.Expr, state scopeState) (string, fact, bool) {
	sel, ok := unparen(x).(*ast.SelectorExpr)
	if !ok {
		return "", fact{}, false
	}
	selection := a.info.Selections[sel]
	if selection == nil || selection.Kind() != types.MethodVal {
		return "", fact{}, false
	}
	obj := selection.Obj()
	if obj.Pkg() == nil || obj.Pkg().Path() != "testing" {
		return "", fact{}, false
	}
	receiver, err := a.expr(sel.X, state, false)
	if err != nil {
		return "", fact{}, false
	}
	return obj.Name(), receiver, true
}

func joinFacts(values []fact) fact {
	if len(values) == 0 {
		return fact{kind: invalidOrigin}
	}
	out := values[0]
	for _, v := range values[1:] {
		if out.kind == tupleOrigin && v.kind == tupleOrigin && len(out.parts) == 2 && len(v.parts) == 2 {
			out = fact{kind: tupleOrigin, parts: []fact{joinFacts([]fact{out.parts[0], v.parts[0]}), joinFacts([]fact{out.parts[1], v.parts[1]})}}
			continue
		}
		if isEncoded(out) && isEncoded(v) {
			out = fact{kind: encodedCloneOrigin}
			continue
		}
		if isDecoded(out) && isDecoded(v) {
			out = fact{kind: decodedCloneOrigin}
			continue
		}
		if isSeed(out) && isSeed(v) {
			out = fact{kind: cloneOrigin}
			continue
		}
		if out.kind != v.kind {
			return fact{kind: otherBytesOrigin}
		}
		if out.kind == otherBytesOrigin {
			out.knownBytes = false
		}
	}
	return out
}

func (a *frame) constantBytes(x ast.Expr) ([]byte, bool) {
	x = unparen(x)
	if id, ok := x.(*ast.Ident); ok && id.Name == "nil" && a.info.Uses[id] == types.Universe.Lookup("nil") {
		return nil, true
	}
	if call, ok := x.(*ast.CallExpr); ok && len(call.Args) == 1 && a.info.Types[call.Fun].IsType() && byteSlice(a.info.TypeOf(call)) {
		arg := unparen(call.Args[0])
		if id, ok := arg.(*ast.Ident); ok && a.info.Uses[id] == types.Universe.Lookup("nil") {
			return nil, true
		}
		if lit, ok := arg.(*ast.BasicLit); ok && lit.Kind == token.STRING {
			s, err := strconv.Unquote(lit.Value)
			if err != nil {
				return nil, false
			}
			return a.stringBytes(lit, s)
		}
		return nil, false
	}
	lit, ok := x.(*ast.CompositeLit)
	if !ok || !byteSlice(a.info.TypeOf(lit)) {
		return nil, false
	}
	if _, err := a.literalLength(lit); err != nil {
		a.exhausted = err
		return nil, false
	}
	var out []byte
	nextIndex := 0
	for _, item := range lit.Elts {
		if a.charge(item, 1) != nil {
			return nil, false
		}
		expr := item
		index := nextIndex
		if kv, ok := item.(*ast.KeyValueExpr); ok {
			n, ok := a.constantInt(kv.Key)
			if !ok || n < 0 || n >= 4096 {
				return nil, false
			}
			index = int(n)
			expr = kv.Value
		}
		value := a.info.Types[expr].Value
		if value == nil || value.Kind() != constant.Int {
			return nil, false
		}
		n, ok := constant.Int64Val(value)
		if !ok || n < 0 || n > 255 || index >= 4096 {
			return nil, false
		}
		for len(out) <= index {
			if a.charge(lit, 1) != nil {
				return nil, false
			}
			out = append(out, 0)
		}
		out[index] = byte(n)
		nextIndex = index + 1
	}
	return out, true
}
func (a *frame) constantInt(x ast.Expr) (int64, bool) {
	v := a.info.Types[x].Value
	if v == nil || v.Kind() != constant.Int {
		return 0, false
	}
	return constant.Int64Val(v)
}
