package reports

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/adapters/exercism"
	"go-practice-bank/internal/catalog"
	curriculumdata "go-practice-bank/internal/curriculum"
	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/sources"
	"go-practice-bank/internal/triage"
)

func TestV4KeepsHistoricalSourceReportsBoundToV3(t *testing.T) {
	repository, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	root := t.TempDir()
	// Metadata only: no source artifact, authored implementation or test is run.
	for _, path := range []string{"curriculum", "sources", "catalog"} {
		if err := isolation.CopyTree(filepath.Join(repository, path), filepath.Join(root, path)); err != nil {
			t.Fatal(err)
		}
	}
	if err := os.MkdirAll(filepath.Join(root, "reports"), 0700); err != nil {
		t.Fatal(err)
	}
	registry, err := sources.LoadRegistry(root)
	if err != nil {
		t.Fatal(err)
	}
	source, err := registry.Find("exercism-go")
	if err != nil {
		t.Fatal(err)
	}
	candidates, err := catalog.ReadJSONL(root, "catalog/candidates/exercism-go.jsonl")
	if err != nil {
		t.Fatal(err)
	}
	inventory, err := curriculumdata.LoadConceptInventory(root)
	if err != nil {
		t.Fatal(err)
	}
	config := exercism.Config{}
	for _, slug := range inventory.Concepts {
		config.Concepts = append(config.Concepts, exercism.ConceptConfig{Slug: slug})
	}
	if _, err := Generate(root, source, candidates, config); err != nil {
		t.Fatal(err)
	}
	bundle, err := triage.Load(root)
	if err != nil {
		t.Fatal(err)
	}
	version, contracts, err := triage.LoadReviewedSkillContracts(root)
	if err != nil {
		t.Fatal(err)
	}
	if version != 3 || len(contracts) != 177 {
		t.Fatalf("historical source context=v%d/%d ids", version, len(contracts))
	}
	if err := triage.Validate(bundle, candidates, contracts, version, source); err != nil {
		t.Fatal(err)
	}
	if err := triage.GenerateReports(root, bundle, contracts); err != nil {
		t.Fatal(err)
	}
	files, err := os.ReadDir(filepath.Join(root, "reports"))
	if err != nil {
		t.Fatal(err)
	}
	if len(files) != 13 {
		t.Fatalf("historical source report inventory=%d, want13", len(files))
	}
	for _, file := range files {
		generated, err := os.ReadFile(filepath.Join(root, "reports", file.Name()))
		if err != nil {
			t.Fatal(err)
		}
		expected, err := os.ReadFile(filepath.Join(repository, "reports", file.Name()))
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(generated, expected) {
			t.Fatalf("v4 reinterpreted historical source report %s", file.Name())
		}
	}
}
