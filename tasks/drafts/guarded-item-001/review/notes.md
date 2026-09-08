# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For a valid zero-based index, return the exact stored integer.
- The first and final indexes are valid.
- For a negative index or an index at or beyond len(values), return -1.
- Nil and empty slices have no valid index; leave the input unchanged.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.GetItem([]int{5,2,10,6,8,7,0,9},4); got != 8 { t.Errorf("GetItem: got %v, want %v", got, 8) }`
- Case02: `if got := subject.GetItem([]int{5,2,9},0); got != 5 { t.Errorf("GetItem: got %v, want %v", got, 5) }`
- Case03: `if got := subject.GetItem([]int{5,2,9},2); got != 9 { t.Errorf("GetItem: got %v, want %v", got, 9) }`
- Case04: `if got := subject.GetItem([]int{5,2,9},-1); got != -1 { t.Errorf("GetItem: got %v, want %v", got, -1) }`
- Case05: `if got := subject.GetItem([]int{5,2,9},3); got != -1 { t.Errorf("GetItem: got %v, want %v", got, -1) }`
- Case06: `if got := subject.GetItem(nil,0); got != -1 { t.Errorf("GetItem: got %v, want %v", got, -1) }`
- Case07: `input:=[]int{4,7};subject.GetItem(input,1);if !slices.Equal(input,[]int{4,7}){t.Errorf("input changed: %v",input)}`

Realistic required defect models:

- negative-clamped
- last-rejected
- one-based
- empty-zero
- upper-zero
- destructive-read

Retained source behavior: card-tricks selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| stored zero-based value | 01–03 | one-based |
| final index valid | 03 | last-rejected |
| negative/upper invalid | 04–05 | negative-clamped; upper-zero |
| empty invalid | 06 | empty-zero |
| input unchanged | 07 | destructive-read |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
