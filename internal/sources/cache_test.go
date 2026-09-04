package sources

import (
	"context"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

type fakeGit struct {
	calls int
	head  string
}

func (f *fakeGit) Run(_ context.Context, _ string, args ...string) (string, error) {
	f.calls++
	for _, arg := range args {
		if arg == "rev-parse" {
			return f.head, nil
		}
	}
	return "", nil
}

func TestPolicyRejectedBeforeGitOrFilesystem(t *testing.T) {
	for _, mode := range []string{"blocked", "concepts-only"} {
		t.Run(mode, func(t *testing.T) {
			runner := &fakeGit{}
			cache := filepath.Join(t.TempDir(), "must-not-exist")
			source := validSource()
			source.UsageMode = mode
			_, err := (Syncer{Git: runner}).Sync(context.Background(), source, cache, false)
			if err == nil {
				t.Fatal("expected policy error")
			}
			if runner.calls != 0 {
				t.Fatalf("git calls = %d", runner.calls)
			}
			if _, statErr := os.Stat(cache); !os.IsNotExist(statErr) {
				t.Fatalf("cache was created: %v", statErr)
			}
		})
	}
}

func TestIncorrectCheckoutSHAIsRejected(t *testing.T) {
	source := validSource()
	cache := t.TempDir()
	root := filepath.Join(cache, source.ID, source.Commit)
	if err := os.MkdirAll(filepath.Join(root, "checkout", ".git"), 0o700); err != nil {
		t.Fatal(err)
	}
	if err := writeJSON(filepath.Join(root, "source-manifest.json"), expectedManifest(source)); err != nil {
		t.Fatal(err)
	}
	runner := &fakeGit{head: strings.Repeat("b", 40)}
	_, err := (Syncer{Git: runner}).Sync(context.Background(), source, cache, false)
	if err == nil || !strings.Contains(err.Error(), "does not match pinned commit") {
		t.Fatalf("expected SHA mismatch, got %v", err)
	}
}

func TestSafeGitEnvironmentDropsInheritedGitVariables(t *testing.T) {
	environment := safeGitEnvironment([]string{"PATH=/bin", "GIT_DIR=/tmp/evil", "GIT_CONFIG_COUNT=1"})
	joined := strings.Join(environment, "\n")
	if strings.Contains(joined, "GIT_DIR=/tmp/evil") || strings.Contains(joined, "GIT_CONFIG_COUNT=1") {
		t.Fatalf("unsafe environment survived: %s", joined)
	}
	if !strings.Contains(joined, "GIT_CONFIG_NOSYSTEM=1") || !strings.Contains(joined, "GIT_CONFIG_GLOBAL=/dev/null") {
		t.Fatalf("safe Git controls missing: %s", joined)
	}
}

func TestSyncRejectsSymlinkedSourceCache(t *testing.T) {
	cache := t.TempDir()
	outside := t.TempDir()
	source := validSource()
	if err := os.Symlink(outside, filepath.Join(cache, source.ID)); err != nil {
		t.Skipf("symlink unavailable: %v", err)
	}
	runner := &fakeGit{}
	_, err := (Syncer{Git: runner}).Sync(context.Background(), source, cache, false)
	if err == nil || !strings.Contains(err.Error(), "source cache is not a real directory") {
		t.Fatalf("expected source-cache symlink rejection, got %v", err)
	}
	if runner.calls != 0 {
		t.Fatalf("git calls = %d", runner.calls)
	}
}

func TestValidateRejectsInvalidRegistryEntry(t *testing.T) {
	source := validSource()
	source.Repository = "https://user:secret@example.com/repo?x=1"
	source.Commit = "short"
	if err := Validate(source); err == nil {
		t.Fatal("expected invalid source entry")
	}
}

func validSource() Source {
	return Source{
		ID:            "fixture-source",
		Repository:    "https://example.com/source",
		DefaultBranch: "main",
		Commit:        strings.Repeat("a", 40),
		License:       SourceLicense{Expression: "MIT", EvidencePath: "LICENSE"},
		UsageMode:     "adapt",
		Adapter:       "fixture-v1",
		Scope:         "test fixture",
		Evidence:      []string{"https://example.com/evidence"},
		Risks:         []string{"test risk"},
	}
}
