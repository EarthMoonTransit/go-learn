# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Write tests for placing every prefix argument before the original slice elements.
- Preserve the order of both the prefix arguments and the original values.
- With no prefix arguments, return values equal to the original input.
- Nil or empty input accepts a prefix; zero total elements may be returned as nil or an empty slice.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.PrependItems([]int{3,2,6},5); !slices.Equal(got,[]int{5,3,2,6}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{5,3,2,6})}`
- Case02: `if got := subject.PrependItems([]int{3,2,6,4,8},5,1); !slices.Equal(got,[]int{5,1,3,2,6,4,8}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{5,1,3,2,6,4,8})}`
- Case03: `if got := subject.PrependItems([]int{3,2,6}); !slices.Equal(got,[]int{3,2,6}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{3,2,6})}`
- Case04: `if got := subject.PrependItems(nil,5,1); !slices.Equal(got,[]int{5,1}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{5,1})}`
- Case05: `if got := subject.PrependItems(nil); !slices.Equal(got,[]int{}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{})}`
- Case06: `if got := subject.PrependItems([]int{},2); !slices.Equal(got,[]int{2}) {t.Errorf("PrependItems: got %v, want %v",got,[]int{2})}`

Realistic required defect models:

- suffix-instead
- reverse-prefix
- first-prefix-only
- no-prefix-clears
- empty-input-ignored
- reverse-original

Retained source behavior: card-tricks selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| prefix before original | 01 | suffix-instead |
| prefix order/all values | 02 | reverse-prefix; first-prefix-only |
| original order | 01–02 | reverse-original |
| absent prefix unchanged | 03 | no-prefix-clears |
| empty original accepts prefix | 04,06 | empty-input-ignored |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
