package subject_test

import (
	"testing"
	"time"
)

func TestFiniteBudgetExtensions(t *testing.T) {
	for count := 0; count <= 3; count++ {
		for errorAt := -1; errorAt < count; errorAt++ {
			checkBudgetCase(t, budgetCase{count: count, budget: 7 * time.Second, hasCeiling: true, ceiling: 4 * time.Second, hasParentDeadline: true, parentDeadline: 9 * time.Second, errorAt: errorAt})
			checkBudgetCase(t, budgetCase{count: count, budget: 3 * time.Second, hasCeiling: true, ceiling: 8 * time.Second, hasParentDeadline: true, parentDeadline: 6 * time.Second, errorAt: errorAt})
			checkBudgetCase(t, budgetCase{count: count, budget: -2 * time.Second, errorAt: errorAt})
			checkBudgetCase(t, budgetCase{count: count, budget: 5 * time.Second, canceled: true, errorAt: errorAt})
		}
	}
}
