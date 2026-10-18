package tasking

import (
	"bytes"
	"fmt"

	"gopkg.in/yaml.v3"
)

// Keep explicit null distinct from omission: fuzz is forbidden outside its
// profile, and null must not disappear through the optional pointer's JSON
// omitempty normalization. Decode the alias with KnownFields to preserve the
// existing closed field contract, including limits and nested descriptors.
func (plan *VerificationPlan) UnmarshalYAML(node *yaml.Node) error {
	type plain VerificationPlan
	data, err := yaml.Marshal(node)
	if err != nil {
		return err
	}
	decoder := yaml.NewDecoder(bytes.NewReader(data))
	decoder.KnownFields(true)
	var value plain
	if err := decoder.Decode(&value); err != nil {
		return err
	}
	var fields map[string]any
	if err := node.Decode(&fields); err != nil {
		return err
	}
	if _, present := fields["fuzz"]; present && value.Fuzz == nil {
		return fmt.Errorf("verification.fuzz must be an object when present")
	}
	*plan = VerificationPlan(value)
	return nil
}
