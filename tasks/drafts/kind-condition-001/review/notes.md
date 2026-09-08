# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return true for the exact kind car.
- Return true for the exact kind truck.
- Return false for every other string, including empty input.
- Matching is case-sensitive and does not trim whitespace.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.RequiresPermit("car"); got != true { t.Errorf("RequiresPermit: got %v, want %v", got, true) }`
- Case02: `if got := subject.RequiresPermit("truck"); got != true { t.Errorf("RequiresPermit: got %v, want %v", got, true) }`
- Case03: `if got := subject.RequiresPermit("bike"); got != false { t.Errorf("RequiresPermit: got %v, want %v", got, false) }`
- Case04: `if got := subject.RequiresPermit(""); got != false { t.Errorf("RequiresPermit: got %v, want %v", got, false) }`
- Case05: `if got := subject.RequiresPermit("CAR"); got != false { t.Errorf("RequiresPermit: got %v, want %v", got, false) }`
- Case06: `if got := subject.RequiresPermit(" car"); got != false { t.Errorf("RequiresPermit: got %v, want %v", got, false) }`
- Case07: `if got := subject.RequiresPermit("truck "); got != false { t.Errorf("RequiresPermit: got %v, want %v", got, false) }`

Realistic required defect models:

- car-only
- truck-only
- default-true
- case-normalized
- space-normalized

Retained source behavior: vehicle-purchase selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| both accepted kinds | 01–02 | car-only; truck-only |
| other/empty false | 03–04 | default-true |
| exact case/whitespace | 05–07 | case-normalized; space-normalized |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.
