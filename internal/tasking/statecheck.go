package tasking

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"go-practice-bank/internal/toolchainpolicy"
)

func checkInactive(root string, located LocatedTask) (QualityReport, error) {
	c := checker{report: QualityReport{
		SchemaVersion: 1, TaskID: located.ID, State: located.State,
		FailingGateIDs: []string{}, StarterStatus: "not-applicable", ReferenceStatus: "not-applicable", RaceStatus: "not-applicable",
		Mutants: MutantSummary{Survived: []string{}},
	}}
	data := packageData{Located: located}
	var loadProblems []string
	if err := inspectTree(located.Path); err != nil {
		loadProblems = append(loadProblems, err.Error())
	}
	for _, item := range []struct {
		name   string
		value  any
		schema string
	}{
		{"task.yaml", &data.Task, "task.schema.json"},
		{"lifecycle.yaml", &data.Lifecycle, "task-lifecycle.schema.json"},
		{"provenance.yaml", &data.Provenance, "task-provenance.schema.json"},
	} {
		if err := decodeStrictFile(filepath.Join(located.Path, filepath.FromSlash(item.name)), item.value); err != nil {
			loadProblems = append(loadProblems, item.name+": "+err.Error())
			continue
		}
		if err := validateSchema(root, item.schema, item.value); err != nil {
			loadProblems = append(loadProblems, item.name+": "+err.Error())
		}
	}
	c.gate("package-metadata", loadProblems)
	c.gate("taxonomy-and-lifecycle", checkMetadata(root, data))
	c.gate("provenance-license", checkProvenance(root, data))
	c.notApplicable("review-evidence", "rejected and retired packages retain core audit metadata only")
	c.notApplicable("execution", "rejected and retired packages are not executable authoring targets")
	c.finish()
	c.report.PromotionEligible = false
	return c.validatedReport(root)
}

func (s *CheckSession) checkLegacy(located LocatedTask) (QualityReport, error) {
	root := s.root
	c := checker{report: QualityReport{
		SchemaVersion: 1, TaskID: located.ID, State: located.State, Legacy: true,
		ReviewStatus: "not-applicable", StarterStatus: "not-applicable", ReferenceStatus: "not-applicable", RaceStatus: "not-applicable",
		FailingGateIDs: []string{}, Mutants: MutantSummary{Survived: []string{}},
	}}
	var task Task
	var metadataProblems []string
	if err := inspectTree(located.Path); err != nil {
		metadataProblems = append(metadataProblems, err.Error())
	}
	if err := decodeStrictFile(filepath.Join(located.Path, "task.yaml"), &task); err != nil {
		metadataProblems = append(metadataProblems, "task.yaml: "+err.Error())
	} else if err := validateSchema(root, "task.schema.json", task); err != nil {
		metadataProblems = append(metadataProblems, err.Error())
	}
	if task.ID != located.ID || task.QualityStatus != "ready" {
		metadataProblems = append(metadataProblems, "legacy task id/status must match its ready directory")
	}
	c.gate("legacy-metadata", metadataProblems)

	policy, err := toolchainpolicy.Load(root)
	if err != nil {
		return QualityReport{}, err
	}
	version, err := policy.VersionForTask(located.ID, true)
	if err != nil {
		return QualityReport{}, err
	}
	goMod, readErr := os.ReadFile(filepath.Join(located.Path, "go.mod"))
	var toolchainProblems []string
	if readErr != nil {
		toolchainProblems = append(toolchainProblems, readErr.Error())
	} else {
		text := string(goMod)
		if !strings.Contains(text, "\ngo "+version+"\n") || strings.Contains(text, "\ntoolchain ") {
			toolchainProblems = append(toolchainProblems, "legacy go.mod must use its pinned Go language version without a toolchain directive")
		}
	}
	c.gate("legacy-toolchain", toolchainProblems)

	var commandProblems []string
	if len(metadataProblems) == 0 && len(toolchainProblems) == 0 {
		results, runErr := s.runner.RunLegacy(located.Path, task.ValidationCommands)
		if runErr != nil {
			commandProblems = append(commandProblems, runErr.Error())
		} else {
			for _, result := range results {
				if result.Err != nil {
					commandProblems = append(commandProblems, fmt.Sprintf("%s: %v", result.Command, result.Err))
				}
			}
		}
	} else {
		commandProblems = append(commandProblems, "execution skipped because legacy metadata/toolchain validation failed")
	}
	c.gate("legacy-commands", commandProblems)
	c.report.TechnicalStatus = "pass"
	for _, gate := range c.report.Gates {
		if gate.Status == "fail" {
			c.report.TechnicalStatus = "fail"
			c.report.FailingGateIDs = append(c.report.FailingGateIDs, gate.ID)
		}
	}
	c.report.FailingGateIDs = uniqueSorted(c.report.FailingGateIDs)
	c.report.Passed = c.report.TechnicalStatus == "pass"
	c.report.PromotionEligible = false
	return c.validatedReport(root)
}
