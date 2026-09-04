package tasking

import (
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"go-practice-bank/internal/isolation"
	"go-practice-bank/internal/safeio"
	"gopkg.in/yaml.v3"
)

type PromoteOptions struct {
	ID     string
	To     State
	Actor  string
	Reason string
	Date   string
	runner executionRunner
	hooks  promotionHooks
}

type promotionHooks struct {
	afterLock   func()
	beforeWrite func(string) error
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

func Promote(root string, options PromoteOptions) (located LocatedTask, err error) {
	var session *CheckSession
	if options.runner == nil {
		session, err = NewCheckSession(root)
	} else {
		session, err = newCheckSessionWithRunner(root, options.runner)
	}
	if err != nil {
		return LocatedTask{}, err
	}
	defer func() { err = closeWithPrimary(err, session.Close) }()
	return promoteWithSession(root, options, session)
}

func promoteWithSession(root string, options PromoteOptions, session *CheckSession) (LocatedTask, error) {
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
	if !taskIDPattern.MatchString(options.ID) {
		return LocatedTask{}, fmt.Errorf("invalid task id %q", options.ID)
	}
	unlock, err := acquirePromotionLock(root, options.ID)
	if err != nil {
		return LocatedTask{}, err
	}
	defer unlock()
	if options.hooks.afterLock != nil {
		options.hooks.afterLock()
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
	lastDate := data.Lifecycle.Transitions[len(data.Lifecycle.Transitions)-1].At
	if options.Date < lastDate {
		return LocatedTask{}, fmt.Errorf("transition date %s predates last lifecycle event %s", options.Date, lastDate)
	}
	if options.To == StateRejected && data.Review.Decision != "reject" {
		return LocatedTask{}, errors.New("rejection requires review decision reject")
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
	if err := ensureRealDirectory(filepath.Dir(destination), 0o755); err != nil {
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
	if err := validateSchema(root, "task.schema.json", data.Task); err != nil {
		return LocatedTask{}, fmt.Errorf("future task metadata: %w", err)
	}
	if err := validateSchema(root, "task-lifecycle.schema.json", data.Lifecycle); err != nil {
		return LocatedTask{}, fmt.Errorf("future lifecycle metadata: %w", err)
	}
	if problems := validateLifecycle(data.Lifecycle); len(problems) != 0 {
		return LocatedTask{}, fmt.Errorf("future lifecycle is invalid: %s", problems[0])
	}
	if err := validateFuturePromotion(session, located, options.To, updatedTask, updatedLifecycle); err != nil {
		return LocatedTask{}, err
	}

	rollback := func() error {
		if _, err := safeio.AtomicWrite(located.Path, "task.yaml", originalTask, 1<<20); err != nil {
			return err
		}
		_, err := safeio.AtomicWrite(located.Path, "lifecycle.yaml", originalLifecycle, 1<<20)
		return err
	}
	if options.hooks.beforeWrite != nil {
		if err := options.hooks.beforeWrite("task.yaml"); err != nil {
			return LocatedTask{}, err
		}
	}
	if _, err := safeio.AtomicWrite(located.Path, "task.yaml", updatedTask, 1<<20); err != nil {
		return LocatedTask{}, err
	}
	if options.hooks.beforeWrite != nil {
		if err := options.hooks.beforeWrite("lifecycle.yaml"); err != nil {
			if rollbackErr := rollback(); rollbackErr != nil {
				return LocatedTask{}, fmt.Errorf("injected promotion failure: %v; rollback failed: %v", err, rollbackErr)
			}
			return LocatedTask{}, err
		}
	}
	if _, err := safeio.AtomicWrite(located.Path, "lifecycle.yaml", updatedLifecycle, 1<<20); err != nil {
		_ = rollback()
		return LocatedTask{}, err
	}
	if err := os.Rename(located.Path, destination); err != nil {
		if rollbackErr := rollback(); rollbackErr != nil {
			return LocatedTask{}, fmt.Errorf("atomic lifecycle rename: %v; rollback failed: %v", err, rollbackErr)
		}
		return LocatedTask{}, fmt.Errorf("atomic lifecycle rename: %w", err)
	}
	return LocatedTask{ID: options.ID, State: options.To, Path: destination}, nil
}

func acquirePromotionLock(root, id string) (func(), error) {
	directory := filepath.Join(root, "tasks", ".locks")
	if err := ensureRealDirectory(directory, 0o700); err != nil {
		return nil, err
	}
	path := filepath.Join(directory, id+".lock")
	file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_EXCL, 0o600)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("promotion for %s is already locked", id)
		}
		return nil, err
	}
	if err := file.Close(); err != nil {
		_ = os.Remove(path)
		return nil, err
	}
	return func() { _ = os.Remove(path) }, nil
}

func ensureRealDirectory(path string, mode os.FileMode) error {
	info, err := os.Lstat(path)
	if errors.Is(err, os.ErrNotExist) {
		if err := os.Mkdir(path, mode); err != nil && !errors.Is(err, os.ErrExist) {
			return err
		}
		info, err = os.Lstat(path)
	}
	if err != nil {
		return err
	}
	if info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
		return fmt.Errorf("promotion directory must be a real directory: %s", path)
	}
	return nil
}

func validateFuturePromotion(session *CheckSession, source LocatedTask, target State, taskYAML, lifecycleYAML []byte) error {
	staging, err := os.MkdirTemp("", "go-task-promotion-*")
	if err != nil {
		return err
	}
	defer os.RemoveAll(staging)
	packageCopy := filepath.Join(staging, source.ID)
	if err := isolation.CopyTree(source.Path, packageCopy); err != nil {
		return err
	}
	if _, err := safeio.AtomicWrite(packageCopy, "task.yaml", taskYAML, 1<<20); err != nil {
		return err
	}
	if _, err := safeio.AtomicWrite(packageCopy, "lifecycle.yaml", lifecycleYAML, 1<<20); err != nil {
		return err
	}
	report, err := session.checkLocated(LocatedTask{ID: source.ID, State: target, Path: packageCopy})
	if err != nil {
		return err
	}
	switch target {
	case StateReady:
		if !report.Passed {
			return fmt.Errorf("future ready package fails gates: %v", report.FailingGateIDs)
		}
	case StateDraft, StateInReview:
		if report.TechnicalStatus != "pass" {
			return fmt.Errorf("future %s package fails technical gates: %v", target, report.FailingGateIDs)
		}
	case StateRejected, StateRetired:
		if report.TechnicalStatus != "pass" {
			return fmt.Errorf("future %s package fails audit gates: %v", target, report.FailingGateIDs)
		}
	}
	return nil
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
