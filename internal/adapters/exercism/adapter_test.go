package exercism

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"

	"go-practice-bank/internal/catalog"
	"go-practice-bank/internal/sources"
)

func TestLoadConfigAndDiscoverExercises(t *testing.T) {
	config, err := LoadConfig("testdata/fixture")
	if err != nil {
		t.Fatal(err)
	}
	if config.Slug != "go" || len(config.Concepts) != 2 || len(config.Exercises.Concept) != 1 || len(config.Exercises.Practice) != 1 {
		t.Fatalf("unexpected config: %+v", config)
	}

	candidates, err := (Adapter{}).Harvest("testdata/fixture", fixtureSource(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if got := []string{candidates[0].Slug, candidates[1].Slug}; !reflect.DeepEqual(got, []string{"slice-lab", "echo"}) {
		t.Fatalf("candidate order = %v", got)
	}
	echo := candidate(t, candidates, "echo")
	if len(echo.StarterFilePaths) != 1 || len(echo.PublicTestPaths) != 2 || len(echo.ReferenceSolutionPaths) != 1 || len(echo.GeneratedTestDataPaths) != 2 {
		t.Fatalf("unexpected echo classification: %+v", echo)
	}
	if !contains(echo.AdditionalFilePaths, "exercises/practice/echo/notes.txt") {
		t.Fatalf("additional files = %v", echo.AdditionalFilePaths)
	}
	foundHash := ""
	for _, file := range echo.FileHashes {
		if file.Path == "exercises/practice/echo/notes.txt" {
			foundHash = file.SHA256
		}
	}
	if foundHash != "64323fe1f83aedac649e8ec87387767ae583f7812505bab4ebaef95dd1b1509d" {
		t.Fatalf("notes hash = %q", foundHash)
	}
	jsonl, err := catalog.MarshalJSONL(candidates)
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(jsonl), "Return the supplied string unchanged") {
		t.Fatal("candidate JSONL leaked upstream/fixture file content")
	}
}

func TestHarvestIsDeterministic(t *testing.T) {
	first, err := (Adapter{}).Harvest("testdata/fixture", fixtureSource(), nil)
	if err != nil {
		t.Fatal(err)
	}
	second, err := (Adapter{}).Harvest("testdata/fixture", fixtureSource(), nil)
	if err != nil {
		t.Fatal(err)
	}
	firstJSON, _ := catalog.MarshalJSONL(first)
	secondJSON, _ := catalog.MarshalJSONL(second)
	if !reflect.DeepEqual(firstJSON, secondJSON) {
		t.Fatal("harvest output is not byte-identical")
	}
}

func TestMissingDeclaredFileFails(t *testing.T) {
	root := copyFixture(t)
	if err := os.Remove(filepath.Join(root, "exercises", "practice", "echo", "echo_test.go")); err != nil {
		t.Fatal(err)
	}
	_, err := (Adapter{}).Harvest(root, fixtureSource(), nil)
	if err == nil || !strings.Contains(err.Error(), "declared test file") {
		t.Fatalf("expected missing declared file error, got %v", err)
	}
}

func TestMissingExerciseBecomesIncompleteCandidate(t *testing.T) {
	root := copyFixture(t)
	if err := os.RemoveAll(filepath.Join(root, "exercises", "practice", "echo")); err != nil {
		t.Fatal(err)
	}
	candidates, err := (Adapter{}).Harvest(root, fixtureSource(), nil)
	if err != nil {
		t.Fatal(err)
	}
	echo := candidate(t, candidates, "echo")
	if !contains(echo.ImportWarnings, "exercise directory is missing") || len(echo.FileHashes) != 0 {
		t.Fatalf("unexpected missing exercise candidate: %+v", echo)
	}
}

func TestUnknownConceptIsReported(t *testing.T) {
	root := copyFixture(t)
	configPath := filepath.Join(root, "config.json")
	b, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	b = []byte(strings.Replace(string(b), `"practices": ["strings"]`, `"practices": ["ghost-concept"]`, 1))
	if err := os.WriteFile(configPath, b, 0o644); err != nil {
		t.Fatal(err)
	}
	candidates, err := (Adapter{}).Harvest(root, fixtureSource(), nil)
	if err != nil {
		t.Fatal(err)
	}
	if !contains(candidate(t, candidates, "echo").ImportWarnings, "unknown upstream concept: ghost-concept") {
		t.Fatalf("warnings = %v", candidate(t, candidates, "echo").ImportWarnings)
	}
}

func TestIntermediateSymlinkIsRejected(t *testing.T) {
	temporary := t.TempDir()
	root := filepath.Join(temporary, "fixture")
	copyTree(t, "testdata/fixture", root)
	outside := filepath.Join(temporary, "outside")
	if err := os.Rename(filepath.Join(root, "exercises"), outside); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(root, "exercises")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	_, err := (Adapter{}).Harvest(root, fixtureSource(), nil)
	if err == nil || !strings.Contains(err.Error(), "symlink component") {
		t.Fatalf("expected symlink rejection, got %v", err)
	}
}

func TestTraversalSlugIsRejected(t *testing.T) {
	root := copyFixture(t)
	configPath := filepath.Join(root, "config.json")
	b, err := os.ReadFile(configPath)
	if err != nil {
		t.Fatal(err)
	}
	b = []byte(strings.Replace(string(b), `"slug": "echo"`, `"slug": "../escape"`, 1))
	if err := os.WriteFile(configPath, b, 0o644); err != nil {
		t.Fatal(err)
	}
	_, err = (Adapter{}).Harvest(root, fixtureSource(), nil)
	if err == nil || !strings.Contains(err.Error(), "invalid exercise slug") {
		t.Fatalf("expected traversal rejection, got %v", err)
	}
}

func TestGeneratedCommitResolvesToExactCompanionRevision(t *testing.T) {
	companion := fixtureCompanion()
	candidates, err := (Adapter{}).Harvest("testdata/fixture", fixtureSource(), &companion)
	if err != nil {
		t.Fatal(err)
	}
	references := candidate(t, candidates, "echo").ProvenanceRefs
	if len(references) != 1 {
		t.Fatalf("provenance refs = %+v", references)
	}
	wantCommit := strings.Repeat("a", 40)
	if references[0].Commit != wantCommit || references[0].AbbreviatedCommit != "aaaaaaa" {
		t.Fatalf("provenance ref = %+v", references[0])
	}
}

func TestGeneratedCommitMustResolveToRegisteredCompanionRevision(t *testing.T) {
	companion := fixtureCompanion()
	companion.ProvenanceCommits = nil
	_, err := (Adapter{}).Harvest("testdata/fixture", fixtureSource(), &companion)
	if err == nil || !strings.Contains(err.Error(), "resolves to 0 registered revisions") {
		t.Fatalf("expected unresolved companion revision error, got %v", err)
	}
}

func fixtureSource() sources.Source {
	return sources.Source{
		ID:         "exercism-go",
		Repository: "https://github.com/exercism/go",
		Commit:     strings.Repeat("a", 40),
		License:    sources.SourceLicense{Expression: "MIT"},
		UsageMode:  "adapt",
		Adapter:    "exercism-go-v1",
	}
}

func fixtureCompanion() sources.Source {
	return sources.Source{
		ID:         "exercism-problem-specifications",
		Repository: "https://github.com/exercism/problem-specifications",
		Commit:     strings.Repeat("b", 40),
		License:    sources.SourceLicense{Expression: "MIT"},
		UsageMode:  "reference-only",
		Adapter:    "problem-specifications-provenance-v1",
		ProvenanceCommits: []sources.ProvenanceCommit{{
			Commit: strings.Repeat("a", 40), Evidence: "https://example.com/commit/" + strings.Repeat("a", 40),
		}},
	}
}

func candidate(t *testing.T, candidates []catalog.Candidate, slug string) catalog.Candidate {
	t.Helper()
	for _, candidate := range candidates {
		if candidate.Slug == slug {
			return candidate
		}
	}
	t.Fatalf("candidate %q not found", slug)
	return catalog.Candidate{}
}

func copyFixture(t *testing.T) string {
	t.Helper()
	root := filepath.Join(t.TempDir(), "fixture")
	copyTree(t, "testdata/fixture", root)
	return root
}

func copyTree(t *testing.T, source, destination string) {
	t.Helper()
	err := filepath.WalkDir(source, func(path string, entry os.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.MkdirAll(target, 0o755)
		}
		b, err := os.ReadFile(path)
		if err != nil {
			return err
		}
		return os.WriteFile(target, b, 0o644)
	})
	if err != nil {
		t.Fatal(err)
	}
}
