# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- For a valid index, remove exactly that element and preserve the order of the remaining values.
- Both the first and final element can be removed.
- For a negative index or an index at or beyond len(values), return unchanged values.
- Nil and empty inputs remain empty; input storage may be reused or changed for a valid removal.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.RemoveItem([]int{3,2,6,4,8},2); !slices.Equal(got,[]int{3,2,4,8}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{3,2,4,8})}`
- Case02: `if got := subject.RemoveItem([]int{3,2,6},0); !slices.Equal(got,[]int{2,6}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{2,6})}`
- Case03: `if got := subject.RemoveItem([]int{3,2,6},2); !slices.Equal(got,[]int{3,2}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{3,2})}`
- Case04: `if got := subject.RemoveItem([]int{3,2,6},-1); !slices.Equal(got,[]int{3,2,6}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{3,2,6})}`
- Case05: `if got := subject.RemoveItem([]int{3,2,6},3); !slices.Equal(got,[]int{3,2,6}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{3,2,6})}`
- Case06: `if got := subject.RemoveItem([]int{3,2,6},10); !slices.Equal(got,[]int{3,2,6}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{3,2,6})}`
- Case07: `if got := subject.RemoveItem(nil,0); !slices.Equal(got,[]int{}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{})}`
- Case08: `if got := subject.RemoveItem([]int{},-1); !slices.Equal(got,[]int{}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{})}`
- Case09: `if got := subject.RemoveItem([]int{7},0); !slices.Equal(got,[]int{}) {t.Errorf("RemoveItem: got %v, want %v",got,[]int{})}`

Realistic required defect models:

- first-unchanged
- last-unchanged
- negative-clamped
- upper-clamped
- empty-placeholder
- remove-last-instead

Retained source behavior: card-tricks selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| interior removal and order | 01 | remove-last-instead |
| first/final removal | 02–03 | first-unchanged; last-unchanged |
| negative/out-of-range unchanged | 04–06 | negative-clamped; upper-clamped |
| nil/empty and singleton result | 07–09 | empty-placeholder; last-unchanged |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
