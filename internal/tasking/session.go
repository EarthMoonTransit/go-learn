package tasking

import (
	"errors"
	"fmt"

	"go-practice-bank/internal/toolchainpolicy"
)

// CheckSession owns the execution resources for one explicit verification pass.
// It is intentionally not cached or stored in package-global state.
type CheckSession struct {
	root   string
	policy toolchainpolicy.Policy
	runner executionRunner
	closed bool
}

func NewCheckSession(root string) (*CheckSession, error) {
	runner, err := newRealExecutionRunner()
	if err != nil {
		return nil, err
	}
	session, err := newCheckSessionWithRunner(root, runner)
	if err != nil {
		_ = runner.Close()
		return nil, err
	}
	return session, nil
}

func newCheckSessionWithRunner(root string, runner executionRunner) (*CheckSession, error) {
	if runner == nil {
		return nil, errors.New("task-check runner is required")
	}
	policy, err := toolchainpolicy.Load(root)
	if err != nil {
		return nil, err
	}
	return &CheckSession{root: root, policy: policy, runner: runner}, nil
}

func (s *CheckSession) Close() error {
	if s == nil || s.closed {
		return nil
	}
	s.closed = true
	return s.runner.Close()
}

func (s *CheckSession) Check(id string) (QualityReport, error) {
	if s == nil || s.closed {
		return QualityReport{}, errors.New("task-check session is closed")
	}
	located, err := FindTask(s.root, id)
	if err != nil {
		return QualityReport{}, err
	}
	return s.checkLocated(located)
}

func (s *CheckSession) CheckAll(states []State) (CheckCollection, error) {
	if s == nil || s.closed {
		return CheckCollection{}, errors.New("task-check session is closed")
	}
	located, err := DiscoverTasks(s.root, states)
	if err != nil {
		return CheckCollection{}, err
	}
	beforeExecution := s.runner.Metrics()
	actualRuntime, err := s.runner.ActualRuntime()
	if err != nil {
		return CheckCollection{}, fmt.Errorf("query actual execution runtime: %w", err)
	}
	toolchainDigest, err := ToolchainPolicyDigest(s.root)
	if err != nil {
		return CheckCollection{}, err
	}
	collection := CheckCollection{
		SchemaVersion:         1,
		ContractVersion:       CheckContractVersion,
		ToolchainPolicyDigest: toolchainDigest,
		VerifierSourceDigest:  VerifierSourceDigest(),
		Runtime:               actualRuntime,
		Tasks:                 make([]CheckedTask, 0, len(located)),
	}
	for _, task := range located {
		before, err := TaskInputDigest(s.root, task)
		if err != nil {
			return CheckCollection{}, fmt.Errorf("digest %s: %w", task.ID, err)
		}
		report, err := s.checkLocated(task)
		if err != nil {
			return CheckCollection{}, fmt.Errorf("check %s: %w", task.ID, err)
		}
		after, err := TaskInputDigest(s.root, task)
		if err != nil {
			return CheckCollection{}, fmt.Errorf("redigest %s: %w", task.ID, err)
		}
		if before != after {
			return CheckCollection{}, fmt.Errorf("task %s changed during verification", task.ID)
		}
		collection.Tasks = append(collection.Tasks, CheckedTask{
			TaskID: task.ID, State: task.State, InputDigest: before, Report: report,
		})
	}
	collection.Execution = executionDelta(beforeExecution, s.runner.Metrics())
	collection.Execution.Sessions = 1
	collection.Execution.TaskChecks = len(located)
	actualRuntime, err = s.runner.ActualRuntime()
	if err != nil {
		return CheckCollection{}, fmt.Errorf("confirm actual execution runtime: %w", err)
	}
	if err := validateCheckCollectionWithRuntime(s.root, collection, states, &actualRuntime); err != nil {
		return CheckCollection{}, fmt.Errorf("generated check collection is invalid: %w", err)
	}
	return collection, nil
}

func executionDelta(before, after ExecutionMetrics) ExecutionMetrics {
	return ExecutionMetrics{
		TaskChecks:     after.TaskChecks - before.TaskChecks,
		Workspaces:     after.Workspaces - before.Workspaces,
		VariantRuns:    after.VariantRuns - before.VariantRuns,
		LegacyCommands: after.LegacyCommands - before.LegacyCommands,
		ChildGoRuns:    after.ChildGoRuns - before.ChildGoRuns,
		RuntimeQueries: after.RuntimeQueries - before.RuntimeQueries,
		FuzzRuns:       after.FuzzRuns - before.FuzzRuns,
		FuzzGoRuns:     after.FuzzGoRuns - before.FuzzGoRuns,
		FuzzWorkspaces: after.FuzzWorkspaces - before.FuzzWorkspaces,
		FuzzHelperRuns: after.FuzzHelperRuns - before.FuzzHelperRuns,
		FuzzNativeRuns: after.FuzzNativeRuns - before.FuzzNativeRuns,
	}
}

func Check(root, id string) (report QualityReport, err error) {
	return checkWithRunnerFactory(root, id, func() (executionRunner, error) {
		return newRealExecutionRunner()
	})
}

func checkWithRunnerFactory(root, id string, factory func() (executionRunner, error)) (report QualityReport, err error) {
	runner, err := factory()
	if err != nil {
		return QualityReport{}, err
	}
	session, err := newCheckSessionWithRunner(root, runner)
	if err != nil {
		_ = runner.Close()
		return QualityReport{}, err
	}
	defer func() { err = closeWithPrimary(err, session.Close) }()
	return session.Check(id)
}

func CheckAll(root string, states []State) (collection CheckCollection, err error) {
	session, err := NewCheckSession(root)
	if err != nil {
		return CheckCollection{}, err
	}
	defer func() { err = closeWithPrimary(err, session.Close) }()
	return session.CheckAll(states)
}
