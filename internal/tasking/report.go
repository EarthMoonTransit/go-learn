package tasking

import (
	"bytes"
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
	located, err := DiscoverTasks(root, nil)
	if err != nil {
		return AggregateReport{}, err
	}
	checks, err := CheckAll(root, nil)
	if err != nil {
		return AggregateReport{}, err
	}
	byID := map[string]QualityReport{}
	for _, check := range checks.Reports {
		byID[check.TaskID] = check
	}
	for _, item := range located {
		var task Task
		if err := decodeStrictFile(filepath.Join(item.Path, "task.yaml"), &task); err != nil {
			return AggregateReport{}, fmt.Errorf("decode %s: %w", item.ID, err)
		}
		check := byID[item.ID]
		report.Tasks = append(report.Tasks, AggregateTask{
			ID: task.ID, State: item.State, TaskType: task.TaskType, Primary: task.Skills.Primary,
			TechnicalStatus: check.TechnicalStatus, ReviewStatus: check.ReviewStatus,
			PromotionEligible: check.PromotionEligible, FailingGateIDs: append([]string{}, check.FailingGateIDs...),
			StarterStatus: check.StarterStatus, ReferenceStatus: check.ReferenceStatus, RaceStatus: check.RaceStatus,
			Mutants: check.Mutants, Legacy: check.Legacy,
		})
		stateCounts[item.State]++
		typeCounts[task.TaskType]++
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
	policyGate := Gate{ID: "task-state-policy", Status: "pass", Details: []string{}}
	for _, check := range checks.Reports {
		if !check.Passed {
			policyGate.Status = "fail"
			policyGate.Details = append(policyGate.Details, check.TaskID)
		}
	}
	sort.Strings(policyGate.Details)
	report.Gates = []Gate{policyGate, diversity}
	return report, nil
}

func MarshalAggregateYAML(report AggregateReport) ([]byte, error) {
	return yaml.Marshal(report)
}

func ValidateAggregateReport(root string, report AggregateReport) error {
	return validateSchema(root, "task-aggregate-report.schema.json", report)
}

func CheckAggregateArtifacts(root, markdownPath, yamlPath string, report AggregateReport) error {
	yamlContent, err := MarshalAggregateYAML(report)
	if err != nil {
		return err
	}
	for _, artifact := range []struct {
		path     string
		expected []byte
	}{
		{markdownPath, MarshalAggregateMarkdown(report)},
		{yamlPath, yamlContent},
	} {
		if err := cleanRelativePath(artifact.path); err != nil {
			return err
		}
		actual, err := os.ReadFile(filepath.Join(root, filepath.FromSlash(artifact.path)))
		if err != nil {
			return err
		}
		if !bytes.Equal(actual, artifact.expected) {
			return fmt.Errorf("%s is stale", artifact.path)
		}
	}
	return nil
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
	output.WriteString("| ID | State | Technical | Review | Eligible | Starter | Reference | Race | Mutants | Compatibility |\n| --- | --- | --- | --- | ---: | --- | --- | --- | --- | --- |\n")
	for _, task := range report.Tasks {
		compatibility := "Phase 3A"
		if task.Legacy {
			compatibility = "legacy v1"
		}
		mutants := fmt.Sprintf("%d/%d", task.Mutants.Detected, task.Mutants.Total)
		fmt.Fprintf(&output, "| `%s` | `%s` | `%s` | `%s` | %t | `%s` | `%s` | `%s` | %s | %s |\n", task.ID, task.State, task.TechnicalStatus, task.ReviewStatus, task.PromotionEligible, task.StarterStatus, task.ReferenceStatus, task.RaceStatus, mutants, compatibility)
	}
	return output.Bytes()
}
