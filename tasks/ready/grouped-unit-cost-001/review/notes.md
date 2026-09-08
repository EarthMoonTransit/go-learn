# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Each complete group of ten units costs 95000.
- Each remaining individual unit costs 10000.
- Use as many complete groups as possible and return their combined cost.
- count is between 0 and 100000; zero units cost zero and the result fits uint.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.BatchCost(0); got != uint(0) { t.Errorf("BatchCost: got %v, want %v", got, uint(0)) }`
- Case02: `if got := subject.BatchCost(1); got != uint(10000) { t.Errorf("BatchCost: got %v, want %v", got, uint(10000)) }`
- Case03: `if got := subject.BatchCost(9); got != uint(90000) { t.Errorf("BatchCost: got %v, want %v", got, uint(90000)) }`
- Case04: `if got := subject.BatchCost(10); got != uint(95000) { t.Errorf("BatchCost: got %v, want %v", got, uint(95000)) }`
- Case05: `if got := subject.BatchCost(21); got != uint(200000) { t.Errorf("BatchCost: got %v, want %v", got, uint(200000)) }`
- Case06: `if got := subject.BatchCost(37); got != uint(355000) { t.Errorf("BatchCost: got %v, want %v", got, uint(355000)) }`
- Case07: `if got := subject.BatchCost(100); got != uint(950000) { t.Errorf("BatchCost: got %v, want %v", got, uint(950000)) }`

Realistic required defect models:

- individual-only
- round-groups-up
- remainder-free
- one-group-only
- zero-one-unit

Retained source behavior: cars-assemble selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| groups of ten | 04,07 | individual-only; one-group-only |
| remainder charge | 02–03,05–06 | round-groups-up; remainder-free |
| zero | 01 | zero-one-unit |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.


## Independent review

Public cases distinguish complete groups, individual remainders, multiple groups and zero cost. Arithmetic is the observed primary skill.

Reviewed the prompt, API, actual assertions, all 5 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 5 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
