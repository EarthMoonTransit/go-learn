# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Write tests for whether the entire byte string is valid UTF-8.
- Accept empty input, ASCII, NUL and every correctly encoded scalar, including U+FFFD.
- Reject isolated continuation bytes, invalid leading bytes, overlong encodings and truncated sequences.
- Reject encodings of surrogate values or values above U+10FFFF, including malformed bytes after a valid prefix.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.ValidUTF8("ASCII"); got != true { 		t.Errorf("ValidUTF8: got %v, want %v", got, true) 	}`
- Case02: `if got := subject.ValidUTF8(""); got != true { 		t.Errorf("ValidUTF8: got %v, want %v", got, true) 	}`
- Case03: `if got := subject.ValidUTF8("\x00é中🔍�"); got != true { 		t.Errorf("ValidUTF8: got %v, want %v", got, true) 	}`
- Case04: `if got := subject.ValidUTF8("\x80"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case05: `if got := subject.ValidUTF8("\xff"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case06: `if got := subject.ValidUTF8("\xc0\xaf"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case07: `if got := subject.ValidUTF8("\xe2\x82"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case08: `if got := subject.ValidUTF8("\xed\xa0\x80"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case09: `if got := subject.ValidUTF8("\xf4\x90\x80\x80"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`
- Case10: `if got := subject.ValidUTF8("ok\xff"); got != false { 		t.Errorf("ValidUTF8: got %v, want %v", got, false) 	}`

Realistic required defect models:

- accept-any
- reject-replacement
- first-only
- ascii-only
- empty-invalid
- overlong-accepted
- surrogate-accepted
- above-max-accepted
- truncated-accepted
- continuation-skipped
- invalid-leading-skipped

Learner-authored tests must distinguish every private defect; only the starting example is exposed.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| accept complete valid encodings including empty/NUL/U+FFFD | 01–03 | ascii-only; empty-invalid; reject-replacement |
| reject malformed, overlong and truncated units | 04–07 | continuation-skipped; invalid-leading-skipped; overlong-accepted; truncated-accepted |
| reject surrogate and out-of-range encodings | 08–09 | surrogate-accepted; above-max-accepted |
| validate after the initial prefix | 10 | first-only |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Category-specific mutation regressions now distinguish missing boundary cases. `review/weak-submission.txt` records a deliberately incomplete suite that must leave required mutants alive in disposable task-check evidence.

Disposable weak-suite regression passed: the saved narrow suite passes reference, leaves at least one required mutant alive (`starter-full` insufficiency gate PASS), and the full reference test submission detects every required defect. Evidence: `.tmp/checks/utf8-validity-tests-weak.yaml`. All runs used task-check on an external disposable copy.
