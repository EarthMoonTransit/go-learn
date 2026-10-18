package tasking

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/santhosh-tekuri/jsonschema/v5"
	"gopkg.in/yaml.v3"
)

func decodeStrictFile(path string, destination any) error {
	content, err := os.ReadFile(path)
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(content))
	decoder.KnownFields(true)
	if err := decoder.Decode(destination); err != nil {
		return err
	}
	var trailing any
	if err := decoder.Decode(&trailing); err == nil {
		return errors.New("multiple YAML documents are not allowed")
	} else if !errors.Is(err, io.EOF) {
		return err
	}
	return nil
}

func validateSchema(root, schemaName string, value any) error {
	schema, err := compileSchema(root, schemaName)
	if err != nil {
		return err
	}
	return validateSchemaValue(schemaName, schema, value)
}

func compileSchema(root, schemaName string) (*jsonschema.Schema, error) {
	path := filepath.Join(root, "schemas", schemaName)
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, fmt.Errorf("read %s: %w", schemaName, err)
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	if err := compiler.AddResource(schemaName, bytes.NewReader(content)); err != nil {
		return nil, fmt.Errorf("load %s: %w", schemaName, err)
	}
	schema, err := compiler.Compile(schemaName)
	if err != nil {
		return nil, fmt.Errorf("compile %s: %w", schemaName, err)
	}
	return schema, nil
}

func validateSchemaValue(schemaName string, schema *jsonschema.Schema, value any) error {
	normalized, err := json.Marshal(value)
	if err != nil {
		return err
	}
	var document any
	if err := json.Unmarshal(normalized, &document); err != nil {
		return err
	}
	if err := schema.Validate(document); err != nil {
		return fmt.Errorf("%s violation: %w", schemaName, err)
	}
	return nil
}

type packageData struct {
	Located           LocatedTask
	ContractVersion   int
	ProvenanceVersion int
	Task              Task
	TaskV2            TaskV2
	Lifecycle         Lifecycle
	Review            Review
	Provenance        Provenance
	ProvenanceV2      ProvenanceV2
	Prompt            []byte
}

func loadPackage(root string, located LocatedTask) (packageData, []string) {
	return loadPackageWithSchemaValidator(root, located, func(name string, value any) error {
		return validateSchema(root, name, value)
	})
}

func loadPackageWithSchemaValidator(root string, located LocatedTask, validate func(string, any) error) (packageData, []string) {
	result := packageData{Located: located}
	var problems []string
	decode := func(path, schema string, value any) error {
		if err := decodeStrictFile(path, value); err != nil {
			return err
		}
		return validate(schema, value)
	}
	if err := inspectTree(located.Path); err != nil {
		return result, []string{err.Error()}
	}
	taskPath := filepath.Join(located.Path, "task.yaml")
	taskVersion, err := documentSchemaVersion(taskPath)
	if err != nil {
		problems = append(problems, "task.yaml: "+err.Error())
	} else {
		result.ContractVersion = taskVersion
		switch taskVersion {
		case 1:
			if err := decode(taskPath, "task.schema.json", &result.Task); err != nil {
				problems = append(problems, "task.yaml: "+err.Error())
			}
		case 2:
			if err := decode(taskPath, "task-v2.schema.json", &result.TaskV2); err != nil {
				problems = append(problems, "task.yaml: "+err.Error())
			} else {
				result.Task = normalizeTaskV2(result.TaskV2)
			}
		default:
			problems = append(problems, fmt.Sprintf("task.yaml: unsupported schema_version %d", taskVersion))
		}
	}

	files := []struct {
		name   string
		value  any
		schema string
	}{
		{"lifecycle.yaml", &result.Lifecycle, "task-lifecycle.schema.json"},
		{"review/review.yaml", &result.Review, "task-review.schema.json"},
	}
	for _, file := range files {
		path := filepath.Join(located.Path, filepath.FromSlash(file.name))
		if err := decodeStrictFile(path, file.value); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", file.name, err))
			continue
		}
		if err := validate(file.schema, file.value); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", file.name, err))
		}
	}
	provenancePath := filepath.Join(located.Path, "provenance.yaml")
	provenanceVersion, err := documentSchemaVersion(provenancePath)
	if err != nil {
		problems = append(problems, "provenance.yaml: "+err.Error())
	} else {
		result.ProvenanceVersion = provenanceVersion
		switch provenanceVersion {
		case 1:
			if err := decode(provenancePath, "task-provenance.schema.json", &result.Provenance); err != nil {
				problems = append(problems, "provenance.yaml: "+err.Error())
			}
		case 2:
			if err := decode(provenancePath, "task-provenance-v2.schema.json", &result.ProvenanceV2); err != nil {
				problems = append(problems, "provenance.yaml: "+err.Error())
			}
		default:
			problems = append(problems, fmt.Sprintf("provenance.yaml: unsupported schema_version %d", provenanceVersion))
		}
	}
	prompt, err := os.ReadFile(filepath.Join(located.Path, "TASK.md"))
	if err != nil {
		problems = append(problems, fmt.Sprintf("TASK.md: %v", err))
	} else {
		result.Prompt = prompt
	}
	return result, problems
}

func decodeAndValidate(root, path, schema string, value any) error {
	if err := decodeStrictFile(path, value); err != nil {
		return err
	}
	return validateSchema(root, schema, value)
}

func documentSchemaVersion(path string) (int, error) {
	var header struct {
		SchemaVersion int `yaml:"schema_version"`
	}
	content, err := os.ReadFile(path)
	if err != nil {
		return 0, err
	}
	if err := yaml.Unmarshal(content, &header); err != nil {
		return 0, err
	}
	if header.SchemaVersion == 0 {
		return 0, errors.New("schema_version is required")
	}
	return header.SchemaVersion, nil
}

func normalizeTaskV2(task TaskV2) Task {
	taskType := map[string]string{
		"implementation": "implement", "debugging": "debug", "refactoring": "refactor",
		"api-design": "implement", "test-writing": "test-writing", "test-repair": "test-writing",
		"fuzz-test-writing": "test-writing", "integration": "implement", "performance": "performance",
		"code-review-patch": "debug",
	}[task.Activity]
	return Task{
		SchemaVersion: task.SchemaVersion, ID: task.ID, Title: task.Title, TaskType: taskType,
		Skills: task.Skills, Prerequisites: task.Prerequisites, NewConcepts: task.NewConcepts,
		Difficulty: task.Difficulty, EstimatedTimeMinutes: task.EstimatedTimeMinutes,
		QualityStatus: task.QualityStatus,
	}
}

func loadTaskIndex(path string) (Task, error) {
	version, err := documentSchemaVersion(path)
	if err != nil {
		return Task{}, err
	}
	switch version {
	case 1:
		var task Task
		if err := decodeStrictFile(path, &task); err != nil {
			return Task{}, err
		}
		return task, nil
	case 2:
		var task TaskV2
		if err := decodeStrictFile(path, &task); err != nil {
			return Task{}, err
		}
		return normalizeTaskV2(task), nil
	default:
		return Task{}, fmt.Errorf("unsupported task schema_version %d", version)
	}
}
