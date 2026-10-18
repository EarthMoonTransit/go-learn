package validation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strings"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"go-practice-bank/internal/catalog"
	curriculumdata "go-practice-bank/internal/curriculum"
	"go-practice-bank/internal/learner"
	sourcepolicy "go-practice-bank/internal/sources"
	"go-practice-bank/internal/tasking"
	"go-practice-bank/internal/taxonomy"
	"go-practice-bank/internal/toolchainpolicy"
	"go-practice-bank/internal/triage"
	"gopkg.in/yaml.v3"
)

var (
	skillIDPattern = regexp.MustCompile(`^[a-z][a-z0-9-]*\.[a-z][a-z0-9-]*$`)
	taskIDPattern  = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*-[0-9]{3}$`)
	shaPattern     = regexp.MustCompile(`^[0-9a-f]{40}$`)
	sha256Pattern  = regexp.MustCompile(`^[0-9a-f]{64}$`)
)

var exerciseTypes = set("implement", "debug", "predict", "refactor", "test-writing", "performance")
var difficultyValues = set("introductory", "basic", "intermediate", "advanced")
var qualityValues = set("draft", "review", "validated", "ready", "quarantined", "retired")
var transformValues = set("original", "adapted", "translated-adaptation", "combined", "metadata-only")
var usageValues = set("adapt", "reference-only", "concepts-only", "blocked")

func ValidateRoot(root string, opts Options) (Result, error) {
	var result Result
	var problems []error
	if err := learner.ValidateTheory(root); err != nil {
		problems = append(problems, fmt.Errorf("theory: %w", err))
	}
	if err := decodeStrict(filepath.Join(root, "curriculum", "skills.yaml"), &result.Curriculum); err != nil {
		return result, err
	}
	taxonomySchema := "schemas/taxonomy-v3.schema.json"
	if result.Curriculum.Version == 4 {
		taxonomySchema = "schemas/taxonomy-v4.schema.json"
	}

	taskSchema, err := compileTaskSchema(filepath.Join(root, "schemas", "task.schema.json"))
	if err != nil {
		problems = append(problems, err)
	}
	taskV2Schema, err := compileTaskSchema(filepath.Join(root, "schemas", "task-v2.schema.json"))
	if err != nil {
		problems = append(problems, err)
	}
	staticSchemas := []struct {
		data, schema, label string
	}{
		{"curriculum/skills.yaml", taxonomySchema, "taxonomy"},
		{"curriculum/migrations/v2-to-v3.yaml", "schemas/taxonomy-migration.schema.json", "taxonomy migration"},
		{"curriculum/mappings/exercism-go.yaml", "schemas/exercism-concept-map.schema.json", "Exercism concept map"},
		{"curriculum/evidence/exercism-go.yaml", "schemas/skill-evidence.schema.json", "Exercism skill evidence"},
		{"catalog/reviews/exercism-go.yaml", "schemas/candidate-review.schema.json", "Exercism candidate reviews"},
		{"catalog/clusters/exercism-go.yaml", "schemas/candidate-clusters.schema.json", "Exercism candidate clusters"},
		{"catalog/seed-plans/exercism-foundation-pilot.yaml", "schemas/seed-plan.schema.json", "Exercism seed plan"},
		{"config/toolchain.yaml", "schemas/toolchain-policy.schema.json", "toolchain policy"},
		{"config/project-identity.yaml", "schemas/project-identity.schema.json", "project identity"},
		{"config/outbound-license.yaml", "schemas/outbound-license-policy.schema.json", "outbound license policy"},
	}
	if result.Curriculum.Version == 4 {
		staticSchemas = append(staticSchemas,
			struct{ data, schema, label string }{taxonomy.V3Path, "schemas/taxonomy-v3.schema.json", "historical v3 taxonomy"},
			struct{ data, schema, label string }{taxonomy.MigrationPath, "schemas/taxonomy-v3-to-v4.schema.json", "taxonomy v3 to v4 migration"},
			struct{ data, schema, label string }{taxonomy.ScopePath, taxonomy.ScopeSchemaPath, "coverage scopes"},
		)
	}
	if _, err := os.Lstat(filepath.Join(root, "curriculum/topics.yaml")); err == nil {
		staticSchemas = append(staticSchemas, struct{ data, schema, label string }{"curriculum/topics.yaml", "schemas/topic.schema.json", "theory topics"})
	}
	if policy, policyErr := toolchainpolicy.Load(root); policyErr != nil {
		problems = append(problems, policyErr)
	} else if policyErr := toolchainpolicy.ValidateRepository(root, policy); policyErr != nil {
		problems = append(problems, policyErr)
	}
	for _, schemaName := range []string{
		"topic.schema.json", "task-v2.schema.json", "task-lifecycle.schema.json", "task-review.schema.json", "task-provenance.schema.json", "task-provenance-v2.schema.json", "project-identity.schema.json", "outbound-license-policy.schema.json", "task-quality-report.schema.json", "task-check-collection.schema.json", "task-aggregate-report.schema.json",
	} {
		if _, schemaErr := compileYAMLSchema(filepath.Join(root, "schemas", schemaName)); schemaErr != nil {
			problems = append(problems, schemaErr)
		}
	}
	for _, item := range staticSchemas {
		schema, schemaErr := compileYAMLSchema(filepath.Join(root, filepath.FromSlash(item.schema)))
		if schemaErr != nil {
			problems = append(problems, schemaErr)
			continue
		}
		if schemaErr := validateYAMLSchema(filepath.Join(root, filepath.FromSlash(item.data)), schema, item.label); schemaErr != nil {
			problems = append(problems, schemaErr)
		}
	}

	skills, domainBySkill, curriculumProblems := validateCurriculum(result.Curriculum)
	problems = append(problems, curriculumProblems...)
	reviewedSkills := skills
	reviewedVersion := result.Curriculum.Version
	if result.Curriculum.Version == 4 {
		path, version, err := taxonomy.ReviewedPath(root)
		if err != nil {
			return result, errors.Join(append(problems, err)...)
		}
		var historical Curriculum
		if err := decodeStrict(filepath.Join(root, filepath.FromSlash(path)), &historical); err != nil {
			return result, err
		}
		var historicalProblems []error
		reviewedSkills, _, historicalProblems = validateCurriculum(historical)
		problems = append(problems, historicalProblems...)
		reviewedVersion = version
		if err := taxonomy.CheckMigrationReport(root); err != nil {
			problems = append(problems, err)
		}
	}
	if err := decodeStrict(filepath.Join(root, "curriculum", "migrations", "v2-to-v3.yaml"), &result.Migration); err != nil {
		return result, err
	}
	migrationIndex, migrationProblems := validateMigration(root, result.Migration, reviewedSkills)
	problems = append(problems, migrationProblems...)

	if err := decodeStrict(filepath.Join(root, "sources", "registry.yaml"), &result.Registry); err != nil {
		return result, err
	}
	sources, registryProblems := validateRegistry(result.Registry)
	problems = append(problems, registryProblems...)
	problems = append(problems, validateOutboundLicenses(root, result.Registry)...)
	if exercismSource, ok := sources["exercism-go"]; ok {
		mapping, mapErr := curriculumdata.LoadConceptMap(root)
		evidence, evidenceErr := curriculumdata.LoadEvidence(root)
		inventory, inventoryErr := curriculumdata.LoadConceptInventory(root)
		candidates, candidatesErr := catalog.ReadJSONL(root, "catalog/candidates/exercism-go.jsonl")
		if mapErr != nil {
			problems = append(problems, mapErr)
		}
		if evidenceErr != nil {
			problems = append(problems, evidenceErr)
		}
		if inventoryErr != nil {
			problems = append(problems, inventoryErr)
		} else {
			problems = append(problems, curriculumdata.ValidateInventory(inventory, exercismSource))
		}
		if candidatesErr != nil {
			problems = append(problems, candidatesErr)
		}
		if mapErr == nil && evidenceErr == nil && inventoryErr == nil && candidatesErr == nil {
			problems = append(problems, curriculumdata.Validate(mapping, evidence, exercismSource, reviewedVersion, skillSet(reviewedSkills), inventory.Concepts, candidates))
			bundle, bundleErr := triage.Load(root)
			taxonomyVersion, contracts, contractsErr := triage.LoadReviewedSkillContracts(root)
			if bundleErr != nil {
				problems = append(problems, bundleErr)
			}
			if contractsErr != nil {
				problems = append(problems, contractsErr)
			}
			if bundleErr == nil && contractsErr == nil {
				problems = append(problems, triage.Validate(bundle, candidates, contracts, taxonomyVersion, exercismSource))
				result.ReviewedCandidates = len(bundle.Reviews.Reviews)
				result.SelectedSeeds = len(bundle.Seeds.SelectedSeeds)
			}
		}
	}

	paths, err := filepath.Glob(filepath.Join(root, "tasks", "ready", "*", "task.yaml"))
	if err != nil {
		problems = append(problems, fmt.Errorf("find task manifests: %w", err))
	}
	sort.Strings(paths)
	seenTaskIDs := map[string]string{}
	for _, path := range paths {
		version, versionErr := taskManifestVersion(path)
		if versionErr != nil {
			problems = append(problems, versionErr)
			continue
		}
		var task Task
		switch version {
		case 1:
			if taskSchema != nil {
				if err := validateTaskSchema(path, taskSchema); err != nil {
					problems = append(problems, err)
				}
			}
			if err := decodeStrict(path, &task); err != nil {
				problems = append(problems, err)
				continue
			}
		case 2:
			if taskV2Schema != nil {
				if err := validateTaskSchema(path, taskV2Schema); err != nil {
					problems = append(problems, err)
				}
			}
			var taskV2 tasking.TaskV2
			if err := decodeStrict(path, &taskV2); err != nil {
				problems = append(problems, err)
				continue
			}
			task = normalizeValidationTaskV2(taskV2)
			problems = append(problems, validateTaskV2(path, taskV2, skills, migrationIndex)...)
		default:
			problems = append(problems, fmt.Errorf("%s: unsupported task schema_version %d", path, version))
			continue
		}
		result.Tasks = append(result.Tasks, task)
		if first, ok := seenTaskIDs[task.ID]; ok {
			problems = append(problems, fmt.Errorf("%s: duplicate task id %q (first in %s)", path, task.ID, first))
		} else {
			seenTaskIDs[task.ID] = path
		}
		if version == 1 {
			problems = append(problems, validateTask(root, path, task, skills, sources, migrationIndex)...)
		}
		if task.QualityStatus == result.Curriculum.CoveragePolicy.CountedStatus {
			result.ReadyTasks++
		}
	}

	if len(paths) == 0 && !opts.AllowEmpty {
		problems = append(problems, errors.New("tasks/ready contains no task.yaml manifests"))
	}
	var aggregate tasking.AggregateReport
	var aggregateErr error
	switch {
	case opts.Checks != nil:
		aggregate, aggregateErr = tasking.BuildAggregateReportFromChecks(root, *opts.Checks)
	case opts.RunCommands:
		var checks tasking.CheckCollection
		checks, aggregateErr = tasking.CheckAll(root, nil)
		if aggregateErr == nil {
			aggregate, aggregateErr = tasking.BuildAggregateReportFromChecks(root, checks)
		}
	default:
		aggregateErr = decodeStrict(filepath.Join(root, "reports", "task-quality.yaml"), &aggregate)
	}
	if aggregateErr != nil {
		problems = append(problems, fmt.Errorf("task lifecycle inventory: %w", aggregateErr))
	} else if err := tasking.ValidateAggregateReport(root, aggregate); err != nil {
		problems = append(problems, fmt.Errorf("task lifecycle report schema: %w", err))
	} else if err := tasking.ValidateAggregateInputs(root, aggregate); err != nil {
		problems = append(problems, fmt.Errorf("task lifecycle report inputs: %w", err))
	} else {
		expectedYAML, marshalErr := tasking.MarshalAggregateYAML(aggregate)
		if marshalErr != nil {
			problems = append(problems, fmt.Errorf("marshal task lifecycle report: %w", marshalErr))
		} else {
			for _, artifact := range []struct {
				path     string
				expected []byte
			}{
				{"reports/task-quality.md", tasking.MarshalAggregateMarkdown(aggregate)},
				{"reports/task-quality.yaml", expectedYAML},
			} {
				actual, readErr := os.ReadFile(filepath.Join(root, filepath.FromSlash(artifact.path)))
				if readErr != nil {
					problems = append(problems, fmt.Errorf("read deterministic %s: %w", artifact.path, readErr))
				} else if !bytes.Equal(actual, artifact.expected) {
					problems = append(problems, fmt.Errorf("%s is stale; run task-report", artifact.path))
				}
			}
		}
	}
	problems = append(problems, validateMigrationTaskReferences(root, result.Migration, result.Tasks)...)
	result.Coverage = calculateCoverage(result.Curriculum, result.Tasks, domainBySkill, migrationIndex)
	if opts.EnforceCoverage {
		problems = append(problems, validateCoverage(result)...)
	}

	return result, errors.Join(problems...)
}

func validateOutboundLicenses(root string, registry Registry) []error {
	var policy struct {
		SchemaVersion           int      `yaml:"schema_version"`
		OriginalContentLicense  string   `yaml:"original_content_license"`
		ThirdPartyRule          string   `yaml:"third_party_rule"`
		CompatibleSourceLicense []string `yaml:"compatible_source_licenses"`
		NoticesFile             string   `yaml:"notices_file"`
	}
	if err := decodeStrict(filepath.Join(root, "config", "outbound-license.yaml"), &policy); err != nil {
		return []error{err}
	}
	if policy.OriginalContentLicense != "Apache-2.0" {
		return []error{errors.New("outbound policy must license original project content as Apache-2.0")}
	}
	licenseText, err := os.ReadFile(filepath.Join(root, "LICENSE"))
	if err != nil || !bytes.Contains(licenseText, []byte("Apache License")) || !bytes.Contains(licenseText, []byte("Version 2.0")) {
		return []error{errors.New("repository LICENSE must contain Apache License 2.0 terms")}
	}
	compatible := set(policy.CompatibleSourceLicense...)
	var problems []error
	for _, source := range registry.Sources {
		if source.UsageMode == "adapt" && !compatible[source.License.Expression] {
			problems = append(problems, fmt.Errorf("outbound policy does not declare compatibility for adapted source %q license %q", source.ID, source.License.Expression))
		}
	}
	return problems
}

func decodeStrict(path string, dst any) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", path, err)
	}
	dec := yaml.NewDecoder(bytes.NewReader(b))
	dec.KnownFields(true)
	if err := dec.Decode(dst); err != nil {
		return fmt.Errorf("decode %s: %w", path, err)
	}
	var extra any
	if err := dec.Decode(&extra); err == nil {
		return fmt.Errorf("decode %s: multiple YAML documents are not allowed", path)
	} else if !errors.Is(err, io.EOF) {
		return fmt.Errorf("decode trailing YAML in %s: %w", path, err)
	}
	return nil
}

func taskManifestVersion(path string) (int, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	var header struct {
		SchemaVersion int `yaml:"schema_version"`
	}
	if err := yaml.Unmarshal(content, &header); err != nil {
		return 0, fmt.Errorf("decode %s schema_version: %w", path, err)
	}
	if header.SchemaVersion == 0 {
		return 0, fmt.Errorf("%s: schema_version is required", path)
	}
	return header.SchemaVersion, nil
}

func normalizeValidationTaskV2(task tasking.TaskV2) Task {
	taskType := map[string]string{
		"implementation": "implement", "debugging": "debug", "refactoring": "refactor", "api-design": "implement",
		"test-writing": "test-writing", "test-repair": "test-writing", "fuzz-test-writing": "test-writing",
		"integration": "implement", "performance": "performance", "code-review-patch": "debug",
	}[task.Activity]
	return Task{
		SchemaVersion: task.SchemaVersion, ID: task.ID, Title: task.Title, TaskType: taskType,
		Skills:        TaskSkills{Primary: task.Skills.Primary, Secondary: task.Skills.Secondary},
		Prerequisites: task.Prerequisites, NewConcepts: task.NewConcepts, Difficulty: task.Difficulty,
		EstimatedTimeMinutes: TimeRange{Min: task.EstimatedTimeMinutes.Min, Max: task.EstimatedTimeMinutes.Max},
		QualityStatus:        task.QualityStatus,
	}
}

func validateTaskV2(path string, task tasking.TaskV2, skills map[string]Skill, migration migrationIndex) []error {
	var problems []error
	if filepath.Base(filepath.Dir(path)) != task.ID {
		problems = append(problems, fmt.Errorf("%s: task id must match parent directory", path))
	}
	if task.EstimatedTimeMinutes.Min > task.EstimatedTimeMinutes.Max {
		problems = append(problems, fmt.Errorf("%s: invalid estimated_time_minutes range", path))
	}
	manifest, err := os.ReadFile(path)
	if err != nil {
		return append(problems, err)
	}
	digest := digestBytes(manifest)
	problems = append(problems, validateTaskSkill(path, task.ID, "skills.primary", task.Skills.Primary, digest, skills, migration)...)
	for field, values := range map[string][]string{"skills.secondary": task.Skills.Secondary, "prerequisites": task.Prerequisites, "new_concepts": task.NewConcepts} {
		for _, duplicate := range duplicates(values) {
			problems = append(problems, fmt.Errorf("%s: duplicate %s skill %q", path, field, duplicate))
		}
		for _, id := range values {
			problems = append(problems, validateTaskSkill(path, task.ID, field, id, digest, skills, migration)...)
		}
	}
	assessed := append([]string{task.Skills.Primary}, task.Skills.Secondary...)
	problems = append(problems, disjoint(path, "primary/secondary", []string{task.Skills.Primary}, task.Skills.Secondary)...)
	problems = append(problems, disjoint(path, "assessed/prerequisite", assessed, task.Prerequisites)...)
	problems = append(problems, disjoint(path, "new/prerequisite", task.NewConcepts, task.Prerequisites)...)
	return problems
}

func compileTaskSchema(path string) (*jsonschema.Schema, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read task schema: %w", err)
	}
	var schema map[string]any
	if err := json.Unmarshal(b, &schema); err != nil {
		return nil, fmt.Errorf("parse task schema: %w", err)
	}
	if schema["$schema"] == nil || schema["required"] == nil || schema["properties"] == nil {
		return nil, errors.New("task schema must declare $schema, required, and properties")
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	if err := compiler.AddResource("task.schema.json", bytes.NewReader(b)); err != nil {
		return nil, fmt.Errorf("load task schema: %w", err)
	}
	compiled, err := compiler.Compile("task.schema.json")
	if err != nil {
		return nil, fmt.Errorf("compile task schema: %w", err)
	}
	return compiled, nil
}

func compileYAMLSchema(path string) (*jsonschema.Schema, error) {
	b, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read schema %s: %w", path, err)
	}
	var document map[string]any
	if err := json.Unmarshal(b, &document); err != nil {
		return nil, fmt.Errorf("parse schema %s: %w", path, err)
	}
	if document["$schema"] == nil || document["type"] == nil {
		return nil, fmt.Errorf("schema %s must declare $schema and type", path)
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	resource := filepath.Base(path)
	if err := compiler.AddResource(resource, bytes.NewReader(b)); err != nil {
		return nil, fmt.Errorf("load schema %s: %w", path, err)
	}
	compiled, err := compiler.Compile(resource)
	if err != nil {
		return nil, fmt.Errorf("compile schema %s: %w", path, err)
	}
	return compiled, nil
}

func validateYAMLSchema(path string, schema *jsonschema.Schema, label string) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s for schema validation: %w", path, err)
	}
	var yamlValue any
	if err := yaml.Unmarshal(b, &yamlValue); err != nil {
		return fmt.Errorf("decode %s for schema validation: %w", path, err)
	}
	jsonBytes, err := json.Marshal(yamlValue)
	if err != nil {
		return fmt.Errorf("normalize %s for schema validation: %w", path, err)
	}
	var jsonValue any
	if err := json.Unmarshal(jsonBytes, &jsonValue); err != nil {
		return fmt.Errorf("normalize %s as JSON: %w", path, err)
	}
	if err := schema.Validate(jsonValue); err != nil {
		return fmt.Errorf("%s: %s schema violation: %w", path, label, err)
	}
	return nil
}

func validateTaskSchema(path string, schema *jsonschema.Schema) error {
	b, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s for schema validation: %w", path, err)
	}
	var yamlValue any
	if err := yaml.Unmarshal(b, &yamlValue); err != nil {
		return fmt.Errorf("decode %s for schema validation: %w", path, err)
	}
	jsonBytes, err := json.Marshal(yamlValue)
	if err != nil {
		return fmt.Errorf("normalize %s for schema validation: %w", path, err)
	}
	var jsonValue any
	if err := json.Unmarshal(jsonBytes, &jsonValue); err != nil {
		return fmt.Errorf("normalize %s as JSON: %w", path, err)
	}
	if err := schema.Validate(jsonValue); err != nil {
		return fmt.Errorf("%s: task schema violation: %w", path, err)
	}
	return nil
}

func validateCurriculum(c Curriculum) (map[string]Skill, map[string]string, []error) {
	var problems []error
	if c.Version != 3 && c.Version != 4 {
		problems = append(problems, fmt.Errorf("curriculum: version must be 3 or 4, got %d", c.Version))
	}
	if c.Track != "go-fundamentals" {
		problems = append(problems, fmt.Errorf("curriculum: track must be go-fundamentals, got %q", c.Track))
	}
	if c.TargetReadyTasks.Min < 350 || c.TargetReadyTasks.Max > 500 || c.TargetReadyTasks.Min > c.TargetReadyTasks.Max {
		problems = append(problems, fmt.Errorf("curriculum: target_ready_tasks must stay within 350..500"))
	}
	if c.CoveragePolicy.CountedStatus != "ready" || c.CoveragePolicy.PrimarySkillPerTask != 1 || c.CoveragePolicy.PrerequisiteCountsOccurrence {
		problems = append(problems, errors.New("curriculum: coverage policy must count one primary skill on ready tasks and exclude prerequisites"))
	}

	skills := map[string]Skill{}
	domainBySkill := map[string]string{}
	domains := map[string]bool{}
	primaryMinimumTotal := 0
	for _, domain := range c.Domains {
		if !regexp.MustCompile(`^[a-z][a-z0-9-]*$`).MatchString(domain.ID) || domain.Title == "" || len(domain.Skills) == 0 {
			problems = append(problems, fmt.Errorf("curriculum: domain %q must have id, title, and skills", domain.ID))
		}
		if domains[domain.ID] {
			problems = append(problems, fmt.Errorf("curriculum: duplicate domain %q", domain.ID))
		}
		domains[domain.ID] = true
		for _, skill := range domain.Skills {
			if !skillIDPattern.MatchString(skill.ID) || skill.Domain != domain.ID || !strings.HasPrefix(skill.ID, domain.ID+".") {
				problems = append(problems, fmt.Errorf("curriculum: invalid skill id %q in domain %q", skill.ID, domain.ID))
			}
			if _, ok := skills[skill.ID]; ok {
				problems = append(problems, fmt.Errorf("curriculum: duplicate skill %q", skill.ID))
			}
			if skill.Assignable == nil {
				problems = append(problems, fmt.Errorf("curriculum: skill %q must explicitly declare assignable", skill.ID))
			}
			if skill.TaxonomyVersion != c.Version {
				problems = append(problems, fmt.Errorf("curriculum: skill %q must declare taxonomy_version %d", skill.ID, c.Version))
			}
			problems = append(problems, duplicateFieldProblems(skill.ID, "prerequisite", skill.Prerequisites)...)
			problems = append(problems, duplicateFieldProblems(skill.ID, "misconception", skill.CommonMisconceptions)...)
			problems = append(problems, duplicateFieldProblems(skill.ID, "exercise type", skill.RecommendedExerciseTypes)...)
			problems = append(problems, duplicateFieldProblems(skill.ID, "mastery criterion", skill.Mastery)...)
			problems = append(problems, duplicateFieldProblems(skill.ID, "evidence mode", skill.EvidenceRequirements.RequiredModes)...)
			problems = append(problems, duplicateFieldProblems(skill.ID, "observable outcome", skill.EvidenceRequirements.ObservableOutcomes)...)
			for _, typ := range skill.RecommendedExerciseTypes {
				if !exerciseTypes[typ] {
					problems = append(problems, fmt.Errorf("curriculum: skill %q has unknown exercise type %q", skill.ID, typ))
				}
			}
			switch skill.Status {
			case "active":
				if !isAssignable(skill) {
					problems = append(problems, fmt.Errorf("curriculum: active skill %q must be assignable", skill.ID))
				}
				if len(skill.Definition) < 10 || len(skill.CommonMisconceptions) < 2 || len(skill.RecommendedExerciseTypes) < 3 || len(skill.Mastery) < 2 {
					problems = append(problems, fmt.Errorf("curriculum: active skill %q lacks the full teaching contract", skill.ID))
				}
				if !validMinimumTier(skill.MinimumOccurrences) {
					problems = append(problems, fmt.Errorf("curriculum: active skill %q must use minimum tier 2/4, 3/5, or 4/6", skill.ID))
				}
				if skill.EvidenceRequirements.MinimumDistinctTaskTypes < 3 || len(skill.EvidenceRequirements.RequiredModes) != 3 || len(skill.EvidenceRequirements.ObservableOutcomes) < 2 {
					problems = append(problems, fmt.Errorf("curriculum: active skill %q lacks reviewed evidence requirements", skill.ID))
				}
				for _, mode := range skill.EvidenceRequirements.RequiredModes {
					if !set("implementation", "debugging", "verification-or-integration")[mode] {
						problems = append(problems, fmt.Errorf("curriculum: active skill %q has invalid evidence mode %q", skill.ID, mode))
					}
				}
				if len(skill.ReplacedBy) != 0 {
					problems = append(problems, fmt.Errorf("curriculum: active skill %q cannot declare replaced_by", skill.ID))
				}
				primaryMinimumTotal += skill.MinimumOccurrences.Primary
			case "deprecated":
				if isAssignable(skill) || skill.Assignable == nil || *skill.Assignable || len(skill.ReplacedBy) == 0 {
					problems = append(problems, fmt.Errorf("curriculum: deprecated skill %q must be non-assignable and declare replacements", skill.ID))
				}
				if len(skill.Prerequisites) != 0 || len(skill.CommonMisconceptions) != 0 || len(skill.RecommendedExerciseTypes) != 0 || len(skill.Mastery) != 0 || skill.MinimumOccurrences != (Minimum{}) || skill.EvidenceRequirements.MinimumDistinctTaskTypes != 0 || len(skill.EvidenceRequirements.RequiredModes) != 0 || len(skill.EvidenceRequirements.ObservableOutcomes) != 0 {
					problems = append(problems, fmt.Errorf("curriculum: deprecated skill %q must have a zero non-teaching contract", skill.ID))
				}
			default:
				problems = append(problems, fmt.Errorf("curriculum: skill %q has invalid status %q", skill.ID, skill.Status))
			}
			skills[skill.ID] = skill
			domainBySkill[skill.ID] = domain.ID
		}
	}
	if primaryMinimumTotal < c.TargetReadyTasks.Min || primaryMinimumTotal > c.TargetReadyTasks.Max {
		problems = append(problems, fmt.Errorf("curriculum: active primary minimum total %d is outside target %d..%d", primaryMinimumTotal, c.TargetReadyTasks.Min, c.TargetReadyTasks.Max))
	}
	for _, skill := range skills {
		for _, prerequisite := range skill.Prerequisites {
			if prerequisite == skill.ID {
				problems = append(problems, fmt.Errorf("curriculum: skill %q cannot require itself", skill.ID))
			} else if value, ok := skills[prerequisite]; !ok {
				problems = append(problems, fmt.Errorf("curriculum: skill %q references unknown prerequisite %q", skill.ID, prerequisite))
			} else if !isAssignable(value) {
				problems = append(problems, fmt.Errorf("curriculum: skill %q references non-assignable prerequisite %q", skill.ID, prerequisite))
			}
		}
		for _, replacement := range skill.ReplacedBy {
			if value, ok := skills[replacement]; !ok || !isAssignable(value) {
				problems = append(problems, fmt.Errorf("curriculum: deprecated skill %q has unknown or non-assignable replacement %q", skill.ID, replacement))
			}
		}
	}
	problems = append(problems, detectCycles(skills)...)
	return skills, domainBySkill, problems
}

func validMinimumTier(value Minimum) bool {
	return value == (Minimum{Primary: 2, Total: 4}) || value == (Minimum{Primary: 3, Total: 5}) || value == (Minimum{Primary: 4, Total: 6})
}

func duplicateFieldProblems(skillID, label string, values []string) []error {
	var problems []error
	for _, value := range duplicates(values) {
		problems = append(problems, fmt.Errorf("curriculum: skill %q has duplicate %s %q", skillID, label, value))
	}
	return problems
}

func detectCycles(skills map[string]Skill) []error {
	const (
		unseen = iota
		visiting
		done
	)
	state := map[string]int{}
	var problems []error
	var visit func(string, []string)
	visit = func(id string, path []string) {
		if state[id] == done {
			return
		}
		if state[id] == visiting {
			problems = append(problems, fmt.Errorf("curriculum: prerequisite cycle: %s -> %s", strings.Join(path, " -> "), id))
			return
		}
		state[id] = visiting
		for _, prerequisite := range skills[id].Prerequisites {
			if _, ok := skills[prerequisite]; ok {
				visit(prerequisite, append(path, id))
			}
		}
		state[id] = done
	}
	for id := range skills {
		visit(id, nil)
	}
	return problems
}

func validateRegistry(r Registry) (map[string]Source, []error) {
	var problems []error
	if r.Version != 1 || r.AuditedAt == "" || r.Policy != "docs/LICENSE_POLICY.md" {
		problems = append(problems, errors.New("registry: version, audited_at, and policy must be set"))
	}
	sources := map[string]Source{}
	for _, source := range r.Sources {
		prefix := fmt.Sprintf("registry source %q", source.ID)
		if err := sourcepolicy.Validate(source); err != nil {
			problems = append(problems, err)
		}
		if source.ID == "" || source.Repository == "" || source.DefaultBranch == "" || source.Adapter == "" || source.Scope == "" {
			problems = append(problems, fmt.Errorf("%s: id, repository, default_branch, adapter, and scope are required", prefix))
		}
		if _, ok := sources[source.ID]; ok {
			problems = append(problems, fmt.Errorf("%s: duplicate id", prefix))
		}
		if u, err := url.Parse(source.Repository); err != nil || u.Scheme != "https" || u.Host == "" {
			problems = append(problems, fmt.Errorf("%s: repository must be an absolute https URL", prefix))
		}
		if !shaPattern.MatchString(source.Commit) {
			problems = append(problems, fmt.Errorf("%s: commit must be an exact lowercase 40-character SHA", prefix))
		}
		if source.License.Expression == "" || source.License.EvidencePath == "" {
			problems = append(problems, fmt.Errorf("%s: license expression and evidence_path are required", prefix))
		}
		if !usageValues[source.UsageMode] {
			problems = append(problems, fmt.Errorf("%s: invalid usage_mode %q", prefix, source.UsageMode))
		}
		if len(source.Evidence) == 0 || len(source.Risks) == 0 {
			problems = append(problems, fmt.Errorf("%s: audit evidence and risks are required", prefix))
		}
		sources[source.ID] = source
	}
	if len(r.Sources) == 0 {
		problems = append(problems, errors.New("registry: at least one audited source is required"))
	}
	return sources, problems
}

func validateTask(root, path string, task Task, skills map[string]Skill, sources map[string]Source, migration migrationIndex) []error {
	var problems []error
	prefix := path
	if task.SchemaVersion != 1 {
		problems = append(problems, fmt.Errorf("%s: schema_version must be 1", prefix))
	}
	if !taskIDPattern.MatchString(task.ID) {
		problems = append(problems, fmt.Errorf("%s: invalid task id %q", prefix, task.ID))
	}
	if filepath.Base(filepath.Dir(path)) != task.ID {
		problems = append(problems, fmt.Errorf("%s: task id must match parent directory", prefix))
	}
	if len(task.Title) < 8 || len(task.Title) > 120 {
		problems = append(problems, fmt.Errorf("%s: title length must be 8..120", prefix))
	}
	if !exerciseTypes[task.TaskType] || !difficultyValues[task.Difficulty] || !qualityValues[task.QualityStatus] || !transformValues[task.TransformationType] {
		problems = append(problems, fmt.Errorf("%s: task_type, difficulty, transformation_type, or quality_status has an invalid value", prefix))
	}
	if task.EstimatedTimeMinutes.Min < 5 || task.EstimatedTimeMinutes.Max > 180 || task.EstimatedTimeMinutes.Min > task.EstimatedTimeMinutes.Max {
		problems = append(problems, fmt.Errorf("%s: invalid estimated_time_minutes range", prefix))
	}
	if len(task.Skills.Secondary) > 3 || len(task.NewConcepts) > 3 || len(task.ValidationCommands) == 0 {
		problems = append(problems, fmt.Errorf("%s: too many secondary/new skills or missing validation commands", prefix))
	}

	allLists := map[string][]string{
		"secondary":     task.Skills.Secondary,
		"prerequisites": task.Prerequisites,
		"new_concepts":  task.NewConcepts,
	}
	manifestBytes, manifestErr := os.ReadFile(path)
	manifestHash := ""
	if manifestErr != nil {
		problems = append(problems, fmt.Errorf("%s: read manifest for migration validation: %w", prefix, manifestErr))
	} else {
		manifestHash = digestBytes(manifestBytes)
	}
	problems = append(problems, validateTaskSkill(prefix, task.ID, "skills.primary", task.Skills.Primary, manifestHash, skills, migration)...)
	for name, ids := range allLists {
		for _, duplicate := range duplicates(ids) {
			problems = append(problems, fmt.Errorf("%s: duplicate %s skill %q", prefix, name, duplicate))
		}
		field := name
		if name == "secondary" {
			field = "skills.secondary"
		}
		for _, id := range ids {
			problems = append(problems, validateTaskSkill(prefix, task.ID, field, id, manifestHash, skills, migration)...)
		}
	}
	primaryAndSecondary := append([]string{task.Skills.Primary}, task.Skills.Secondary...)
	problems = append(problems, disjoint(prefix, "primary/secondary", []string{task.Skills.Primary}, task.Skills.Secondary)...)
	problems = append(problems, disjoint(prefix, "assessed/prerequisite", primaryAndSecondary, task.Prerequisites)...)
	problems = append(problems, disjoint(prefix, "new/prerequisite", task.NewConcepts, task.Prerequisites)...)

	if !shaPattern.MatchString(task.Source.Commit) || task.Source.Path == "" {
		problems = append(problems, fmt.Errorf("%s: source path and exact commit are required", prefix))
	}
	registered, ok := sources[task.Source.ID]
	if !ok {
		problems = append(problems, fmt.Errorf("%s: unknown source id %q", prefix, task.Source.ID))
	} else {
		if task.Source.Repository != registered.Repository || task.Source.License != registered.License.Expression {
			problems = append(problems, fmt.Errorf("%s: source repository/license must match registry entry %q", prefix, task.Source.ID))
		}
		if task.TransformationType != "original" && task.Source.Commit != registered.Commit {
			problems = append(problems, fmt.Errorf("%s: adapted source commit must match audited registry entry %q", prefix, task.Source.ID))
		}
		if task.TransformationType == "original" && task.Source.ID != "go-practice-bank" {
			problems = append(problems, fmt.Errorf("%s: original tasks must use the internal go-practice-bank source", prefix))
		}
		if registered.UsageMode == "blocked" {
			problems = append(problems, fmt.Errorf("%s: blocked source %q cannot back a task", prefix, task.Source.ID))
		}
		if (registered.UsageMode == "reference-only" || registered.UsageMode == "concepts-only") && task.TransformationType != "metadata-only" {
			problems = append(problems, fmt.Errorf("%s: source mode %q only permits metadata-only candidate handling", prefix, registered.UsageMode))
		}
	}

	if task.QualityStatus == "ready" {
		dir := filepath.Dir(path)
		_, lifecycleErr := os.Stat(filepath.Join(dir, "lifecycle.yaml"))
		if errors.Is(lifecycleErr, os.ErrNotExist) {
			for _, required := range []string{"README.md", "provenance.yaml", "starter", "tests/public", "tests/grader", "solution/reference"} {
				if _, err := os.Stat(filepath.Join(dir, filepath.FromSlash(required))); err != nil {
					problems = append(problems, fmt.Errorf("%s: ready task is missing %s", prefix, required))
				}
			}
		} else if lifecycleErr != nil {
			problems = append(problems, fmt.Errorf("%s: inspect lifecycle.yaml: %w", prefix, lifecycleErr))
		}
	}
	_ = root
	return problems
}

func validateTaskSkill(prefix, taskID, field, rawID, manifestHash string, skills map[string]Skill, migration migrationIndex) []error {
	value, ok := skills[rawID]
	if ok && isAssignable(value) {
		return nil
	}
	if resolved, allowed := migrationResolution(migration, taskID, field, rawID, manifestHash); allowed {
		if target, exists := skills[resolved]; exists && isAssignable(target) {
			return nil
		}
	}
	if !ok {
		return []error{fmt.Errorf("%s: unknown %s skill %q", prefix, field, rawID)}
	}
	return []error{fmt.Errorf("%s: %s skill %q is deprecated or non-assignable without an exact legacy resolution", prefix, field, rawID)}
}

func calculateCoverage(c Curriculum, tasks []Task, domainBySkill map[string]string, migration migrationIndex) []Coverage {
	byID := map[string]*Coverage{}
	for _, domain := range c.Domains {
		for _, skill := range domain.Skills {
			if !isAssignable(skill) {
				continue
			}
			byID[skill.ID] = &Coverage{SkillID: skill.ID, DomainID: domainBySkill[skill.ID], RequiredPrimary: skill.MinimumOccurrences.Primary, RequiredTotal: skill.MinimumOccurrences.Total}
		}
	}
	for _, task := range tasks {
		if task.QualityStatus != c.CoveragePolicy.CountedStatus {
			continue
		}
		primary := coverageSkillID(migration, task.ID, "skills.primary", task.Skills.Primary)
		if item := byID[primary]; item != nil {
			item.ActualPrimary++
			item.ActualTotal++
		}
		for _, id := range task.Skills.Secondary {
			resolved := coverageSkillID(migration, task.ID, "skills.secondary", id)
			if item := byID[resolved]; item != nil {
				item.ActualTotal++
			}
		}
	}
	coverage := make([]Coverage, 0, len(byID))
	for _, item := range byID {
		coverage = append(coverage, *item)
	}
	sort.Slice(coverage, func(i, j int) bool { return coverage[i].SkillID < coverage[j].SkillID })
	return coverage
}

func coverageSkillID(index migrationIndex, taskID, field, rawID string) string {
	ref, ok := index.legacyByKey[strings.Join([]string{taskID, field, rawID}, "\x00")]
	if !ok {
		return rawID
	}
	if ref.CoverageBehavior == "count-resolved" {
		return ref.ResolvedSkillID
	}
	return ""
}

func validateCoverage(result Result) []error {
	var problems []error
	for _, item := range result.Coverage {
		if item.ActualPrimary < item.RequiredPrimary || item.ActualTotal < item.RequiredTotal {
			problems = append(problems, fmt.Errorf("coverage: %s has primary %d/%d and total %d/%d", item.SkillID, item.ActualPrimary, item.RequiredPrimary, item.ActualTotal, item.RequiredTotal))
		}
	}
	target := result.Curriculum.TargetReadyTasks
	if result.ReadyTasks < target.Min || result.ReadyTasks > target.Max {
		problems = append(problems, fmt.Errorf("coverage: ready task count %d is outside %d..%d", result.ReadyTasks, target.Min, target.Max))
	}
	return problems
}

func duplicates(values []string) []string {
	seen := map[string]bool{}
	duplicates := map[string]bool{}
	for _, value := range values {
		if seen[value] {
			duplicates[value] = true
		}
		seen[value] = true
	}
	var result []string
	for value := range duplicates {
		result = append(result, value)
	}
	sort.Strings(result)
	return result
}

func disjoint(prefix, label string, left, right []string) []error {
	seen := map[string]bool{}
	for _, value := range left {
		seen[value] = true
	}
	var problems []error
	for _, value := range right {
		if seen[value] {
			problems = append(problems, fmt.Errorf("%s: %s sets overlap at %q", prefix, label, value))
		}
	}
	return problems
}

func set(values ...string) map[string]bool {
	result := make(map[string]bool, len(values))
	for _, value := range values {
		result[value] = true
	}
	return result
}

func skillSet(skills map[string]Skill) map[string]curriculumdata.SkillContract {
	result := make(map[string]curriculumdata.SkillContract, len(skills))
	for id, skill := range skills {
		result[id] = curriculumdata.SkillContract{Status: skill.Status, Assignable: isAssignable(skill)}
	}
	return result
}
