package learner

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"go-practice-bank/internal/tasking"
)

var (
	ErrOperationBusy = errors.New("another learner operation is running or an operation lock needs manual recovery")
	ErrStaleView     = errors.New("learner state changed; refresh before trying again")
)

// A lease spans the complete operation, including verification. The smaller
// progress lock remains inside this lease; read-only requests take neither.
func (s *Service) acquireOperation() (func() error, error) {
	path := filepath.Join(s.home, ".operation.lock")
	file, err := os.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, 0600)
	if err != nil {
		if errors.Is(err, os.ErrExist) {
			return nil, fmt.Errorf("%w: %s (see docs/LEARNER_CLI.md)", ErrOperationBusy, path)
		}
		return nil, err
	}
	info, statErr := file.Stat()
	_, writeErr := fmt.Fprintln(file, os.Getpid())
	syncErr := file.Sync()
	release := func() error {
		current, err := os.Lstat(path)
		if err != nil {
			return errors.Join(err, file.Close())
		}
		if info == nil || !current.Mode().IsRegular() || !os.SameFile(info, current) {
			return errors.Join(errors.New("operation lock changed; refusing to remove it"), file.Close())
		}
		// Keep the descriptor open until identity is checked so an unlinked
		// original inode cannot be recycled as somebody else's lock.
		if err := file.Close(); err != nil {
			return err
		}
		return os.Remove(path)
	}
	if err := errors.Join(statErr, writeErr, syncErr); err != nil {
		return nil, errors.Join(err, release())
	}
	return release, nil
}

// Include immutable workspace markers so reset/start of the same ID cannot
// recreate an old revision, even when both operations share a clock tick.
// Missing legacy generations are supported without rewriting existing files.
func (s *Service) stateRevision(progress Progress) (string, error) {
	markers := map[string]string{}
	if len(progress.Tasks) > 0 {
		root, err := os.OpenRoot(s.home)
		if err != nil {
			return "", err
		}
		defer root.Close()
		for id := range progress.Tasks {
			markers[id] = "unavailable"
			if !learnerTaskIDPattern.MatchString(id) {
				continue
			}
			dir := filepath.Join("workspaces", id)
			path := filepath.Join(dir, learnerWorkspaceMarker)
			valid := true
			for _, name := range []string{"workspaces", dir, path} {
				info, err := root.Lstat(name)
				if err != nil || info.Mode()&os.ModeSymlink != 0 || (name != path && !info.IsDir()) || (name == path && (!info.Mode().IsRegular() || info.Size() > 64<<10)) {
					valid = false
					break
				}
			}
			if !valid {
				continue
			}
			file, err := root.Open(path)
			if err != nil {
				continue
			}
			content, readErr := io.ReadAll(io.LimitReader(file, (64<<10)+1))
			closeErr := file.Close()
			if readErr == nil && closeErr == nil && len(content) <= 64<<10 {
				sum := sha256.Sum256(content)
				markers[id] = hex.EncodeToString(sum[:])
			}
		}
	}
	data, err := json.Marshal(struct {
		Progress Progress
		Markers  map[string]string
	}{progress, markers})
	if err != nil {
		return "", err
	}
	sum := sha256.Sum256(data)
	return hex.EncodeToString(sum[:]), nil
}

func withOperation[T any](s *Service, taskID, revision string, requireCurrent bool, action func() (T, error)) (value T, err error) {
	release, err := s.acquireOperation()
	if err != nil {
		return value, err
	}
	defer func() { err = errors.Join(err, release()) }()
	if revision != "" || requireCurrent {
		progress, loadErr := loadProgress(s.home)
		if loadErr != nil {
			return value, loadErr
		}
		currentRevision, revisionErr := s.stateRevision(progress)
		if revisionErr != nil {
			return value, revisionErr
		}
		if (revision != "" && currentRevision != revision) || (requireCurrent && progress.CurrentTaskID != taskID) {
			return value, ErrStaleView
		}
	}
	return action()
}

