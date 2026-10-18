// Package isolation prepares disposable workspaces for authored task code.
package isolation

import (
	"context"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

const maxCopiedFileSize = 8 << 20

var ErrCommandTimeout = errors.New("Go command timed out")

// ExecutionMode selects the minimum Go build environment required by a
// verifier operation. Callers choose this mode from trusted task metadata;
// RunGo never infers it by inspecting command text.
type ExecutionMode string

const (
	ExecutionModeNormal ExecutionMode = "normal"
	ExecutionModeRace   ExecutionMode = "race"
)

// RunGoOptions is the typed execution policy for one Go command.
type RunGoOptions struct {
	Mode ExecutionMode
}

// Environment owns a disposable directory and isolated Go caches.
type Environment struct {
	Root string
}

func NewEnvironment(prefix string) (*Environment, error) {
	if prefix == "" || strings.ContainsAny(prefix, `/\\`) {
		return nil, errors.New("temporary environment prefix must be a simple name")
	}
	root, err := os.MkdirTemp("", prefix+"-*")
	if err != nil {
		return nil, fmt.Errorf("create isolated environment: %w", err)
	}
	for _, name := range []string{"cache", "home", "modcache", "tmp", "workspaces"} {
		if err := os.Mkdir(filepath.Join(root, name), 0o700); err != nil {
			_ = os.RemoveAll(root)
			return nil, fmt.Errorf("prepare isolated environment: %w", err)
		}
	}
	environment := &Environment{Root: root}
	// GOTELEMETRY is a read-only `go env` value, not a mode override.
	// Disable collection before the first Go command can open counters or
	// start a telemetry sidecar that outlives the disposable environment.
	mode := filepath.Join(environment.configDir(), "go", "telemetry", "mode")
	if err := os.MkdirAll(filepath.Dir(mode), 0o700); err != nil {
		_ = environment.Close()
		return nil, fmt.Errorf("prepare isolated telemetry configuration: %w", err)
	}
	if err := os.WriteFile(mode, []byte("off\n"), 0o600); err != nil {
		_ = environment.Close()
		return nil, fmt.Errorf("disable isolated Go telemetry: %w", err)
	}
	return environment, nil
}

// Match os.UserConfigDir for the child environment's isolated home/AppData.
// No path or setting from the user's actual configuration is consulted.
func (e *Environment) configDir() string {
	home := filepath.Join(e.Root, "home")
	switch runtime.GOOS {
	case "darwin", "ios":
		return filepath.Join(home, "Library", "Application Support")
	case "windows":
		return filepath.Join(home, "AppData", "Roaming")
	case "plan9":
		return filepath.Join(home, "lib")
	default:
		return filepath.Join(home, ".config")
	}
}

func (e *Environment) Close() error {
	if e == nil || e.Root == "" {
		return nil
	}
	return os.RemoveAll(e.Root)
}

func (e *Environment) NewWorkspace(name string) (string, error) {
	if e == nil || e.Root == "" {
		return "", errors.New("isolated environment is not initialized")
	}
	if name == "" || strings.ContainsAny(name, `/\\`) || name == "." || name == ".." {
		return "", fmt.Errorf("invalid workspace name %q", name)
	}
	path, err := os.MkdirTemp(filepath.Join(e.Root, "workspaces"), name+"-*")
	if err != nil {
		return "", fmt.Errorf("create workspace: %w", err)
	}
	return path, nil
}

func (e *Environment) GoEnv(options RunGoOptions) ([]string, error) {
	cgoEnabled := "0"
	switch options.Mode {
	case ExecutionModeNormal:
	case ExecutionModeRace:
		cgoEnabled = "1"
	default:
		return nil, fmt.Errorf("unknown Go execution mode %q", options.Mode)
	}
	values := map[string]string{
		"CGO_ENABLED": cgoEnabled,
		"GOENV":       "off",
		"GOCACHE":     filepath.Join(e.Root, "cache"),
		"GOMODCACHE":  filepath.Join(e.Root, "modcache"),
		"GOFLAGS":     "-mod=readonly",
		"GOPROXY":     "off",
		"GOTOOLCHAIN": "local",
		"GOWORK":      "off",
		"HOME":        filepath.Join(e.Root, "home"),
		"PATH":        strings.Join([]string{filepath.Join(runtime.GOROOT(), "bin"), "/usr/bin", "/bin"}, string(filepath.ListSeparator)),
		"TMPDIR":      filepath.Join(e.Root, "tmp"),
	}
	if runtime.GOOS == "windows" {
		values["APPDATA"] = e.configDir()
		for _, key := range []string{"PATHEXT", "SYSTEMROOT", "WINDIR"} {
			if value, ok := os.LookupEnv(key); ok {
				values[key] = value
			}
		}
	}
	if runtime.GOOS == "plan9" {
		values["home"] = filepath.Join(e.Root, "home")
	}
	keys := make([]string, 0, len(values))
	for key := range values {
		keys = append(keys, key)
	}
	sort.Strings(keys)
	result := make([]string, 0, len(keys))
	for _, key := range keys {
		result = append(result, key+"="+values[key])
	}
	return result, nil
}

// RaceDetectorSupported reports whether the active Go toolchain documents
// race-detector support for a GOOS/GOARCH pair. Compiler availability is a
// separate infrastructure property and is diagnosed by the Go command.
func RaceDetectorSupported(goos, goarch string) bool {
	switch goos {
	case "linux":
		switch goarch {
		case "amd64", "arm64", "loong64", "ppc64le", "riscv64", "s390x":
			return true
		}
	case "darwin":
		return goarch == "amd64" || goarch == "arm64"
	case "freebsd", "netbsd", "windows":
		return goarch == "amd64"
	}
	return false
}

// RunGo executes the Go binary from the active runtime in a disposable
// workspace. It provides process and filesystem hygiene for trusted repository
// inputs, not an operating-system security sandbox for hostile code.
func (e *Environment) RunGo(dir string, timeout time.Duration, options RunGoOptions, args ...string) ([]byte, error) {
	canonicalRoot, err := filepath.EvalSymlinks(e.Root)
	if err != nil {
		return nil, fmt.Errorf("resolve isolated root: %w", err)
	}
	canonicalDir, err := filepath.EvalSymlinks(dir)
	if err != nil {
		return nil, fmt.Errorf("resolve command directory: %w", err)
	}
	relative, err := filepath.Rel(canonicalRoot, canonicalDir)
	if err != nil || relative == ".." || strings.HasPrefix(relative, ".."+string(filepath.Separator)) {
		return nil, errors.New("command directory escapes isolated environment")
	}
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	executable := filepath.Join(runtime.GOROOT(), "bin", "go")
	if runtime.GOOS == "windows" {
		executable += ".exe"
	}
	cmd := exec.CommandContext(ctx, executable, args...)
	cmd.Dir = canonicalDir
	goEnv, err := e.GoEnv(options)
	if err != nil {
		return nil, err
	}
	cmd.Env = goEnv
	output, runErr := cmd.CombinedOutput()
	if errors.Is(ctx.Err(), context.DeadlineExceeded) {
		return output, fmt.Errorf("%w after %s", ErrCommandTimeout, timeout)
	}
	return output, runErr
}

// CopyTree copies a real directory without following symlinks or accepting
// devices, sockets, named pipes, or oversized files.
func CopyTree(source, destination string) error {
	rootInfo, err := os.Lstat(source)
	if err != nil {
		return fmt.Errorf("inspect source tree: %w", err)
	}
	if rootInfo.Mode()&os.ModeSymlink != 0 || !rootInfo.IsDir() {
		return errors.New("source tree must be a real directory")
	}
	if err := os.MkdirAll(destination, 0o700); err != nil {
		return fmt.Errorf("create destination tree: %w", err)
	}
	return filepath.WalkDir(source, func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		relative, err := filepath.Rel(source, path)
		if err != nil {
			return err
		}
		if relative == "." {
			return nil
		}
		if !fs.ValidPath(filepath.ToSlash(relative)) {
			return fmt.Errorf("invalid source path %q", relative)
		}
		info, err := os.Lstat(path)
		if err != nil {
			return err
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return fmt.Errorf("symlink is not allowed: %s", relative)
		}
		target := filepath.Join(destination, relative)
		if entry.IsDir() {
			return os.Mkdir(target, 0o700)
		}
		if !info.Mode().IsRegular() {
			return fmt.Errorf("non-regular file is not allowed: %s", relative)
		}
		return CopyFile(path, target)
	})
}

