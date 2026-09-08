# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Input is valid UTF-8 and replacement is a valid scalar.
- Treat index as a zero-based rune index and replace exactly that rune.
- Return the resulting UTF-8 string and true, preserving all other code points.
- For negative or out-of-range index, return the original string and false; do not alter the caller's input value.

Public/reference assertions (each reports actual/expected values):

- Case01: `if a, b := subject.EditRune("abc", 1, 'X'); a != "aXc" || b != true { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "aXc", true) 	}`
- Case02: `if a, b := subject.EditRune("Aé🔍", 1, '中'); a != "A中🔍" || b != true { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "A中🔍", true) 	}`
- Case03: `if a, b := subject.EditRune("éé", 1, 'a'); a != "éa" || b != true { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "éa", true) 	}`
- Case04: `if a, b := subject.EditRune("", 0, 'x'); a != "" || b != false { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "", false) 	}`
- Case05: `if a, b := subject.EditRune("é", 1, 'x'); a != "é" || b != false { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "é", false) 	}`
- Case06: `if a, b := subject.EditRune("ab", -1, 'x'); a != "ab" || b != false { 		t.Errorf("EditRune: got (%q,%v), want (%q,%v)", a, b, "ab", false) 	}`

Realistic required defect models:

- byte-index
- replace-all
- invalid-clears
- last-rejected

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| replace exactly one rune and preserve others | 01–03 | byte-index; replace-all; last-rejected |
| return original for invalid positions | 04–06 | invalid-clears; byte-index |
| caller string value is immutable | 01–03 | Go value contract, no additional runtime prohibition |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.
