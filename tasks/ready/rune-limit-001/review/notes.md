# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Input is valid UTF-8; each Unicode code point counts once, including combining marks.
- Return true when the code-point count is at most limit.
- Empty input fits limit zero; no input fits a negative limit.
- The result depends on code points, not UTF-8 bytes or displayed grapheme clusters.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.WithinLimit("exercism❗", 9); got != true { 		t.Errorf("WithinLimit: got %v, want %v", got, true) 	}`
- Case02: `if got := subject.WithinLimit("exercism🔍", 8); got != false { 		t.Errorf("WithinLimit: got %v, want %v", got, false) 	}`
- Case03: `if got := subject.WithinLimit("", 0); got != true { 		t.Errorf("WithinLimit: got %v, want %v", got, true) 	}`
- Case04: `if got := subject.WithinLimit("", -1); got != false { 		t.Errorf("WithinLimit: got %v, want %v", got, false) 	}`
- Case05: `if got := subject.WithinLimit("e\u0301", 1); got != false { 		t.Errorf("WithinLimit: got %v, want %v", got, false) 	}`
- Case06: `if got := subject.WithinLimit("中🔍", 3); got != true { 		t.Errorf("WithinLimit: got %v, want %v", got, true) 	}`

Realistic required defect models:

- byte-count
- exclusive-limit
- negative-unlimited
- ignore-combining
- reversed-comparison

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| code-point count, inclusive boundary | 01,06 | byte-count; exclusive-limit |
| reject over limit | 02 | byte-count boundary variants |
| zero and negative limits | 03–04 | exclusive-limit; negative-unlimited |
| combining marks are separate | 05 | ignore-combining |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Additional public regression defect: reversed-comparison violates the existing public empty/boundary requirement with actual/expected diagnostics.
