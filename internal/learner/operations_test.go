package learner

import (
	"bytes"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"sync"
	"testing"
	"time"

	"go-practice-bank/internal/tasking"
)

func uiTestService(t *testing.T, home string) *Service {
	t.Helper()
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	if home == "" {
		home = t.TempDir()
	}
	s, err := New(Config{Root: root, Home: home})
	if err != nil {
		t.Fatal(err)
	}
	// Fast tests exercise service policy, never execute authored task code.
	s.verify = func(_, _, _ string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		return tasking.LearnerVerificationResult{Passed: true, Scope: scope, Summary: "Fixture verification passed."}, nil
	}
	s.launch = func(string) error { return ErrIDEUnavailable }
	s.ideAvailable = false
	return s
}
func persistedBytes(t *testing.T, s *Service) []byte {
	t.Helper()
	data, err := os.ReadFile(filepath.Join(s.home, "progress.json"))
	if err != nil {
		t.Fatal(err)
	}
	return data
}

func TestOperationLeaseCoordinatesServicesAndCLI(t *testing.T) {
	first := uiTestService(t, "")
	second := uiTestService(t, first.home)
	workspacePath, err := first.Start("literal-bindings-001")
	if err != nil {
		t.Fatal(err)
	}
	entered, release := make(chan struct{}), make(chan struct{})
	first.verify = func(_, id, workspace string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		if id != "literal-bindings-001" || workspace != workspacePath {
			t.Error("wrong captured workspace")
		}
		close(entered)
		<-release
		return tasking.LearnerVerificationResult{Passed: true, Scope: scope, Summary: "Fixture pass"}, nil
	}
	result := make(chan error, 1)
	go func() { _, err := first.Check(); result <- err }()
	<-entered
	var releaseOnce sync.Once
	unblock := func() { releaseOnce.Do(func() { close(release) }) }
	t.Cleanup(unblock)
	before := persistedBytes(t, first)
	mutations := []struct {
		name string
		run  func() error
	}{
		{"start", func() error { _, err := second.Start("zero-results-001"); return err }},
		{"resume", func() error { _, err := second.Resume("literal-bindings-001"); return err }},
		{"check", func() error { _, err := second.Check(); return err }},
		{"complete", func() error { _, err := second.Complete(); return err }},
		{"hint", func() error { _, _, err := second.Hint(); return err }},
		{"pause", second.Pause},
		{"feedback", func() error { return second.Feedback("fixture") }},
		{"solution", func() error { _, err := second.Solution(true); return err }},
		{"reset", func() error { return second.ResetTask("literal-bindings-001", true) }},
		{"export", func() error { _, err := second.ExportReviewContext(); return err }},
	}
	for _, mutation := range mutations {
		t.Run(mutation.name, func(t *testing.T) {
			if err := mutation.run(); !errors.Is(err, ErrOperationBusy) {
				t.Fatalf("want busy, got %v", err)
			}
		})
	}
	var stdout, stderr bytes.Buffer
	if code := RunCLI(Config{Root: first.root, Home: first.home}, []string{"resume", "literal-bindings-001"}, &stdout, &stderr); code != 2 || !bytes.Contains(stderr.Bytes(), []byte("operation")) {
		t.Fatalf("CLI ignored lease: %d %s", code, stderr.String())
	}
	if _, err := second.CourseView(); err != nil {
		t.Fatal(err)
	}
	if _, err := second.TaskView("zero-results-001"); err != nil {
		t.Fatal(err)
	}
	if _, err := second.Progress(); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(before, persistedBytes(t, first)) {
		t.Fatal("read-only or conflicting operation changed progress")
	}
	unblock()
	if err := <-result; err != nil {
		t.Fatal(err)
	}
	progress, err := second.Progress()
	if err != nil {
		t.Fatal(err)
	}
	if progress.CurrentTaskID != "literal-bindings-001" || progress.Tasks[progress.CurrentTaskID].Attempts != 1 {
		t.Fatalf("wrong applied result: %+v", progress)
	}
	if _, err := os.Lstat(filepath.Join(first.home, ".operation.lock")); !errors.Is(err, os.ErrNotExist) {
		t.Fatalf("lease not released: %v", err)
	}
}

