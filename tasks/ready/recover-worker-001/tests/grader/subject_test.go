package subject_test

import (
	subject "task.local/recover-worker-001/subject"
	"testing"
	"testing/synctest"
)

func TestLiteralPanicDiagnostics(t *testing.T) {
	for _, tc := range []struct {
		value any
		text  string
	}{{-71, "worker panic: -71"}, {"%v\nΩ", "worker panic: %v\nΩ"}} {
		synctest.Test(t, func(t *testing.T) {
			checkWorkerResult(t, drainWorker(t, subject.Launch(func() int { panic(tc.value) })), 0, tc.text, true)
		})
	}
}
