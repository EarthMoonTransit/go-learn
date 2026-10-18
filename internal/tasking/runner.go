package tasking

import (
	"bufio"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"time"
	"unicode"

	"go-practice-bank/internal/fuzznative"
	"go-practice-bank/internal/isolation"
)

const taskCommandTimeout = 60 * time.Second

const maxVariantDiagnosticRunes = 2048

type variantFailureKind string
type variantFailurePhase string

const (
	variantFailureNone            variantFailureKind = "none"
	variantFailureAssembly        variantFailureKind = "assembly"
	variantFailureCompile         variantFailureKind = "compile"
	variantFailureTest            variantFailureKind = "test"
	variantFailureTimeout         variantFailureKind = "timeout"
	variantFailureRaceUnsupported variantFailureKind = "race-unsupported"
	variantFailureRaceCGO         variantFailureKind = "race-cgo-incompatible"

	variantPhaseNone           variantFailurePhase = "none"
	variantPhaseAssembly       variantFailurePhase = "assembly"
	variantPhaseCompile        variantFailurePhase = "compile"
	variantPhaseTest           variantFailurePhase = "test"
	variantPhaseInfrastructure variantFailurePhase = "infrastructure"
)

type variantRequest struct {
	TaskID         string
	GoVersion      string
	Race           bool
	Vet            bool
	TimeoutSeconds int
	Implementation []string
	Tests          []string
}

type mappedFile struct {
	Source   string
	Relative string
}

type workspaceRequest struct {
	TaskID         string
	GoVersion      string
	Race           bool
	Vet            bool
	TimeoutSeconds int
	Skeleton       []mappedFile
	Overlay        []mappedFile
	Tests          []mappedFile
	Checks         []string
}

type variantResult struct {
	Compiled     bool
	Passed       bool
	Canonical    string
	FailureKind  variantFailureKind
	FailurePhase variantFailurePhase
	Diagnostic   string
}

type executionRunner interface {
	RunVariant(variantRequest) variantResult
	RunFuzz(fuzzRequest) fuzznative.Result
	ActualRuntime() (isolation.RuntimeInfo, error)
	RunWorkspace(workspaceRequest) variantResult
	RunLegacy(string, []string) ([]isolation.CommandResult, error)
	Metrics() ExecutionMetrics
	Close() error
}

type realExecutionRunner struct {
	environment *isolation.Environment
	metrics     ExecutionMetrics
	fuzz        *fuzznative.Engine
}

func newRealExecutionRunner() (*realExecutionRunner, error) {
	environment, err := isolation.NewEnvironment("go-task-check")
	if err != nil {
		return nil, err
	}
	return &realExecutionRunner{
		environment: environment,
		metrics:     ExecutionMetrics{Sessions: 1},
	}, nil
}

func (r *realExecutionRunner) RunVariant(request variantRequest) variantResult {
	if request.Race && !isolation.RaceDetectorSupported(runtime.GOOS, runtime.GOARCH) {
		return variantResult{
			FailureKind: variantFailureRaceUnsupported, FailurePhase: variantPhaseInfrastructure,
			Diagnostic: fmt.Sprintf("race detector is unsupported on %s/%s", runtime.GOOS, runtime.GOARCH),
		}
	}
	r.metrics.VariantRuns++
	workspace, err := r.environment.NewWorkspace(request.TaskID)
	if err != nil {
		return assemblyFailure("create isolated workspace", err, r.environment.Root, "")
	}
	r.metrics.Workspaces++
	goMod := []byte("module task.local/" + request.TaskID + "\n\ngo " + request.GoVersion + "\n")
	if err := os.WriteFile(filepath.Join(workspace, "go.mod"), goMod, 0o600); err != nil {
		return assemblyFailure("write isolated go.mod", err, r.environment.Root, workspace)
	}
	subject := filepath.Join(workspace, "subject")
	if err := os.Mkdir(subject, 0o700); err != nil {
		return assemblyFailure("create isolated subject package", err, r.environment.Root, workspace)
	}
	for index, path := range request.Implementation {
		name := fmt.Sprintf("impl_%03d_%s", index, filepath.Base(path))
		if err := isolation.CopyFile(path, filepath.Join(subject, name)); err != nil {
			return assemblyFailure("copy implementation", err, r.environment.Root, workspace, path)
		}
	}
	for index, path := range request.Tests {
		name := fmt.Sprintf("check_%03d_%s", index, filepath.Base(path))
		if err := isolation.CopyFile(path, filepath.Join(subject, name)); err != nil {
			return assemblyFailure("copy tests", err, r.environment.Root, workspace, path)
		}
	}
	mode := isolation.ExecutionModeNormal
	args := []string{"test", "-json", "-count=1"}
	if request.Vet {
		args = []string{"vet"}
	}
	if request.Race {
		mode = isolation.ExecutionModeRace
		args = append(args, "-race")
	}
	args = append(args, "./subject")
	r.metrics.ChildGoRuns++
	output, runErr := r.environment.RunGo(workspace, variantTimeout(request.TimeoutSeconds), isolation.RunGoOptions{Mode: mode}, args...)
	canonical, compiled := canonicalTestEvents(output, request.TaskID)
	if request.Vet && runErr == nil {
		compiled = true
	}
	result := variantResult{Compiled: compiled, Passed: runErr == nil, Canonical: canonical, FailureKind: variantFailureNone, FailurePhase: variantPhaseNone}
	if runErr != nil {
		result.FailureKind, result.FailurePhase = classifyVariantFailure(request.Race, compiled, output, runErr)
		result.Diagnostic = sanitizeVariantDiagnostic(output, runErr, r.environment.Root, workspace, request.TaskID)
	}
	return result
}

