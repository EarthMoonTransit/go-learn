package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"go-practice-bank/internal/isolation"
)

func TestFuzzLayoutRequiresExactCorrectVariantsAndPublicContract(t *testing.T) {
	for _, test := range []struct {
		name, want string
		mutate     func(*testing.T, string, *packageData)
	}{
		{"complete", "", func(*testing.T, string, *packageData) {}},
		{"missing-correct-variant", "exactly match", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			if err := os.RemoveAll(filepath.Join(d.Located.Path, "verification/correct/loop")); err != nil {
				t.Fatal(err)
			}
		}},
		{"extra-correct-variant", "exactly match", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			if err := isolation.CopyTree(filepath.Join(d.Located.Path, "verification/correct/loop"), filepath.Join(d.Located.Path, "verification/correct/extra")); err != nil {
				t.Fatal(err)
			}
		}},
		{"missing-public-contract", "complete finite source contract", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			if err := os.Remove(filepath.Join(d.Located.Path, filepath.FromSlash(fuzzSourceContractPath))); err != nil {
				t.Fatal(err)
			}
		}},
		{"stale-public-contract", "complete finite source contract", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			path := filepath.Join(d.Located.Path, filepath.FromSlash(fuzzSourceContractPath))
			content, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, append(content, '\n'), 0600); err != nil {
				t.Fatal(err)
			}
		}},
		{"unlinked-public-contract", "must link", func(_ *testing.T, _ string, d *packageData) {
			d.Prompt = []byte(strings.ReplaceAll(string(d.Prompt), "](FUZZ_SOURCE.md)", "](missing.md)"))
		}},
		{"changed-correct-api", "correct variant loop", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			path := filepath.Join(d.Located.Path, "verification/correct/loop/hex.go")
			body, err := os.ReadFile(path)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(path, []byte(strings.ReplaceAll(string(body), "func Decode(", "func DecodeOther(")), 0600); err != nil {
				t.Fatal(err)
			}
		}},
		{"symlink-correct-file", "symlink", func(t *testing.T, _ string, d *packageData) {
			t.Helper()
			path := filepath.Join(d.Located.Path, "verification/correct/loop/hex.go")
			outside := filepath.Join(t.TempDir(), "outside.go")
			if err := os.WriteFile(outside, []byte("package outside"), 0600); err != nil {
				t.Fatal(err)
			}
			if err := os.Remove(path); err != nil {
				t.Fatal(err)
			}
			if err := os.Symlink(outside, path); err != nil {
				t.Skip(err)
			}
		}},
	} {
		t.Run(test.name, func(t *testing.T) {
			root, data := fuzzFixtureRepository(t)
			test.mutate(t, root, &data)
			problems, layout := checkPackageContractV2(root, data)
			problems = append(problems, checkPublicAPIV2(data, layout)...)
			if test.want == "" {
				if len(problems) > 0 {
					t.Fatal(problems)
				}
				if len(layout.VerificationCorrect) != 2 || len(layout.Mutants) != 2 {
					t.Fatalf("missing finite variants: %+v", layout)
				}
			} else if !containsText(problems, test.want) {
				t.Fatalf("want %q diagnostic, got %v", test.want, problems)
			}
		})
	}
}
