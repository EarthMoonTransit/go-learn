package frame

import (
	"errors"
	"go/ast"
	"go/token"
	"go/types"
)

func booleanType(t types.Type) bool { b, ok := t.(*types.Basic); return ok && b.Kind() == types.Bool }
func tupleSignature(sig *types.Signature) bool {
	return sig.Results().Len() == 2 && byteSlice(sig.Results().At(0).Type()) && booleanType(sig.Results().At(1).Type())
}
func isEncoded(v fact) bool { return v.kind == encodedOrigin || v.kind == encodedCloneOrigin }
func isDecoded(v fact) bool { return v.kind == decodedOrigin || v.kind == decodedCloneOrigin }
func isResult(v fact) bool  { return isEncoded(v) || isDecoded(v) || v.kind == resultOrigin }

// subjectAPIName binds both package path and the exact nonvariadic Go
// signatures consumed by the observer wrappers. Merely exporting a matching
// name is not sufficient evidence that the checked function is the contract.
func (a *frame) subjectAPIName(fn *types.Func) string {
	if fn == nil || fn.Pkg() == nil || fn.Pkg().Path() != a.identity.SubjectPath() {
		return ""
	}
	sig, ok := fn.Type().(*types.Signature)
	if !ok || sig.Recv() != nil || sig.Variadic() || (sig.TypeParams() != nil && sig.TypeParams().Len() != 0) || sig.Params().Len() != 1 || !byteSlice(sig.Params().At(0).Type()) {
		return ""
	}
	switch fn.Name() {
	case "Encode":
		if sig.Results().Len() == 1 && byteSlice(sig.Results().At(0).Type()) {
			return "Encode"
		}
	case "Decode":
		if sig.Results().Len() == 2 && byteSlice(sig.Results().At(0).Type()) && booleanType(sig.Results().At(1).Type()) {
			return "Decode"
		}
	}
	return ""
}

func (a *frame) isNamedSubjectAPI(fn *types.Func) bool {
	return fn != nil && fn.Pkg() != nil && fn.Pkg().Path() == a.identity.SubjectPath() && (fn.Name() == "Encode" || fn.Name() == "Decode")
}
func (a *frame) tupleLocal(stmt *ast.AssignStmt, state scopeState) error {
	if stmt.Tok != token.DEFINE || len(stmt.Rhs) != 1 {
		return a.fail(stmt, "frame.tuple", "two fresh local bindings from one exact tuple call required")
	}
	if _, ok := unparen(stmt.Rhs[0]).(*ast.CallExpr); !ok {
		return a.fail(stmt, "frame.tuple", "tuple RHS must be static Decode/helper call")
	}
	v, err := a.expr(stmt.Rhs[0], state, false)
	if err != nil {
		return err
	}
	if v.kind != tupleOrigin || len(v.parts) != 2 {
		return a.fail(stmt, "frame.tuple", "exact([]byte,bool) tuple required")
	}
	for i, lhs := range stmt.Lhs {
		id, ok := lhs.(*ast.Ident)
		if !ok {
			return a.fail(lhs, "frame.write", "fresh tuple identifiers required")
		}
		if id.Name == "_" {
			continue
		}
		obj := a.info.Defs[id]
		if obj == nil {
			return a.fail(id, "frame.write", "tuple bindings must be fresh")
		}
		if (i == 0 && !byteSlice(obj.Type())) || (i == 1 && !booleanType(obj.Type())) {
			return a.fail(id, "frame.tuple", "tuple binding types must be []byte, bool")
		}
		state[obj] = binding{value: v.parts[i], mutableScalar: i == 1}
	}
	return nil
}

// Each complete role attempt uses ordinary strict evaluation, including nested
// actual helper arguments. The shared budget/resource failure is never reset.
func (a *frame) helperRoles(obj *types.Func, site ast.Node) error {
	sig := obj.Type().(*types.Signature)
	args := make([]fact, sig.Params().Len())
	var positions []int
	for i := range args {
		t := sig.Params().At(i).Type()
		args[i] = fact{kind: scalarOrigin}
		if byteSlice(t) {
			positions = append(positions, i)
		}
		if testingPointer(t, "T") {
			args[i].kind = testingTOrigin
		}
	}
	roles := []origin{seedOrigin, encodedOrigin, decodedOrigin, otherBytesOrigin}
	var first error
	var attempt func(int) (bool, error)
	attempt = func(index int) (bool, error) {
		if index == len(positions) {
			_, err := a.helper(obj, args, site)
			if err == nil {
				return true, nil
			}
			if a.exhausted != nil {
				return false, a.exhausted
			}
			var fe *frameError
			if errors.As(err, &fe) && (fe.Code == "frame.resources" || fe.Code == "frame.budget") {
				a.exhausted = err
				return false, err
			}
			if first == nil {
				first = err
			}
			return false, nil
		}
		for _, role := range roles {
			args[positions[index]] = fact{kind: role}
			ok, err := attempt(index + 1)
			if ok || err != nil {
				return ok, err
			}
		}
		return false, nil
	}
	ok, err := attempt(0)
	if err != nil {
		return err
	}
	if !ok {
		return first
	}
	return nil
}

// The allowed standard-library oracle cannot mutate caller-owned storage: only
// these two explicit nil-destination expressions are admitted. Its result never
// gains Seed/Encoded lineage, even when byte values match a SUT result.
func (a *frame) hexNilDestination(e ast.Expr) bool {
	e = unparen(e)
	if id, ok := e.(*ast.Ident); ok {
		return a.info.Uses[id] == types.Universe.Lookup("nil")
	}
	call, ok := e.(*ast.CallExpr)
	if !ok || len(call.Args) != 1 || call.Ellipsis != token.NoPos || !a.info.Types[call.Fun].IsType() || !byteSlice(a.info.TypeOf(call)) {
		return false
	}
	id, ok := unparen(call.Args[0]).(*ast.Ident)
	return ok && a.info.Uses[id] == types.Universe.Lookup("nil")
}
func hexOracleSignature(fn *types.Func) bool {
	if fn == nil || fn.Pkg() == nil || fn.Pkg().Path() != "encoding/hex" || fn.Name() != "AppendEncode" {
		return false
	}
	sig, ok := fn.Type().(*types.Signature)
	return ok && sig.Recv() == nil && !sig.Variadic() && (sig.TypeParams() == nil || sig.TypeParams().Len() == 0) && sig.Params().Len() == 2 && byteSlice(sig.Params().At(0).Type()) && byteSlice(sig.Params().At(1).Type()) && sig.Results().Len() == 1 && byteSlice(sig.Results().At(0).Type())
}
