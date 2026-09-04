package sources

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

const CacheEnvironment = "GO_LEARN_SOURCE_CACHE"

const maxGitOutputBytes = 4 << 20

type Manifest struct {
	SchemaVersion         int    `json:"schema_version"`
	SourceID              string `json:"source_id"`
	Repository            string `json:"repository"`
	Commit                string `json:"commit"`
	License               string `json:"license"`
	UsageMode             string `json:"usage_mode"`
	Adapter               string `json:"adapter"`
	CheckoutDirectory     string `json:"checkout_directory"`
	GitHead               string `json:"git_head"`
	DetachedHead          bool   `json:"detached_head"`
	SubmodulesInitialized bool   `json:"submodules_initialized"`
	GitmodulesPresent     bool   `json:"gitmodules_present"`
}

type Checkout struct {
	Source       Source
	CacheRoot    string
	Directory    string
	ManifestPath string
	Manifest     Manifest
}

type GitRunner interface {
	Run(ctx context.Context, directory string, args ...string) (string, error)
}

type ExecGit struct{}

func (ExecGit) Run(ctx context.Context, directory string, args ...string) (string, error) {
	cmd := exec.CommandContext(ctx, "git", args...)
	cmd.Dir = directory
	cmd.Env = safeGitEnvironment(os.Environ())
	var stdout, stderr limitedBuffer
	stdout.limit = maxGitOutputBytes
	stderr.limit = maxGitOutputBytes
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr
	if err := cmd.Run(); err != nil {
		return "", fmt.Errorf("git %s: %w: %s", strings.Join(args, " "), err, strings.TrimSpace(stderr.String()))
	}
	return strings.TrimSpace(stdout.String()), nil
}

type Syncer struct {
	Git GitRunner
}

func DefaultCacheRoot() (string, error) {
	if configured := os.Getenv(CacheEnvironment); configured != "" {
		absolute, err := filepath.Abs(configured)
		if err != nil {
			return "", fmt.Errorf("resolve %s: %w", CacheEnvironment, err)
		}
		return filepath.Clean(absolute), nil
	}
	home, err := os.UserHomeDir()
	if err != nil {
		return "", fmt.Errorf("resolve home directory: %w", err)
	}
	return filepath.Join(home, ".cache", "go-learn", "sources"), nil
}

func (s Syncer) Sync(ctx context.Context, source Source, cacheRoot string, allowConceptsOnly bool) (Checkout, error) {
	if err := Validate(source); err != nil {
		return Checkout{}, err
	}
	if source.UsageMode == "blocked" {
		return Checkout{}, fmt.Errorf("source %q is blocked by license policy", source.ID)
	}
	if source.UsageMode == "concepts-only" && !allowConceptsOnly {
		return Checkout{}, fmt.Errorf("source %q is concepts-only; pass the explicit allow flag to synchronize it", source.ID)
	}
	if s.Git == nil {
		s.Git = ExecGit{}
	}
	cacheRoot, err := secureCacheRoot(cacheRoot)
	if err != nil {
		return Checkout{}, fmt.Errorf("resolve cache root: %w", err)
	}
	sourceRoot := filepath.Join(cacheRoot, source.ID)
	if info, err := os.Lstat(sourceRoot); err == nil && (info.Mode()&os.ModeSymlink != 0 || !info.IsDir()) {
		return Checkout{}, fmt.Errorf("source cache is not a real directory: %s", sourceRoot)
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Checkout{}, fmt.Errorf("inspect source cache: %w", err)
	}
	if err := os.MkdirAll(sourceRoot, 0o700); err != nil {
		return Checkout{}, fmt.Errorf("create source cache: %w", err)
	}
	_ = os.Chmod(sourceRoot, 0o700)
	finalRoot := filepath.Join(sourceRoot, source.Commit)
	if info, err := os.Lstat(finalRoot); err == nil {
		if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
			return Checkout{}, fmt.Errorf("cache target is not a real directory: %s", finalRoot)
		}
		return s.verifyExisting(ctx, source, cacheRoot, finalRoot)
	} else if err != nil && !errors.Is(err, os.ErrNotExist) {
		return Checkout{}, fmt.Errorf("inspect cache target: %w", err)
	}
	temporary, err := os.MkdirTemp(sourceRoot, ".sync-")
	if err != nil {
		return Checkout{}, fmt.Errorf("create temporary checkout: %w", err)
	}
	defer os.RemoveAll(temporary)
	checkoutDir := filepath.Join(temporary, "checkout")
	hooksDir := filepath.Join(temporary, "empty-hooks")
	if err := os.MkdirAll(hooksDir, 0o700); err != nil {
		return Checkout{}, fmt.Errorf("create empty hooks directory: %w", err)
	}
	gitArgs := func(args ...string) []string {
		return safeGitArgs(hooksDir, args...)
	}
	if _, err := s.Git.Run(ctx, temporary, gitArgs("init", "--quiet", checkoutDir)...); err != nil {
		return Checkout{}, err
	}
	if _, err := s.Git.Run(ctx, checkoutDir, gitArgs("remote", "add", "origin", source.Repository)...); err != nil {
		return Checkout{}, err
	}
	if _, err := s.Git.Run(ctx, checkoutDir, gitArgs("fetch", "--quiet", "--no-tags", "--depth=1", "origin", source.Commit)...); err != nil {
		return Checkout{}, err
	}
	if _, err := s.Git.Run(ctx, checkoutDir, gitArgs("checkout", "--quiet", "--detach", source.Commit)...); err != nil {
		return Checkout{}, err
	}
	manifest, err := s.verifyGitState(ctx, source, checkoutDir)
	if err != nil {
		return Checkout{}, err
	}
	manifestPath := filepath.Join(temporary, "source-manifest.json")
	if err := writeJSON(manifestPath, manifest); err != nil {
		return Checkout{}, err
	}
	if err := os.Rename(temporary, finalRoot); err != nil {
		if info, statErr := os.Stat(finalRoot); statErr == nil && info.IsDir() {
			return s.verifyExisting(ctx, source, cacheRoot, finalRoot)
		}
		return Checkout{}, fmt.Errorf("publish source cache: %w", err)
	}
	return Checkout{
		Source:       source,
		CacheRoot:    cacheRoot,
		Directory:    filepath.Join(finalRoot, "checkout"),
		ManifestPath: filepath.Join(finalRoot, "source-manifest.json"),
		Manifest:     manifest,
	}, nil
}

