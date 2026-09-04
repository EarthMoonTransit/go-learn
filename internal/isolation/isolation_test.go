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

func TestLegacyParserReturnsTypedRaceMode(t *testing.T) {
	normal, err := ParseLegacyCommand("go test ./...")
	if err != nil {
		t.Fatal(err)
	}
	race, err := ParseLegacyCommand("go test -race ./...")
	if err != nil {
		t.Fatal(err)
	}
	if normal.Mode != ExecutionModeNormal || race.Mode != ExecutionModeRace {
		t.Fatalf("unexpected modes: normal=%q race=%q", normal.Mode, race.Mode)
	}
}

func TestEnvironmentDoesNotInheritSecrets(t *testing.T) {
	t.Setenv("TASK_SECRET_FOR_TEST", "must-not-leak")
	environment, err := NewEnvironment("minimal-env")
	if err != nil {
		t.Fatal(err)
	}
	defer environment.Close()
	values, err := environment.GoEnv(RunGoOptions{Mode: ExecutionModeNormal})
	if err != nil {
		t.Fatal(err)
	}
	joined := strings.Join(values, "\n")
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

func TestGoEnvironmentUsesCGOOnlyForTypedRaceMode(t *testing.T) {
	t.Setenv("CGO_CFLAGS", "-DUNTRUSTED_NATIVE_FLAG")
	t.Setenv("CGO_LDFLAGS", "-luntrusted")
	t.Setenv("CC", "/untrusted/compiler")
	environment, err := NewEnvironment("execution-mode")
	if err != nil {
		t.Fatal(err)
	}
	defer environment.Close()

	normal, err := environment.GoEnv(RunGoOptions{Mode: ExecutionModeNormal})
	if err != nil {
		t.Fatal(err)
	}
	race, err := environment.GoEnv(RunGoOptions{Mode: ExecutionModeRace})
	if err != nil {
		t.Fatal(err)
	}
	if !containsEnvironmentValue(normal, "CGO_ENABLED=0") {
		t.Fatalf("normal environment does not disable cgo: %v", normal)
	}
	if !containsEnvironmentValue(race, "CGO_ENABLED=1") {
		t.Fatalf("race environment does not enable cgo: %v", race)
	}
	joinedRace := strings.Join(race, "\n")
	if strings.Contains(joinedRace, "CGO_CFLAGS=") || strings.Contains(joinedRace, "CGO_LDFLAGS=") || strings.Contains(joinedRace, "CC=") || strings.Contains(joinedRace, "UNTRUSTED") {
		t.Fatalf("race environment inherited native compiler or linker flags: %v", race)
	}
	if _, err := environment.GoEnv(RunGoOptions{Mode: ExecutionMode("invalid")}); err == nil {
		t.Fatal("unknown execution mode was accepted")
	}
}

func TestRaceDetectorSupportMatrixIncludesGitHubLinuxAMD64(t *testing.T) {
	if !RaceDetectorSupported("linux", "amd64") {
		t.Fatal("linux/amd64 must never be classified as race-unsupported")
	}
	if RaceDetectorSupported("js", "wasm") {
		t.Fatal("js/wasm must be classified as race-unsupported")
	}
}

func containsEnvironmentValue(values []string, want string) bool {
	for _, value := range values {
		if value == want {
			return true
		}
	}
	return false
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
