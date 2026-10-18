package main

import (
	"bytes"
	"testing"
)

func TestExplicitScopeAndChecksRequiredBeforeRepositoryAccess(t *testing.T) {
	for _, args := range [][]string{nil, {"-scope", "full"}, {"-checks", "missing"}, {"-scope", "specialist", "-checks", "missing"}, {"-scope", "foundation", "-checks", "missing", "positional"}} {
		var out, diagnostic bytes.Buffer
		if code := run(append([]string{"-root", t.TempDir()}, args...), &out, &diagnostic); code != 2 || out.Len() != 0 || diagnostic.Len() == 0 {
			t.Fatalf("%v: exit%d stdout%q stderr%q", args, code, out.String(), diagnostic.String())
		}
	}
}