func (s Syncer) Open(ctx context.Context, source Source, cacheRoot string) (Checkout, error) {
	if err := Validate(source); err != nil {
		return Checkout{}, err
	}
	if s.Git == nil {
		s.Git = ExecGit{}
	}
	absolute, err := secureCacheRoot(cacheRoot)
	if err != nil {
		return Checkout{}, fmt.Errorf("resolve cache root: %w", err)
	}
	return s.verifyExisting(ctx, source, absolute, filepath.Join(absolute, source.ID, source.Commit))
}

func (s Syncer) verifyExisting(ctx context.Context, source Source, cacheRoot, finalRoot string) (Checkout, error) {
	manifestPath := filepath.Join(finalRoot, "source-manifest.json")
	if err := requireRealDirectory(finalRoot); err != nil {
		return Checkout{}, err
	}
	if err := requireRegularFile(manifestPath); err != nil {
		return Checkout{}, err
	}
	b, err := readBounded(manifestPath, 1<<20)
	if err != nil {
		return Checkout{}, fmt.Errorf("read cached source manifest: %w", err)
	}
	var manifest Manifest
	dec := json.NewDecoder(bytes.NewReader(b))
	dec.DisallowUnknownFields()
	if err := dec.Decode(&manifest); err != nil {
		return Checkout{}, fmt.Errorf("decode cached source manifest: %w", err)
	}
	var trailing any
	if err := dec.Decode(&trailing); !errors.Is(err, io.EOF) {
		return Checkout{}, errors.New("cached source manifest contains trailing JSON")
	}
	expected := expectedManifest(source)
	if manifest.SchemaVersion != expected.SchemaVersion || manifest.SourceID != expected.SourceID || manifest.Repository != expected.Repository || manifest.Commit != expected.Commit || manifest.License != expected.License || manifest.UsageMode != expected.UsageMode || manifest.Adapter != expected.Adapter || manifest.CheckoutDirectory != "checkout" {
		return Checkout{}, fmt.Errorf("cached manifest for %q does not match registry", source.ID)
	}
	checkoutDir := filepath.Join(finalRoot, "checkout")
	if err := requireRealDirectory(checkoutDir); err != nil {
		return Checkout{}, err
	}
	if err := requireRealDirectory(filepath.Join(checkoutDir, ".git")); err != nil {
		return Checkout{}, fmt.Errorf("invalid cached git directory: %w", err)
	}
	verified, err := s.verifyGitState(ctx, source, checkoutDir)
	if err != nil {
		return Checkout{}, err
	}
	if manifest != verified {
		return Checkout{}, fmt.Errorf("cached manifest for %q does not match actual checkout state", source.ID)
	}
	return Checkout{Source: source, CacheRoot: cacheRoot, Directory: checkoutDir, ManifestPath: manifestPath, Manifest: manifest}, nil
}

