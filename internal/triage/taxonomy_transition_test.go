package triage

import (
	"os"
	"path/filepath"
	"reflect"
	"strings"
	"testing"
)

func TestLiveTaxonomyLoaderRejectsUnsupportedVersions(t *testing.T) {
	root := t.TempDir()
	if err := os.MkdirAll(filepath.Join(root, "curriculum"), 0700); err != nil {
		t.Fatal(err)
	}
	for _, version := range []string{"3", "4", "5"} {
		t.Run(version, func(t *testing.T) {
			if err := os.WriteFile(filepath.Join(root, "curriculum/skills.yaml"), []byte("version: "+version+"\n"), 0600); err != nil {
				t.Fatal(err)
			}
			_, _, err := LoadSkillContracts(root)
			switch version {
			case "3":
				if err != nil {
					t.Fatalf("v3 loader behavior changed: %v", err)
				}
			case "4":
				if err == nil || !strings.Contains(err.Error(), "v3-to-v4.yaml") {
					t.Fatalf("v4 used without finite transition: %v", err)
				}
			default:
				if err == nil || !strings.Contains(err.Error(), "unsupported taxonomy version 5") {
					t.Fatalf("unknown live taxonomy accepted: %v", err)
				}
			}
		})
	}
}

func TestReviewedSourceContractsStayV3WhileTaskContractsAreV4(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	liveVersion, live, err := LoadSkillContracts(root)
	if err != nil {
		t.Fatal(err)
	}
	oldVersion, old, err := LoadReviewedSkillContracts(root)
	if err != nil {
		t.Fatal(err)
	}
	if liveVersion != 4 || oldVersion != 3 || len(live) != 180 || len(old) != 177 {
		t.Fatalf("live=%d/%d historical=%d/%d", liveVersion, len(live), oldVersion, len(old))
	}
	for id, contract := range old {
		if !reflect.DeepEqual(live[id], contract) {
			t.Fatalf("old reviewed contract reinterpreted: %s", id)
		}
	}
	for _, id := range []string{"expressions.bitwise-masks", "expressions.integer-shifts", "functions.recursion"} {
		if _, ok := old[id]; ok {
			t.Fatalf("new skill %s retroactively entered source review context", id)
		}
	}
}

func TestLiveContractOutputsCannotPoisonWarmTaxonomyValidation(t *testing.T) {
	root, err := filepath.Abs(filepath.Join("..", ".."))
	if err != nil {
		t.Fatal(err)
	}
	_, first, err := LoadSkillContracts(root)
	if err != nil {
		t.Fatal(err)
	}
	_, want, err := LoadSkillContracts(root)
	if err != nil {
		t.Fatal(err)
	}
	contract := first["functions.recursion"]
	if len(contract.Prerequisites) == 0 {
		t.Fatal("fixture lost recursion prerequisites")
	}
	contract.Prerequisites[0] = "poisoned"
	first["functions.recursion"] = contract
	delete(first, "expressions.bitwise-masks")
	_, got, err := LoadSkillContracts(root)
	if err != nil || !reflect.DeepEqual(got, want) {
		t.Fatalf("mutable public contracts poisoned subsequent load: %v", err)
	}
}
