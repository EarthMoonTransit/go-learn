package validation

import (
	"path/filepath"
	"reflect"
	"slices"
	"strings"
	"testing"

	"go-practice-bank/internal/taxonomy"
)

func numericFixture(t *testing.T) (Curriculum, []Coverage, taxonomy.CoverageScopes) {
	t.Helper()
	root := repositoryRoot(t)
	var curriculum Curriculum
	if err := decodeStrict(filepath.Join(root, "curriculum/skills.yaml"), &curriculum); err != nil {
		t.Fatal(err)
	}
	scopes, err := taxonomy.LoadCoverageScopes(root)
	if err != nil {
		t.Fatal(err)
	}
	var rows []Coverage
	for _, domain := range curriculum.Domains {
		for _, skill := range domain.Skills {
			if isAssignable(skill) {
				minimum := skill.MinimumOccurrences
				rows = append(rows, Coverage{SkillID: skill.ID, DomainID: domain.ID, RequiredPrimary: minimum.Primary, RequiredTotal: minimum.Total, ActualPrimary: minimum.Primary, ActualTotal: minimum.Total})
			}
		}
	}
	return curriculum, rows, scopes
}

func TestNumericCoverageFullMinimaAndExplicitScopes(t *testing.T) {
	c, rows, scopes := numericFixture(t)
	got, err := evaluateNumericCoverage(c, rows, 418, scopes)
	if err != nil {
		t.Fatal(err)
	}
	if !got.FullNumericPassed || !got.Foundation.MinimaMet || !got.Specialist.MinimaMet || !got.ReadyTarget.Met || len(got.Skills) != 145 {
		t.Fatalf("exact minima did not pass: %+v", got)
	}
	for _, group := range []struct {
		got                   NumericScopeCoverage
		count, primary, total int
	}{{got.Full, 145, 418, 708}, {got.Foundation, 141, 410, 692}, {got.Specialist, 4, 8, 16}} {
		if len(group.got.SkillIDs) != group.count || !slices.IsSorted(group.got.SkillIDs) || group.got.RequiredPrimary != group.primary || group.got.ActualPrimary != group.primary || group.got.RequiredTotal != group.total || group.got.ActualTotal != group.total || group.got.PrimaryDeficit != 0 || group.got.TotalDeficit != 0 {
			t.Fatalf("incorrect scope arithmetic: %+v", group.got)
		}
	}
	// Ordering of valid metadata/collection rows cannot change the arithmetic.
	slices.Reverse(rows)
	slices.Reverse(scopes.Foundation)
	slices.Reverse(scopes.Specialist)
	again, err := evaluateNumericCoverage(c, rows, 418, scopes)
	if err != nil || !reflect.DeepEqual(got, again) {
		t.Fatalf("row order changed evaluation: %v", err)
	}
}

func TestNumericCoverageFoundationCanPassWithSpecialistDeficits(t *testing.T) {
	c, rows, scopes := numericFixture(t)
	for i := range rows {
		if slices.Contains(scopes.Specialist, rows[i].SkillID) {
			rows[i].ActualPrimary, rows[i].ActualTotal = 0, 0
		}
	}
	got, err := evaluateNumericCoverage(c, rows, 410, scopes)
	if err != nil {
		t.Fatal(err)
	}
	if !got.Foundation.MinimaMet || got.Specialist.MinimaMet || got.Full.MinimaMet || got.FullNumericPassed || !got.ReadyTarget.Met || got.Specialist.PrimaryDeficit != 8 || got.Specialist.TotalDeficit != 16 || got.Full.PrimaryDeficit != 8 || got.Full.TotalDeficit != 16 {
		t.Fatalf("scope selection hid specialist/full deficits: %+v", got)
	}
	if len(got.Skills) != 145 || got.ReadyTarget.Min != 350 || got.ReadyTarget.Max != 500 {
		t.Fatal("foundation arithmetic filtered the full inventory or invented a target")
	}
}

