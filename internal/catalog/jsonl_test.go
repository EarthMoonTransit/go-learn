package catalog

import (
	"bytes"
	"os"
	"path/filepath"
	"testing"
)

func TestJSONLStableSortingAndNoOpRewrite(t *testing.T) {
	candidates := []Candidate{
		{SchemaVersion: 1, SourceID: "source", SourcePath: "z/path", Slug: "z", ExerciseKind: "practice", UpstreamConcepts: []string{}, FileHashes: []FileHash{}, ImportWarnings: []string{}},
		{SchemaVersion: 1, SourceID: "source", SourcePath: "a/path", Slug: "a", ExerciseKind: "concept", UpstreamConcepts: []string{}, FileHashes: []FileHash{}, ImportWarnings: []string{}},
	}
	first, err := MarshalJSONL(candidates)
	if err != nil {
		t.Fatal(err)
	}
	candidates[0], candidates[1] = candidates[1], candidates[0]
	second, err := MarshalJSONL(candidates)
	if err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(first, second) {
		t.Fatal("input ordering changed JSONL")
	}
	root := t.TempDir()
	changed, err := WriteJSONL(root, "catalog/candidates/source.jsonl", candidates)
	if err != nil || !changed {
		t.Fatalf("first write: changed=%v err=%v", changed, err)
	}
	changed, err = WriteJSONL(root, "catalog/candidates/source.jsonl", candidates)
	if err != nil || changed {
		t.Fatalf("second write: changed=%v err=%v", changed, err)
	}
	decoded, err := ReadJSONL(root, "catalog/candidates/source.jsonl")
	if err != nil || len(decoded) != 2 || decoded[0].Slug != "a" {
		t.Fatalf("decoded=%+v err=%v", decoded, err)
	}
}

func TestWriteRejectsSymlinkedOutputAncestor(t *testing.T) {
	root := t.TempDir()
	outside := t.TempDir()
	if err := os.Symlink(outside, filepath.Join(root, "catalog")); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	_, err := WriteJSONL(root, "catalog/candidates/source.jsonl", []Candidate{})
	if err == nil {
		t.Fatal("expected symlinked ancestor rejection")
	}
}
