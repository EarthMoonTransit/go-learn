package isolation

import (
	"encoding/json"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"
)

func TestIsolatedGoTelemetryIsOffBeforeFirstCommand(t *testing.T) {
	environment, err := NewEnvironment("telemetry-off")
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { _ = environment.Close() })
	workspace, err := environment.NewWorkspace("metadata")
	if err != nil {
		t.Fatal(err)
	}
	output, err := environment.RunGo(workspace, 10*time.Second, RunGoOptions{Mode: ExecutionModeNormal}, "env", "-json", "GOTELEMETRY", "GOTELEMETRYDIR")
	if err != nil {
		t.Fatal(err)
	}
	var info struct{ GOTELEMETRY, GOTELEMETRYDIR string }
	if err := json.Unmarshal(output, &info); err != nil {
		t.Fatal(err)
	}
	if info.GOTELEMETRY != "off" {
		t.Fatalf("first isolated Go command has telemetry mode %q; want off", info.GOTELEMETRY)
	}
	relative, err := filepath.Rel(environment.Root, info.GOTELEMETRYDIR)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) || filepath.IsAbs(relative) {
		t.Fatalf("telemetry configuration escaped disposable root: %q (%v)", info.GOTELEMETRYDIR, err)
	}
	entries, err := os.ReadDir(info.GOTELEMETRYDIR)
	if err != nil || len(entries) != 1 || entries[0].Name() != "mode" {
		t.Fatalf("disabled telemetry created counters or sidecar files: %v (%v)", entries, err)
	}
	if err := environment.Close(); err != nil {
		t.Fatalf("close after trusted metadata query: %v", err)
	}
	if _, err := os.Stat(environment.Root); !os.IsNotExist(err) {
		t.Fatalf("disposable telemetry configuration survived Close: %v", err)
	}
}

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
