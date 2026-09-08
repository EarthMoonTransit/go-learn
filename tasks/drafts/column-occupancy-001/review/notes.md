# Author review notes

Independent content review pending. Estimates are uncalibrated.

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

Retained source behavior: chessboard selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| selected column and all positions | 01–02 | all-columns; skip-first; skip-last |
| false values contribute zero | 06 | length-only |
| missing/empty zero | 03–05 | missing-one; length-only |
| preserve input | 07 | destructive-count |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.