func (s *Service) Start(id string) (string, error) {
	return withOperation(s, "", "", false, func() (string, error) { return s.start(id) })
}
func (s *Service) Resume(id string) (string, error) {
	return withOperation(s, "", "", false, func() (string, error) { return s.resume(id) })
}
func (s *Service) Check() (tasking.LearnerVerificationResult, error) {
	return withOperation(s, "", "", false, s.check)
}
func (s *Service) Complete() (tasking.LearnerVerificationResult, error) {
	return withOperation(s, "", "", false, s.complete)
}
func (s *Service) Pause() error {
	_, err := withOperation(s, "", "", false, func() (struct{}, error) { return struct{}{}, s.pause() })
	return err
}
func (s *Service) Feedback(text string) error {
	_, err := withOperation(s, "", "", false, func() (struct{}, error) { return struct{}{}, s.feedback(text) })
	return err
}
func (s *Service) ResetTask(id string, confirmed bool) error {
	_, err := withOperation(s, "", "", false, func() (struct{}, error) { return struct{}{}, s.resetTask(id, confirmed) })
	return err
}
func (s *Service) Solution(reveal bool) ([]tasking.LearnerWorkspaceFile, error) {
	return withOperation(s, "", "", false, func() ([]tasking.LearnerWorkspaceFile, error) { return s.solution(reveal) })
}
func (s *Service) ExportReviewContext() (string, error) {
	return withOperation(s, "", "", false, s.exportReviewContext)
}

type UsedHint struct {
	Number int    `json:"number"`
	Text   string `json:"text"`
}

func (s *Service) Hint() (int, string, error) {
	hint, err := withOperation(s, "", "", false, func() (UsedHint, error) { n, text, err := s.hint(); return UsedHint{n, text}, err })
	return hint.Number, hint.Text, err
}

type ActionResult struct {
	Workspace    string                             `json:"workspace,omitempty"`
	Verification *tasking.LearnerVerificationResult `json:"verification,omitempty"`
	Hint         *UsedHint                          `json:"hint,omitempty"`
	Message      string                             `json:"message,omitempty"`
}

// PerformAction shares the exact implementation used by CLI operations. A UI
// action names its intended task and the persisted snapshot the user saw.
func (s *Service) PerformAction(action, id, revision string) (ActionResult, error) {
	if !learnerTaskIDPattern.MatchString(id) || len(revision) != 64 {
		return ActionResult{}, ErrStaleView
	}
	switch action {
	case "start", "resume", "check", "complete", "hint", "pause", "open":
	default:
		return ActionResult{}, errors.New("unsupported learner action")
	}
	requireCurrent := action == "check" || action == "complete" || action == "hint" || action == "pause"
	return withOperation(s, id, revision, requireCurrent, func() (ActionResult, error) {
		switch action {
		case "start":
			path, err := s.start(id)
			return ActionResult{Workspace: path}, err
		case "resume":
			path, err := s.resume(id)
			return ActionResult{Workspace: path}, err
		case "check":
			result, err := s.check()
			return ActionResult{Verification: &result}, err
		case "complete":
			result, err := s.complete()
			return ActionResult{Verification: &result}, err
		case "hint":
			n, text, err := s.hint()
			return ActionResult{Hint: &UsedHint{n, text}}, err
		case "pause":
			return ActionResult{Message: "Work session paused."}, s.pause()
		case "open":
			progress, err := loadProgress(s.home)
			if err != nil {
				return ActionResult{}, err
			}
			task := progress.Tasks[id]
			if task == nil {
				return ActionResult{}, ErrStaleView
			}
			info, err := s.Show(id)
			if err != nil {
				return ActionResult{}, err
			}
			workspace, err := s.validateWorkspace(task, info)
			if err != nil {
				return ActionResult{}, err
			}
			if err := s.launch(workspace); err != nil {
				return ActionResult{}, err
			}
			return ActionResult{Workspace: workspace, Message: "Workspace sent to GoLand."}, nil
		}
		return ActionResult{}, errors.New("unsupported learner action")
	})
}
