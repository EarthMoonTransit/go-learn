package subject_test

import (
	"task.local/owned-values-001/subject"
	"testing"
	"testing/synctest"
)

func TestGraderCase01(t *testing.T) {
	synctest.Test(t, func(t *testing.T) { _ = subject.OwnedValues([]int{1, 2, 3}); synctest.Wait() })
}
