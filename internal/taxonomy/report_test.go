package taxonomy

import (
	"bytes"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func TestMigrationReportIsDeterministicAndCheckIsReadOnly(t *testing.T) {
	root := fixture(t)
	first, err := MigrationReport(root)
	if err != nil {
		t.Fatal(err)
	}
	second, err := MigrationReport(root)
	if err != nil || !bytes.Equal(first, second) {
		t.Fatalf("nondeterministic metadata report: %v", err)
	}
	for _, needle := range []string{"145 active skills", "35 deprecated stubs", "180 total IDs", "418 primary / 708 total", "not actual task counts", "Recursion implementation/debugging authoring remains blocked"} {
		if !bytes.Contains(first, []byte(needle)) {
			t.Fatalf("missing report claim %q", needle)
		}
	}
	path := filepath.Join(root, filepath.FromSlash(ReportPath))
	if err := CheckMigrationReport(root); err == nil {
		t.Fatal("missing report accepted")
	}
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		t.Fatal("read-only report check wrote output")
	}
	if err := os.MkdirAll(filepath.Dir(path), 0700); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, first, 0600); err != nil {
		t.Fatal(err)
	}
	if err := CheckMigrationReport(root); err != nil {
		t.Fatal(err)
	}
	stale := bytes.Replace(first, []byte("145 active skills"), []byte("142 active skills"), 1)
	if err := os.WriteFile(path, stale, 0600); err != nil {
		t.Fatal(err)
	}
	if err := CheckMigrationReport(root); err == nil || !strings.Contains(err.Error(), "is stale") {
		t.Fatalf("stale report accepted: %v", err)
	}
	after, err := os.ReadFile(path)
	if err != nil || !bytes.Equal(after, stale) {
		t.Fatal("report check repaired file implicitly")
	}
}
