package learner

import (
	"context"
	"errors"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"time"
)

var ErrIDEUnavailable = errors.New("GoLand could not be opened; copy the workspace path and open it manually")

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
				return "/usr/bin/open", []string{"-a", app}
			}
		}
	}
	return "", nil
}

func launchGoLand(workspace string) error {
	executable, args := findGoLand()
	if executable == "" || !filepath.IsAbs(workspace) {
		return ErrIDEUnavailable
	}
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// No shell and no HTTP-supplied executable, flags or filesystem path.
	args = append(args, workspace)
	if err := exec.CommandContext(ctx, executable, args...).Run(); err != nil {
		return ErrIDEUnavailable
	}
	return nil
}
