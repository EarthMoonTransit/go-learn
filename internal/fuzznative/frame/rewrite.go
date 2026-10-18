package frame

import (
	"bytes"
	"fmt"
	"go/ast"
	"go/format"
	"go/token"
	"go/types"
	"strconv"
)

type RewriteResult struct {
	Sources     map[string]string `json:"sources"`
	Methods     map[string]int    `json:"methods"`
	ImportAlias string            `json:"import_alias"`
}
type plannedCall struct {
	call     *ast.CallExpr
	receiver ast.Expr
	method   string
	file     *ast.File
}
type plannedAPICall struct {
	call     *ast.CallExpr
	function ast.Expr
	method   string
	file     *ast.File
}
type plannedAPIAlias struct {
	expr   ast.Expr
	method string
	file   *ast.File
}

// Rewrite consumes this typed inspection exactly once. Admission was checked
// against original source; all callee/receiver decisions below use those objects.
func (p Inspection) Rewrite(observerImport string) (RewriteResult, error) {
	a := p.engine
	if a == nil || a.rewritten {
		return RewriteResult{}, fmt.Errorf("typed inspection missing or already consumed")
	}
	if observerImport != a.identity.ObserverPath() {
		return RewriteResult{}, fmt.Errorf("unexpected synthetic observer import")
	}
	names := map[string]bool{}
	for _, file := range a.files {
		ast.Inspect(file, func(node ast.Node) bool {
			if id, ok := node.(*ast.Ident); ok {
				names[id.Name] = true
			}
			return true
		})
	}
	alias := ""
	for i := 0; ; i++ {
		candidate := fmt.Sprintf("__native_fuzz_observer%d", i)
		if !names[candidate] {
			alias = candidate
			break
		}
	}
	var plans []plannedCall
	var apiPlans []plannedAPICall
	var aliasPlans []plannedAPIAlias
	var planErr error
	for _, file := range a.files {
		ast.Inspect(file, func(node ast.Node) bool {
			if planErr != nil {
				return false
			}
			call, ok := node.(*ast.CallExpr)
			if !ok {
				return true
			}
			if fn := a.apiCalls[call]; fn != nil {
				method := a.subjectAPIName(fn)
				if method == "" {
					planErr = fmt.Errorf("admitted subject API lost exact signature")
					return false
				}
				aliased := false
				if id, ok := unparen(call.Fun).(*ast.Ident); ok {
					aliased = a.apiAliases[a.info.Uses[id]] == fn
				}
				if !aliased {
					apiPlans = append(apiPlans, plannedAPICall{call, call.Fun, method, file})
				}
			}
			sel, ok := unparen(call.Fun).(*ast.SelectorExpr)
			if !ok {
				return true
			}
			selection := a.info.Selections[sel]
			if selection == nil || selection.Kind() != types.MethodVal {
				return true
			}
			obj := selection.Obj()
			if obj.Pkg() == nil || obj.Pkg().Path() != "testing" {
				return true
			}
			method := obj.Name()
			receiverType := a.info.TypeOf(sel.X)
			if testingPointer(receiverType, "F") {
				switch method {
				case "Add":
					if len(call.Args) != 1 || call.Ellipsis != token.NoPos || !byteSlice(a.info.TypeOf(call.Args[0])) {
						planErr = fmt.Errorf("admitted F.Add argument lacks original []byte type")
						return false
					}
					method = "Add"
				case "Fuzz":
					method = "Register"
				default:
					planErr = fmt.Errorf("admitted unsupported native F method %s", method)
					return false
				}
			} else if testingPointer(receiverType, "T") {
				switch method {
				case "Helper":
					return true
				case "Error", "Errorf", "Fatal", "Fatalf":
				default:
					planErr = fmt.Errorf("admitted unsupported native T method %s", method)
					return false
				}
			} else {
				planErr = fmt.Errorf("admitted unexpected testing receiver")
				return false
			}
			plans = append(plans, plannedCall{call, sel.X, method, file})
			return true
		})
	}
	for expr, fn := range a.apiAliasInit {
		method := a.subjectAPIName(fn)
		if method == "" {
			return RewriteResult{}, fmt.Errorf("admitted API alias lost exact signature")
		}
		for _, file := range a.files {
			if expr.Pos() >= file.Pos() && expr.End() <= file.End() {
				aliasPlans = append(aliasPlans, plannedAPIAlias{expr, method, file})
				break
			}
		}
	}
	if planErr != nil {
		return RewriteResult{}, planErr
	}
	a.rewritten = true
	changed := map[*ast.File]bool{}
	result := RewriteResult{Sources: map[string]string{}, Methods: map[string]int{}, ImportAlias: alias}
	for _, plan := range plans {
		plan.call.Fun = &ast.SelectorExpr{X: ast.NewIdent(alias), Sel: ast.NewIdent(plan.method)}
		plan.call.Args = append([]ast.Expr{plan.receiver}, plan.call.Args...)
		changed[plan.file] = true
		result.Methods[plan.method]++
	}
	for _, plan := range apiPlans {
		plan.call.Fun = &ast.SelectorExpr{X: ast.NewIdent(alias), Sel: ast.NewIdent(plan.method)}
		plan.call.Args = append([]ast.Expr{plan.function}, plan.call.Args...)
		changed[plan.file] = true
	}
	for _, plan := range aliasPlans {
		wrapped := &ast.CallExpr{Fun: &ast.SelectorExpr{X: ast.NewIdent(alias), Sel: ast.NewIdent("Wrap" + plan.method)}, Args: []ast.Expr{plan.expr}}
		replaced := false
		// Replace the original expression in its containing AST node. The
		// package is small and exact: all aliases are single-value local RHSs.
		for _, file := range a.files {
			if file == plan.file {
				ast.Inspect(file, func(node ast.Node) bool {
					switch parent := node.(type) {
					case *ast.AssignStmt:
						for i, rhs := range parent.Rhs {
							if rhs.Pos() == plan.expr.Pos() && rhs.End() == plan.expr.End() {
								parent.Rhs[i] = wrapped
								replaced = true
								return false
							}
						}
					case *ast.ValueSpec:
						for i, rhs := range parent.Values {
							if rhs.Pos() == plan.expr.Pos() && rhs.End() == plan.expr.End() {
								parent.Values[i] = wrapped
								replaced = true
								return false
							}
						}
					}
					return true
				})
			}
		}
		if !replaced {
			return RewriteResult{}, fmt.Errorf("API alias initializer was not found in its original binding")
		}
		changed[plan.file] = true
		result.Methods["Wrap"+plan.method]++
	}
	for _, fn := range a.apiCalls {
		result.Methods[a.subjectAPIName(fn)]++
	}
	for _, file := range a.files {
		if changed[file] {
			spec := &ast.ImportSpec{Name: ast.NewIdent(alias), Path: &ast.BasicLit{Kind: token.STRING, Value: strconv.Quote(observerImport)}}
			decl := &ast.GenDecl{Tok: token.IMPORT, Specs: []ast.Spec{spec}}
			file.Decls = append([]ast.Decl{decl}, file.Decls...)
			file.Imports = append(file.Imports, spec)
		}
		var source bytes.Buffer
		if err := format.Node(&source, a.fset, file); err != nil {
			return RewriteResult{}, fmt.Errorf("format instrumented source: %w", err)
		}
		name := a.fset.File(file.Pos()).Name()
		result.Sources[name] = source.String()
	}
	return result, nil
}
