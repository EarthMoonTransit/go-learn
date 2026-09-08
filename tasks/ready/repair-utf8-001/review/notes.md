# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Read arbitrary bytes as UTF-8 units from left to right.
- Preserve every valid scalar, including encoded U+FFFD and NUL.
- Replace each malformed one-byte decoding unit with the UTF-8 encoding of U+FFFD.
- Do not combine adjacent malformed units; empty input returns an empty string.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.RepairUTF8("plain é"); got != "plain é" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "plain é") 	}`
- Case02: `if got := subject.RepairUTF8("A\xffB"); got != "A�B" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "A�B") 	}`
- Case03: `if got := subject.RepairUTF8("\xff\xfe"); got != "��" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "��") 	}`
- Case04: `if got := subject.RepairUTF8("�\x00"); got != "�\x00" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "�\x00") 	}`
- Case05: `if got := subject.RepairUTF8(""); got != "" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "") 	}`
- Case06: `if got := subject.RepairUTF8("\xe2\x82"); got != "��" { 		t.Errorf("RepairUTF8: got %q, want %q", got, "��") 	}`

Realistic required defect models:

- collapse-invalid
- drop-invalid
- return-original
- drop-replacement
- empty-replacement

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| preserve valid scalars and NUL | 01,04 | drop-replacement |
| replace invalid units | 02,06 | return-original; drop-invalid |
| keep adjacent units separate | 03,06 | collapse-invalid |
| empty string | 05 | empty-replacement |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Additional public regression defect: empty-replacement violates the existing public empty/boundary requirement with actual/expected diagnostics.
