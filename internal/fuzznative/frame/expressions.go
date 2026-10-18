package frame

import (
	"go/ast"
	"go/token"
	"go/types"
)

func (a *frame) expr(x ast.Expr, state scopeState, parent bool) (value fact, err error) {
	defer func() {
		if a.exhausted != nil {
			value = fact{}
			err = a.exhausted
		}
	}()
	if err := a.charge(x, 1); err != nil {
		return fact{}, err
	}
	x = unparen(x)
	switch e := x.(type) {
	case *ast.Ident:
		obj := a.info.Uses[e]
		if obj == types.Universe.Lookup("nil") {
			return fact{kind: otherBytesOrigin, knownBytes: true}, nil
		}
		if v, ok := state[obj]; ok {
			return v.value, nil
		}
		if _, ok := obj.(*types.Const); ok && obj.Parent() == types.Universe {
			return fact{kind: scalarOrigin}, nil
		}
		if fn, ok := obj.(*types.Func); ok {
			return a.functionValue(e, fn)
		}
		return fact{}, a.fail(e, "frame.capture", "identifier is not a local binding/native parameter or allowed static function")
	case *ast.BasicLit:
		if !scalarType(a.info.TypeOf(e)) {
			return fact{}, a.fail(e, "frame.expression", "unsupported scalar literal")
		}
		return fact{kind: scalarOrigin}, nil
	case *ast.SelectorExpr:
		if a.info.Selections[e] != nil {
			return fact{}, a.fail(e, "frame.call", "method/field values unsupported")
		}
		fn, ok := a.info.Uses[e.Sel].(*types.Func)
		if !ok {
			return fact{}, a.fail(e, "frame.expression", "only allowed package function selectors")
		}
		return a.functionValue(e, fn)
	case *ast.UnaryExpr:
		if e.Op != token.ADD && e.Op != token.SUB && e.Op != token.NOT && e.Op != token.XOR {
			return fact{}, a.fail(e, "frame.expression", "address/dereference/channel operation unsupported")
		}
		v, err := a.expr(e.X, state, parent)
		if err != nil {
			return fact{}, err
		}
		if v.kind != scalarOrigin {
			return fact{}, a.fail(e, "frame.expression", "unary scalar operand required")
		}
		return fact{kind: scalarOrigin}, nil
	case *ast.BinaryExpr:
		left, err := a.expr(e.X, state, parent)
		if err != nil {
			return fact{}, err
		}
		right, err := a.expr(e.Y, state, parent)
		if err != nil {
			return fact{}, err
		}
		if left.kind == scalarOrigin && right.kind == scalarOrigin {
			return fact{kind: scalarOrigin}, nil
		}
		// []byte-versus-nil representation assertions are syntactically allowed;
		// later correct value-equivalent implementation variants must reject any
		// harness which imposes an uncontracted nil/ownership requirement.
		if (e.Op == token.EQL || e.Op == token.NEQ) && byteFact(left) && byteFact(right) {
			return fact{kind: scalarOrigin}, nil
		}
		return fact{}, a.fail(e, "frame.expression", "unsupported non-scalar binary operands")
	case *ast.IndexExpr:
		base, err := a.expr(e.X, state, parent)
		if err != nil {
			return fact{}, err
		}
		index, err := a.expr(e.Index, state, parent)
		if err != nil {
			return fact{}, err
		}
		if !byteFact(base) || index.kind != scalarOrigin || !byteSlice(a.info.TypeOf(e.X)) {
			return fact{}, a.fail(e, "frame.expression", "read-only []byte scalar index required")
		}
		return fact{kind: scalarOrigin}, nil
	case *ast.SliceExpr:
		if e.Slice3 {
			return fact{}, a.fail(e, "frame.expression", "three-index slices unsupported")
		}
		base, err := a.expr(e.X, state, parent)
		if err != nil {
			return fact{}, err
		}
		if !byteFact(base) {
			return fact{}, a.fail(e, "frame.expression", "read-only []byte slice required")
		}
		for _, bound := range []ast.Expr{e.Low, e.High} {
			if bound == nil {
				continue
			}
			v, err := a.expr(bound, state, parent)
			if err != nil {
				return fact{}, err
			}
			if v.kind != scalarOrigin {
				return fact{}, a.fail(bound, "frame.expression", "scalar slice bound required")
			}
		}
		if e.Low == nil && e.High == nil {
			return base, nil
		}
		if low, ok := a.constantInt(e.Low); ok && low == 0 {
			id, xok := unparen(e.X).(*ast.Ident)
			call, cok := unparen(e.High).(*ast.CallExpr)
			if xok && cok && len(call.Args) == 1 {
				fun, fok := unparen(call.Fun).(*ast.Ident)
				arg, aok := unparen(call.Args[0]).(*ast.Ident)
				if fok && aok && a.info.Uses[fun] == types.Universe.Lookup("len") && a.info.Uses[id] == a.info.Uses[arg] {
					return base, nil
				}
			}
		}
		return fact{kind: otherBytesOrigin}, nil
	case *ast.CompositeLit:
		if _, err := a.literalLength(e); err != nil {
			return fact{}, err
		}
		if seedTable(a.info.TypeOf(e)) && parent {
			var table [][]byte
			nextIndex := 0
			for _, item := range e.Elts {
				index := nextIndex
				if keyed, ok := item.(*ast.KeyValueExpr); ok {
					n, known := a.constantInt(keyed.Key)
					if !known || n < 0 || n >= MaxLiteralLength {
						return fact{}, a.fail(item, "frame.seeds", "finite seed table index required")
					}
					index = int(n)
					item = keyed.Value
				}
				v, err := a.expr(item, state, true)
				if err != nil {
					return fact{}, err
				}
				if v.kind != otherBytesOrigin || !v.knownBytes {
					return fact{}, a.fail(item, "frame.seeds", "finite seed-list element required")
				}
				for len(table) <= index {
					if err := a.charge(e, 1); err != nil {
						return fact{}, err
					}
					table = append(table, nil)
				}
				table[index] = v.bytes
				nextIndex = index + 1
			}
			return fact{kind: seedTableOrigin, table: table}, nil
		}
		if !byteSlice(a.info.TypeOf(e)) {
			return fact{}, a.fail(e, "frame.expression", "only []byte literals; [][]byte in parent only")
		}
		for _, item := range e.Elts {
			if kv, ok := item.(*ast.KeyValueExpr); ok {
				if _, err := a.expr(kv.Key, state, parent); err != nil {
					return fact{}, err
				}
				item = kv.Value
			}
			v, err := a.expr(item, state, parent)
			if err != nil {
				return fact{}, err
			}
			if v.kind != scalarOrigin {
				return fact{}, a.fail(item, "frame.expression", "scalar byte literal element required")
			}
		}
		b, known := a.constantBytes(e)
		return fact{kind: otherBytesOrigin, bytes: b, knownBytes: known}, nil
	case *ast.CallExpr:
		return a.call(e, state, parent)
	default:
		return fact{}, a.fail(e, "frame.expression", "unsupported expression %T", e)
	}
}

