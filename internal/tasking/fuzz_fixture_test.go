package tasking

import (
	"os"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/isolation"
)

const syntheticFuzzID = "synthetic-hex-v2-001"

func fuzzFixtureRepository(t *testing.T) (string, packageData) {
	t.Helper()
	root := fixtureRepository(t)
	repository := repositoryRoot(t)
	if err := os.MkdirAll(filepath.Join(root, "docs"), 0700); err != nil {
		t.Fatal(err)
	}
	if err := isolation.CopyFile(filepath.Join(repository, filepath.FromSlash(fuzzCanonicalContractPath)), filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath))); err != nil {
		t.Fatal(err)
	}
	source := filepath.Join(repository, "internal/testdata/tasking/v2/fuzz", syntheticFuzzID)
	destination := filepath.Join(root, "tasks/drafts", syntheticFuzzID)
	if err := isolation.CopyTree(source, destination); err != nil {
		t.Fatal(err)
	}
	data, problems := loadPackage(root, LocatedTask{ID: syntheticFuzzID, State: StateDraft, Path: destination})
	if len(problems) > 0 {
		t.Fatal(problems)
	}
	return root, data
}
