package tasking

import (
	"os"
	"runtime"
	"sort"
	"strings"

	"go-practice-bank/internal/fuzznative"
	"go-practice-bank/internal/isolation"
)

// fakeExecutionRunner never executes authored verification in ordinary unit
// tests. Real task subprocess behavior uses taskintegration-tagged tests;
// collection validation may query only the trusted Go runtime metadata.
type fakeExecutionRunner struct {
	metrics         ExecutionMetrics
	workspaceIDs    []string
	cacheIDs        []string
	raceUnsupported bool
	closed          bool
	runtimeRead     bool
}

func (r *fakeExecutionRunner) ActualRuntime() (isolation.RuntimeInfo, error) {
	if !r.runtimeRead {
		r.runtimeRead = true
		r.metrics.Workspaces++
		r.metrics.ChildGoRuns++
		r.metrics.RuntimeQueries++
	}
	return isolation.RuntimeInfo{Version: runtime.Version(), GOOS: runtime.GOOS, GOARCH: runtime.GOARCH}, nil
}

func newFakeExecutionRunner() *fakeExecutionRunner {
	return &fakeExecutionRunner{metrics: ExecutionMetrics{Sessions: 1}}
}

func (r *fakeExecutionRunner) RunFuzz(fuzzRequest) fuzznative.Result {
	r.metrics.FuzzRuns++
	return fuzznative.Result{Phase: "infrastructure", Diagnostic: "synthetic fuzz verification requires an explicit scripted outcome"}
}

func (r *fakeExecutionRunner) RunVariant(request variantRequest) variantResult {
	if request.Race && r.raceUnsupported {
		return variantResult{
			FailureKind: variantFailureRaceUnsupported, FailurePhase: variantPhaseInfrastructure,
			Diagnostic: "race detector is unsupported on synthetic/architecture",
		}
	}
	r.metrics.VariantRuns++
	r.metrics.Workspaces++
	r.metrics.ChildGoRuns++
	r.workspaceIDs = append(r.workspaceIDs, request.TaskID+"-"+string(rune('a'+len(r.workspaceIDs))))
	r.cacheIDs = append(r.cacheIDs, "shared-session-cache")
	implementation := readTestFiles(request.Implementation)
	tests := readTestFiles(request.Tests)
	if strings.Contains(implementation, "synthetic compile failure") {
		return variantResult{FailureKind: variantFailureCompile, FailurePhase: variantPhaseCompile, Diagnostic: "synthetic compile failure"}
	}
	meaningful := strings.Contains(tests, "subject.Identity")
	passed := true
	if meaningful {
		passed = strings.Contains(implementation, "return value") &&
			!strings.Contains(implementation, "return value + 1") &&
			!strings.Contains(implementation, "return 0")
		if strings.Contains(tests, "Identity(0)") && strings.Contains(implementation, "return 0") {
			passed = true
		}
	}
	result := variantResult{Compiled: true, Passed: passed, Canonical: "run|task/subject|Test\npass|task/subject|Test", FailureKind: variantFailureNone, FailurePhase: variantPhaseNone}
	if !passed {
		result.FailureKind = variantFailureTest
		result.FailurePhase = variantPhaseTest
		result.Diagnostic = "synthetic test failure"
	}
	return result
}

func (r *fakeExecutionRunner) RunWorkspace(request workspaceRequest) variantResult {
	if request.Race && r.raceUnsupported {
		return variantResult{FailureKind: variantFailureRaceUnsupported, FailurePhase: variantPhaseInfrastructure, Diagnostic: "race detector is unsupported on synthetic/architecture"}
	}
	r.metrics.VariantRuns++
	r.metrics.Workspaces++
	r.metrics.ChildGoRuns++
	r.workspaceIDs = append(r.workspaceIDs, request.TaskID+"-workspace")
	r.cacheIDs = append(r.cacheIDs, "shared-session-cache")
	var implementationPaths, testPaths []string
	for _, file := range append(append([]mappedFile{}, request.Skeleton...), request.Overlay...) {
		implementationPaths = append(implementationPaths, file.Source)
	}
	for _, file := range request.Tests {
		testPaths = append(testPaths, file.Source)
	}
	implementation := readTestFiles(implementationPaths)
	tests := readTestFiles(testPaths)
	meaningful := strings.Contains(tests, "subject.Identity") || strings.Contains(tests, "adapter.Transform")
	passed := request.Vet || !meaningful || strings.Contains(implementation, "return value") && !strings.Contains(implementation, "return value + 1") && !strings.Contains(implementation, "return 0")
	result := variantResult{Compiled: true, Passed: passed, Canonical: "run|task/subject|Test\npass|task/subject|Test", FailureKind: variantFailureNone, FailurePhase: variantPhaseNone}
	if !passed {
		result.FailureKind, result.FailurePhase, result.Diagnostic = variantFailureTest, variantPhaseTest, "synthetic test failure"
	}
	return result
}

func (r *fakeExecutionRunner) RunLegacy(_ string, commands []string) ([]isolation.CommandResult, error) {
	r.metrics.LegacyCommands += len(commands)
	r.metrics.Workspaces += len(commands)
	r.metrics.ChildGoRuns += len(commands)
	results := make([]isolation.CommandResult, len(commands))
	for index, command := range commands {
		results[index] = isolation.CommandResult{Command: command}
		r.workspaceIDs = append(r.workspaceIDs, "legacy-"+command)
		r.cacheIDs = append(r.cacheIDs, "shared-session-cache")
	}
	return results, nil
}

func (r *fakeExecutionRunner) Metrics() ExecutionMetrics { return r.metrics }
func (r *fakeExecutionRunner) Close() error {
	r.closed = true
	return nil
}

func readTestFiles(paths []string) string {
	paths = append([]string(nil), paths...)
	sort.Strings(paths)
	var output strings.Builder
	for _, path := range paths {
		content, _ := os.ReadFile(path)
		output.Write(content)
	}
	return output.String()
}
