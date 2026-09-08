# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- For a present key return its stored integer and true.
- A present zero is different from a missing key.
- For a missing key, including reads from a nil map, return zero and false.
- Use exact string keys, including the empty string, and leave all entries unchanged.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if a,b := subject.Quantity(map[string]int{"item":12},"item"); a != 12 || b != true {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,12,true)}`
- Case02: `if a,b := subject.Quantity(map[string]int{"item":0},"item"); a != 0 || b != true {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,0,true)}`
- Case03: `if a,b := subject.Quantity(map[string]int{"item":12},"absent"); a != 0 || b != false {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,0,false)}`
- Case04: `if a,b := subject.Quantity(nil,"item"); a != 0 || b != false {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,0,false)}`
- Case05: `if a,b := subject.Quantity(map[string]int{"":-2},""); a != -2 || b != true {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,-2,true)}`
- Case06: `m:=map[string]int{"a":3,"b":4};subject.Quantity(m,"a");if len(m)!=2||m["a"]!=3||m["b"]!=4{t.Errorf("input changed: %v",m)}`
- Case07: `if a,b := subject.Quantity(map[string]int{"A":3},"a"); a != 0 || b != false {t.Errorf("Quantity: got (%v,%v), want (%v,%v)",a,b,0,false)}`

Realistic required defect models:

- zero-is-absent
- always-present
- value-discarded
- empty-rejected
- destructive-read
- case-folded

Retained source behavior: gross-store selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| stored quantity/presence | 01–02 | value-discarded; zero-is-absent |
| missing/nil absent | 03–04 | always-present |
| empty/exact keys | 05,07 | empty-rejected; case-folded |
| preserve input | 06 | destructive-read |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.


## Independent review

The contract explicitly returns both stored quantity and key presence, with public zero, missing, nil, empty-key, case-sensitive and preservation checks. Its route role is practice after the earlier fallback lookup lesson.

Reviewed the prompt, API, actual assertions, all 6 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 6 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
