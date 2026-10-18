package tasking

import (
	"fmt"
	"path/filepath"

	"go-practice-bank/internal/fuzznative"
	"go-practice-bank/internal/fuzznative/frame"
	"go-practice-bank/internal/isolation"
)

type fuzzRequest struct {
	variantRequest
	ModulePath string
	Plan       FuzzPlan
}

func (r *realExecutionRunner) ActualRuntime() (isolation.RuntimeInfo, error) {
	if r.fuzz == nil {
		r.fuzz = fuzznative.NewEngine(r.environment)
	}
	before := r.fuzz.Metrics()
	info, err := r.fuzz.ActualRuntime(taskCommandTimeout)
	after := r.fuzz.Metrics()
	r.metrics.Workspaces += after.Workspaces - before.Workspaces
	r.metrics.ChildGoRuns += after.GoRuns - before.GoRuns
	r.metrics.RuntimeQueries += after.GoRuns - before.GoRuns
	return info, err
}

func (r *realExecutionRunner) RunFuzz(request fuzzRequest) fuzznative.Result {
	r.metrics.FuzzRuns++
	fail := func(err error) fuzznative.Result {
		return fuzznative.Result{Phase: "assembly", Diagnostic: sanitizeVariantDiagnostic(nil, err, r.environment.Root, "", request.TaskID, append(append([]string{}, request.Implementation...), request.Tests...)...)}
	}
	if len(request.Tests) < 1 || len(request.Tests) > frame.MaxFiles {
		return fail(fmt.Errorf("frame.resources: require 1..8 submitted test files"))
	}
	read := func(paths []string, limit int64) (map[string]string, error) {
		files := map[string]string{}
		for _, path := range paths {
			name := filepath.Base(path)
			if _, exists := files[name]; exists {
				return nil, fmt.Errorf("duplicate native source filename %s", name)
			}
			data, err := readRegularFile(path, limit)
			if err != nil {
				return nil, err
			}
			files[name] = string(data)
		}
		return files, nil
	}
	// Keep original basenames. Prefixing would turn Go-ignored _hidden files
	// into selected files and hide the finite grammar's file-selection rule.
	tests, err := read(request.Tests, frame.MaxSourceBytes)
	if err != nil {
		return fail(err)
	}
	implementation, err := read(request.Implementation, 8<<20)
	if err != nil {
		return fail(err)
	}
	if r.fuzz == nil {
		r.fuzz = fuzznative.NewEngine(r.environment)
	}
	before := r.fuzz.Metrics()
	result := r.fuzz.Verify(fuzznative.Request{
		Identity:  fuzznative.Identity{ModulePath: request.ModulePath, Grammar: request.Plan.Grammar, Target: request.Plan.Target},
		GoVersion: request.GoVersion, Implementation: implementation, Sources: tests, Timeout: variantTimeout(request.TimeoutSeconds), Race: request.Race, Vet: request.Vet,
	})
	after := r.fuzz.Metrics()
	r.metrics.Workspaces += after.Workspaces - before.Workspaces
	r.metrics.ChildGoRuns += after.GoRuns - before.GoRuns
	r.metrics.FuzzGoRuns += after.GoRuns - before.GoRuns
	r.metrics.FuzzWorkspaces += after.Workspaces - before.Workspaces
	r.metrics.FuzzHelperRuns += after.HelperRuns - before.HelperRuns
	r.metrics.FuzzNativeRuns += after.NativeRuns - before.NativeRuns
	return result
}
