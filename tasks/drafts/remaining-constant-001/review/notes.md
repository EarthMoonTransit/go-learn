# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Declare targetMinutes as an integer constant with value 40.
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
