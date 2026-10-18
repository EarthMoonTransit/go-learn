package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestV4CollectionBindsMigrationAndReviewedInputs(t *testing.T) {
	root := fixtureRepository(t)
	collection, err := fakeSession(t, root).CheckAll(nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := ValidateCheckCollection(root, collection, nil); err != nil {
		t.Fatal(err)
	}
	contentBefore, err := TaskContentDigest(filepath.Join(root, "tasks", "drafts", syntheticID))
	if err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{
		"schemas/taxonomy-v4.schema.json", "schemas/taxonomy-v3-to-v4.schema.json",
		"schemas/coverage-scopes.schema.json", "curriculum/coverage-scopes.yaml",
		"schemas/coverage-report.schema.json",
		"curriculum/migrations/v3-to-v4.yaml", "curriculum/migrations/snapshots/skills-v3.yaml",
		"curriculum/mappings/exercism-go.yaml", "curriculum/evidence/exercism-go.yaml",
		"catalog/reviews/exercism-go.yaml", "catalog/seed-plans/exercism-foundation-pilot.yaml",
		"curriculum/migrations/snapshots/exercism-map-v3.yaml", "curriculum/migrations/snapshots/exercism-evidence-v3.yaml",
		"curriculum/migrations/snapshots/exercism-reviews-v3.yaml", "curriculum/migrations/snapshots/exercism-seeds-v3.yaml",
	} {
		t.Run(path, func(t *testing.T) {
			absolute := filepath.Join(root, filepath.FromSlash(path))
			before, err := os.ReadFile(absolute)
			if err != nil {
				t.Fatal(err)
			}
			if err := os.WriteFile(absolute, append(append([]byte(nil), before...), '\n'), 0600); err != nil {
				t.Fatal(err)
			}
			t.Cleanup(func() {
				if err := os.WriteFile(absolute, before, 0600); err != nil {
					t.Error(err)
				}
			})
			if err := ValidateCheckCollection(root, collection, nil); err == nil || !strings.Contains(err.Error(), "digest changed") {
				t.Fatalf("changed %s accepted: %v", path, err)
			}
			if path == "schemas/coverage-report.schema.json" {
				contentAfter, err := TaskContentDigest(filepath.Join(root, "tasks", "drafts", syntheticID))
				if err != nil || contentAfter != contentBefore {
					t.Fatalf("report schema changed learner content identity: %v", err)
				}
			}
		})
	}
}

func TestV3CollectionRetainsOriginalFreshnessContract(t *testing.T) {
	root := fixtureRepository(t)
	baseline, err := os.ReadFile(filepath.Join(root, "curriculum/migrations/snapshots/skills-v3.yaml"))
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "curriculum/skills.yaml"), baseline, 0600); err != nil {
		t.Fatal(err)
	}
	collection, err := fakeSession(t, root).CheckAll(nil)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(filepath.Join(root, "curriculum/migrations/v3-to-v4.yaml"), []byte("not part of a v3 collection\n"), 0600); err != nil {
		t.Fatal(err)
	}
	for _, path := range []string{"curriculum/coverage-scopes.yaml", "schemas/coverage-scopes.schema.json", "schemas/coverage-report.schema.json"} {
		if err := os.WriteFile(filepath.Join(root, filepath.FromSlash(path)), []byte("not part of a v3 collection\n"), 0600); err != nil {
			t.Fatal(err)
		}
	}
	if err := ValidateCheckCollection(root, collection, nil); err != nil {
		t.Fatalf("v4 metadata reinterpreted v3 collection: %v", err)
	}
}
