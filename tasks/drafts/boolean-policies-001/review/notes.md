# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- FastAllowed returns the negation of blocked.
- AnyActive returns true when at least one of its three inputs is true.
- SignalAllowed requires ready to be true and blocked to be false.
- ReleaseAllowed requires b to be false and either override to be true or both a false and ready true.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.FastAllowed(true); got != false { t.Errorf("FastAllowed: got %v, want %v", got, false) }`
- Case02: `if got := subject.FastAllowed(false); got != true { t.Errorf("FastAllowed: got %v, want %v", got, true) }`
- Case03: `if got := subject.AnyActive(false,false,false); got != false { t.Errorf("AnyActive: got %v, want %v", got, false) }`
- Case04: `if got := subject.AnyActive(true,false,false); got != true { t.Errorf("AnyActive: got %v, want %v", got, true) }`
- Case05: `if got := subject.AnyActive(false,true,false); got != true { t.Errorf("AnyActive: got %v, want %v", got, true) }`
- Case06: `if got := subject.AnyActive(false,false,true); got != true { t.Errorf("AnyActive: got %v, want %v", got, true) }`
- Case07: `if got := subject.SignalAllowed(false,true); got != true { t.Errorf("SignalAllowed: got %v, want %v", got, true) }`
- Case08: `if got := subject.SignalAllowed(true,true); got != false { t.Errorf("SignalAllowed: got %v, want %v", got, false) }`
- Case09: `if got := subject.SignalAllowed(false,false); got != false { t.Errorf("SignalAllowed: got %v, want %v", got, false) }`
- Case10: `if got := subject.ReleaseAllowed(true,false,false,true); got != true { t.Errorf("ReleaseAllowed: got %v, want %v", got, true) }`
- Case11: `if got := subject.ReleaseAllowed(false,true,true,true); got != false { t.Errorf("ReleaseAllowed: got %v, want %v", got, false) }`
- Case12: `if got := subject.ReleaseAllowed(false,false,true,false); got != true { t.Errorf("ReleaseAllowed: got %v, want %v", got, true) }`
- Case13: `if got := subject.ReleaseAllowed(true,false,true,false); got != false { t.Errorf("ReleaseAllowed: got %v, want %v", got, false) }`
- Case14: `if got := subject.ReleaseAllowed(false,false,false,false); got != false { t.Errorf("ReleaseAllowed: got %v, want %v", got, false) }`

Realistic required defect models:

- fast-inverted
- all-active-required
- third-input-ignored
- signal-or
- override-bypasses-block
- ready-always-required
- first-block-ignored

Retained source behavior: annalyns-infiltration selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| negation | 01–02 | fast-inverted |
| any of three inputs | 03–06 | all-active-required; third-input-ignored |
| signal conjunction | 07–09 | signal-or |
| both release paths and shared block | 10–14 | override-bypasses-block; ready-always-required; first-block-ignored |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
