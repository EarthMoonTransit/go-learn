package tasking

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestNativeV4CollectionRejectsOldContractChangedVerifierAndRuntime(t *testing.T) {
	root := fixtureRepository(t)
	collection, err := fakeSession(t, root).CheckAll(nil)
	if err != nil {
		t.Fatal(err)
	}
	if collection.ContractVersion != "task-check/v4" || collection.VerifierSourceDigest != VerifierSourceDigest() || collection.Runtime.Version == "" || collection.Execution.RuntimeQueries != 1 {
		t.Fatalf("missing v4 execution identity: %+v", collection)
	}
	for _, test := range []struct {
		name, want string
		mutate     func(*CheckCollection)
	}{
		{"v3", "require \"task-check/v4\"", func(c *CheckCollection) { c.ContractVersion = "task-check/v3" }},
		{"changed-verifier", "verifier source digest changed", func(c *CheckCollection) { c.VerifierSourceDigest = strings.Repeat("0", 64) }},
		{"changed-runtime-version", "actual Go runtime changed", func(c *CheckCollection) { c.Runtime.Version = "go1.27.999" }},
		{"changed-runtime-platform", "actual Go runtime changed", func(c *CheckCollection) { c.Runtime.GOOS = "other" }},
		{"unobserved-native-count", "accounting", func(c *CheckCollection) { c.Execution.FuzzNativeRuns = 1 }},
	} {
		t.Run(test.name, func(t *testing.T) {
			changed := collection
			test.mutate(&changed)
			err := ValidateCheckCollection(root, changed, nil)
			if err == nil || !strings.Contains(err.Error(), test.want) {
				t.Fatalf("want %q rejection, got %v", test.want, err)
			}
		})
	}
}

func TestFuzzCollectionBindsPublicGrammarWithoutChangingOldTaskContent(t *testing.T) {
	root, data := fuzzFixtureRepository(t)
	oldPath := filepath.Join(root, "tasks/drafts", syntheticID)
	oldBefore, err := TaskContentDigest(oldPath)
	if err != nil {
		t.Fatal(err)
	}
	before, err := TaskInputDigest(root, data.Located)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath))
	body, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, append(body, '\n'), 0600); err != nil {
		t.Fatal(err)
	}
	after, err := TaskInputDigest(root, data.Located)
	if err != nil || before == after {
		t.Fatalf("public grammar changed without freshness invalidation: %v", err)
	}
	oldAfter, err := TaskContentDigest(oldPath)
	if err != nil || oldAfter != oldBefore {
		t.Fatalf("new profile grammar changed existing learner identity: %v", err)
	}
}
