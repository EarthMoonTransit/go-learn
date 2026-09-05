//go:build unix

package tasking

import (
	"os"
	"path/filepath"
	"syscall"
	"testing"
)

func TestLearnerIDERejectsNonRegularMetadata(t *testing.T) {
	for _, kind := range []string{"implementation", "tests", "workspace-patch"} {
		t.Run(kind, func(t *testing.T) {
			loaded, workspace := ideaWorkspaceFixture(t, kind)
			if err := os.Mkdir(filepath.Join(workspace, ".idea"), 0700); err != nil {
				t.Fatal(err)
			}
			if err := syscall.Mkfifo(filepath.Join(workspace, ".idea", "pipe"), 0600); err != nil {
				t.Fatal(err)
			}
			if _, err := inspectLearnerSubmission(loaded, workspace); err == nil {
				t.Fatal("nonregular IDE metadata accepted")
			}
		})
	}
}