func (s Syncer) verifyGitState(ctx context.Context, source Source, checkoutDir string) (Manifest, error) {
	gitArgs := func(args ...string) []string { return safeGitArgs("/dev/null", args...) }
	head, err := s.Git.Run(ctx, checkoutDir, gitArgs("rev-parse", "HEAD")...)
	if err != nil {
		return Manifest{}, fmt.Errorf("verify checkout HEAD: %w", err)
	}
	if head != source.Commit {
		return Manifest{}, fmt.Errorf("checkout HEAD %q does not match pinned commit %q", head, source.Commit)
	}
	branch, err := s.Git.Run(ctx, checkoutDir, gitArgs("branch", "--show-current")...)
	if err != nil {
		return Manifest{}, fmt.Errorf("verify detached HEAD: %w", err)
	}
	if branch != "" {
		return Manifest{}, fmt.Errorf("checkout is attached to branch %q", branch)
	}
	remote, err := s.Git.Run(ctx, checkoutDir, gitArgs("remote", "get-url", "origin")...)
	if err != nil {
		return Manifest{}, fmt.Errorf("verify checkout remote: %w", err)
	}
	if remote != source.Repository {
		return Manifest{}, fmt.Errorf("checkout remote %q does not match registry %q", remote, source.Repository)
	}
	status, err := s.Git.Run(ctx, checkoutDir, gitArgs("status", "--porcelain=v1", "--untracked-files=all")...)
	if err != nil {
		return Manifest{}, fmt.Errorf("verify checkout cleanliness: %w", err)
	}
	if status != "" {
		return Manifest{}, errors.New("cached checkout has tracked or untracked modifications")
	}
	if fileExists(filepath.Join(checkoutDir, ".git", "modules")) {
		return Manifest{}, errors.New("cached checkout contains initialized submodules")
	}
	if fileExists(filepath.Join(checkoutDir, ".git", "objects", "info", "alternates")) {
		return Manifest{}, errors.New("cached checkout uses forbidden alternate object storage")
	}
	manifest := expectedManifest(source)
	manifest.GitmodulesPresent = fileExists(filepath.Join(checkoutDir, ".gitmodules"))
	return manifest, nil
}

func expectedManifest(source Source) Manifest {
	return Manifest{
		SchemaVersion:         1,
		SourceID:              source.ID,
		Repository:            source.Repository,
		Commit:                source.Commit,
		License:               source.License.Expression,
		UsageMode:             source.UsageMode,
		Adapter:               source.Adapter,
		CheckoutDirectory:     "checkout",
		GitHead:               source.Commit,
		DetachedHead:          true,
		SubmodulesInitialized: false,
	}
}

func writeJSON(path string, value any) error {
	b, err := json.MarshalIndent(value, "", "  ")
	if err != nil {
		return fmt.Errorf("encode source manifest: %w", err)
	}
	b = append(b, '\n')
	if err := os.WriteFile(path, b, 0o644); err != nil {
		return fmt.Errorf("write source manifest: %w", err)
	}
	return nil
}

func safeGitEnvironment(environment []string) []string {
	result := make([]string, 0, len(environment)+5)
	for _, item := range environment {
		key, _, _ := strings.Cut(item, "=")
		if !strings.HasPrefix(key, "GIT_") {
			result = append(result, item)
		}
	}
	return append(result,
		"GIT_CONFIG_NOSYSTEM=1",
		"GIT_CONFIG_GLOBAL=/dev/null",
		"GIT_TERMINAL_PROMPT=0",
		"GIT_ASKPASS=true",
		"GIT_LFS_SKIP_SMUDGE=1",
	)
}

func safeGitArgs(hooksDirectory string, args ...string) []string {
	prefix := []string{
		"-c", "core.hooksPath=" + hooksDirectory,
		"-c", "core.fsmonitor=false",
		"-c", "core.autocrlf=false",
		"-c", "credential.helper=",
		"-c", "protocol.file.allow=never",
		"-c", "submodule.recurse=false",
		"-c", "fetch.fsckObjects=true",
		"-c", "transfer.fsckObjects=true",
	}
	return append(prefix, args...)
}

func secureCacheRoot(root string) (string, error) {
	absolute, err := filepath.Abs(root)
	if err != nil {
		return "", err
	}
	if err := os.MkdirAll(absolute, 0o700); err != nil {
		return "", fmt.Errorf("create cache root: %w", err)
	}
	resolved, err := filepath.EvalSymlinks(absolute)
	if err != nil {
		return "", fmt.Errorf("resolve cache root: %w", err)
	}
	if err := requireRealDirectory(resolved); err != nil {
		return "", err
	}
	_ = os.Chmod(resolved, 0o700)
	return resolved, nil
}

func requireRealDirectory(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("inspect directory %s: %w", path, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("path is not a real directory: %s", path)
	}
	return nil
}

func requireRegularFile(path string) error {
	info, err := os.Lstat(path)
	if err != nil {
		return fmt.Errorf("inspect file %s: %w", path, err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return fmt.Errorf("path is not a regular file: %s", path)
	}
	return nil
}

type limitedBuffer struct {
	buffer bytes.Buffer
	limit  int
}

func (w *limitedBuffer) Write(p []byte) (int, error) {
	remaining := w.limit - w.buffer.Len()
	if remaining <= 0 {
		return len(p), nil
	}
	if len(p) > remaining {
		_, _ = w.buffer.Write(p[:remaining])
		return len(p), nil
	}
	return w.buffer.Write(p)
}

func (w *limitedBuffer) String() string {
	return w.buffer.String()
}

func fileExists(path string) bool {
	_, err := os.Lstat(path)
	return err == nil
}
