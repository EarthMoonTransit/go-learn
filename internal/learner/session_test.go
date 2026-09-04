package learner

import "testing"

func TestVerificationTimeDoesNotCountAsLearnerWork(t *testing.T) {
	progress := &TaskProgress{ActivitySeconds: 15, ActiveSessionStartedAt: "2026-09-04T10:00:00Z"}
	finishVerificationSession(progress, "2026-09-04T10:01:00Z", "2026-09-04T10:07:00Z", true)
	endActiveSession(progress, "2026-09-04T10:08:00Z")
	if progress.ActivitySeconds != 135 {
		t.Fatalf("verification time counted: %d", progress.ActivitySeconds)
	}
	finishVerificationSession(progress, "2026-09-04T11:00:00Z", "2026-09-04T11:01:00Z", true)
	if progress.ActiveSessionStartedAt != "" || progress.ActivitySeconds != 135 {
		t.Fatal("verification resumed a paused task")
	}
	progress.ActiveSessionStartedAt = "2026-09-04T12:00:00Z"
	finishVerificationSession(progress, "2026-09-04T12:01:00Z", "2026-09-04T12:05:00Z", false)
	if progress.ActiveSessionStartedAt != "" || progress.ActivitySeconds != 195 {
		t.Fatal("completion did not stop the timer")
	}
}
