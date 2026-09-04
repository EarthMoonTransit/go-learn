package tasking

import (
	"bytes"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"gopkg.in/yaml.v3"
)

func BuildAggregateReport(root string) (AggregateReport, error) {
	report := AggregateReport{SchemaVersion: 1}
	stateCounts := map[State]int{}
	typeCounts := map[string]int{}
	seen := map[string]State{}
	for _, state := range orderedStates {
		directory := filepath.Join(root, "tasks", stateDirectories[state])
		entries, err := os.ReadDir(directory)
		if errors.Is(err, os.ErrNotExist) {
			continue
		}
		if err != nil {
			return AggregateReport{}, err
		}
		for _, entry := range entries {
			if !taskIDPattern.MatchString(entry.Name()) {
				continue
			}
			path := filepath.Join(directory, entry.Name())
			info, err := os.Lstat(path)
			if err != nil || info.Mode()&os.ModeSymlink != 0 || !info.IsDir() {
				return AggregateReport{}, fmt.Errorf("task path must be a real directory: %s", path)
			}
			if first, exists := seen[entry.Name()]; exists {
				return AggregateReport{}, fmt.Errorf("duplicate task id %q in %s and %s", entry.Name(), first, state)
			}
			seen[entry.Name()] = state
			if err := inspectTree(path); err != nil {
				return AggregateReport{}, fmt.Errorf("inspect %s: %w", entry.Name(), err)
			}
			var task Task
			if err := decodeStrictFile(filepath.Join(path, "task.yaml"), &task); err != nil {
				return AggregateReport{}, fmt.Errorf("decode %s: %w", entry.Name(), err)
			}
			if task.ID != entry.Name() {
				return AggregateReport{}, fmt.Errorf("task id %q does not match directory %q", task.ID, entry.Name())
			}
			legacy := false
			if _, err := os.Lstat(filepath.Join(path, "lifecycle.yaml")); errors.Is(err, os.ErrNotExist) && state == StateReady {
				legacy = true
			} else if err != nil {
				return AggregateReport{}, fmt.Errorf("inspect lifecycle for %s: %w", entry.Name(), err)
			} else {
				var lifecycle Lifecycle
				if err := decodeStrictFile(filepath.Join(path, "lifecycle.yaml"), &lifecycle); err != nil {
					return AggregateReport{}, fmt.Errorf("decode lifecycle for %s: %w", entry.Name(), err)
				}
				if lifecycle.State != state || task.QualityStatus != qualityByState[state] {
					return AggregateReport{}, fmt.Errorf("task %s lifecycle metadata does not match %s directory", entry.Name(), state)
				}
			}
			report.Tasks = append(report.Tasks, AggregateTask{
				ID: task.ID, State: state, TaskType: task.TaskType, Primary: task.Skills.Primary, Legacy: legacy,
			})
			stateCounts[state]++
			typeCounts[task.TaskType]++
		}
	}
	sort.Slice(report.Tasks, func(i, j int) bool { return report.Tasks[i].ID < report.Tasks[j].ID })
	for _, state := range orderedStates {
		report.States = append(report.States, StateCount{State: state, Count: stateCounts[state]})
	}
	types := make([]string, 0, len(typeCounts))
	for taskType := range typeCounts {
		types = append(types, taskType)
	}
	sort.Strings(types)
	for _, taskType := range types {
		report.TaskTypes = append(report.TaskTypes, TypeCount{TaskType: taskType, Count: typeCounts[taskType]})
	}
	diversity := Gate{ID: "task-type-diversity", Status: "not-applicable", Details: []string{"fewer than 10 lifecycle tasks"}}
	if len(report.Tasks) >= 10 {
		diversity.Status = "pass"
		diversity.Details = []string{}
		maximum := 0
		for _, count := range typeCounts {
			if count > maximum {
				maximum = count
			}
		}
		if len(typeCounts) < 3 || maximum*100 > len(report.Tasks)*60 {
			diversity.Status = "fail"
			diversity.Details = []string{"at least 3 task types are required and no type may exceed 60% once 10 tasks exist"}
		}
	}
	report.Gates = []Gate{diversity}
	return report, nil
}

func MarshalAggregateYAML(report AggregateReport) ([]byte, error) {
	return yaml.Marshal(report)
}

func ValidateAggregateReport(root string, report AggregateReport) error {
	return validateSchema(root, "task-aggregate-report.schema.json", report)
}

func MarshalAggregateMarkdown(report AggregateReport) []byte {
	var output bytes.Buffer
	output.WriteString("# Task quality report\n\n")
	output.WriteString("Generated deterministically by `task-report`; it contains no timestamps.\n\n")
	output.WriteString("## Lifecycle\n\n")
	output.WriteString("| State | Count |\n| --- | ---: |\n")
	for _, item := range report.States {
		fmt.Fprintf(&output, "| `%s` | %d |\n", item.State, item.Count)
	}
	output.WriteString("\n## Task types\n\n")
	if len(report.TaskTypes) == 0 {
		output.WriteString("No task packages found.\n")
	} else {
		output.WriteString("| Type | Count |\n| --- | ---: |\n")
		for _, item := range report.TaskTypes {
			fmt.Fprintf(&output, "| `%s` | %d |\n", item.TaskType, item.Count)
		}
	}
	output.WriteString("\n## Aggregate gates\n\n")
	for _, gate := range report.Gates {
		fmt.Fprintf(&output, "- `%s`: **%s**", gate.ID, gate.Status)
		if len(gate.Details) != 0 {
			fmt.Fprintf(&output, " — %s", strings.Join(gate.Details, "; "))
		}
		output.WriteString("\n")
	}
	output.WriteString("\n## Packages\n\n")
	output.WriteString("| ID | State | Type | Primary skill | Compatibility |\n| --- | --- | --- | --- | --- |\n")
	for _, task := range report.Tasks {
		compatibility := "Phase 3A"
		if task.Legacy {
			compatibility = "legacy v1"
		}
		fmt.Fprintf(&output, "| `%s` | `%s` | `%s` | `%s` | %s |\n", task.ID, task.State, task.TaskType, task.Primary, compatibility)
	}
	return output.Bytes()
}
