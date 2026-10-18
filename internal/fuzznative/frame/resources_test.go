package frame_test

import (
	"fmt"
	"go/token"
	"strings"
	"testing"

	"go-practice-bank/internal/fuzznative/frame"
)

func TestFiniteAdmissionResourceBoundaries(t *testing.T) {
	identity := frame.Identity{ModulePath: "task.local/typed-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	base := source(identity, completeBody, "")
	seedBlock := `f.Add([]byte(nil));f.Add([]byte{0,128});`
	for _, test := range []struct{ name, source, want string }{
		{"64-seeds", strings.Replace(base, seedBlock, strings.Repeat(`f.Add([]byte(nil));`, 64), 1), ""},
		{"65-seeds", strings.Replace(base, seedBlock, strings.Repeat(`f.Add([]byte(nil));`, 65), 1), "frame.seeds"},
		{"no-seeds", strings.Replace(base, seedBlock, "", 1), "frame.seeds"},
		{"4096-byte-keyed-seed", strings.Replace(base, seedBlock, `f.Add([]byte{4095:0});`, 1), ""},
		{"4097-byte-keyed-seed", strings.Replace(base, seedBlock, `f.Add([]byte{4096:0});`, 1), "frame.resources"},
		{"huge-keyed-seed", strings.Replace(base, seedBlock, `f.Add([]byte{1<<28:0});`, 1), "frame.resources"},
		{"4096-byte-string", strings.Replace(base, seedBlock, `f.Add([]byte("`+strings.Repeat("a", 4096)+`"));`, 1), ""},
		{"4097-byte-string", strings.Replace(base, seedBlock, `f.Add([]byte("`+strings.Repeat("a", 4097)+`"));`, 1), "frame.resources"},
	} {
		t.Run(test.name, func(t *testing.T) {
			_, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{"fuzz_test.go": test.source}, identity)
			if test.want == "" {
				if err != nil {
					t.Fatal(err)
				}
			} else if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("want %s, got %v", test.want, err)
			}
		})
	}
	for _, test := range []struct {
		name    string
		sources map[string]string
	}{
		{"empty", nil},
		{"over-total-bytes", map[string]string{"fuzz_test.go": base + strings.Repeat(" ", 256<<10)}},
		{"not-UTF8", map[string]string{"fuzz_test.go": base + string([]byte{255})}},
		{"nine-files", func() map[string]string {
			out := map[string]string{"fuzz_test.go": base}
			for i := range 8 {
				out[fmt.Sprintf("extra%d_test.go", i)] = "package subject_test"
			}
			return out
		}()},
	} {
		t.Run(test.name, func(t *testing.T) {
			if _, err := frame.Analyze(token.NewFileSet(), testImports(identity), test.sources, identity); err == nil || !strings.Contains(err.Error(), "frame.resources") {
				t.Fatalf("resource overflow lacked public diagnostic: %v", err)
			}
		})
	}
}

func TestSharedAnalysisBudgetBoundsExpandedHelperGraph(t *testing.T) {
	identity := frame.Identity{ModulePath: "task.local/typed-hex-001", Grammar: "hex-codec-v1", Target: "FuzzHex"}
	graph := func(count int) string {
		var helpers strings.Builder
		for i := range count {
			fmt.Fprintf(&helpers, "func h%02d(input []byte)[]byte{", i)
			if i == count-1 {
				helpers.WriteString("return input")
			} else {
				fmt.Fprintf(&helpers, "h%02d(input);return h%02d(input)", i+1, i+1)
			}
			helpers.WriteString("}\n")
		}
		return source(identity, `h00(input);`+completeBody, helpers.String())
	}
	if _, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{"fuzz_test.go": graph(4)}, identity); err != nil {
		t.Fatalf("small acyclic helper graph rejected: %v", err)
	}
	if _, err := frame.Analyze(token.NewFileSet(), testImports(identity), map[string]string{"fuzz_test.go": graph(15)}, identity); err == nil || !strings.Contains(err.Error(), "frame.budget") {
		t.Fatalf("expanded acyclic graph escaped shared analysis budget: %v", err)
	}
}