func TestUIActionRejectsStaleStateWithoutChangingSelection(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	view, err := s.TaskView("literal-bindings-001")
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.Start("zero-results-001"); err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	if _, err := s.PerformAction("check", "literal-bindings-001", view.Revision); !errors.Is(err, ErrStaleView) {
		t.Fatalf("stale view accepted: %v", err)
	}
	course, err := s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	if _, err := s.PerformAction("check", "literal-bindings-001", course.Revision); !errors.Is(err, ErrStaleView) {
		t.Fatalf("wrong current task accepted: %v", err)
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("stale action mutated progress")
	}
}

func TestHintsAreChosenWithinSharedOperation(t *testing.T) {
	a := uiTestService(t, "")
	b := uiTestService(t, a.home)
	if _, err := a.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	start := make(chan struct{})
	results := make(chan int, 2)
	failures := make(chan error, 2)
	for _, s := range []*Service{a, b} {
		go func() {
			<-start
			n, _, err := s.Hint()
			if err != nil && !errors.Is(err, ErrOperationBusy) {
				failures <- err
			}
			results <- n
		}()
	}
	close(start)
	seen := map[int]bool{}
	for i := 0; i < 2; i++ {
		n := <-results
		if n > 0 {
			if seen[n] {
				t.Fatal("duplicate hint selection")
			}
			seen[n] = true
		}
	}
	close(failures)
	for err := range failures {
		t.Fatal(err)
	}
	next, _, err := a.Hint()
	if err != nil {
		t.Fatal(err)
	}
	if seen[next] {
		t.Fatal("repeated previously accepted hint")
	}
}

