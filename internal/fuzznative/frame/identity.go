package frame

import (
	"fmt"
	"regexp"
	"strings"
)

// Identity is a closed task binding for the finite native seed profile.
// Package and observer paths derive from one validated generated module.
type Identity struct {
	ModulePath string `json:"module_path"`
	Grammar    string `json:"grammar"`
	Target     string `json:"target"`
}

var taskIDPattern = regexp.MustCompile(`^[a-z0-9]+(?:-[a-z0-9]+)*-[0-9]{3}$`)

func (i Identity) Validate() error {
	id, ok := strings.CutPrefix(i.ModulePath, "task.local/")
	if !ok || len(id) < 8 || len(id) > 80 || !taskIDPattern.MatchString(id) || i.Grammar != "hex-codec-v1" || i.Target != "FuzzHex" {
		return fmt.Errorf("invalid finite task identity")
	}
	return nil
}

func (i Identity) SubjectPath() string  { return i.ModulePath + "/subject" }
func (i Identity) ObserverPath() string { return i.ModulePath + "/internal/observe" }
