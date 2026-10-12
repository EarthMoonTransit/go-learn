//go:build taskintegration

package learner

import (
	"fmt"
	"path/filepath"
	"testing"

	"go-practice-bank/internal/tasking"
)

// One explicit top-level test owns the disposable compiler cache. Subtests
// retain their own learner homes and every invocation gets a fresh module.
// This invokes the same verifier as the one-shot CLI, with no cached results.
func integrationVerifier(t *testing.T, root string) func(string, string, string, tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
	t.Helper()
	// Service.New resolves repository symlinks (including macOS /var).
	root, err := filepath.EvalSymlinks(root)
	if err != nil {
		t.Fatal(err)
	}
	session, err := tasking.NewCheckSession(root)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		if err := session.Close(); err != nil {
			t.Error(err)
		}
	})
	return func(actualRoot, id, workspace string, scope tasking.LearnerVerificationScope) (tasking.LearnerVerificationResult, error) {
		if actualRoot != root {
			return tasking.LearnerVerificationResult{}, fmt.Errorf("integration verifier repository changed: %s", actualRoot)
		}
		return session.VerifyLearnerWorkspace(id, workspace, scope)
	}
}
