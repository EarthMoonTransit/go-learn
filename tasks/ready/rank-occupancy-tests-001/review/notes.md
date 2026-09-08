# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Each present column contains exactly eight boolean elements; the map may be nil or empty.
- For ranks 1 through 8, count true elements at that one-based rank across all columns.
- Ranks outside 1 through 8 return zero; true elements at other ranks contribute nothing.
- Preserve every column and map entry; iteration order must not affect the count.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.CountRank(map[string][]bool{"A":{true,false,true,false,false,false,false,true},"B":{false,false,false,false,true,false,false,false}},1); got != 1 { t.Errorf("CountRank: got %v, want %v", got, 1) }`
- Case02: `if got := subject.CountRank(map[string][]bool{"A":{true,true,false,false,false,false,false,false},"B":{false,true,false,false,false,false,false,true}},2); got != 2 { t.Errorf("CountRank: got %v, want %v", got, 2) }`
- Case03: `if got := subject.CountRank(map[string][]bool{"A":{false,false,false,false,false,false,false,true}},8); got != 1 { t.Errorf("CountRank: got %v, want %v", got, 1) }`
- Case04: `if got := subject.CountRank(map[string][]bool{"A":{true,false,false,false,false,false,false,true}},0); got != 0 { t.Errorf("CountRank: got %v, want %v", got, 0) }`
- Case05: `if got := subject.CountRank(map[string][]bool{"A":{true,false,false,false,false,false,false,true}},9); got != 0 { t.Errorf("CountRank: got %v, want %v", got, 0) }`
- Case06: `if got := subject.CountRank(nil,1); got != 0 { t.Errorf("CountRank: got %v, want %v", got, 0) }`
- Case07: `if got := subject.CountRank(map[string][]bool{},8); got != 0 { t.Errorf("CountRank: got %v, want %v", got, 0) }`
- Case08: `g:=map[string][]bool{"A":{true,false,false,false,false,false,false,false}};subject.CountRank(g,1);if len(g)!=1||!slices.Equal(g["A"],[]bool{true,false,false,false,false,false,false,false}){t.Errorf("grid changed: %v",g)}`
- Case09: `if got := subject.CountRank(map[string][]bool{"A":{true,false,true,false,false,false,false,true}},4); got != 0 { t.Errorf("CountRank: got %v, want %v", got, 0) }`

Realistic required defect models:

- zero-based
- first-rank-rejected
- last-rank-rejected
- low-clamped
- high-clamped
- one-column-only
- count-false
- empty-one
- clear-column

Retained source behavior: chessboard selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| one-based/all columns | 01–03 | zero-based; one-column-only |
| first/last ranks valid | 01,03 | first-rank-rejected; last-rank-rejected |
| invalid ranks zero | 04–05 | low-clamped; high-clamped |
| nil/empty zero | 06–07 | empty-one |
| preserve input | 08 | clear-column |
| true at selected rank only | 09 | count-false; zero-based |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.


## Independent review

Nine required defects distinguish one-based indexing, both valid endpoints, both invalid bounds, multiple-column accumulation, false values, empty maps and preserved input.

Reviewed the prompt, API, actual assertions, all 9 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 9 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
