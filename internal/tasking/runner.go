package tasking

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go-practice-bank/internal/isolation"
)

const taskCommandTimeout = 60 * time.Second

type variantRequest struct {
	TaskID         string
	GoVersion      string
	Race           bool
	Implementation []string
	Tests          []string
}

type variantResult struct {
	Compiled  bool
	Passed    bool
	Canonical string
}

type executionRunner interface {
	RunVariant(variantRequest) variantResult
	RunLegacy(string, []string) ([]isolation.CommandResult, error)
	Metrics() ExecutionMetrics
	Close() error
}

type realExecutionRunner struct {
	environment *isolation.Environment
	metrics     ExecutionMetrics
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
	r.metrics.VariantRuns++
	workspace, err := r.environment.NewWorkspace(request.TaskID)
	if err != nil {
		return variantResult{}
	}
	r.metrics.Workspaces++
	goMod := []byte("module task.local/" + request.TaskID + "\n\ngo " + request.GoVersion + "\n")
	if err := os.WriteFile(filepath.Join(workspace, "go.mod"), goMod, 0o600); err != nil {
		return variantResult{}
	}
	subject := filepath.Join(workspace, "subject")
	if err := os.Mkdir(subject, 0o700); err != nil {
		return variantResult{}
	}
	for index, path := range request.Implementation {
		name := fmt.Sprintf("impl_%03d_%s", index, filepath.Base(path))
		if err := isolation.CopyFile(path, filepath.Join(subject, name)); err != nil {
			return variantResult{}
		}
	}
	for index, path := range request.Tests {
		name := fmt.Sprintf("check_%03d_%s", index, filepath.Base(path))
		if err := isolation.CopyFile(path, filepath.Join(subject, name)); err != nil {
			return variantResult{}
		}
	}
	args := []string{"test", "-json", "-count=1"}
	if request.Race {
		args = append(args, "-race")
	}
	args = append(args, "./subject")
	r.metrics.ChildGoRuns++
	output, runErr := r.environment.RunGo(workspace, taskCommandTimeout, args...)
	canonical, compiled := canonicalTestEvents(output, request.TaskID)
	return variantResult{Compiled: compiled, Passed: runErr == nil, Canonical: canonical}
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