func TestUnknownAndReplacedOperationLocksArePreserved(t *testing.T) {
	s := uiTestService(t, "")
	path := filepath.Join(s.home, ".operation.lock")
	if err := os.WriteFile(path, []byte("unknown owner"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Start("literal-bindings-001"); !errors.Is(err, ErrOperationBusy) {
		t.Fatalf("unknown lock bypassed: %v", err)
	}
	data, err := os.ReadFile(path)
	if err != nil || string(data) != "unknown owner" {
		t.Fatal("unknown lock changed")
	}
	if err := os.Remove(path); err != nil {
		t.Fatal(err)
	} // This test owns the synthetic lock.
	release, err := s.acquireOperation()
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Rename(path, path+".original"); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, []byte("replacement"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := release(); err == nil {
		t.Fatal("replaced lock silently removed")
	}
	data, err = os.ReadFile(path)
	if err != nil || string(data) != "replacement" {
		t.Fatal("foreign replacement was removed")
	}
}

func TestOperationLeaseReleasedAfterVerifierError(t *testing.T) {
	s := uiTestService(t, "")
	if _, err := s.Start("literal-bindings-001"); err != nil {
		t.Fatal(err)
	}
	s.verify = func(string, string, string, tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		return tasking.LearnerVerificationResult{}, errors.New("fixture verifier failure")
	}
	if _, err := s.Check(); err == nil {
		t.Fatal("missing verifier error")
	}
	if err := s.Pause(); err != nil {
		t.Fatalf("failed verification held lease: %v", err)
	}
}

func TestSameClockWorkspaceRecreationInvalidatesUIRevision(t *testing.T) {
	s := uiTestService(t, "")
	s.now = func() time.Time { return time.Date(2026, 9, 5, 12, 0, 0, 0, time.UTC) }
	id := "literal-bindings-001"
	if _, err := s.Start(id); err != nil {
		t.Fatal(err)
	}
	before := persistedBytes(t, s)
	view, err := s.TaskView(id)
	if err != nil {
		t.Fatal(err)
	}
	if err := s.ResetTask(id, true); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Start(id); err != nil {
		t.Fatal(err)
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("fixture must recreate identical progress")
	}
	course, err := s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	if course.Revision == view.Revision {
		t.Fatal("replacement workspace retained old revision")
	}
	for _, action := range []string{"check", "complete", "hint", "pause", "resume", "open"} {
		if _, err := s.PerformAction(action, id, view.Revision); !errors.Is(err, ErrStaleView) {
			t.Fatalf("stale %s accepted: %v", action, err)
		}
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("stale actions changed replacement progress")
	}
	if _, err := s.PerformAction("check", id, course.Revision); err != nil {
		t.Fatal(err)
	}
}

func TestLegacyWorkspaceMarkerRemainsReadableWithoutRewrite(t *testing.T) {
	s := uiTestService(t, "")
	id := "literal-bindings-001"
	workspace, err := s.Start(id)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(workspace, learnerWorkspaceMarker)
	data, err := os.ReadFile(path)
	if err != nil {
		t.Fatal(err)
	}
	var marker WorkspaceManifest
	if err := json.Unmarshal(data, &marker); err != nil {
		t.Fatal(err)
	}
	if marker.Generation == "" {
		t.Fatal("new workspace lacks generation")
	}
	marker.Generation = ""
	legacy, err := json.Marshal(marker)
	if err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.WriteFile(path, legacy, 0400); err != nil {
		t.Fatal(err)
	}
	view, err := s.TaskView(id)
	if err != nil || view.WorkspaceUnavailable {
		t.Fatalf("legacy workspace rejected: %+v %v", view, err)
	}
	if _, err := s.PerformAction("check", id, view.Revision); err != nil {
		t.Fatal(err)
	}
	after, err := os.ReadFile(path)
	if err != nil || !bytes.Equal(after, legacy) {
		t.Fatal("legacy marker rewritten")
	}
}

func TestCourseViewDistinguishesPreparedBlocksAndContentGap(t *testing.T) {
	s := uiTestService(t, "")
	course, err := s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	if course.Blocks[0].Status != "current" || course.Blocks[1].Status != "ready" {
		t.Fatalf("prepared route mislabeled: %+v", course.Blocks[:2])
	}
	tasks, _, err := s.ListTasks()
	if err != nil {
		t.Fatal(err)
	}
	// Synthetic completions only exercise route policy; no authored code runs.
	err = withProgressLock(s.home, func(p *Progress) error {
		for _, task := range tasks {
			p.Tasks[task.ID] = &TaskProgress{TaskID: task.ID, Completed: true, Workspace: "workspaces/" + task.ID}
		}
		return nil
	})
	if err != nil {
		t.Fatal(err)
	}
	course, err = s.CourseView()
	if err != nil {
		t.Fatal(err)
	}
	if course.NextTaskID != "" || course.NextMessage == "" {
		t.Fatalf("content gap unexplained: %+v", course)
	}
}

func TestVerifierInfrastructureFailureExcludesWaitingTime(t *testing.T) {
	for _, paused := range []bool{false, true} {
		for _, full := range []bool{false, true} {
			s := uiTestService(t, "")
			now := time.Date(2026, 9, 5, 12, 0, 0, 0, time.UTC)
			s.now = func() time.Time { return now }
			if _, err := s.Start("literal-bindings-001"); err != nil {
				t.Fatal(err)
			}
			now = now.Add(time.Minute)
			if paused {
				if err := s.Pause(); err != nil {
					t.Fatal(err)
				}
			}
			s.verify = func(string, string, string, tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
				now = now.Add(5 * time.Minute)
				return tasking.LearnerVerificationResult{}, errors.New("fixture failure")
			}
			var err error
			if full {
				_, err = s.Complete()
			} else {
				_, err = s.Check()
			}
			if err == nil {
				t.Fatal("infrastructure error lost")
			}
			p, err := loadProgress(s.home)
			if err != nil {
				t.Fatal(err)
			}
			task := p.Tasks["literal-bindings-001"]
			if task.ActivitySeconds != 60 || task.Completed || task.Attempts != 0 || (task.ActiveSessionStartedAt == "") != paused {
				t.Fatalf("failed verifier changed time/result: %+v", task)
			}
		}
	}
}
