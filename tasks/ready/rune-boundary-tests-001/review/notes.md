# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Input is valid UTF-8; pos is a zero-based rune position.
- Return the byte offset at that position and true.
- The position immediately after the final rune is valid and returns len(s).
- Negative positions and positions beyond the rune count return zero and false; position zero in empty input is valid.

Public/reference assertions (each reports actual/expected values):

- Case01: `if a, b := subject.RuneBoundary("abc", 1); a != 1 || b != true { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 1, true) 	}`
- Case02: `if a, b := subject.RuneBoundary("Aé🔍中", 3); a != 7 || b != true { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 7, true) 	}`
- Case03: `if a, b := subject.RuneBoundary("é", 1); a != 2 || b != true { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 2, true) 	}`
- Case04: `if a, b := subject.RuneBoundary("", 0); a != 0 || b != true { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, true) 	}`
- Case05: `if a, b := subject.RuneBoundary("é", 2); a != 0 || b != false { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, false) 	}`
- Case06: `if a, b := subject.RuneBoundary("x", -1); a != 0 || b != false { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 0, false) 	}`
- Case07: `if a, b := subject.RuneBoundary("e\u0301", 2); a != 3 || b != true { 		t.Errorf("RuneBoundary: got (%v,%v), want (%v,%v)", a, b, 3, true) 	}`

Realistic required defect models:

- rune-index
- reject-end
- negative-clamped
- byte-bounds
- upper-bound-clamped
- empty-rejected

Learner-authored tests must distinguish every private defect; only the starting example is exposed.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| rune position maps to byte offset | 01–02,07 | rune-index; byte-bounds |
| end position is valid | 03 | reject-end |
| empty and both invalid bounds | 04–06 | empty-rejected; negative-clamped; upper-bound-clamped |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Category-specific mutation regressions now distinguish missing boundary cases. `review/weak-submission.txt` records a deliberately incomplete suite that must leave required mutants alive in disposable task-check evidence.

Disposable weak-suite regression passed: the saved narrow suite passes reference, leaves at least one required mutant alive (`starter-full` insufficiency gate PASS), and the full reference test submission detects every required defect. Evidence: `.tmp/checks/rune-boundary-tests-weak.yaml`. All runs used task-check on an external disposable copy.
