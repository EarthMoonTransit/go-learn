package isolation

import (
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestRunShellInCopyDoesNotMutateSource(t *testing.T) {
	source := t.TempDir()
	if err := os.WriteFile(filepath.Join(source, "original"), []byte("unchanged"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := RunShellInCopy(source, "printf changed > touched", 5*time.Second); err != nil {
		t.Fatal(err)
	}
	if _, err := os.Stat(filepath.Join(source, "touched")); !os.IsNotExist(err) {
		t.Fatalf("command mutated source directory: %v", err)
	}
}

func TestCopyTreeRejectsSymlink(t *testing.T) {
	source := t.TempDir()
	outside := filepath.Join(t.TempDir(), "outside")
	if err := os.WriteFile(outside, []byte("outside"), 0o644); err != nil {
		t.Fatal(err)
	}
	if err := os.Symlink(outside, filepath.Join(source, "escape")); err != nil {
		t.Fatal(err)
	}
	if err := CopyTree(source, filepath.Join(t.TempDir(), "copy")); err == nil {
		t.Fatal("symlink was accepted")
	}
}
