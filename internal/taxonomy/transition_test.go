package taxonomy

import (
	"os"
	"path/filepath"
	"strings"
	"testing"

	"gopkg.in/yaml.v3"
)

func fixture(t *testing.T) string {
	t.Helper()
	root := t.TempDir()
	source := filepath.Join("..", "..")
	paths := []string{"curriculum/skills.yaml", MigrationPath, V3Path}
	for path, pin := range reviewedV3 {
		paths = append(paths, path, pin.snapshot)
	}
	for _, path := range paths {
		b, err := os.ReadFile(filepath.Join(source, filepath.FromSlash(path)))
		if err != nil {
			t.Fatal(err)
		}
		dst := filepath.Join(root, filepath.FromSlash(path))
		if err := os.MkdirAll(filepath.Dir(dst), 0700); err != nil {
			t.Fatal(err)
		}
		if err := os.WriteFile(dst, b, 0600); err != nil {
			t.Fatal(err)
		}
	}
	return root
}

func rewrite(t *testing.T, root, path string, dst any, mutate func()) {
	t.Helper()
	if _, err := readStrict(root, path, dst); err != nil {
		t.Fatal(err)
	}
	mutate()
	b, err := yaml.Marshal(dst)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, filepath.FromSlash(path)), b, 0600); err != nil {
		t.Fatal(err)
	}
}

func TestActualAdditiveTransitionAndReviewedSelection(t *testing.T) {
	root := fixture(t)
	transition, err := ValidateTransition(root)
	if err != nil {
		t.Fatal(err)
	}
	if len(transition.Entries) != 177 || len(transition.Additions) != 3 || len(transition.ReviewedDocuments) != 4 {
		t.Fatal("incomplete transition inventory")
	}
	path, version, err := ReviewedPath(root)
	if err != nil || path != V3Path || version != 3 {
		t.Fatalf("reviewed selection = %q v%d, %v", path, version, err)
	}
}

func TestRejectsChangedHistoryAndUnreviewedCompatibility(t *testing.T) {
	cases := []struct {
		name, want string
		change     func(*testing.T, string)
	}{
		{"old teaching contract", "old contract changed", func(t *testing.T, r string) {
			var d document
			rewrite(t, r, "curriculum/skills.yaml", &d, func() { d.Domains[0].Skills[0]["definition"] = "silently changed old teaching outcome" })
		}},
		{"deprecated assignment", "old contract changed", func(t *testing.T, r string) {
			var d document
			rewrite(t, r, "curriculum/skills.yaml", &d, func() {
				for _, dom := range d.Domains {
					for _, s := range dom.Skills {
						if s["status"] == "deprecated" {
							s["assignable"] = true
							return
						}
					}
				}
			})
		}},
		{"missing old id", "complete177-old/180-live", func(t *testing.T, r string) {
			var d document
			rewrite(t, r, "curriculum/skills.yaml", &d, func() { d.Domains[0].Skills = d.Domains[0].Skills[1:] })
		}},
		{"duplicate inventory entry", "incorrect preserved entry", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.Entries[1] = d.Entries[0] })
		}},
		{"missing deprecated migration entry", "complete177-old/180-live", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.Entries = d.Entries[:176] })
		}},
		{"new unapproved id", "invalid approved addition", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.Additions[0].ID = "expressions.unreviewed" })
		}},
		{"changed task target", "task target mismatch", func(t *testing.T, r string) {
			var d document
			rewrite(t, r, "curriculum/skills.yaml", &d, func() {
				var n yaml.Node
				if err := yaml.Unmarshal([]byte("min: 100\nmax: 500\n"), &n); err != nil {
					t.Fatal(err)
				}
				d.Target = *n.Content[0]
			})
		}},
		{"tampered baseline pin", "baseline identity", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.Baseline.TaxonomySHA256 = strings.Repeat("0", 64) })
		}},
		{"tampered baseline bytes", "taxonomy hash/version", func(t *testing.T, r string) {
			p := filepath.Join(r, filepath.FromSlash(V3Path))
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(p, append(b, '\n'), 0600); err != nil {
				t.Fatal(err)
			}
		}},
		{"missing review record", "exactly4 reviewed", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.ReviewedDocuments = d.ReviewedDocuments[:3] })
		}},
		{"review inventory omitted", "reviewed skill inventory mismatch", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.ReviewedDocuments[0].SkillIDs = nil })
		}},
		{"review bytes repinned", "reviewed document identity", func(t *testing.T, r string) {
			var d Transition
			rewrite(t, r, MigrationPath, &d, func() { d.ReviewedDocuments[0].SHA256 = strings.Repeat("0", 64) })
		}},
	}
	for _, tc := range cases {
		t.Run(tc.name, func(t *testing.T) {
			root := fixture(t)
			tc.change(t, root)
			_, err := ValidateTransition(root)
			if err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("got %v, want diagnostic containing %q", err, tc.want)
			}
		})
	}
	for path := range reviewedV3 {
		t.Run("reviewed bytes "+path, func(t *testing.T) {
			root := fixture(t)
			p := filepath.Join(root, filepath.FromSlash(path))
			b, err := os.ReadFile(p)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(p, append(b, '\n'), 0600); err != nil {
				t.Fatal(err)
			}
			if _, _, err := ReviewedPath(root); err == nil || !strings.Contains(err.Error(), "reviewed v3 bytes changed") {
				t.Fatalf("got %v", err)
			}
		})
	}
}

