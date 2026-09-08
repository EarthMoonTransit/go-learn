# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Map two through nine to their corresponding integers, 2 through 9.
- Map ace to 11.
- Map ten, jack, queen and king to 10.
- All other strings return zero; matching is exact and case-sensitive.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.TokenValue("ace"); got != 11 { t.Errorf("TokenValue: got %v, want %v", got, 11) }`
- Case02: `if got := subject.TokenValue("two"); got != 2 { t.Errorf("TokenValue: got %v, want %v", got, 2) }`
- Case03: `if got := subject.TokenValue("three"); got != 3 { t.Errorf("TokenValue: got %v, want %v", got, 3) }`
- Case04: `if got := subject.TokenValue("four"); got != 4 { t.Errorf("TokenValue: got %v, want %v", got, 4) }`
- Case05: `if got := subject.TokenValue("five"); got != 5 { t.Errorf("TokenValue: got %v, want %v", got, 5) }`
- Case06: `if got := subject.TokenValue("six"); got != 6 { t.Errorf("TokenValue: got %v, want %v", got, 6) }`
- Case07: `if got := subject.TokenValue("seven"); got != 7 { t.Errorf("TokenValue: got %v, want %v", got, 7) }`
- Case08: `if got := subject.TokenValue("eight"); got != 8 { t.Errorf("TokenValue: got %v, want %v", got, 8) }`
- Case09: `if got := subject.TokenValue("nine"); got != 9 { t.Errorf("TokenValue: got %v, want %v", got, 9) }`
- Case10: `if got := subject.TokenValue("ten"); got != 10 { t.Errorf("TokenValue: got %v, want %v", got, 10) }`
- Case11: `if got := subject.TokenValue("jack"); got != 10 { t.Errorf("TokenValue: got %v, want %v", got, 10) }`
- Case12: `if got := subject.TokenValue("queen"); got != 10 { t.Errorf("TokenValue: got %v, want %v", got, 10) }`
- Case13: `if got := subject.TokenValue("king"); got != 10 { t.Errorf("TokenValue: got %v, want %v", got, 10) }`
- Case14: `if got := subject.TokenValue("joker"); got != 0 { t.Errorf("TokenValue: got %v, want %v", got, 0) }`
- Case15: `if got := subject.TokenValue("ACE"); got != 0 { t.Errorf("TokenValue: got %v, want %v", got, 0) }`
- Case16: `if got := subject.TokenValue(""); got != 0 { t.Errorf("TokenValue: got %v, want %v", got, 0) }`

Realistic required defect models:

- ace-one
- face-missing
- case-folded
- unknown-ten
- digit-shift

Retained source behavior: blackjack selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| ace value | 01 | ace-one |
| numeric word mapping | 02–09 | digit-shift |
| grouped ten-value cases | 10–13 | face-missing |
| default/exact case | 14–16 | unknown-ten; case-folded |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