func (r *realExecutionRunner) RunWorkspace(request workspaceRequest) variantResult {
	if request.Race && !isolation.RaceDetectorSupported(runtime.GOOS, runtime.GOARCH) {
		return variantResult{FailureKind: variantFailureRaceUnsupported, FailurePhase: variantPhaseInfrastructure,
			Diagnostic: fmt.Sprintf("race detector is unsupported on %s/%s", runtime.GOOS, runtime.GOARCH)}
	}
	r.metrics.VariantRuns++
	workspace, err := r.environment.NewWorkspace(request.TaskID)
	if err != nil {
		return assemblyFailure("create isolated workspace", err, r.environment.Root, "")
	}
	r.metrics.Workspaces++
	goMod := []byte("module task.local/" + request.TaskID + "\n\ngo " + request.GoVersion + "\n")
	if err := os.WriteFile(filepath.Join(workspace, "go.mod"), goMod, 0o600); err != nil {
		return assemblyFailure("write isolated go.mod", err, r.environment.Root, workspace)
	}
	for _, group := range [][]mappedFile{request.Skeleton, request.Overlay, request.Tests} {
		for _, file := range group {
			if err := cleanRelativePath(file.Relative); err != nil {
				return assemblyFailure("reject unsafe workspace path", err, r.environment.Root, workspace, file.Source)
			}
			destination := filepath.Join(workspace, filepath.FromSlash(file.Relative))
			if err := os.MkdirAll(filepath.Dir(destination), 0o700); err != nil {
				return assemblyFailure("create workspace package", err, r.environment.Root, workspace)
			}
			if err := isolation.CopyFile(file.Source, destination); err != nil {
				return assemblyFailure("copy workspace file", err, r.environment.Root, workspace, file.Source)
			}
		}
	}
	mode := isolation.ExecutionModeNormal
	args := []string{"test", "-json", "-count=1"}
	if request.Vet {
		args = []string{"vet"}
	}
	if request.Race {
		mode = isolation.ExecutionModeRace
		args = append(args, "-race")
	}
	args = append(args, request.Checks...)
	r.metrics.ChildGoRuns++
	output, runErr := r.environment.RunGo(workspace, variantTimeout(request.TimeoutSeconds), isolation.RunGoOptions{Mode: mode}, args...)
	canonical, compiled := canonicalTestEvents(output, request.TaskID)
	if request.Vet && runErr == nil {
		compiled = true
	}
	result := variantResult{Compiled: compiled, Passed: runErr == nil, Canonical: canonical, FailureKind: variantFailureNone, FailurePhase: variantPhaseNone}
	if runErr != nil {
		result.FailureKind, result.FailurePhase = classifyVariantFailure(request.Race, compiled, output, runErr)
		result.Diagnostic = sanitizeVariantDiagnostic(output, runErr, r.environment.Root, workspace, request.TaskID)
	}
	return result
}

func variantTimeout(seconds int) time.Duration {
	if seconds <= 0 || seconds > int(taskCommandTimeout/time.Second) {
		return taskCommandTimeout
	}
	return time.Duration(seconds) * time.Second
}

