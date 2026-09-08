# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Use the supplied BandRate result as the percentage for value.
- Return value multiplied by that float32 rate widened to float64, then divided by 100.
- Preserve the sign and fractional input; zero returns zero.
- Inputs are finite and between -1000000000 and 1000000000; allow absolute error 0.00000001 or relative error 0.000000000001.

Public/reference assertions (actual/expected diagnostics):

- Case01: `for _,value:=range []float64{-10000,0,555.55,4999.99,34600.80}{want:=value*float64(subject.BandRate(value))/100;got:=subject.BandAmount(value);if math.IsNaN(got)||math.Abs(got-want)>math.Max(1e-8,math.Abs(want)*1e-12){t.Errorf("BandAmount(%v) got %.15g want %.15g",value,got,want)}}`

Realistic required defect models:

- narrow-input
- rate-truncated
- missing-percent
- absolute-output
- literal-wide-rate
- input-truncated

Retained source behavior: interest-is-interesting selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Defect |
|---|---|---|
| widen actual supplied rate | 01 table | rate-truncated; literal-wide-rate |
| preserve fractional input | 01 table | narrow-input; input-truncated |
| sign/percent | 01 table | absolute-output; missing-percent |

For test-writing tasks cases name the reference learner suite; mutants remain private. Public inputs cover basic observed requirements; additional grader cases remain within the stated domain.


## Independent review

The existing workspace-tests profile preserves the supplied BandRate implementation. Public checks distinguish widening its actual rounded float32 value from decimal reconstruction, retain fractional float64 input, and verify sign and percentage scaling with stated tolerance and NaN rejection.

Reviewed the prompt, API, actual checks, all 6 required defects, meaningful primary skill and route dependencies. Exact pinned hashes for the consulted source instructions, starter definitions, hand-authored tests, metadata and LICENSE match the local Exercism checkout, as do authors/contributors and the MIT notice. These are substantive permitted adaptations with private reference/mutant/review material excluded from learner payloads.

The author-provided targeted task-check report passes and the strong reference detects all 6 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full repository gates and allocation reconciliation remain separate. Automated fixtures are not learner calibration.
