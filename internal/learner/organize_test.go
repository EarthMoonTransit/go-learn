package learner

import (
	"bytes"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"os"
	"path/filepath"
	"strings"
	"testing"
)

func legacyOrganizeService(t *testing.T) *Service {
	t.Helper()
	s := uiTestService(t, "")
	if err := saveProgress(s.home, Progress{SchemaVersion: 1, Tasks: map[string]*TaskProgress{}}); err != nil {
		t.Fatal(err)
	}
	for _, id := range []string{"literal-bindings-001", "zero-results-001"} {
		if _, err := s.Start(id); err != nil {
			t.Fatal(err)
		}
	}
	if err := withProgressLock(s.home, func(p *Progress) error {
		task := p.Tasks["literal-bindings-001"]
		task.Completed = true
		task.CompletedAt = "2026-09-08T12:00:00Z"
		task.Attempts = 4
		task.UsedHints = []int{1}
		task.ActivitySeconds = 123
		return nil
	}); err != nil {
		t.Fatal(err)
	}
	return s
}
func TestOrganizePreservesFilesAndHistoryAndIsIdempotent(t *testing.T) {
	s := legacyOrganizeService(t)
	before := persistedBytes(t, s)
	plan, err := s.Organize(true)
	if err != nil || len(plan.Moves) != 2 {
		t.Fatalf("plan=%+v err=%v", plan, err)
	}
	if !bytes.Equal(before, persistedBytes(t, s)) {
		t.Fatal("preview changed progress")
	}
	if _, err := os.Stat(filepath.Join(s.home, "backups")); !os.IsNotExist(err) {
		t.Fatal("preview created backup")
	}
	oldMarker, err := os.ReadFile(filepath.Join(s.home, "workspaces/literal-bindings-001", learnerWorkspaceMarker))
	if err != nil {
		t.Fatal(err)
	}
	plan, err = s.Organize(false)
	if err != nil {
		t.Fatal(err)
	}
	var expected Progress
	if err = json.Unmarshal(before, &expected); err != nil {
		t.Fatal(err)
	}
	upgradeLearningProgress(&expected, progressSchemaVersion)
	for _, m := range plan.Moves {
		expected.Tasks[m.TaskID].Workspace = m.To
		afterDigest, e := workspaceTree(filepath.Join(s.home, m.To), "")
		if e != nil || afterDigest != m.Digest {
			t.Fatalf("changed workspace: %s %v", m.TaskID, e)
		}
		backupDigest, e := workspaceTree(filepath.Join(s.home, plan.Backup, m.From), "")
		if e != nil || backupDigest != m.Digest {
			t.Fatal("bad backup", e)
		}
	}
	actual, err := loadProgress(s.home)
	if err != nil || !sameProgress(expected, actual) {
		t.Fatalf("history changed: %+v %v", actual, err)
	}
	marker, err := os.ReadFile(filepath.Join(s.home, actual.Tasks["literal-bindings-001"].Workspace, learnerWorkspaceMarker))
	if err != nil || !bytes.Equal(oldMarker, marker) {
		t.Fatal("marker changed", err)
	}
	backupProgress, err := os.ReadFile(filepath.Join(s.home, plan.Backup, "progress.json"))
	if err != nil || !bytes.Equal(before, backupProgress) {
		t.Fatal("original progress not backed up", err)
	}
	again, err := s.Organize(false)
	if err != nil || len(again.Moves) != 0 {
		t.Fatalf("second organize=%+v %v", again, err)
	}
	view, err := s.TaskView("literal-bindings-001")
	if err != nil || view.WorkspaceUnavailable || !view.Completed {
		t.Fatalf("view=%+v err=%v", view, err)
	}
	if _, err = s.Resume("zero-results-001"); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Check(); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Complete(); err != nil {
		t.Fatal(err)
	}
	if err = s.ResetTask("zero-results-001", true); err != nil {
		t.Fatal(err)
	}
	if _, err = os.Stat(filepath.Join(s.home, actual.Tasks["literal-bindings-001"].Workspace)); err != nil {
		t.Fatal("reset removed another workspace", err)
	}
}
func TestNewWorkspacesUseExplicitRouteNumbers(t *testing.T) {
	s := uiTestService(t, "")
	path, err := s.Start("literal-bindings-001")
	if err != nil {
		t.Fatal(err)
	}
	if path != filepath.Join(s.home, "workspaces/01-first-steps/01-literal-bindings") {
		t.Fatal(path)
	}
	p, err := loadProgress(s.home)
	if err != nil || p.SchemaVersion != progressSchemaVersion {
		t.Fatal(p, err)
	}
	if err := s.ResetTask("literal-bindings-001", true); err != nil {
		t.Fatal(err)
	}
}
func TestOrganizeRefusesConflictsAndSymlinkParentsWithoutMutation(t *testing.T) {
	for _, kind := range []string{"conflict", "parent symlink", "source symlink"} {
		t.Run(kind, func(t *testing.T) {
			s := legacyOrganizeService(t)
			before := persistedBytes(t, s)
			target := filepath.Join(s.home, "workspaces/01-first-steps")
			switch kind {
			case "conflict":
				if err := os.MkdirAll(filepath.Join(target, "01-literal-bindings"), 0700); err != nil {
					t.Fatal(err)
				}
			case "parent symlink":
				if err := os.Symlink(t.TempDir(), target); err != nil {
					t.Fatal(err)
				}
			case "source symlink":
				if err := os.Symlink(t.TempDir(), filepath.Join(s.home, "workspaces/literal-bindings-001/escape")); err != nil {
					t.Fatal(err)
				}
			}
			if _, err := s.Organize(false); err == nil {
				t.Fatal("unsafe organization accepted")
			}
			if !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatal("failure changed progress")
			}
			if _, err := os.Stat(filepath.Join(s.home, organizeJournalName)); !os.IsNotExist(err) {
				t.Fatal("unsafe plan published a journal")
			}
		})
	}
}
func stageOrganizeRecovery(t *testing.T, s *Service) organizeJournal {
	t.Helper()
	p, err := loadProgress(s.home)
	if err != nil {
		t.Fatal(err)
	}
	plan, err := s.planOrganization(p)
	if err != nil {
		t.Fatal(err)
	}
	plan.Backup = "backups/organize-fixture"
	raw, _ := json.Marshal(p)
	if err = writePrivateFile(filepath.Join(s.home, plan.Backup, "progress.json"), raw, 0600); err != nil {
		t.Fatal(err)
	}
	var after Progress
	_ = json.Unmarshal(raw, &after)
	upgradeLearningProgress(&after, progressSchemaVersion)
	for _, m := range plan.Moves {
		_, err = workspaceTree(filepath.Join(s.home, m.From), filepath.Join(s.home, plan.Backup, m.From))
		if err != nil {
			t.Fatal(err)
		}
		after.Tasks[m.TaskID].Workspace = m.To
	}
	sum := sha256.Sum256(raw)
	j := organizeJournal{Version: 1, Plan: plan, Before: p, After: after, ProgressDigest: hex.EncodeToString(sum[:])}
	raw, _ = json.Marshal(j)
	if err = writePrivateFile(filepath.Join(s.home, organizeJournalName), raw, 0600); err != nil {
		t.Fatal(err)
	}
	return j
}
func TestOrganizeRecoversPartialMoveAndCommittedProgress(t *testing.T) {
	for _, committed := range []bool{false, true} {
		t.Run(map[bool]string{false: "partial rename", true: "progress saved"}[committed], func(t *testing.T) {
			s := legacyOrganizeService(t)
			j := stageOrganizeRecovery(t, s)
			moves := j.Plan.Moves[:1]
			if committed {
				moves = j.Plan.Moves
			}
			for _, m := range moves {
				if err := ensurePrivateDirectory(s.home, filepath.Dir(m.To)); err != nil {
					t.Fatal(err)
				}
				if err := os.Rename(filepath.Join(s.home, m.From), filepath.Join(s.home, m.To)); err != nil {
					t.Fatal(err)
				}
			}
			if committed {
				if err := saveProgress(s.home, j.After); err != nil {
					t.Fatal(err)
				}
			}
			before := persistedBytes(t, s)
			if _, err := s.Check(); err == nil || !strings.Contains(err.Error(), "unfinished") {
				t.Fatal("ordinary operation ignored migration", err)
			}
			preview, err := s.Organize(true)
			if err != nil || !preview.Recovery || !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatal("recovery preview mutated", err)
			}
			plan, err := s.Organize(false)
			if err != nil || !plan.Recovery {
				t.Fatal("recovery failed", err)
			}
			p, err := loadProgress(s.home)
			if err != nil || !sameProgress(p, j.After) {
				t.Fatal("recovery lost state", err)
			}
		})
	}
}
func TestOrganizeRejectsChangedRecoveryDataAndBusyOperation(t *testing.T) {
	s := legacyOrganizeService(t)
	release, err := s.acquireOperation()
	if err != nil {
		t.Fatal(err)
	}
	if _, err = s.Organize(false); !errors.Is(err, ErrOperationBusy) {
		t.Fatal(err)
	}
	if err = release(); err != nil {
		t.Fatal(err)
	}
	j := stageOrganizeRecovery(t, s)
	first := j.Plan.Moves[0]
	if err = os.WriteFile(filepath.Join(s.home, first.From, "extra.txt"), []byte("edited while interrupted"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err = s.Organize(false); err == nil {
		t.Fatal("changed source accepted")
	}
	if _, err = os.Stat(filepath.Join(s.home, first.From, "extra.txt")); err != nil {
		t.Fatal("edit lost", err)
	}
}

func TestOrganizeRecoveryRequiresOriginalProgressBackup(t *testing.T) {
	for _, remove := range []bool{false, true} {
		t.Run(map[bool]string{false: "corrupt", true: "missing"}[remove], func(t *testing.T) {
			s := legacyOrganizeService(t)
			j := stageOrganizeRecovery(t, s)
			before := persistedBytes(t, s)
			path := filepath.Join(s.home, j.Plan.Backup, "progress.json")
			var err error
			if remove {
				err = os.Remove(path)
			} else {
				err = os.WriteFile(path, []byte("corrupt"), 0600)
			}
			if err != nil {
				t.Fatal(err)
			}
			if _, err = s.Organize(false); err == nil {
				t.Fatal("lost original progress backup accepted")
			}
			if !bytes.Equal(before, persistedBytes(t, s)) {
				t.Fatal("recovery changed progress")
			}
			for _, m := range j.Plan.Moves {
				if _, err = os.Stat(filepath.Join(s.home, m.From)); err != nil {
					t.Fatal("source moved without a valid progress backup", err)
				}
			}
		})
	}
}
func TestPinnedWorkspaceTreeDoesNotFollowReplacedSource(t *testing.T) {
	s := legacyOrganizeService(t)
	relative := "workspaces/literal-bindings-001"
	handle, err := openPrivateDirectory(s.home, relative)
	if err != nil {
		t.Fatal(err)
	}
	defer handle.Close()
	expected, err := workspaceTreeRoot(handle, nil)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(s.home, relative)
	if err = os.Rename(path, path+"-retained"); err != nil {
		t.Fatal(err)
	}
	outside := t.TempDir()
	if err = os.WriteFile(filepath.Join(outside, "secret"), []byte("outside"), 0600); err != nil {
		t.Fatal(err)
	}
	if err = os.Symlink(outside, path); err != nil {
		t.Fatal(err)
	}
	actual, err := workspaceTreeRoot(handle, nil)
	if err != nil || actual != expected {
		t.Fatal("pinned traversal followed replacement", err)
	}
}
func TestUnpublishedJournalTemporaryFileDoesNotBlockOrganization(t *testing.T) {
	s := legacyOrganizeService(t)
	if err := os.WriteFile(filepath.Join(s.home, ".write-interrupted"), []byte("partial journal"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err := s.Organize(false); err != nil {
		t.Fatal(err)
	}
}

func TestChangedWorkspaceBeforeBackupDoesNotPublishJournal(t *testing.T) {
	s := legacyOrganizeService(t)
	before, err := loadProgress(s.home)
	if err != nil {
		t.Fatal(err)
	}
	plan, err := s.planOrganization(before)
	if err != nil {
		t.Fatal(err)
	}
	path := filepath.Join(s.home, plan.Moves[0].From, "new-edit.txt")
	if err = os.WriteFile(path, []byte("changed after plan"), 0600); err != nil {
		t.Fatal(err)
	}
	if _, err = s.prepareOrganization(before, plan); err == nil {
		t.Fatal("changed snapshot accepted")
	}
	if _, err = os.Stat(filepath.Join(s.home, organizeJournalName)); !os.IsNotExist(err) {
		t.Fatal("inconsistent journal published")
	}
	if _, err = os.Stat(path); err != nil {
		t.Fatal("edit lost", err)
	}
}

func TestOrganizePreservesExplicitFilePermissions(t *testing.T) {
	s := legacyOrganizeService(t)
	relative := "workspaces/literal-bindings-001/ide-settings.xml"
	path := filepath.Join(s.home, relative)
	if err := os.WriteFile(path, []byte("IDE metadata"), 0600); err != nil {
		t.Fatal(err)
	}
	if err := os.Chmod(path, 0666); err != nil {
		t.Fatal(err)
	}
	plan, err := s.Organize(false)
	if err != nil {
		t.Fatal(err)
	}
	for _, m := range plan.Moves {
		if m.TaskID == "literal-bindings-001" {
			for _, p := range []string{filepath.Join(s.home, m.To, "ide-settings.xml"), filepath.Join(s.home, plan.Backup, relative)} {
				info, e := os.Stat(p)
				if e != nil || info.Mode().Perm() != 0666 {
					t.Fatalf("permissions changed: %s %v", p, e)
				}
			}
		}
	}
}
