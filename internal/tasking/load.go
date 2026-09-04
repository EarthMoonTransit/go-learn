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
	path := filepath.Join(root, "schemas", schemaName)
	content, err := os.ReadFile(path)
	if err != nil {
		return fmt.Errorf("read %s: %w", schemaName, err)
	}
	compiler := jsonschema.NewCompiler()
	compiler.AssertFormat = true
	if err := compiler.AddResource(schemaName, bytes.NewReader(content)); err != nil {
		return fmt.Errorf("load %s: %w", schemaName, err)
	}
	schema, err := compiler.Compile(schemaName)
	if err != nil {
		return fmt.Errorf("compile %s: %w", schemaName, err)
	}
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
	Located    LocatedTask
	Task       Task
	Lifecycle  Lifecycle
	Review     Review
	Provenance Provenance
	Prompt     []byte
}

func loadPackage(root string, located LocatedTask) (packageData, []string) {
	result := packageData{Located: located}
	var problems []string
	if err := inspectTree(located.Path); err != nil {
		return result, []string{err.Error()}
	}
	files := []struct {
		name   string
		value  any
		schema string
	}{
		{"task.yaml", &result.Task, "task.schema.json"},
		{"lifecycle.yaml", &result.Lifecycle, "task-lifecycle.schema.json"},
		{"review/review.yaml", &result.Review, "task-review.schema.json"},
		{"provenance.yaml", &result.Provenance, "task-provenance.schema.json"},
	}
	for _, file := range files {
		path := filepath.Join(located.Path, filepath.FromSlash(file.name))
		if err := decodeStrictFile(path, file.value); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", file.name, err))
			continue
		}
		if err := validateSchema(root, file.schema, file.value); err != nil {
			problems = append(problems, fmt.Sprintf("%s: %v", file.name, err))
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