func (a *frame) functionValue(site ast.Node, fn *types.Func) (fact, error) {
	if a.functions[fn] != nil && a.functions[fn] != a.parent {
		return fact{kind: functionOrigin, function: fn}, nil
	}
	if fn.Pkg() != nil {
		path, name := fn.Pkg().Path(), fn.Name()
		if a.isNamedSubjectAPI(fn) {
			if a.subjectAPIName(fn) == "" {
				return fact{}, a.fail(site, "frame.api-signature", "subject Encode/Decode must have the exact finite byte API signature")
			}
			return fact{kind: functionOrigin, function: fn}, nil
		}
		if path == "encoding/hex" && name == "AppendEncode" && hexOracleSignature(fn) {
			return fact{kind: functionOrigin, function: fn}, nil
		}
		if path == "bytes" && (name == "Equal" || name == "Clone") {
			return fact{kind: functionOrigin, function: fn}, nil
		}
	}
	return fact{}, a.fail(site, "frame.call", "unsupported static function")
}

func (a *frame) call(call *ast.CallExpr, state scopeState, parent bool) (fact, error) {
	if call.Ellipsis != token.NoPos {
		return fact{}, a.fail(call, "frame.call", "variadic expansion unsupported")
	}
	if a.info.Types[call.Fun].IsType() {
		if len(call.Args) != 1 {
			return fact{}, a.fail(call, "frame.expression", "one conversion operand required")
		}
		if byteSlice(a.info.TypeOf(call)) {
			if b, ok := a.constantBytes(call); ok {
				return fact{kind: otherBytesOrigin, bytes: b, knownBytes: true}, nil
			}
			return fact{}, a.fail(call, "frame.expression", "[]byte conversion only of string literal or nil constant")
		}
		b, ok := a.info.TypeOf(call).(*types.Basic)
		if !ok || b.Info()&types.IsInteger == 0 || b.Kind() == types.Uintptr {
			return fact{}, a.fail(call, "frame.expression", "only allowed integer scalar conversions")
		}
		v, err := a.expr(call.Args[0], state, parent)
		if err != nil {
			return fact{}, err
		}
		arg, ok := a.info.TypeOf(call.Args[0]).(*types.Basic)
		if !ok || arg.Info()&types.IsInteger == 0 || arg.Kind() == types.Uintptr || v.kind != scalarOrigin {
			return fact{}, a.fail(call, "frame.expression", "integer scalar conversion operand required")
		}
		return fact{kind: scalarOrigin}, nil
	}
	if id, ok := unparen(call.Fun).(*ast.Ident); ok {
		if b, ok := a.info.Uses[id].(*types.Builtin); ok {
			if b != types.Universe.Lookup("len") || len(call.Args) != 1 {
				return fact{}, a.fail(call, "frame.call", "only len builtin is supported")
			}
			v, err := a.expr(call.Args[0], state, parent)
			if err != nil {
				return fact{}, err
			}
			if !byteFact(v) && v.kind != seedTableOrigin && v.kind != scalarOrigin {
				return fact{}, a.fail(call, "frame.expression", "unsupported len operand")
			}
			return fact{kind: scalarOrigin}, nil
		}
	}
	if method, receiver, ok := a.testingMethod(call.Fun, state); ok {
		if parent || receiver.kind != testingTOrigin {
			return fact{}, a.fail(call, "frame.testing-receiver", "native callback T required")
		}
		if method != "Helper" && method != "Error" && method != "Errorf" && method != "Fatal" && method != "Fatalf" {
			return fact{}, a.fail(call, "frame.call", "unsupported testing.T method %s", method)
		}
		for _, arg := range call.Args {
			v, err := a.expr(arg, state, false)
			if err != nil {
				return fact{}, err
			}
			if v.kind != scalarOrigin && !byteFact(v) {
				return fact{}, a.fail(arg, "frame.call", "diagnostics accept scalar/read-only byte values only")
			}
		}
		return fact{kind: invalidOrigin}, nil
	}
	fun, err := a.expr(call.Fun, state, parent)
	if err != nil {
		return fact{}, err
	}
	if fun.kind != functionOrigin || fun.function == nil {
		return fact{}, a.fail(call, "frame.call", "statically resolved immutable function target required")
	}
	if parent {
		return fact{}, a.fail(call, "frame.call", "SUT/library/helper computation in parent unsupported")
	}
	var args []fact
	for _, arg := range call.Args {
		v, err := a.expr(arg, state, false)
		if err != nil {
			return fact{}, err
		}
		args = append(args, v)
	}
	fn := fun.function
	if a.functions[fn] != nil {
		return a.helper(fn, args, call)
	}
	path, name := fn.Pkg().Path(), fn.Name()
	apiName := a.subjectAPIName(fn)
	if a.isNamedSubjectAPI(fn) && apiName == "" {
		return fact{}, a.fail(call, "frame.api-signature", "subject Encode/Decode must have the exact finite byte API signature")
	}
	if path == a.identity.SubjectPath() && name == "Encode" && apiName == "Encode" {
		if len(args) != 1 || !isSeed(args[0]) {
			return fact{}, a.fail(call, "origin.argument", "Encode requires current Seed/SeedClone")
		}
		a.apiCalls[call] = fn
		return fact{kind: encodedOrigin}, nil
	}
	if path == a.identity.SubjectPath() && name == "Decode" && apiName == "Decode" {
		if len(args) != 1 || !isEncoded(args[0]) {
			return fact{}, a.fail(call, "origin.argument", "Decode requires current Encoded/EncodedClone")
		}
		a.apiCalls[call] = fn
		return fact{kind: tupleOrigin, parts: []fact{{kind: decodedOrigin}, {kind: scalarOrigin}}}, nil
	}
	if path == "encoding/hex" && name == "AppendEncode" && hexOracleSignature(fn) {
		if parent || len(args) != 2 || !a.hexNilDestination(call.Args[0]) || !isSeed(args[1]) {
			return fact{}, a.fail(call, "origin.oracle", "hex oracle requires literal nil/[]byte(nil) destination and current Seed/SeedClone input")
		}
		return fact{kind: otherBytesOrigin}, nil
	}
	if path == "bytes" && name == "Equal" {
		if len(args) != 2 || !byteFact(args[0]) || !byteFact(args[1]) {
			return fact{}, a.fail(call, "frame.call", "Equal requires two read-only byte values")
		}
		return fact{kind: scalarOrigin}, nil
	}
	if path == "bytes" && name == "Clone" {
		if len(args) != 1 || !byteFact(args[0]) {
			return fact{}, a.fail(call, "frame.call", "Clone requires a read-only byte value")
		}
		v := args[0]
		if isSeed(v) {
			v.kind = cloneOrigin
		} else if isEncoded(v) {
			v.kind = encodedCloneOrigin
		} else if isDecoded(v) {
			v.kind = decodedCloneOrigin
		}
		return v, nil
	}
	return fact{}, a.fail(call, "frame.call", "unsupported function")
}
