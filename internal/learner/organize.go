package learner

import (
	"bytes"
	"crypto/rand"
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const organizeJournalName = ".organize.json"

type OrganizeMove struct {
	TaskID string `json:"task_id"`
	From   string `json:"from"`
	To     string `json:"to"`
	Digest string `json:"digest"`
}
type OrganizePlan struct {
	FromVersion int            `json:"from_version"`
	ToVersion   int            `json:"to_version"`
	Recovery    bool           `json:"recovery"`
	Backup      string         `json:"backup,omitempty"`
	Moves       []OrganizeMove `json:"moves"`
}
type organizeJournal struct {
	ProgressDigest string       `json:"progress_digest"`
	Version        int          `json:"version"`
	Plan           OrganizePlan `json:"plan"`
	Before         Progress     `json:"before"`
	After          Progress     `json:"after"`
}

// Organize only changes workspace locations and the schema/path fields of
// progress. It does not pause sessions, reset tasks or claim learning evidence.
func (s *Service) Organize(dryRun bool) (plan OrganizePlan, err error) {
	if dryRun {
		j, exists, e := s.readOrganizeJournal()
		if e != nil {
			return plan, e
		}
		if exists {
			j.Plan.Recovery = true
			return j.Plan, nil
		}
		p, e := loadProgress(s.home)
		if e != nil {
			return plan, e
		}
		return s.planOrganization(p)
	}
	release, err := s.acquireOperation()
	if err != nil {
		return plan, err
	}
	defer func() { err = errors.Join(err, release()) }()
	j, exists, err := s.readOrganizeJournal()
	if err != nil {
		return plan, err
	}
	if !exists {
		before, e := loadProgress(s.home)
		if e != nil {
			return plan, e
		}
		plan, e = s.planOrganization(before)
		if e != nil {
			return plan, e
		}
		if len(plan.Moves) == 0 && before.SchemaVersion == progressSchemaVersion {
			return plan, nil
		}
		j, e = s.prepareOrganization(before, plan)
		if e != nil {
			return plan, e
		}
		plan = j.Plan
		data, e := json.MarshalIndent(j, "", "  ")
		if e != nil {
			return plan, e
		}
		if e = atomicPrivateWrite(filepath.Join(s.home, organizeJournalName), append(data, '\n')); e != nil {
			return plan, e
		}
	} else {
		plan = j.Plan
		plan.Recovery = true
	}
	if err = s.applyOrganization(j); err != nil {
		return plan, err
	}
	return plan, nil
}

func (s *Service) planOrganization(p Progress) (OrganizePlan, error) {
	plan := OrganizePlan{FromVersion: p.SchemaVersion, ToVersion: progressSchemaVersion, Moves: []OrganizeMove{}}
	ids := make([]string, 0, len(p.Tasks))
	for id := range p.Tasks {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	destinations := map[string]bool{}
	for _, id := range ids {
		task := p.Tasks[id]
		if task == nil || task.TaskID != id {
			return plan, errors.New("invalid task identity in progress")
		}
		if err := s.checkWorkspaceRelative(id, task.Workspace); err != nil {
			return plan, err
		}
		target, err := s.numberedWorkspaceRelative(id)
		if err != nil {
			return plan, err
		}
		if destinations[target] {
			return plan, errors.New("duplicate organization destination")
		}
		destinations[target] = true
		if target == task.Workspace {
			continue
		}
		handle, err := openPrivateDirectory(s.home, filepath.FromSlash(task.Workspace))
		if err != nil {
			return plan, err
		}
		defer handle.Close()
		if _, err := os.Lstat(filepath.Join(s.home, filepath.FromSlash(target))); err == nil {
			return plan, fmt.Errorf("destination already exists for %s", id)
		} else if !errors.Is(err, os.ErrNotExist) {
			return plan, err
		}
		// Reject existing symlinked parents before creating a journal or backup.
		for _, parent := range pathAncestors(filepath.Dir(filepath.FromSlash(target))) {
			h, e := openPrivateDirectory(s.home, parent)
			if errors.Is(e, os.ErrNotExist) {
				break
			}
			if e != nil {
				return plan, e
			}
			h.Close()
		}
		marker, err := readRootRegular(handle, learnerWorkspaceMarker, 64<<10)
		if err != nil {
			return plan, err
		}
		var identity WorkspaceManifest
		if err = json.Unmarshal(marker, &identity); err != nil || identity.TaskID != id {
			return plan, errors.New("workspace marker does not match progress")
		}
		digest, err := workspaceTreeRoot(handle, nil)
		if err != nil {
			return plan, err
		}
		plan.Moves = append(plan.Moves, OrganizeMove{TaskID: id, From: task.Workspace, To: target, Digest: digest})
	}
	return plan, nil
}
func (s *Service) readOrganizeJournal() (organizeJournal, bool, error) {
	var j organizeJournal
	data, err := readPrivateRegular(filepath.Join(s.home, organizeJournalName), 3*maxProgressBytes)
	if errors.Is(err, os.ErrNotExist) {
		return j, false, nil
	}
	if err != nil {
		return j, false, err
	}
	d := json.NewDecoder(bytes.NewReader(data))
	d.DisallowUnknownFields()
	if err = d.Decode(&j); err != nil {
		return j, false, err
	}
	var extra any
	if err = d.Decode(&extra); err != io.EOF {
		return j, false, errors.New("invalid organization journal trailer")
	}
	if j.Version != 1 || (j.After.SchemaVersion < 2 || j.After.SchemaVersion > progressSchemaVersion || j.After.SchemaVersion != j.Plan.ToVersion) || j.Before.Tasks == nil || j.After.Tasks == nil {
		return j, false, errors.New("unsupported organization journal")
	}
	if !strings.HasPrefix(j.Plan.Backup, "backups/organize-") || filepath.ToSlash(filepath.Clean(j.Plan.Backup)) != j.Plan.Backup || len(strings.Split(j.Plan.Backup, "/")) != 2 {
		return j, false, errors.New("invalid organization backup path")
	}
	expectedRaw, err := json.Marshal(j.Before)
	if err != nil {
		return j, false, err
	}
	var expected Progress
	_ = json.Unmarshal(expectedRaw, &expected)
	upgradeLearningProgress(&expected, j.Plan.ToVersion)
	seen := map[string]bool{}
	for _, m := range j.Plan.Moves {
		before, after := j.Before.Tasks[m.TaskID], j.After.Tasks[m.TaskID]
		if seen[m.TaskID] || before == nil || after == nil || m.From == m.To || before.TaskID != m.TaskID || after.TaskID != m.TaskID || before.Workspace != m.From || after.Workspace != m.To {
			return j, false, errors.New("inconsistent organization move")
		}
		seen[m.TaskID] = true
		if err = s.checkWorkspaceRelative(m.TaskID, m.From); err != nil {
			return j, false, err
		}
		to, e := s.numberedWorkspaceRelative(m.TaskID)
		if e != nil || to != m.To {
			return j, false, errors.New("organization destination differs from route")
		}
		if _, e = hex.DecodeString(m.Digest); e != nil || len(m.Digest) != 64 {
			return j, false, errors.New("invalid workspace digest")
		}
		expected.Tasks[m.TaskID].Workspace = m.To
	}
	if !sameProgress(expected, j.After) {
		return j, false, errors.New("organization journal changes learning history")
	}
	if _, e := hex.DecodeString(j.ProgressDigest); e != nil || len(j.ProgressDigest) != 64 {
		return j, false, errors.New("invalid progress backup digest")
	}
	return j, true, nil
}
func sameProgress(a, b Progress) bool {
	left, _ := json.Marshal(a)
	right, _ := json.Marshal(b)
	return bytes.Equal(left, right)
}
func (s *Service) applyOrganization(j organizeJournal) error {
	return withProgressLock(s.home, func(current *Progress) error {
		if !sameProgress(*current, j.Before) && !sameProgress(*current, j.After) {
			return errors.New("progress changed since organization began; refusing recovery")
		}
		backupRoot, e := openPrivateDirectory(s.home, filepath.FromSlash(j.Plan.Backup))
		if e != nil {
			return e
		}
		raw, e := readRootRegular(backupRoot, "progress.json", maxProgressBytes)
		backupRoot.Close()
		if e != nil {
			return e
		}
		sum := sha256.Sum256(raw)
		if hex.EncodeToString(sum[:]) != j.ProgressDigest {
			return errors.New("original progress backup is missing or changed")
		}
		for _, m := range j.Plan.Moves {
			// Verify the recovery copy before moving anything or committing paths.
			h, e := openPrivateDirectory(s.home, filepath.Join(filepath.FromSlash(j.Plan.Backup), filepath.FromSlash(m.From)))
			if e != nil {
				return e
			}
			digest, e := workspaceTreeRoot(h, nil)
			h.Close()
			if e != nil {
				return e
			}
			if digest != m.Digest {
				return errors.New("organization backup does not match source")
			}
			from, to := filepath.Join(s.home, filepath.FromSlash(m.From)), filepath.Join(s.home, filepath.FromSlash(m.To))
			_, fromErr := os.Lstat(from)
			_, toErr := os.Lstat(to)
			if fromErr == nil && toErr == nil {
				return fmt.Errorf("both organization paths exist for %s", m.TaskID)
			}
			relative := m.From
			if errors.Is(fromErr, os.ErrNotExist) && toErr == nil {
				relative = m.To
			} else if fromErr != nil {
				return fromErr
			} else if !errors.Is(toErr, os.ErrNotExist) {
				return toErr
			}
			h, e = openPrivateDirectory(s.home, filepath.FromSlash(relative))
			if e != nil {
				return e
			}
			digest, e = workspaceTreeRoot(h, nil)
			h.Close()
			if e != nil {
				return e
			}
			if digest != m.Digest {
				return errors.New("workspace changed during organization; files retained for recovery")
			}
			if relative == m.From {
				if e = ensurePrivateDirectory(s.home, filepath.Dir(filepath.FromSlash(m.To))); e != nil {
					return e
				}
				root, e := openPrivateDirectory(s.home, "workspaces")
				if e != nil {
					return e
				}
				e = root.Rename(strings.TrimPrefix(m.From, "workspaces/"), strings.TrimPrefix(m.To, "workspaces/"))
				root.Close()
				if e != nil {
					return e
				}
			}
		}
		*current = j.After
		// Persist before clearing the journal. Replaying after this save is safe.
		if e := saveProgress(s.home, *current); e != nil {
			return e
		}
		return os.Remove(filepath.Join(s.home, organizeJournalName))
	})
}

// Digest and optionally copy a bounded tree. Source files are only read as data.
// The root handle confines reads even if a component is replaced concurrently.
func workspaceTree(source, destination string) (string, error) {
	root, err := os.OpenRoot(source)
	if err != nil {
		return "", err
	}
	defer root.Close()
	var target *os.Root
	if destination != "" {
		if err = os.MkdirAll(destination, 0700); err != nil {
			return "", err
		}
		target, err = os.OpenRoot(destination)
		if err != nil {
			return "", err
		}
		defer target.Close()
	}
	return workspaceTreeRoot(root, target)
}
func readRootRegular(root *os.Root, path string, limit int64) ([]byte, error) {
	before, err := root.Lstat(path)
	if err != nil {
		return nil, err
	}
	if !before.Mode().IsRegular() || before.Size() > limit {
		return nil, errors.New("backup input must be a bounded regular file")
	}
	file, err := root.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	opened, err := file.Stat()
	if err != nil || !os.SameFile(before, opened) {
		return nil, errors.New("backup input changed")
	}
	data, err := io.ReadAll(io.LimitReader(file, limit+1))
	if err != nil {
		return nil, err
	}
	if int64(len(data)) > limit {
		return nil, errors.New("backup input exceeds limit")
	}
	return data, nil
}
func workspaceTreeRoot(root, target *os.Root) (string, error) {
	hash := sha256.New()
	var total int64
	err := fs.WalkDir(root.FS(), ".", func(path string, entry fs.DirEntry, walkErr error) error {
		if walkErr != nil {
			return walkErr
		}
		if path == "." {
			return nil
		}
		info, e := entry.Info()
		if e != nil {
			return e
		}
		if info.Mode()&os.ModeSymlink != 0 {
			return errors.New("workspace backup refuses symlinks")
		}
		if info.IsDir() {
			if target != nil {
				return target.Mkdir(path, 0700)
			}
			return nil
		}
		if !info.Mode().IsRegular() {
			return errors.New("workspace backup requires regular files")
		}
		content, e := readRootRegular(root, path, (128<<20)-total)
		if e != nil {
			return e
		}
		total += int64(len(content))
		fmt.Fprintf(hash, "%d:%s:%o:%d:", len(path), path, info.Mode().Perm(), len(content))
		hash.Write(content)
		if target != nil {
			file, e := target.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, info.Mode().Perm())
			if e != nil {
				return e
			}
			_, e = file.Write(content)
			e = errors.Join(e, file.Chmod(info.Mode().Perm()), file.Sync(), file.Close())
			return e
		}
		return nil
	})
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(hash.Sum(nil)), nil
}

func writeRootFile(root *os.Root, path string, content []byte, mode os.FileMode) error {
	file, err := root.OpenFile(path, os.O_CREATE|os.O_EXCL|os.O_WRONLY, mode)
	if err != nil {
		return err
	}
	_, err = file.Write(content)
	return errors.Join(err, file.Sync(), file.Close())
}

func (s *Service) prepareOrganization(before Progress, plan OrganizePlan) (organizeJournal, error) {
	var err error
	if err = ensurePrivateDirectory(s.home, "backups"); err != nil {
		return organizeJournal{}, err
	}
	backupParent, e := openPrivateDirectory(s.home, "backups")
	if e != nil {
		return organizeJournal{}, e
	}
	suffix := make([]byte, 16)
	if _, e = rand.Read(suffix); e != nil {
		backupParent.Close()
		return organizeJournal{}, e
	}
	backupName := "organize-" + hex.EncodeToString(suffix)
	if e = backupParent.Mkdir(backupName, 0700); e != nil {
		backupParent.Close()
		return organizeJournal{}, e
	}
	backupRoot, e := backupParent.OpenRoot(backupName)
	backupParent.Close()
	if e != nil {
		return organizeJournal{}, e
	}
	defer backupRoot.Close()
	plan.Backup = filepath.ToSlash(filepath.Join("backups", backupName))
	homeRoot, e := openPrivateDirectory(s.home, ".")
	if e != nil {
		return organizeJournal{}, e
	}
	raw, e := readRootRegular(homeRoot, "progress.json", maxProgressBytes)
	homeRoot.Close()
	if errors.Is(e, os.ErrNotExist) {
		raw, e = json.Marshal(before)
	}
	if e != nil {
		return organizeJournal{}, e
	}
	var copiedProgress Progress
	if e = json.Unmarshal(raw, &copiedProgress); e != nil || !sameProgress(copiedProgress, before) {
		return organizeJournal{}, errors.New("progress changed while backing up")
	}
	if e = writeRootFile(backupRoot, "progress.json", raw, 0600); e != nil {
		return organizeJournal{}, e
	}
	progressSum := sha256.Sum256(raw)
	for _, m := range plan.Moves {
		source, e := openPrivateDirectory(s.home, filepath.FromSlash(m.From))
		if e != nil {
			return organizeJournal{}, e
		}
		backupRelative := filepath.Join(plan.Backup, m.From)
		if e = ensurePrivateDirectory(s.home, backupRelative); e != nil {
			source.Close()
			return organizeJournal{}, e
		}
		target, e := openPrivateDirectory(s.home, backupRelative)
		if e != nil {
			source.Close()
			return organizeJournal{}, e
		}
		digest, e := workspaceTreeRoot(source, target)
		copiedDigest, copyErr := workspaceTreeRoot(target, nil)
		source.Close()
		target.Close()
		e = errors.Join(e, copyErr)
		if e == nil && copiedDigest != digest {
			e = errors.New("backup bytes or permissions differ from source")
		}

		if e != nil {
			return organizeJournal{}, e
		}
		if digest != m.Digest {
			return organizeJournal{}, errors.New("workspace changed while backing up; no journal published")
		}
	}
	raw, e = json.Marshal(before)
	if e != nil {
		return organizeJournal{}, e
	}
	var after Progress
	if e = json.Unmarshal(raw, &after); e != nil {
		return organizeJournal{}, e
	}
	upgradeLearningProgress(&after, progressSchemaVersion)
	for _, m := range plan.Moves {
		after.Tasks[m.TaskID].Workspace = m.To
	}
	j := organizeJournal{Version: 1, Plan: plan, Before: before, After: after, ProgressDigest: hex.EncodeToString(progressSum[:])}

	return j, nil
}
