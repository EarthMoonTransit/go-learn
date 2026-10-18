package validation

import (
	"fmt"
	"sort"

	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/taxonomy"
)

// NumericSkillCoverage counts encounters only. It establishes no semantic or
// learner-mastery claim. Every row belongs to the complete active inventory.
type NumericSkillCoverage struct {
	SkillID         string `yaml:"skill_id" json:"skill_id"`
	DomainID        string `yaml:"domain_id" json:"domain_id"`
	Scope           string `yaml:"scope" json:"scope"`
	ActualPrimary   int    `yaml:"actual_primary" json:"actual_primary"`
	RequiredPrimary int    `yaml:"required_primary" json:"required_primary"`
	PrimaryDeficit  int    `yaml:"primary_deficit" json:"primary_deficit"`
	ActualTotal     int    `yaml:"actual_total" json:"actual_total"`
	RequiredTotal   int    `yaml:"required_total" json:"required_total"`
	TotalDeficit    int    `yaml:"total_deficit" json:"total_deficit"`
}

type NumericScopeCoverage struct {
	SkillIDs        []string `yaml:"skill_ids" json:"skill_ids"`
	ActualPrimary   int      `yaml:"actual_primary" json:"actual_primary"`
	RequiredPrimary int      `yaml:"required_primary" json:"required_primary"`
	PrimaryDeficit  int      `yaml:"primary_deficit" json:"primary_deficit"`
	ActualTotal     int      `yaml:"actual_total" json:"actual_total"`
	RequiredTotal   int      `yaml:"required_total" json:"required_total"`
	TotalDeficit    int      `yaml:"total_deficit" json:"total_deficit"`
	MinimaMet       bool     `yaml:"minima_met" json:"minima_met"`
}

type NumericReadyTarget struct {
	Actual int  `yaml:"actual" json:"actual"`
	Min    int  `yaml:"min" json:"min"`
	Max    int  `yaml:"max" json:"max"`
	Met    bool `yaml:"met" json:"met"`
}

// NumericCoverage is arithmetic, not an evidence-eligibility verdict. Root
// validation may publish it only with successful ordinary validation and a
// current collection whose package reports all pass.
type NumericCoverage struct {
	TaxonomyVersion   int                    `yaml:"taxonomy_version" json:"taxonomy_version"`
	Skills            []NumericSkillCoverage `yaml:"skills" json:"skills"`
	Foundation        NumericScopeCoverage   `yaml:"foundation" json:"foundation"`
	Specialist        NumericScopeCoverage   `yaml:"specialist" json:"specialist"`
	Full              NumericScopeCoverage   `yaml:"full" json:"full"`
	ReadyTarget       NumericReadyTarget     `yaml:"ready_target" json:"ready_target"`
	FullNumericPassed bool                   `yaml:"full_numeric_passed" json:"full_numeric_passed"`
}

func (result *Result) populateNumericCoverage(root string, checks *tasking.CheckCollection, ordinaryValid bool) {
	result.NumericCoverage = nil
	result.SelectedSkillIDs = nil
	result.NumericUnavailable = ""
	if !ordinaryValid {
		result.NumericUnavailable = "ordinary repository validation failed"
		return
	}
	if result.Curriculum.Version != 4 {
		result.NumericUnavailable = "scoped numeric evaluation requires taxonomy v4"
		return
	}
	if checks == nil {
		result.NumericUnavailable = "a current successful task-check collection is required"
		return
	}
	// ValidateRoot has already validated this effective collection's freshness
	// while constructing the aggregate. A failed report can still be fresh.
	if err := tasking.RequirePassedCheckReports(*checks); err != nil {
		result.NumericUnavailable = err.Error()
		return
	}
	scopes, err := taxonomy.LoadCoverageScopes(root)
	if err != nil {
		result.NumericUnavailable = err.Error()
		return
	}
	value, err := evaluateNumericCoverage(result.Curriculum, result.Coverage, result.ReadyTasks, scopes)
	if err != nil {
		result.NumericUnavailable = err.Error()
		return
	}
	result.NumericCoverage = &value
	ids := value.Full.SkillIDs
	if result.SelectedScope == "foundation" {
		ids = value.Foundation.SkillIDs
	}
	result.SelectedSkillIDs = append([]string(nil), ids...)
}

