# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Return 3.213 for value below zero.
- Return 0.5 for values from zero inclusive to 1000 exclusive.
- Return 1.621 for values from 1000 inclusive to 5000 exclusive; return 2.475 at or above 5000.
- Classify the original float64 input without narrowing it first; return the corresponding float32 rate. Inputs are finite and between -1000000000 and 1000000000.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.BandRate(0); got != float32(0.5) { t.Errorf("BandRate: got %v, want %v", got, float32(0.5)) }`
- Case02: `if got := subject.BandRate(-0.000001); got != float32(3.213) { t.Errorf("BandRate: got %v, want %v", got, float32(3.213)) }`
- Case03: `if got := subject.BandRate(999.99999); got != float32(0.5) { t.Errorf("BandRate: got %v, want %v", got, float32(0.5)) }`
- Case04: `if got := subject.BandRate(1000); got != float32(1.621) { t.Errorf("BandRate: got %v, want %v", got, float32(1.621)) }`
- Case05: `if got := subject.BandRate(4999.99999); got != float32(1.621) { t.Errorf("BandRate: got %v, want %v", got, float32(1.621)) }`
- Case06: `if got := subject.BandRate(5000); got != float32(2.475) { t.Errorf("BandRate: got %v, want %v", got, float32(2.475)) }`

Realistic required defect models:

- narrow-input
- zero-negative
- middle-bound-inclusive
- upper-bound-inclusive
- truncate-rate
- absolute-input

Retained source behavior: interest-is-interesting selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| all thresholds and original precision | 01–06 | narrow-input; zero-negative; middle-bound-inclusive; upper-bound-inclusive |
| negative rate | 02 | absolute-input |
| fractional return value | 04–05 | truncate-rate |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

Public values immediately below thresholds expose narrowing before comparison; exact boundaries, negative/zero inputs and fractional float32 rates are explicit.

Reviewed the prompt, API, actual checks, all 6 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
