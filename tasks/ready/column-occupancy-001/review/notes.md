# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return the number of true elements in the column stored under key.
- False elements contribute nothing; count both first and final elements.
- Missing keys, nil maps and empty columns return zero.
- Keys are exact strings and columns may have any length; preserve the map and its column elements.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.CountColumn(map[string][]bool{"A":{true,false,true,false,false,false,false,true}},"A"); got != 3 { t.Errorf("CountColumn: got %v, want %v", got, 3) }`
- Case02: `if got := subject.CountColumn(map[string][]bool{"A":{true},"B":{true,true}},"A"); got != 1 { t.Errorf("CountColumn: got %v, want %v", got, 1) }`
- Case03: `if got := subject.CountColumn(map[string][]bool{"A":{true}},"Z"); got != 0 { t.Errorf("CountColumn: got %v, want %v", got, 0) }`
- Case04: `if got := subject.CountColumn(nil,"A"); got != 0 { t.Errorf("CountColumn: got %v, want %v", got, 0) }`
- Case05: `if got := subject.CountColumn(map[string][]bool{"A":{}},"A"); got != 0 { t.Errorf("CountColumn: got %v, want %v", got, 0) }`
- Case06: `if got := subject.CountColumn(map[string][]bool{"A":{false,false}},"A"); got != 0 { t.Errorf("CountColumn: got %v, want %v", got, 0) }`
- Case07: `g:=map[string][]bool{"A":{true,false,true}};subject.CountColumn(g,"A");if len(g)!=1||!slices.Equal(g["A"],[]bool{true,false,true}){t.Errorf("grid changed: got %v want A:[true false true]",g)}`

Realistic required defect models:

- length-only
- all-columns
- skip-first
- skip-last
- missing-one
- destructive-count
- case-normalized

Retained source behavior: chessboard selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| selected column and all positions | 01–02 | all-columns; skip-first; skip-last |
| false values contribute zero | 06 | length-only |
| missing/empty zero | 03–05 | missing-one; length-only |
| preserve input | 07 | destructive-count |
| exact string keys | TestPublicExactKey | case-normalized |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.

Public regression TestPublicExactKey rejects case normalization; case-normalized is a required defect with quoted input and actual/expected count.


## Independent review

The exact-key feedback gap is closed by TestPublicExactKey and the case-normalized required defect. Public cases also cover selected-column range counting, endpoints, missing/empty input and preservation.

Reviewed the prompt, API, actual assertions, all 7 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 7 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