// evaluateNumericCoverage uses the already validated curriculum and explicit
// scopes, with the complete rows from calculateCoverage. It does not filter
// tasks, award prerequisites, or replace ordinary repository validation.
func evaluateNumericCoverage(c Curriculum, coverage []Coverage, ready int, scopes taxonomy.CoverageScopes) (NumericCoverage, error) {
	var result NumericCoverage
	if c.Version != 4 || scopes.SchemaVersion != 1 || scopes.TaxonomyVersion != 4 || scopes.Decision != taxonomy.Decision {
		return result, fmt.Errorf("numeric coverage: explicit ADR0031 v4 scopes required")
	}
	if ready < 0 || c.TargetReadyTasks.Min < 0 || c.TargetReadyTasks.Max < c.TargetReadyTasks.Min {
		return result, fmt.Errorf("numeric coverage: invalid ready count or target")
	}
	expected := map[string]Coverage{}
	for _, domain := range c.Domains {
		for _, skill := range domain.Skills {
			if !isAssignable(skill) {
				continue
			}
			if _, found := expected[skill.ID]; found {
				return result, fmt.Errorf("numeric coverage: duplicate curriculum skill %s", skill.ID)
			}
			expected[skill.ID] = Coverage{SkillID: skill.ID, DomainID: domain.ID, RequiredPrimary: skill.MinimumOccurrences.Primary, RequiredTotal: skill.MinimumOccurrences.Total}
		}
	}
	if len(expected) != 145 || len(coverage) != len(expected) || len(scopes.Foundation) != 141 || len(scopes.Specialist) != 4 {
		return result, fmt.Errorf("numeric coverage: complete 145-row / 141-foundation / 4-specialist inventory required")
	}
	classification := map[string]string{}
	for _, group := range []struct {
		name string
		ids  []string
	}{{"foundation", scopes.Foundation}, {"specialist", scopes.Specialist}} {
		for _, id := range group.ids {
			if _, exists := expected[id]; !exists || classification[id] != "" {
				return result, fmt.Errorf("numeric coverage: unknown or duplicate scope skill %s", id)
			}
			classification[id] = group.name
		}
	}
	rows := map[string]Coverage{}
	for _, row := range coverage {
		want, exists := expected[row.SkillID]
		_, duplicate := rows[row.SkillID]
		if !exists || duplicate {
			return result, fmt.Errorf("numeric coverage: unknown or duplicate coverage skill %s", row.SkillID)
		}
		if row.DomainID != want.DomainID || row.RequiredPrimary != want.RequiredPrimary || row.RequiredTotal != want.RequiredTotal {
			return result, fmt.Errorf("numeric coverage: inconsistent required inventory for %s", row.SkillID)
		}
		if row.RequiredPrimary < 0 || row.RequiredTotal < row.RequiredPrimary || row.ActualPrimary < 0 || row.ActualTotal < row.ActualPrimary {
			return result, fmt.Errorf("numeric coverage: invalid counts for %s", row.SkillID)
		}
		rows[row.SkillID] = row
	}
	ids := make([]string, 0, len(expected))
	for id := range expected {
		ids = append(ids, id)
	}
	sort.Strings(ids)
	result.TaxonomyVersion = c.Version
	for _, id := range ids {
		row := rows[id]
		item := NumericSkillCoverage{
			SkillID: id, DomainID: row.DomainID, Scope: classification[id],
			ActualPrimary: row.ActualPrimary, RequiredPrimary: row.RequiredPrimary, PrimaryDeficit: max(0, row.RequiredPrimary-row.ActualPrimary),
			ActualTotal: row.ActualTotal, RequiredTotal: row.RequiredTotal, TotalDeficit: max(0, row.RequiredTotal-row.ActualTotal),
		}
		result.Skills = append(result.Skills, item)
		addNumericRow(&result.Full, item)
		if item.Scope == "foundation" {
			addNumericRow(&result.Foundation, item)
		} else {
			addNumericRow(&result.Specialist, item)
		}
	}
	target := c.TargetReadyTasks
	result.ReadyTarget = NumericReadyTarget{Actual: ready, Min: target.Min, Max: target.Max, Met: ready >= target.Min && ready <= target.Max}
	result.FullNumericPassed = result.Full.MinimaMet && result.ReadyTarget.Met
	return result, nil
}

func addNumericRow(scope *NumericScopeCoverage, row NumericSkillCoverage) {
	scope.SkillIDs = append(scope.SkillIDs, row.SkillID)
	scope.ActualPrimary += row.ActualPrimary
	scope.RequiredPrimary += row.RequiredPrimary
	scope.PrimaryDeficit += row.PrimaryDeficit
	scope.ActualTotal += row.ActualTotal
	scope.RequiredTotal += row.RequiredTotal
	scope.TotalDeficit += row.TotalDeficit
	scope.MinimaMet = scope.PrimaryDeficit == 0 && scope.TotalDeficit == 0
}
