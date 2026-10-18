package tasking

import (
	"fmt"
	"sort"
	"strings"
)

// RequirePassedCheckReports checks executable success, separately from schema
// and freshness. Callers must also validate the collection against the root.
// Aggregate-report construction alone does not enforce this requirement.
func RequirePassedCheckReports(checks CheckCollection) error {
	var failed []string
	for _, checked := range checks.Tasks {
		if !checked.Report.Passed {
			failed = append(failed, checked.TaskID)
		}
	}
	if len(failed) != 0 {
		sort.Strings(failed)
		return fmt.Errorf("task-check: failed package reports: %s", strings.Join(failed, ", "))
	}
	return nil
}
