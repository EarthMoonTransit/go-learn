# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Declare targetMinutes as a constant with value 40 representable as int.
- Return targetMinutes minus elapsed.
- elapsed is between 0 and 40 inclusive; zero elapsed returns the full target and 40 returns zero.
- The constant must remain usable in constant expressions.

Public/reference assertions (actual/expected diagnostics):

- Case01: `const copied=subject.targetMinutes;if copied!=40{t.Errorf("targetMinutes got %d want 40",copied)}`
- Case02: `if got := subject.RemainingMinutes(15); got != 25 { t.Errorf("RemainingMinutes: got %v, want %v", got, 25) }`
- Case03: `if got := subject.RemainingMinutes(30); got != 10 { t.Errorf("RemainingMinutes: got %v, want %v", got, 10) }`
- Case04: `if got := subject.RemainingMinutes(0); got != 40 { t.Errorf("RemainingMinutes: got %v, want %v", got, 40) }`
- Case05: `if got := subject.RemainingMinutes(40); got != 0 { t.Errorf("RemainingMinutes: got %v, want %v", got, 0) }`

Realistic required defect models:

- wrong-target
- elapsed-instead
- reversed-subtraction
- zero-completed
- end-off-by-one

Retained source behavior: lasagna selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


The constant is package-private, keeping its initializer editable within the fixed public function API. The fixed workspace skeleton requires a constant expression, and public external-package tests assert its numeric value through Target(); no checker contract change is required.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| constant-expression and value | 01 | wrong-target; variable declaration rejected by compile-time use |
| remaining subtraction | 02–03 | elapsed-instead; reversed-subtraction |
| zero/full endpoints | 04–05 | zero-completed; end-off-by-one |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.

Uses the existing workspace-tests profile with one editable file and a fixed accessor/constant-expression skeleton. No verifier extension or API fingerprint relaxation.

The contract accepts any constant expression with value 40 representable as int, including 40.0. It assesses const versus var and remaining arithmetic, not an untyped constant kind.


## Independent review

The existing workspace-tests profile confines edits to constants.go and retains the fixed accessor and constant-expression skeleton. The contract intentionally allows a constant value 40 representable as int, including 40.0; public checks cover its value and remainder endpoints.

Reviewed the prompt, API, actual assertions, all 5 required defect implementations and route prerequisites. The source operation is substantively retained with permitted adaptation: the exact pinned artifact hashes, author/contributor names and MIT license text match the local Exercism checkout. No generated or reference-only artifact is claimed as adapted lineage. Private reference/mutant/review files remain outside learner payloads.

The author-provided targeted task-check report passes with all 5 required defects detected. The reviewer read source files and reports without executing authored code or accessing the production learner home. This is content approval for the lifecycle process; full batch acceptance and source-allocation reconciliation remain separate. Automated fixtures are not learner calibration.
