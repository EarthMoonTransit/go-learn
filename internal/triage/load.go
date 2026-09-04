package triage

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"go-practice-bank/internal/safeio"
	"gopkg.in/yaml.v3"
)

func Load(root string) (Bundle, error) {
	var result Bundle
	for _, item := range []struct {
		path string
		dst  any
	}{
		{"catalog/reviews/exercism-go.yaml", &result.Reviews},
		{"catalog/clusters/exercism-go.yaml", &result.Clusters},
		{"catalog/seed-plans/exercism-foundation-pilot.yaml", &result.Seeds},
	} {
		if err := decodeStrict(root, item.path, item.dst); err != nil {
			return Bundle{}, err
		}
	}
	return result, nil
}

func LoadSkillContracts(root string) (int, map[string]SkillContract, error) {
	var document struct {
		Version          int       `yaml:"version"`
		Track            string    `yaml:"track"`
		TargetReadyTasks yaml.Node `yaml:"target_ready_tasks"`
		CoveragePolicy   yaml.Node `yaml:"coverage_policy"`
		Domains          []struct {
			ID     string `yaml:"id"`
			Title  string `yaml:"title"`
			Skills []struct {
				ID                       string    `yaml:"id"`
				Domain                   string    `yaml:"domain"`
				Definition               string    `yaml:"definition"`
				Prerequisites            []string  `yaml:"prerequisites"`
				CommonMisconceptions     []string  `yaml:"common_misconceptions"`
				RecommendedExerciseTypes []string  `yaml:"recommended_exercise_types"`
				MinimumOccurrences       yaml.Node `yaml:"minimum_occurrences"`
				Mastery                  []string  `yaml:"mastery"`
				EvidenceRequirements     yaml.Node `yaml:"evidence_requirements"`
				Status                   string    `yaml:"status"`
				Assignable               bool      `yaml:"assignable"`
				TaxonomyVersion          int       `yaml:"taxonomy_version"`
				ReplacedBy               []string  `yaml:"replaced_by,omitempty"`
			} `yaml:"skills"`
		} `yaml:"domains"`
	}
	if err := decodeStrict(root, "curriculum/skills.yaml", &document); err != nil {
		return 0, nil, err
	}
	contracts := map[string]SkillContract{}
	for _, domain := range document.Domains {
		for _, skill := range domain.Skills {
			contracts[skill.ID] = SkillContract{
				ID: skill.ID, Domain: skill.Domain, Definition: skill.Definition,
				Prerequisites: skill.Prerequisites, CommonMisconceptions: skill.CommonMisconceptions,
				RecommendedExerciseTypes: skill.RecommendedExerciseTypes,
				Status:                   skill.Status, Assignable: skill.Assignable,
			}
		}
	}
	return document.Version, contracts, nil
}

func decodeStrict(root, relative string, destination any) error {
	b, err := safeio.ReadFile(root, relative, 16<<20)
	if err != nil {
		return fmt.Errorf("read %s: %w", relative, err)
	}
	decoder := yaml.NewDecoder(bytes.NewReader(b))
	decoder.KnownFields(true)
	if err := decoder.Decode(destination); err != nil {
		return fmt.Errorf("decode %s: %w", relative, err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return fmt.Errorf("decode %s: multiple YAML documents are not allowed", relative)
	} else if !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode trailing YAML in %s: %w", relative, err)
	}
	return nil
}
