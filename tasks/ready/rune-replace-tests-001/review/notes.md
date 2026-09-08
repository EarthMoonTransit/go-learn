# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Write tests for Replace; inputs are valid UTF-8 and both rune arguments are valid scalars.
- Replace every occurrence of oldRune with newRune, preserving all other code points and their order.
- Replacement may increase or decrease UTF-8 byte length.
- Empty input, absent oldRune and equal rune arguments leave the text unchanged.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.Replace("❗ recommended product", '❗', '?'); got != "? recommended product" { 		t.Errorf("Replace: got %q, want %q", got, "? recommended product") 	}`
- Case02: `if got := subject.Replace("❗x❗", '❗', '?'); got != "?x?" { 		t.Errorf("Replace: got %q, want %q", got, "?x?") 	}`
- Case03: `if got := subject.Replace("aéa", 'a', '🔍'); got != "🔍é🔍" { 		t.Errorf("Replace: got %q, want %q", got, "🔍é🔍") 	}`
- Case04: `if got := subject.Replace("", 'a', 'b'); got != "" { 		t.Errorf("Replace: got %q, want %q", got, "") 	}`
- Case05: `if got := subject.Replace("中", 'a', 'b'); got != "中" { 		t.Errorf("Replace: got %q, want %q", got, "中") 	}`
- Case06: `if got := subject.Replace("éé", 'é', 'é'); got != "éé" { 		t.Errorf("Replace: got %q, want %q", got, "éé") 	}`

Realistic required defect models:

- first-only
- replace-all-runes
- byte-narrowing
- drop-absent
- equal-drops-text
- replacement-narrowed
- empty-as-one-rune
- fixed-byte-buffer

Learner-authored tests must distinguish every private defect; only the starting example is exposed.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| replace all and preserve other runes | 01–02 | first-only; replace-all-runes |
| growth and shrink in encoded width | 01,03 | byte-narrowing; replacement-narrowed; fixed-byte-buffer |
| empty/absent/equal unchanged | 04–06 | drop-absent; equal-drops-text |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Category-specific mutation regressions now distinguish missing boundary cases. `review/weak-submission.txt` records a deliberately incomplete suite that must leave required mutants alive in disposable task-check evidence.

Disposable weak-suite regression passed: the saved narrow suite passes reference, leaves at least one required mutant alive (`starter-full` insufficiency gate PASS), and the full reference test submission detects every required defect. Evidence: `.tmp/checks/rune-replace-tests-weak.yaml`. All runs used task-check on an external disposable copy.
