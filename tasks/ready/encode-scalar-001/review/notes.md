# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Accept scalar values from U+0000 through U+10FFFF, excluding U+D800 through U+DFFF.
- For a valid scalar, return its UTF-8 string, byte count and true.
- For any other rune value, return an empty string, zero and false.
- NUL and the replacement character U+FFFD are valid scalars.

Public/reference assertions (each reports actual/expected values):

- Case01: `if a, b, c := subject.EncodeScalar('A'); a != "A" || b != 1 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "A", 1, true) 	}`
- Case02: `if a, b, c := subject.EncodeScalar('é'); a != "é" || b != 2 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "é", 2, true) 	}`
- Case03: `if a, b, c := subject.EncodeScalar('中'); a != "中" || b != 3 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "中", 3, true) 	}`
- Case04: `if a, b, c := subject.EncodeScalar('🔍'); a != "🔍" || b != 4 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "🔍", 4, true) 	}`
- Case05: `if a, b, c := subject.EncodeScalar(0); a != "\x00" || b != 1 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "\x00", 1, true) 	}`
- Case06: `if a, b, c := subject.EncodeScalar(0xfffd); a != "�" || b != 3 || c != true { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "�", 3, true) 	}`
- Case07: `if a, b, c := subject.EncodeScalar(-1); a != "" || b != 0 || c != false { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false) 	}`
- Case08: `if a, b, c := subject.EncodeScalar(0xd800); a != "" || b != 0 || c != false { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false) 	}`
- Case09: `if a, b, c := subject.EncodeScalar(0x110000); a != "" || b != 0 || c != false { 		t.Errorf("EncodeScalar: got (%q,%v,%v), want (%q,%v,%v)", a, b, c, "", 0, false) 	}`

Realistic required defect models:

- surrogates-accepted
- always-one-byte
- reject-nul
- reject-replacement
- invalid-replaced

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| correct UTF-8 string and byte width | 01–04 | always-one-byte |
| invalid scalar result tuple | 07–09 | surrogates-accepted; invalid-replaced |
| valid NUL and replacement scalar | 05–06 | reject-nul; reject-replacement |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.
