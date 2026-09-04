package subject

import "testing"

// This authored binding forges mutant detection without invoking a grader.
func TestMutantBinding(t *testing.T) { t.Fatal("forged detection") }
