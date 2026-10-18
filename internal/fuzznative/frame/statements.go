package frame

import (
	"go/ast"
	"go/token"
	"go/types"
)

func (a *frame) local(stmt ast.Stmt, state scopeState, parent bool) error {
	if assignment, ok := stmt.(*ast.AssignStmt); ok && !parent && len(assignment.Lhs) == 2 {
		return a.tupleLocal(assignment, state)
	}
	var id *ast.Ident
	var rhs ast.Expr
	zero := false
	switch s := stmt.(type) {
	case *ast.AssignStmt:
		if len(s.Lhs) != 1 || len(s.Rhs) != 1 {
			return a.fail(s, "frame.write", "single-name assignments only")
		}
		var ok bool
		id, ok = s.Lhs[0].(*ast.Ident)
		if !ok {
			return a.fail(s, "frame.write", "element/address writes unsupported")
		}
		rhs = s.Rhs[0]
		if s.Tok != token.DEFINE {
			if s.Tok != token.ASSIGN {
				return a.fail(s, "frame.write", "compound assignment unsupported")
			}
			v, err := a.expr(rhs, state, parent)
			if err != nil {
				return err
			}
			if id.Name == "_" {
				return nil
			}
			obj := a.info.Uses[id]
			b, ok := state[obj]
			if !ok || !b.mutableScalar || parent || v.kind != scalarOrigin {
				return a.fail(id, "frame.write", "only mutable local scalar accumulators may be reassigned")
			}
			return nil
		}
	case *ast.DeclStmt:
		gen, ok := s.Decl.(*ast.GenDecl)
		if !ok || gen.Tok != token.VAR || len(gen.Specs) != 1 || gen.Lparen != token.NoPos {
			return a.fail(s, "frame.declaration", "single var declaration only")
		}
		value, ok := gen.Specs[0].(*ast.ValueSpec)
		if !ok || len(value.Names) != 1 || len(value.Values) > 1 || value.Type == nil {
			return a.fail(s, "frame.declaration", "single var declaration only")
		}
		id = value.Names[0]
		if len(value.Values) == 1 {
			rhs = value.Values[0]
		} else {
			zero = true
		}
	default:
		return a.fail(stmt, "frame.declaration", "unsupported local declaration")
	}
	obj := a.info.Defs[id]
	if obj == nil {
		return a.fail(id, "frame.declaration", "new named local required")
	}
	var v fact
	if zero {
		if scalarType(obj.Type()) {
			v.kind = scalarOrigin
		} else if byteSlice(obj.Type()) {
			v = fact{kind: otherBytesOrigin, knownBytes: true}
		} else {
			return a.fail(id, "frame.declaration", "zero value only for scalar or constant-nil []byte")
		}
	} else {
		var err error
		v, err = a.expr(rhs, state, parent)
		if err != nil {
			return err
		}
	}
	if parent {
		if v.kind != otherBytesOrigin && v.kind != seedTableOrigin && v.kind != testingFOrigin {
			return a.fail(id, "frame.declaration", "parent locals are finite seed values/tables or native F aliases")
		}
		if (v.kind == otherBytesOrigin && !byteSlice(obj.Type())) || (v.kind == seedTableOrigin && !seedTable(obj.Type())) || (v.kind == testingFOrigin && !testingPointer(obj.Type(), "F")) {
			return a.fail(id, "frame.declaration", "parent aliases retain exact []byte/[][]byte/native F types; interface boxing unsupported")
		}
	} else {
		_, signatureAlias := obj.Type().(*types.Signature)
		if !allowedValueType(obj.Type()) && !(v.kind == functionOrigin && signatureAlias) {
			return a.fail(id, "frame.declaration", "unsupported local type")
		}
		if v.kind == functionOrigin && a.isNamedSubjectAPI(v.function) {
			if a.subjectAPIName(v.function) == "" {
				return a.fail(rhs, "frame.api-signature", "subject Encode/Decode must have the exact finite byte API signature")
			}
			a.apiAliases[obj] = v.function
			switch expr := unparen(rhs).(type) {
			case *ast.SelectorExpr:
				if fn, ok := a.info.Uses[expr.Sel].(*types.Func); ok && fn == v.function {
					a.apiAliasInit[rhs] = fn
				}
			case *ast.Ident:
				if alias := a.info.Uses[expr]; alias == nil || a.apiAliases[alias] != v.function {
					return a.fail(rhs, "frame.api-alias", "API function aliases must derive from the exact imported function or an earlier checked alias")
				}
			default:
				return a.fail(rhs, "frame.api-alias", "unsupported subject API function alias expression")
			}
		}
	}
	state[obj] = binding{value: v, mutableScalar: !parent && v.kind == scalarOrigin && scalarType(obj.Type())}
	return nil
}

