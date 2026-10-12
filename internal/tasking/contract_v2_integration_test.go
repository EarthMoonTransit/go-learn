//go:build taskintegration

package tasking

import (
	"path/filepath"
	"testing"
)

func TestTaskV2ProfilesRunInIsolatedWorkspaces(t *testing.T) {
	root := repositoryRoot(t)
	session, err := NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := session.Close(); err != nil {
			t.Error(err)
		}
	})
	for _, fixture := range []struct{ profile, id string }{
		{"fixed", "synthetic-fixed-v2-001"},
		{"mutation", "synthetic-mutation-v2-001"},
		{"workspace", "synthetic-workspace-v2-001"},
	} {
		t.Run(fixture.profile, func(t *testing.T) {
			path := filepath.Join(root, "internal", "testdata", "tasking", "v2", fixture.profile, fixture.id)
			report, err := session.checkLocated(LocatedTask{ID: fixture.id, State: StateDraft, Path: path})
			if err != nil {
				t.Fatal(err)
			}
			if !report.Passed {
				t.Fatalf("isolated %s profile failed: %+v", fixture.profile, report.Gates)
			}
		})
	}
}