func TestV3NeedsNoV4FilesAndUnknownVersionFails(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "curriculum"), 0700); err != nil {
		t.Fatal(err)
	}
	for _, version := range []string{"3", "5"} {
		t.Run(version, func(t *testing.T) {
			if err := os.WriteFile(filepath.Join(root, "curriculum", "skills.yaml"), []byte("version: "+version+"\n"), 0600); err != nil {
				t.Fatal(err)
			}
			p, v, err := ReviewedPath(root)
			if version == "3" {
				if err != nil || p != "curriculum/skills.yaml" || v != 3 {
					t.Fatalf("got %q v%d %v", p, v, err)
				}
			} else if err == nil || !strings.Contains(err.Error(), "unsupported version") {
				t.Fatalf("got %v", err)
			}
		})
	}
}

func TestStandaloneTransitionPreservesCoveragePolicy(t *testing.T) {
	for key, value := range map[string]string{"counted_status": "draft", "primary_skill_per_task": "2", "prerequisite_counts_as_occurrence": "true"} {
		t.Run(key, func(t *testing.T) {
			root := fixture(t)
			var d document
			rewrite(t, root, "curriculum/skills.yaml", &d, func() {
				for i := 0; i < len(d.Policy.Content); i += 2 {
					if d.Policy.Content[i].Value == key {
						d.Policy.Content[i+1].Value = value
						return
					}
				}
				t.Fatalf("missing policy field %s", key)
			})
			if _, err := ValidateTransition(root); err == nil || !strings.Contains(err.Error(), "coverage policy mismatch") {
				t.Fatalf("standalone transition accepted changed %s: %v", key, err)
			}
			if _, _, err := ReviewedPath(root); err == nil {
				t.Fatalf("review loader accepted changed %s", key)
			}
		})
	}
}

func TestRepinningAdditionCannotChangeApprovedTeachingContract(t *testing.T) {
	root := fixture(t)
	var d document
	changed := ""
	rewrite(t, root, "curriculum/skills.yaml", &d, func() {
		for _, dom := range d.Domains {
			for _, s := range dom.Skills {
				if s["id"] == "functions.recursion" {
					s["definition"] = "A materially different teaching contract"
					changed = contractHash(s)
					return
				}
			}
		}
	})
	var transition Transition
	rewrite(t, root, MigrationPath, &transition, func() { transition.Additions[2].ContractSHA256 = changed })
	if _, err := ValidateTransition(root); err == nil || !strings.Contains(err.Error(), "invalid approved addition functions.recursion") {
		t.Fatalf("self-repinned addition accepted: %v", err)
	}
}

func TestLiveTeachingContractsDoNotDependOnInternalReviewBytes(t *testing.T) {
	root := fixture(t)
	for path, pin := range reviewedV3 {
		for _, relative := range []string{path, pin.snapshot} {
			if err := os.Remove(filepath.Join(root, filepath.FromSlash(relative))); err != nil {
				t.Fatal(err)
			}
		}
	}
	if _, err := ValidateLiveTransition(root); err != nil {
		t.Fatalf("learner-only metadata acquired private review dependency: %v", err)
	}
	if _, _, err := ReviewedPath(root); err == nil {
		t.Fatal("reviewed source context accepted absent review bytes")
	}
}
