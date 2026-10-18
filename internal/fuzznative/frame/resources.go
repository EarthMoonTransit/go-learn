package frame

import (
	"go/ast"
	"unicode/utf8"
)

const MaxFiles = 8
const MaxSourceBytes = 256 << 10
const MaxSteps = 100000
const MaxLiteralLength = 4096

func preflight(sources map[string]string) error {
	if len(sources) == 0 || len(sources) > MaxFiles {
		return &frameError{"frame.resources", "package", "require 1..8 source files"}
	}
	size := 0
	for _, s := range sources {
		size += len(s)
		if size > MaxSourceBytes {
			return &frameError{"frame.resources", "package", "combined source exceeds 256 KiB"}
		}
		if !utf8.ValidString(s) {
			return &frameError{"frame.resources", "package", "source must be UTF-8"}
		}
	}
	return nil
}
func (a *frame) charge(site ast.Node, n int) error {
	if a.exhausted != nil {
		return a.exhausted
	}
	if n < 0 || n > MaxSteps-a.steps {
		a.steps = MaxSteps + 1
		a.exhausted = a.fail(site, "frame.budget", "analysis step limit 100000 exceeded")
		return a.exhausted
	}
	a.steps += n
	return nil
}

// literalLength is independent of whether byte values are statically known.
// It charges repeated element and gap processing before any slice allocation.
func (a *frame) literalLength(lit *ast.CompositeLit) (int, error) {
	next, length := 0, 0
	for _, item := range lit.Elts {
		if err := a.charge(item, 1); err != nil {
			return 0, err
		}
		index := next
		if kv, ok := item.(*ast.KeyValueExpr); ok {
			n, known := a.constantInt(kv.Key)
			if !known || n < 0 || n >= MaxLiteralLength {
				return 0, a.fail(kv.Key, "frame.resources", "literal keyed index must be 0..4095")
			}
			index = int(n)
		}
		if index >= MaxLiteralLength {
			return 0, a.fail(item, "frame.resources", "expanded literal length exceeds 4096")
		}
		if index >= length {
			if err := a.charge(item, index+1-length); err != nil {
				return 0, err
			}
			length = index + 1
		}
		next = index + 1
	}
	return length, nil
}
func (a *frame) stringBytes(site ast.Node, s string) ([]byte, bool) {
	if len(s) > MaxLiteralLength {
		a.exhausted = a.fail(site, "frame.resources", "byte string literal length exceeds 4096")
		return nil, false
	}
	if a.charge(site, len(s)) != nil {
		return nil, false
	}
	return []byte(s), true
}
