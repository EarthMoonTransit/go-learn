package taxonomy

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"sync"
	"testing"
)

const smallDocument = "version: 3\ntrack: go\ntarget_ready_tasks: {min: 350, max: 500}\ncoverage_policy: {counted_status: ready}\ndomains:\n- id: basics\n  title: Basics\n  skills:\n  - id: basics.value\n    definition: original\n    prerequisites: [basics.other]\n"

func writeDocument(t *testing.T, root, relative string, content []byte) {
	t.Helper()
	path := filepath.Join(root, filepath.FromSlash(relative))
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, content, 0600); err != nil {
		t.Fatal(err)
	}
}

func TestImmutableDocumentObservesChangedBytesAndFreshDecodeErrors(t *testing.T) {
	for _, path := range []string{V3Path, "curriculum/skills.yaml"} {
		t.Run(path, func(t *testing.T) {
			root := t.TempDir()
			original := []byte(smallDocument)
			writeDocument(t, root, path, original)
			_, want, err := readImmutableDocument(root, path)
			if err != nil {
				t.Fatal(err)
			}
			for _, tc := range []struct{ name, content, diagnostic string }{
				{"same length semantic change", strings.Replace(smallDocument, "version: 3", "version: 5", 1), ""},
				{"malformed", "version: [\n", "taxonomy: decode "},
				{"unknown field", smallDocument + "future: true\n", "field future not found"},
				{"trailing document", smallDocument + "---\n{}\n", "taxonomy: trailing data in "},
			} {
				t.Run(tc.name, func(t *testing.T) {
					writeDocument(t, root, path, []byte(tc.content))
					var fresh document
					freshBytes, freshErr := readStrict(root, path, &fresh)
					actualBytes, actual, actualErr := readImmutableDocument(root, path)
					if fmt.Sprint(actualErr) != fmt.Sprint(freshErr) {
						t.Fatalf("cached read error = %v, fresh decode = %v", actualErr, freshErr)
					}
					if tc.diagnostic != "" {
						if actualErr == nil || !strings.Contains(actualErr.Error(), tc.diagnostic) {
							t.Fatalf("got %v, want %q", actualErr, tc.diagnostic)
						}
					} else if actualErr != nil || actual.Version != 5 || !reflect.DeepEqual(actual, fresh) || !bytes.Equal(actualBytes, freshBytes) {
						t.Fatalf("changed bytes did not produce the fresh document: v%d, %v", actual.Version, actualErr)
					}
					writeDocument(t, root, path, original)
					b, restored, err := readImmutableDocument(root, path)
					if err != nil || !bytes.Equal(b, original) || !reflect.DeepEqual(restored, want) {
						t.Fatalf("valid bytes not restored after error/change: %v", err)
					}
				})
			}
		})
	}
}

func TestImmutableDocumentChecksFilesystemBeforeContentHit(t *testing.T) {
	for _, path := range []string{V3Path, "curriculum/skills.yaml"} {
		for _, kind := range []string{"missing", "symlink", "ancestor symlink", "directory", "oversize"} {
			t.Run(path+"/"+kind, func(t *testing.T) {
				warm, other := t.TempDir(), t.TempDir()
				writeDocument(t, warm, path, []byte(smallDocument))
				writeDocument(t, other, path, []byte(smallDocument))
				if _, _, err := readImmutableDocument(warm, path); err != nil {
					t.Fatal(err)
				}
				file := filepath.Join(other, filepath.FromSlash(path))
				if kind == "ancestor symlink" {
					ancestor := filepath.Join(other, "curriculum")
					if err := os.Rename(ancestor, ancestor+"-real"); err != nil {
						t.Fatal(err)
					}
					if err := os.Symlink(ancestor+"-real", ancestor); err != nil {
						t.Fatal(err)
					}
				} else {
					if err := os.Remove(file); err != nil {
						t.Fatal(err)
					}
					switch kind {
					case "symlink":
						if err := os.Symlink(filepath.Join(warm, filepath.FromSlash(path)), file); err != nil {
							t.Fatal(err)
						}
					case "directory":
						if err := os.Mkdir(file, 0700); err != nil {
							t.Fatal(err)
						}
					case "oversize":
						writeDocument(t, other, path, bytes.Repeat([]byte("x"), (4<<20)+1))
					}
				}
				var fresh document
				_, freshErr := readStrict(other, path, &fresh)
				_, _, err := readImmutableDocument(other, path)
				if err == nil || !strings.HasPrefix(err.Error(), "taxonomy: read ") || fmt.Sprint(err) != fmt.Sprint(freshErr) {
					t.Fatalf("warm content bypassed filesystem guard: got %v, fresh read %v", err, freshErr)
				}
			})
		}
	}
}

