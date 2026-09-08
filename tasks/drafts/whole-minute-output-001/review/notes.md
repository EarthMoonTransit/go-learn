# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return the whole successful units per minute from an hourly rate and percentage.
- Apply the percentage and divide by 60 before discarding the fractional result toward zero.
- Preserve fractional percentages during the calculation.
- rate is between 0 and 1000000; percent is finite and between 0 and 100 inclusive.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.SuccessfulPerMinute(0,100); got != 0 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 0) }`
- Case02: `if got := subject.SuccessfulPerMinute(221,100); got != 3 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 3) }`
- Case03: `if got := subject.SuccessfulPerMinute(426,80); got != 5 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 5) }`
- Case04: `if got := subject.SuccessfulPerMinute(6824,20.5); got != 23 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 23) }`
- Case05: `if got := subject.SuccessfulPerMinute(8000,0); got != 0 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 0) }`
- Case06: `if got := subject.SuccessfulPerMinute(119,100); got != 1 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 1) }`
- Case07: `if got := subject.SuccessfulPerMinute(359,50.5); got != 3 { t.Errorf("SuccessfulPerMinute: got %v, want %v", got, 3) }`

Realistic required defect models:

- hour-truncated-first
- round-up
- percent-truncated
- missing-minute
- zero-percent-full

Retained source behavior: cars-assemble selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| per-minute scaling and truncation | 01–07 | missing-minute; round-up; hour-truncated-first |
| fractional percent | 04,07 | percent-truncated |
| zero rate/percent | 01,05 | zero-percent-full |

Public diagnostics include actual/expected values. For test-writing these are reference learner-suite assertions, not exposed private mutants. Input constraints define the tested domain. Automated fixtures are not learner calibration.