func (a *frame) block(body *ast.BlockStmt, state scopeState, parent bool) (flow, error) {
	out := flow{falls: true}
	for _, stmt := range body.List {
		part, err := a.statement(stmt, state, parent)
		if err != nil {
			return flow{}, err
		}
		// All syntax/call sites are checked, even following an earlier return.
		// Only reachable abstract returns contribute to a helper return join.
		if out.falls {
			out.returns = append(out.returns, part.returns...)
			out.falls = part.falls
		}
	}
	return out, nil
}

func (a *frame) statement(stmt ast.Stmt, state scopeState, parent bool) (flow, error) {
	if err := a.charge(stmt, 1); err != nil {
		return flow{}, err
	}
	out := flow{falls: true}
	switch s := stmt.(type) {
	case *ast.AssignStmt, *ast.DeclStmt:
		return out, a.local(stmt, state, parent)
	case *ast.EmptyStmt:
		return out, nil
	case *ast.BlockStmt:
		return a.block(s, cloneState(state), parent)
	case *ast.ExprStmt:
		if _, ok := s.X.(*ast.CallExpr); !ok {
			return flow{}, a.fail(s, "frame.statement", "only synchronous call expression statements")
		}
		_, err := a.expr(s.X, state, parent)
		return out, err
	case *ast.ReturnStmt:
		if len(s.Results) > 2 {
			return flow{}, a.fail(s, "frame.return", "zero/one expression or exact tuple return")
		}
		if len(s.Results) == 2 {
			left, err := a.expr(s.Results[0], state, parent)
			if err != nil {
				return flow{}, err
			}
			right, err := a.expr(s.Results[1], state, parent)
			if err != nil {
				return flow{}, err
			}
			if !byteFact(left) || right.kind != scalarOrigin || !booleanType(a.info.TypeOf(s.Results[1])) {
				return flow{}, a.fail(s, "frame.return", "tuple return requires readonly bytes and bool")
			}
			return flow{returns: []fact{{kind: tupleOrigin, parts: []fact{left, right}}}, falls: false}, nil
		}
		v := fact{kind: invalidOrigin}
		if len(s.Results) == 1 {
			var err error
			v, err = a.expr(s.Results[0], state, parent)
			if err != nil {
				return flow{}, err
			}
		}
		return flow{returns: []fact{v}, falls: false}, nil
	case *ast.IfStmt:
		sub := cloneState(state)
		if s.Init != nil {
			if _, err := a.statement(s.Init, sub, parent); err != nil {
				return flow{}, err
			}
		}
		v, err := a.expr(s.Cond, sub, parent)
		if err != nil {
			return flow{}, err
		}
		if v.kind != scalarOrigin {
			return flow{}, a.fail(s.Cond, "frame.expression", "scalar condition required")
		}
		yes, err := a.block(s.Body, cloneState(sub), parent)
		if err != nil {
			return flow{}, err
		}
		no := flow{falls: true}
		if s.Else != nil {
			no, err = a.statement(s.Else, cloneState(sub), parent)
			if err != nil {
				return flow{}, err
			}
		}
		return flow{returns: append(yes.returns, no.returns...), falls: yes.falls || no.falls}, nil
	case *ast.RangeStmt:
		if s.Key != nil && s.Tok != token.DEFINE {
			return flow{}, a.fail(s, "frame.write", "range bindings must be fresh")
		}
		v, err := a.expr(s.X, state, parent)
		if err != nil {
			return flow{}, err
		}
		if !isSeed(v) && !isResult(v) {
			return flow{}, a.fail(s, "frame.range", "callback/helper range over read-only seed/result []byte only")
		}
		if !byteSlice(a.info.TypeOf(s.X)) {
			return flow{}, a.fail(s, "frame.range", "only []byte range supported")
		}
		sub := cloneState(state)
		for _, expr := range []ast.Expr{s.Key, s.Value} {
			if expr == nil {
				continue
			}
			id, ok := expr.(*ast.Ident)
			if !ok {
				return flow{}, a.fail(expr, "frame.write", "fresh named range bindings required")
			}
			if id.Name != "_" {
				sub[a.info.Defs[id]] = binding{value: fact{kind: scalarOrigin}}
			}
		}
		part, err := a.block(s.Body, sub, parent)
		if err != nil {
			return flow{}, err
		}
		out.returns = part.returns
		return out, nil
	default:
		return flow{}, a.fail(stmt, "frame.statement", "unsupported statement %T", stmt)
	}
}

