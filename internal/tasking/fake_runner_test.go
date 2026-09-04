package tasking

import (
	"os"
	"sort"
	"strings"

	"go-practice-bank/internal/isolation"
)

// fakeExecutionRunner keeps ordinary unit tests process-free. The real Go
// subprocess behavior is exercised only by taskintegration-tagged tests.
type fakeExecutionRunner struct {
	metrics      ExecutionMetrics
	workspaceIDs []string
	cacheIDs     []string
	closed       bool
}

func newFakeExecutionRunner() *fakeExecutionRunner {
	return &fakeExecutionRunner{metrics: ExecutionMetrics{Sessions: 1}}
}

func (r *fakeExecutionRunner) RunVariant(request variantRequest) variantResult {
	r.metrics.VariantRuns++
	r.metrics.Workspaces++
	r.metrics.ChildGoRuns++
	r.workspaceIDs = append(r.workspaceIDs, request.TaskID+"-"+string(rune('a'+len(r.workspaceIDs))))
	r.cacheIDs = append(r.cacheIDs, "shared-session-cache")
	implementation := readTestFiles(request.Implementation)
	tests := readTestFiles(request.Tests)
	if strings.Contains(implementation, "synthetic compile failure") {
		return variantResult{}
	}
	meaningful := strings.Contains(tests, "subject.Identity")
	passed := true
	if meaningful {
		passed = strings.Contains(implementation, "return value") &&
			!strings.Contains(implementation, "return value + 1") &&
			!strings.Contains(implementation, "return 0")
	}
	return variantResult{Compiled: true, Passed: passed, Canonical: "run|task/subject|Test\npass|task/subject|Test"}
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
