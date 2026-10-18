package isolation

import (
	"io/fs"
	"os"
	"slices"
	"strings"
	"testing"
)

func TestExecutionSourceIdentityIncludesEveryProductionFile(t *testing.T) {
	entries, err := os.ReadDir(".")
	if err != nil {
		t.Fatal(err)
	}
	var actual []string
	for _, entry := range entries {
		if !entry.IsDir() && strings.HasSuffix(entry.Name(), ".go") && !strings.HasSuffix(entry.Name(), "_test.go") {
			actual = append(actual, entry.Name())
		}
	}
	embedded, err := fs.ReadDir(policySources, ".")
	if err != nil {
		t.Fatal(err)
	}
	var names []string
	for _, entry := range embedded {
		names = append(names, entry.Name())
	}
	if !slices.Equal(actual, names) {
		t.Fatalf("update the explicit production source embed: source=%v embedded=%v", actual, names)
	}
	if digest := SourceDigest(); len(digest) != 64 {
		t.Fatalf("invalid execution source digest %q", digest)
	}
}