func (a *frame) helper(obj *types.Func, args []fact, site ast.Node) (fact, error) {
	if err := a.charge(site, 1); err != nil {
		return fact{}, err
	}
	fn := a.functions[obj]
	if fn == nil || fn == a.parent {
		return fact{}, a.fail(site, "frame.call", "submitted synchronous helper required")
	}
	if a.stack[obj] {
		return fact{}, a.fail(site, "frame.recursion", "helper graph must be acyclic")
	}
	sig := obj.Type().(*types.Signature)
	if len(args) != sig.Params().Len() {
		return fact{}, a.fail(site, "frame.call", "helper argument count mismatch")
	}
	state := scopeState{}
	for i, arg := range args {
		t := sig.Params().At(i).Type()
		if testingPointer(t, "T") && arg.kind != testingTOrigin {
			return fact{}, a.fail(site, "frame.testing-receiver", "helper requires native callback T origin")
		}
		if scalarType(t) && arg.kind != scalarOrigin {
			return fact{}, a.fail(site, "frame.call", "scalar argument required")
		}
		if byteSlice(t) && !byteFact(arg) {
			return fact{}, a.fail(site, "frame.call", "read-only []byte argument required")
		}
		state[sig.Params().At(i)] = binding{value: arg}
	}
	for i := 0; i < sig.Results().Len(); i++ {
		res := sig.Results().At(i)
		if res.Name() != "" {
			v := fact{kind: scalarOrigin}
			if byteSlice(res.Type()) {
				v = fact{kind: otherBytesOrigin, knownBytes: true}
			}
			state[res] = binding{value: v, mutableScalar: scalarType(res.Type())}
		}
	}
	a.stack[obj] = true
	part, err := a.block(fn.Body, state, false)
	delete(a.stack, obj)
	if err != nil {
		return fact{}, err
	}
	if sig.Results().Len() == 0 {
		return fact{kind: invalidOrigin}, nil
	}
	if sig.Results().Len() == 2 {
		for _, v := range part.returns {
			if v.kind != tupleOrigin || len(v.parts) != 2 || !byteFact(v.parts[0]) || v.parts[1].kind != scalarOrigin {
				return fact{}, a.fail(fn, "frame.return", "explicit exact tuple return required")
			}
		}
		return joinFacts(part.returns), nil
	}
	// A bare return in a named-result helper uses that tracked result's origin.
	for i, v := range part.returns {
		if v.kind == invalidOrigin {
			res := sig.Results().At(0)
			if b, ok := state[res]; ok {
				part.returns[i] = b.value
			} else {
				return fact{}, a.fail(fn, "frame.return", "missing helper return origin")
			}
		}
	}
	return joinFacts(part.returns), nil
}

func byteFact(v fact) bool { return isSeed(v) || isResult(v) || v.kind == otherBytesOrigin }
