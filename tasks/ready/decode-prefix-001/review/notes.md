# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Decode only the first UTF-8 unit; later bytes do not affect validity of that unit.
- For a valid leading scalar return the rune, its consumed byte count and true.
- For malformed leading bytes return U+FFFD, one and false.
- For empty input return zero, zero and false; encoded U+FFFD is valid and consumes three bytes.

Public/reference assertions (each reports actual/expected values):

- Case01: `if a, b, c := subject.DecodePrefix(""); a != 0 || b != 0 || c != false { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0, 0, false) 	}`
- Case02: `if a, b, c := subject.DecodePrefix("A\xff"); a != 'A' || b != 1 || c != true { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 'A', 1, true) 	}`
- Case03: `if a, b, c := subject.DecodePrefix("éx"); a != 'é' || b != 2 || c != true { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 'é', 2, true) 	}`
- Case04: `if a, b, c := subject.DecodePrefix("🔍"); a != '🔍' || b != 4 || c != true { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, '🔍', 4, true) 	}`
- Case05: `if a, b, c := subject.DecodePrefix("\xffx"); a != 0xfffd || b != 1 || c != false { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 1, false) 	}`
- Case06: `if a, b, c := subject.DecodePrefix("�"); a != 0xfffd || b != 3 || c != true { 		t.Errorf("DecodePrefix: got (%U,%v,%v), want (%U,%v,%v)", a, b, c, 0xfffd, 3, true) 	}`

Realistic required defect models:

- replacement-invalid
- validate-whole-string
- decoded-width
- empty-is-replacement
- one-byte-width

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| first unit only | 02 | validate-whole-string |
| valid width and rune | 03–04 | decoded-width expands same decoding distinction |
| malformed unit consumes one byte | 05 | decoded-width |
| empty and encoded replacement distinguished | 01,06 | empty-is-replacement; replacement-invalid |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Additional public regression defect: one-byte-width violates the existing public empty/boundary requirement with actual/expected diagnostics.
