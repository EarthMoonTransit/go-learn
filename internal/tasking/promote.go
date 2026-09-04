package tasking

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go-practice-bank/internal/safeio"
	"gopkg.in/yaml.v3"
)

type PromoteOptions struct {
	ID     string
	To     State
	Actor  string
	Reason string
	Date   string
}

func allowedTransition(from, to State) bool {
	switch from {
	case StateDraft:
		return to == StateInReview
	case StateInReview:
		return to == StateDraft || to == StateReady || to == StateRejected
	case StateRejected:
		return to == StateDraft
	case StateReady:
		return to == StateRetired
	default:
		return false
	}
}

func Promote(root string, options PromoteOptions) (LocatedTask, error) {
	if options.Actor == "" || options.Reason == "" {
		return LocatedTask{}, errors.New("promotion requires actor and reason")
	}
	if len(options.Actor) > 100 || len(options.Reason) > 300 {
		return LocatedTask{}, errors.New("promotion actor/reason exceeds lifecycle schema limits")
	}
	if options.Date == "" {
		options.Date = time.Now().Format("2006-01-02")
	}
	if _, err := time.Parse("2006-01-02", options.Date); err != nil {
		return LocatedTask{}, fmt.Errorf("invalid date: %w", err)
	}
	located, err := FindTask(root, options.ID)
	if err != nil {
		return LocatedTask{}, err
	}
	if !allowedTransition(located.State, options.To) {
		return LocatedTask{}, fmt.Errorf("transition %s -> %s is not allowed", located.State, options.To)
	}

	data, loadProblems := loadPackage(root, located)
	if len(loadProblems) != 0 {
		return LocatedTask{}, fmt.Errorf("cannot promote package with invalid metadata: %s", loadProblems[0])
	}
	switch options.To {
	case StateInReview:
		report, err := checkLocated(root, located)
		if err != nil {
			return LocatedTask{}, err
		}
		if problems := failingGates(report, "review-evidence"); len(problems) != 0 {
			return LocatedTask{}, fmt.Errorf("cannot enter review: failing gates: %v", problems)
		}
	case StateReady:
		report, err := checkLocated(root, located)
		if err != nil {
			return LocatedTask{}, err
		}
		if !report.Passed {
			return LocatedTask{}, fmt.Errorf("cannot enter ready: failing gates: %v", failingGates(report))
		}
	case StateRejected:
		if data.Review.Decision != "reject" {
			return LocatedTask{}, errors.New("rejection requires review decision reject")
		}
	}

	destination, err := packagePath(root, options.To, options.ID)
	if err != nil {
		return LocatedTask{}, err
	}
	if _, err := os.Lstat(destination); err == nil {
		return LocatedTask{}, fmt.Errorf("destination already exists; refusing to overwrite %s", destination)
	} else if !errors.Is(err, os.ErrNotExist) {
		return LocatedTask{}, err
	}
	if err := os.MkdirAll(filepath.Dir(destination), 0o755); err != nil {
		return LocatedTask{}, err
	}

	originalTask, err := os.ReadFile(filepath.Join(located.Path, "task.yaml"))
	if err != nil {
		return LocatedTask{}, err
	}
	originalLifecycle, err := os.ReadFile(filepath.Join(located.Path, "lifecycle.yaml"))
	if err != nil {
		return LocatedTask{}, err
	}
	data.Task.QualityStatus = qualityByState[options.To]
	data.Lifecycle.State = options.To
	data.Lifecycle.UpdatedAt = options.Date
	data.Lifecycle.Transitions = append(data.Lifecycle.Transitions, Transition{
		From: located.State, To: options.To, At: options.Date, Actor: options.Actor, Reason: options.Reason,
	})
	updatedTask, err := yaml.Marshal(data.Task)
	if err != nil {
		return LocatedTask{}, err
	}
	updatedLifecycle, err := yaml.Marshal(data.Lifecycle)
	if err != nil {
		return LocatedTask{}, err
	}
	rollback := func() {
		_, _ = safeio.AtomicWrite(located.Path, "task.yaml", originalTask, 1<<20)
		_, _ = safeio.AtomicWrite(located.Path, "lifecycle.yaml", originalLifecycle, 1<<20)
	}
	if _, err := safeio.AtomicWrite(located.Path, "task.yaml", updatedTask, 1<<20); err != nil {
		return LocatedTask{}, err
	}
	if _, err := safeio.AtomicWrite(located.Path, "lifecycle.yaml", updatedLifecycle, 1<<20); err != nil {
		rollback()
		return LocatedTask{}, err
	}
	if err := os.Rename(located.Path, destination); err != nil {
		rollback()
		return LocatedTask{}, fmt.Errorf("atomic lifecycle rename: %w", err)
	}
	return LocatedTask{ID: options.ID, State: options.To, Path: destination}, nil
}

func failingGates(report QualityReport, ignored ...string) []string {
	ignore := map[string]bool{}
	for _, id := range ignored {
		ignore[id] = true
	}
	var result []string
	for _, gate := range report.Gates {
		if gate.Status == "fail" && !ignore[gate.ID] {
			result = append(result, gate.ID)
		}
	}
	return result
}
