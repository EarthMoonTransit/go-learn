package isolation

import (
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestLegacyParserRejectsShellAndOtherProgramsBeforeExecution(t *testing.T) {
	marker := filepath.Join(t.TempDir(), "marker")
	commands := []string{
		"curl https://example.com",
		"env",
		"go test ./... $(touch " + marker + ")",
		"go test ./... > " + marker,
		"go test ./... | tee " + marker,
		"go test ../../...",
	}
	for _, command := range commands {
		if _, err := ParseLegacyCommand(command); err == nil {
			t.Errorf("accepted malicious command %q", command)
		}
	}
	source := t.TempDir()
	if err := os.WriteFile(filepath.Join(source, "go.mod"), []byte("module task.local/preflight\n\ngo 1.27\n"), 0o644); err != nil {
		t.Fatal(err)
	}
	if _, err := RunLegacyCommandsInCopies(source, []string{"go test ./...", "go test ./... > " + marker}, 5*time.Second); err == nil {
		t.Fatal("legacy runner did not reject the complete command list before execution")
	}
	if _, err := os.Stat(marker); !os.IsNotExist(err) {
		t.Fatalf("parser caused an external write: %v", err)
	}
}

func TestLegacyParserAcceptsPinnedManifestGrammar(t *testing.T) {
	for _, command := range []string{
		"go test -run '^$' ./...",
		"! go test ./tests/public",
		"go vet ./...",
		"go test -race -count=1 ./solution/reference",
	} {
		if _, err := ParseLegacyCommand(command); err != nil {
			t.Errorf("rejected %q: %v", command, err)
		}
	}
}

func TestEnvironmentDoesNotInheritSecrets(t *testing.T) {
	t.Setenv("TASK_SECRET_FOR_TEST", "must-not-leak")
	environment, err := NewEnvironment("minimal-env")
	if err != nil {
		t.Fatal(err)
	}
	defer environment.Close()
	joined := strings.Join(environment.GoEnv(), "\n")
	if strings.Contains(joined, "TASK_SECRET_FOR_TEST") || strings.Contains(joined, "must-not-leak") {
		t.Fatal("isolated environment inherited an unrelated variable")
	}
	for _, key := range []string{"HOME=", "GOCACHE=", "GOMODCACHE=", "TMPDIR="} {
		index := strings.Index(joined, key)
		if index < 0 || !strings.Contains(joined[index:], environment.Root) {
			t.Fatalf("%s is not rooted in isolated environment: %s", key, joined)
		}
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
