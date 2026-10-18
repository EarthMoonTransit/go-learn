package tasking

import (
	"bytes"
	"fmt"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

const fuzzSourceContractPath = "verification/FUZZ_SOURCE.md"
const fuzzPublicContractPath = "FUZZ_SOURCE.md"
const fuzzCanonicalContractPath = "docs/FUZZ_SOURCE.md"

func checkFuzzLayout(root string, data packageData) ([]string, packageLayout) {
	problems, layout := checkMutationLayout(root, data)
	if data.TaskV2.Verification.Fuzz == nil {
		return append(problems, "fuzz-seed-tests requires verification.fuzz"), layout
	}
	contract, contractErr := readRegularFile(filepath.Join(data.Located.Path, filepath.FromSlash(fuzzSourceContractPath)), 8<<20)
	canonical, canonicalErr := readRegularFile(filepath.Join(root, filepath.FromSlash(fuzzCanonicalContractPath)), 8<<20)
	if contractErr != nil || canonicalErr != nil || !bytes.Equal(contract, canonical) {
		problems = append(problems, "verification/FUZZ_SOURCE.md must contain the current complete finite source contract")
	}
	if !bytes.Contains(data.Prompt, []byte("]("+fuzzPublicContractPath+")")) {
		problems = append(problems, "TASK.md must link to the learner-visible FUZZ_SOURCE.md contract")
	}
	variantsRoot := filepath.Join(data.Located.Path, "verification", "correct")
	info, err := os.Lstat(variantsRoot)
	if err != nil || !info.IsDir() || info.Mode()&os.ModeSymlink != 0 {
		return append(problems, "verification/correct must be a real directory"), layout
	}
	entries, err := os.ReadDir(variantsRoot)
	if err != nil {
		return append(problems, "cannot read verification/correct: "+err.Error()), layout
	}
	var actual []string
	for _, entry := range entries {
		if !entry.IsDir() || !slugPattern.MatchString(entry.Name()) {
			problems = append(problems, "invalid correct variant entry: "+entry.Name())
			continue
		}
		actual = append(actual, entry.Name())
		if err := inspectTree(filepath.Join(variantsRoot, entry.Name())); err != nil {
			problems = append(problems, "correct variant "+entry.Name()+": "+err.Error())
			continue
		}
		files, next := collectGoDirectory(data.Located.Path, "verification/correct/"+entry.Name(), directoryImplementation, problems)
		problems = next
		layout.VerificationCorrect = append(layout.VerificationCorrect, correctVariantLayout{ID: entry.Name(), Files: files})
	}
	expected := append([]string(nil), data.TaskV2.Verification.Fuzz.CorrectVariants...)
	sort.Strings(expected)
	sort.Strings(actual)
	if strings.Join(expected, "\n") != strings.Join(actual, "\n") {
		problems = append(problems, fmt.Sprintf("correct_variants %v must exactly match verification/correct %v", expected, actual))
	}
	return uniqueSorted(problems), layout
}

func checkFuzzPublicAPI(layout packageLayout) []string {
	canonical, err := apiFingerprint(layout.VerificationReference)
	if err != nil {
		return []string{"cannot extract verification reference API: " + err.Error()}
	}
	var problems []string
	for _, variant := range layout.VerificationCorrect {
		fingerprint, err := apiFingerprint(variant.Files)
		if err != nil || fingerprint != canonical {
			problems = append(problems, "correct variant "+variant.ID+" does not preserve the public API")
		}
	}
	for _, variant := range layout.Mutants {
		fingerprint, err := apiFingerprint(variant.Files)
		if err != nil || fingerprint != canonical {
			problems = append(problems, "verification mutant "+variant.ID+" does not preserve the public API")
		}
	}
	return uniqueSorted(problems)
}
