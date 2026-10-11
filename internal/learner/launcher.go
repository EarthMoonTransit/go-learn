package learner

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"

	"go-practice-bank/internal/tasking"
)

var ErrIDEUnavailable = errors.New("GoLand could not be opened; copy the project path and open it manually")

type ideTarget struct {
	Project string
	File    string
}

// Keep all started lessons in one IDE project. Both paths come from validated
// learner metadata, never from HTTP parameters. Legacy workspace paths work too.
func (s *Service) goLandTarget(id, workspace string) (ideTarget, error) {
	files, err := tasking.LearnerWorkspaceFiles(s.root, id)
	if err != nil {
		return ideTarget{}, err
	}
	for _, file := range files {
		if !file.Editable || filepath.Ext(file.Path) != ".go" {
			continue
		}
		relative, err := filepath.Rel(s.home, workspace)
		if err != nil || !safeWorkspaceRelative(filepath.ToSlash(relative)) {
			return ideTarget{}, errors.New("invalid IDE workspace")
		}
		parent, err := openPrivateDirectory(s.home, filepath.Join(relative, filepath.Dir(file.Path)))
		if err != nil {
			return ideTarget{}, err
		}
		info, err := parent.Lstat(filepath.Base(file.Path))
		parent.Close()
		if err != nil || !info.Mode().IsRegular() {
			return ideTarget{}, errors.New("IDE target must be an existing regular lesson file")
		}
		return ideTarget{Project: filepath.Join(s.home, "workspaces"), File: filepath.Join(workspace, file.Path)}, nil
	}
	return ideTarget{}, errors.New("lesson has no editable Go file")
}

func findGoLand() (string, []string) {
	if launcher, err := exec.LookPath("goland"); err == nil {
		absolute, err := filepath.Abs(launcher)
		if err == nil {
			return absolute, nil
		}
	}
	if runtime.GOOS == "darwin" {
		apps := []string{"/Applications/GoLand.app"}
		if home, err := os.UserHomeDir(); err == nil {
			apps = append(apps, filepath.Join(home, "Applications", "GoLand.app"))
		}
		for _, app := range apps {
			if info, err := os.Stat(app); err == nil && info.IsDir() {
				// LaunchServices only forwards --args to a newly launched process.
				// GoLand's own single-instance handoff forwards project + file to
				// the running IDE; the same project reuses its existing window.
				return "/usr/bin/open", []string{"-n", "-a", app, "--args"}
			}
		}
	}
	return "", nil
}

func launchGoLand(target ideTarget) error {
	executable, args := findGoLand()
	if executable == "" || !filepath.IsAbs(target.Project) || !filepath.IsAbs(target.File) {
		return ErrIDEUnavailable
	}
	relative, err := filepath.Rel(target.Project, target.File)
	if err != nil || !filepath.IsLocal(relative) || relative == "." {
		return ErrIDEUnavailable
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// No shell and no HTTP-supplied executable, flags or filesystem path.
	args = append(args, target.Project, target.File)
	if err := exec.CommandContext(ctx, executable, args...).Run(); err != nil {
		return ErrIDEUnavailable
	}
	return nil
}