func TestNumericCoverageSurplusAndTotalCannotCompensateDeficits(t *testing.T) {
	base, original, scopes := numericFixture(t)
	for _, missing := range []string{"primary", "total", "both"} {
		t.Run(missing, func(t *testing.T) {
			c, rows := base, slices.Clone(original)
			deficitID, surplusID := scopes.Foundation[0], scopes.Foundation[1]
			for i := range rows {
				switch rows[i].SkillID {
				case deficitID:
					if missing != "total" {
						rows[i].ActualPrimary--
					}
					if missing != "primary" {
						rows[i].ActualTotal--
					}
				case surplusID:
					rows[i].ActualPrimary += 100
					rows[i].ActualTotal += 100
				}
			}
			got, err := evaluateNumericCoverage(c, rows, 418, scopes)
			if err != nil {
				t.Fatal(err)
			}
			primary, total := 1, 1
			if missing == "primary" {
				total = 0
			}
			if missing == "total" {
				primary = 0
			}
			if got.Foundation.MinimaMet || got.FullNumericPassed || got.Foundation.PrimaryDeficit != primary || got.Foundation.TotalDeficit != total || got.Full.ActualPrimary < got.Full.RequiredPrimary || got.Full.ActualTotal < got.Full.RequiredTotal {
				t.Fatalf("surplus hid per-skill %s deficit: %+v", missing, got.Foundation)
			}
		})
	}
}

func TestNumericCoverageRetainsIndependentWholeBankTarget(t *testing.T) {
	c, rows, scopes := numericFixture(t)
	// 501 credited primary encounters can meet all minima but exceed the target.
	rows[0].ActualPrimary += 83
	rows[0].ActualTotal += 83
	high, err := evaluateNumericCoverage(c, rows, 501, scopes)
	if err != nil || !high.Full.MinimaMet || high.ReadyTarget.Met || high.FullNumericPassed {
		t.Fatalf("high ready count accepted: %+v, %v", high.ReadyTarget, err)
	}
	// Below 350, the 418 primary minima cannot all be met in a real bank.
	// This synthetic arithmetic fixture makes both deficits explicit.
	for i := range rows {
		rows[i].ActualPrimary, rows[i].ActualTotal = 0, 0
	}
	low, err := evaluateNumericCoverage(c, rows, 349, scopes)
	if err != nil || low.Full.MinimaMet || low.ReadyTarget.Met || low.FullNumericPassed || low.Full.PrimaryDeficit != 418 {
		t.Fatalf("low ready target lost: %+v, %v", low.ReadyTarget, err)
	}
}

func TestNumericCoverageRejectsIncompleteOrInconsistentInputs(t *testing.T) {
	base, original, originalScopes := numericFixture(t)
	for _, tc := range []struct {
		name, want string
		change     func(*Curriculum, *[]Coverage, *taxonomy.CoverageScopes, *int)
	}{
		{"missing row", "complete 145-row", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { *r = (*r)[:144] }},
		{"duplicate row", "duplicate coverage skill", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { (*r)[1] = (*r)[0] }},
		{"extra id", "unknown or duplicate coverage skill", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) {
			(*r)[0].SkillID = "functions.future"
		}},
		{"changed minimum", "inconsistent required inventory", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { (*r)[0].RequiredPrimary++ }},
		{"changed domain", "inconsistent required inventory", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { (*r)[0].DomainID = "other" }},
		{"negative count", "invalid counts", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { (*r)[0].ActualPrimary = -1 }},
		{"primary above total", "invalid counts", func(_ *Curriculum, r *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) {
			(*r)[0].ActualPrimary = (*r)[0].ActualTotal + 1
		}},
		{"negative ready", "invalid ready count", func(_ *Curriculum, _ *[]Coverage, _ *taxonomy.CoverageScopes, n *int) { *n = -1 }},
		{"invalid target", "invalid ready count or target", func(c *Curriculum, _ *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { c.TargetReadyTasks.Min = 501 }},
		{"v3", "explicit ADR0031 v4 scopes", func(c *Curriculum, _ *[]Coverage, _ *taxonomy.CoverageScopes, _ *int) { c.Version = 3 }},
		{"scope version", "explicit ADR0031 v4 scopes", func(_ *Curriculum, _ *[]Coverage, s *taxonomy.CoverageScopes, _ *int) { s.TaxonomyVersion = 3 }},
		{"scope missing", "complete 145-row", func(_ *Curriculum, _ *[]Coverage, s *taxonomy.CoverageScopes, _ *int) {
			s.Foundation = s.Foundation[:140]
		}},
		{"scope duplicate", "unknown or duplicate scope skill", func(_ *Curriculum, _ *[]Coverage, s *taxonomy.CoverageScopes, _ *int) {
			s.Foundation[1] = s.Foundation[0]
		}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			c, rows, scopes := base, slices.Clone(original), originalScopes
			scopes.Foundation, scopes.Specialist = slices.Clone(originalScopes.Foundation), slices.Clone(originalScopes.Specialist)
			ready := 418
			tc.change(&c, &rows, &scopes, &ready)
			if _, err := evaluateNumericCoverage(c, rows, ready, scopes); err == nil || !strings.Contains(err.Error(), tc.want) {
				t.Fatalf("got %v, want %q", err, tc.want)
			}
		})
	}
}