func TestMutableReadsAndReturnedBytesCannotPoisonImmutableDocuments(t *testing.T) {
	root := t.TempDir()
	for _, path := range []string{V3Path, "curriculum/skills.yaml"} {
		writeDocument(t, root, path, []byte(smallDocument))
		b, _, err := readImmutableDocument(root, path)
		if err != nil {
			t.Fatal(err)
		}
		var want document
		if _, err := readStrict(root, path, &want); err != nil {
			t.Fatal(err)
		}
		for i := range b {
			b[i] = 'x'
		}
		var mutable document
		if _, err := readStrict(root, path, &mutable); err != nil {
			t.Fatal(err)
		}
		mutable.Domains[0].Skills[0]["definition"] = "poisoned"
		mutable.Domains[0].Skills[0]["prerequisites"].([]any)[0] = "poisoned"
		mutable.Target.Content[0].Value = "poisoned"
		mutable.Policy.Content[0].Value = "poisoned"
		actualBytes, actual, err := readImmutableDocument(root, path)
		if err != nil || !bytes.Equal(actualBytes, []byte(smallDocument)) || !reflect.DeepEqual(actual, want) || actual.Domains[0].Skills[0]["definition"] != "original" {
			t.Fatalf("mutable data poisoned immutable parse: %v", err)
		}
	}
}

func TestConcurrentImmutableDocumentReplacementPreservesReaders(t *testing.T) {
	type input struct {
		root, path string
		want       document
	}
	var inputs []input
	for i := range 4 {
		root := t.TempDir()
		for _, path := range []string{V3Path, "curriculum/skills.yaml"} {
			// Two roots share content; the other roots force role-slot replacement.
			content := strings.Replace(smallDocument, "definition: original", fmt.Sprintf("definition: value%d", i/2), 1)
			writeDocument(t, root, path, []byte(content))
			var want document
			if _, err := readStrict(root, path, &want); err != nil {
				t.Fatal(err)
			}
			inputs = append(inputs, input{root, path, want})
		}
	}
	var wg sync.WaitGroup
	for _, in := range inputs {
		wg.Go(func() {
			for range 16 {
				_, got, err := readImmutableDocument(in.root, in.path)
				if err != nil {
					t.Errorf("concurrent read: %v", err)
					return
				}
				// These production readers must only read shared maps and YAML nodes.
				all, err := skills(got)
				wantSkills, _ := skills(in.want)
				if err != nil || contractHash(all["basics.value"]) != contractHash(wantSkills["basics.value"]) || !nodeEqual(got.Target, in.want.Target) || !policyEqual(got.Policy, in.want.Policy) || !reflect.DeepEqual(got, in.want) {
					t.Errorf("published document changed during replacement: %v", err)
					return
				}
			}
		})
	}
	wg.Wait()
}

func TestWarmTransitionStillDecodesBaselineBeforeHashCheck(t *testing.T) {
	root := fixture(t)
	if _, err := ValidateTransition(root); err != nil {
		t.Fatal(err)
	}
	writeDocument(t, root, V3Path, []byte("version: [\n"))
	if _, err := ValidateTransition(root); err == nil || !strings.HasPrefix(err.Error(), "taxonomy: decode "+V3Path) {
		t.Fatalf("baseline error order changed: %v", err)
	}
}

func TestWarmReviewedSelectionReadsSnapshotBytesAgain(t *testing.T) {
	root := fixture(t)
	if _, _, err := ReviewedPath(root); err != nil {
		t.Fatal(err)
	}
	for _, pin := range reviewedV3 {
		path := filepath.Join(root, filepath.FromSlash(pin.snapshot))
		original, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		writeDocument(t, root, pin.snapshot, append(append([]byte(nil), original...), '\n'))
		if _, _, err := ReviewedPath(root); err == nil || !strings.Contains(err.Error(), "reviewed v3 bytes changed") {
			t.Fatalf("warm reviewed selection accepted changed snapshot %s: %v", pin.snapshot, err)
		}
		if _, err := ValidateLiveTransition(root); err != nil {
			t.Fatalf("live validation acquired private snapshot dependency: %v", err)
		}
		writeDocument(t, root, pin.snapshot, original)
		if _, _, err := ReviewedPath(root); err != nil {
			t.Fatalf("restored snapshot rejected: %v", err)
		}
	}
}

func TestWarmTransitionPublicResultsAreIndependent(t *testing.T) {
	root := fixture(t)
	first, err := ValidateTransition(root)
	if err != nil {
		t.Fatal(err)
	}
	want, err := ValidateTransition(root)
	if err != nil {
		t.Fatal(err)
	}
	first.Entries[0].ContractSHA256 = "poisoned"
	first.Additions[0].ID = "poisoned"
	first.ReviewedDocuments[0].SkillIDs[0] = "poisoned"
	got, err := ValidateTransition(root)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("public transition result poisoned subsequent validation: %v", err)
	}
	var wg sync.WaitGroup
	for range 4 {
		wg.Go(func() {
			for range 3 {
				got, err := ValidateLiveTransition(root)
				if err != nil || !reflect.DeepEqual(got, want) {
					t.Errorf("concurrent semantic validation: %v", err)
				}
			}
		})
	}
	wg.Wait()
}