func (r *realExecutionRunner) RunLegacy(source string, commands []string) ([]isolation.CommandResult, error) {
	results, err := r.environment.RunLegacyCommandsInCopies(source, commands, taskCommandTimeout)
	r.metrics.LegacyCommands += len(results)
	r.metrics.Workspaces += len(results)
	r.metrics.ChildGoRuns += len(results)
	return results, err
}

func (r *realExecutionRunner) Metrics() ExecutionMetrics { return r.metrics }

func (r *realExecutionRunner) Close() error {
	if r == nil || r.environment == nil {
		return nil
	}
	return r.environment.Close()
}

func closeWithPrimary(primary error, closeFn func() error) error {
	closeErr := closeFn()
	if primary == nil {
		return closeErr
	}
	if closeErr == nil {
		return primary
	}
	return errors.Join(primary, fmt.Errorf("close task-check session: %w", closeErr))
}

func assemblyFailure(label string, err error, root, workspace string, taskPaths ...string) variantResult {
	return variantResult{
		FailureKind:  variantFailureAssembly,
		FailurePhase: variantPhaseAssembly,
		Diagnostic:   sanitizeVariantDiagnostic(nil, fmt.Errorf("%s: %w", label, err), root, workspace, "", taskPaths...),
	}
}

func classifyVariantFailure(race, compiled bool, output []byte, runErr error) (variantFailureKind, variantFailurePhase) {
	if errors.Is(runErr, isolation.ErrCommandTimeout) {
		return variantFailureTimeout, variantPhaseInfrastructure
	}
	lower := strings.ToLower(string(output) + "\n" + runErr.Error())
	if race && (strings.Contains(lower, "-race requires cgo") ||
		strings.Contains(lower, "cgo: c compiler") ||
		strings.Contains(lower, "c compiler") && strings.Contains(lower, "not found") ||
		strings.Contains(lower, "cannot use -race")) {
		return variantFailureRaceCGO, variantPhaseCompile
	}
	if !compiled {
		return variantFailureCompile, variantPhaseCompile
	}
	return variantFailureTest, variantPhaseTest
}

func sanitizeVariantDiagnostic(output []byte, runErr error, root, workspace, taskID string, taskPaths ...string) string {
	var lines []string
	scanner := bufio.NewScanner(bytes.NewReader(output))
	for scanner.Scan() {
		line := scanner.Text()
		var event struct {
			Output string `json:"Output"`
		}
		if json.Unmarshal([]byte(line), &event) == nil {
			if strings.TrimSpace(event.Output) != "" {
				lines = append(lines, strings.TrimSpace(event.Output))
			}
			continue
		}
		if strings.TrimSpace(line) != "" {
			lines = append(lines, strings.TrimSpace(line))
		}
	}
	if len(lines) == 0 && runErr != nil {
		lines = append(lines, runErr.Error())
	}
	diagnostic := strings.Join(lines, "\n")
	for _, replacement := range []struct{ old, new string }{
		{workspace, "<workspace>"},
		{root, "<isolation-root>"},
		{filepath.ToSlash(workspace), "<workspace>"},
		{filepath.ToSlash(root), "<isolation-root>"},
	} {
		if replacement.old != "" {
			diagnostic = strings.ReplaceAll(diagnostic, replacement.old, replacement.new)
		}
	}
	if taskID != "" {
		diagnostic = strings.ReplaceAll(diagnostic, "task.local/"+taskID, "task.local/<task>")
	}
	for _, taskPath := range taskPaths {
		if taskPath != "" {
			diagnostic = strings.ReplaceAll(diagnostic, taskPath, "<task-file>")
			diagnostic = strings.ReplaceAll(diagnostic, filepath.ToSlash(taskPath), "<task-file>")
		}
	}
	diagnostic = strings.Map(func(character rune) rune {
		if character == '\n' || character == '\t' || !unicode.IsControl(character) {
			return character
		}
		return -1
	}, diagnostic)
	runes := []rune(strings.TrimSpace(diagnostic))
	if len(runes) > maxVariantDiagnosticRunes {
		diagnostic = string(runes[:maxVariantDiagnosticRunes-1]) + "…"
	}
	return diagnostic
}

func describeVariantFailure(label string, result variantResult) string {
	detail := fmt.Sprintf("%s failed during %s (%s)", label, result.FailurePhase, result.FailureKind)
	if result.Diagnostic != "" {
		detail += ": " + result.Diagnostic
	}
	return detail
}