// CopyFile copies one regular, non-symlink file to a new destination.
func CopyFile(source, destination string) error {
	info, err := os.Lstat(source)
	if err != nil {
		return fmt.Errorf("inspect source file: %w", err)
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.Mode().IsRegular() {
		return fmt.Errorf("source is not a regular file: %s", source)
	}
	if info.Size() > maxCopiedFileSize {
		return fmt.Errorf("source file exceeds %d bytes: %s", maxCopiedFileSize, source)
	}
	input, err := os.Open(source)
	if err != nil {
		return err
	}
	defer input.Close()
	if err := os.MkdirAll(filepath.Dir(destination), 0o700); err != nil {
		return err
	}
	output, err := os.OpenFile(destination, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		return err
	}
	remove := true
	defer func() {
		_ = output.Close()
		if remove {
			_ = os.Remove(destination)
		}
	}()
	written, err := io.Copy(output, io.LimitReader(input, maxCopiedFileSize+1))
	if err != nil {
		return err
	}
	if written > maxCopiedFileSize {
		return fmt.Errorf("source file exceeds %d bytes: %s", maxCopiedFileSize, source)
	}
	if err := output.Sync(); err != nil {
		return err
	}
	if err := output.Close(); err != nil {
		return err
	}
	remove = false
	return nil
}

type CommandResult struct {
	Command string
	Output  []byte
	Err     error
}

// RunLegacyCommandsInCopies accepts only the legacy go test/go vet manifest
// grammar. Every command receives a new package copy; no shell is involved.
func RunLegacyCommandsInCopies(sourceDirectory string, commands []string, timeout time.Duration) ([]CommandResult, error) {
	environment, err := NewEnvironment("go-task-legacy")
	if err != nil {
		return nil, err
	}
	results, runErr := environment.RunLegacyCommandsInCopies(sourceDirectory, commands, timeout)
	closeErr := environment.Close()
	return results, errors.Join(runErr, closeErr)
}

// RunLegacyCommandsInCopies reuses this environment's compiler caches while
// retaining a fresh workspace for every command.
func (e *Environment) RunLegacyCommandsInCopies(sourceDirectory string, commands []string, timeout time.Duration) ([]CommandResult, error) {
	parsed := make([]LegacyCommand, len(commands))
	for index, command := range commands {
		value, err := ParseLegacyCommand(command)
		if err != nil {
			return nil, fmt.Errorf("legacy command %q: %w", command, err)
		}
		parsed[index] = value
	}
	results := make([]CommandResult, 0, len(commands))
	for index, command := range commands {
		workspace, err := e.NewWorkspace("legacy")
		if err != nil {
			return results, err
		}
		if err := CopyTree(sourceDirectory, workspace); err != nil {
			return results, err
		}
		output, runErr := e.RunGo(workspace, timeout, RunGoOptions{Mode: parsed[index].Mode}, parsed[index].Args...)
		if parsed[index].ExpectFailure {
			if errors.Is(runErr, ErrCommandTimeout) {
				// A timeout is an infrastructure failure, never an expected test failure.
			} else if runErr == nil {
				runErr = errors.New("command unexpectedly succeeded")
			} else {
				runErr = nil
			}
		}
		results = append(results, CommandResult{Command: command, Output: output, Err: runErr})
	}
	return results, nil
}
