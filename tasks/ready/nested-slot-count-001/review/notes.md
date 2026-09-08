# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Count every element stored in every column of the map.
- Both true and false elements contribute one.
- Columns may have different lengths; missing, nil or empty columns contribute zero.
- Nil or empty maps return zero; preserve all input entries and elements.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.CountSlots(map[string][]bool{"A":{true,false,true},"B":{false,false}}); got != 5 { t.Errorf("CountSlots: got %v, want %v", got, 5) }`
- Case02: `if got := subject.CountSlots(map[string][]bool{"A":{false,false,false}}); got != 3 { t.Errorf("CountSlots: got %v, want %v", got, 3) }`
- Case03: `if got := subject.CountSlots(nil); got != 0 { t.Errorf("CountSlots: got %v, want %v", got, 0) }`
- Case04: `if got := subject.CountSlots(map[string][]bool{"A":nil,"B":{}}); got != 0 { t.Errorf("CountSlots: got %v, want %v", got, 0) }`
- Case05: `g:=map[string][]bool{"A":{true,false}};subject.CountSlots(g);if len(g)!=1||!slices.Equal(g["A"],[]bool{true,false}){t.Errorf("grid changed: %v",g)}`
- Case06: `if got := subject.CountSlots(map[string][]bool{"A":{true},"B":{true,true},"C":{true,true,true}}); got != 6 { t.Errorf("CountSlots: got %v, want %v", got, 6) }`

Realistic required defect models:

- columns-only
- occupied-only
- fixed-eight
- overwrite-total
- clear-map
- empty-one

Retained source behavior: chessboard selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| all lengths and columns | 01,06 | columns-only; fixed-eight; overwrite-total |
| false included | 02 | occupied-only |
| empty zero | 03–04 | empty-one |
| preserve input | 05 | clear-map |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.


## Independent review

The cases distinguish outer-map size from nested element totals, variable column lengths, inclusion of false elements, accumulation across columns and preserved input.

Reviewed the prompt, API, actual assertions, all 6 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 6 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
