# Author review notes

Independent content review approved on 2026-09-08 by /root/ui_goal_reviewer. Estimates are uncalibrated.

Requirements and public feedback:

- Input is valid UTF-8.
- Return the byte index at which each decoded rune starts, in input order.
- Include ASCII, combining code points and multibyte code points separately.
- Empty input returns a slice with zero elements; either nil or empty is accepted.

Public/reference assertions (each reports actual/expected values):

- Case01: `if got := subject.RuneOffsets(""); !slices.Equal(got, []int{}) { 		t.Errorf("RuneOffsets: got %v, want %v", got, []int{}) 	}`
- Case02: `if got := subject.RuneOffsets("abc"); !slices.Equal(got, []int{0, 1, 2}) { 		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1, 2}) 	}`
- Case03: `if got := subject.RuneOffsets("Aé🔍中"); !slices.Equal(got, []int{0, 1, 3, 7}) { 		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1, 3, 7}) 	}`
- Case04: `if got := subject.RuneOffsets("e\u0301"); !slices.Equal(got, []int{0, 1}) { 		t.Errorf("RuneOffsets: got %v, want %v", got, []int{0, 1}) 	}`

Realistic required defect models:

- rune-indices
- byte-indices
- skip-ascii
- skip-final
- empty-origin

Every listed public requirement is exercised above; confirm each defect also fails public assertions during review. Full checks add boundary inputs for the same stated behavior.


## Requirement → public assertion → defect evidence

| Requirement | Public case (reference submission for tests activity) | Defect model |
|---|---|---|
| byte offsets, ordered and complete | 02–03 | rune-indices; byte-indices; skip-final |
| ASCII and combining runes | 02–04 | skip-ascii; byte-indices |
| empty result | 01 | empty-origin |

Assertions report function name and actual/expected result. Input-domain promises are constraints, not extra hidden requirements. Technical evidence is generated only by task-check in disposable workspaces.

Additional public regression defect: empty-origin violates the existing public empty/boundary requirement with actual/expected diagnostics.
