# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Add one to elements at indexes 0, 2, 4 and every subsequent even index.
- Leave odd-indexed values unchanged and preserve element order.
- Apply the correction to the input slice itself and return the corrected values.
- Empty or nil input has zero output elements; increments fit int.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.CorrectAlternating([]int{2,5,0,7,4,1}); !slices.Equal(got,[]int{3,5,1,7,5,1}) {t.Errorf("CorrectAlternating: got %v, want %v",got,[]int{3,5,1,7,5,1})}`
- Case02: `if got := subject.CorrectAlternating([]int{4}); !slices.Equal(got,[]int{5}) {t.Errorf("CorrectAlternating: got %v, want %v",got,[]int{5})}`
- Case03: `if got := subject.CorrectAlternating([]int{0,0,0}); !slices.Equal(got,[]int{1,0,1}) {t.Errorf("CorrectAlternating: got %v, want %v",got,[]int{1,0,1})}`
- Case04: `if got := subject.CorrectAlternating(nil); !slices.Equal(got,[]int{}) {t.Errorf("CorrectAlternating: got %v, want %v",got,[]int{})}`
- Case05: `input:=[]int{2,3,4};subject.CorrectAlternating(input);if !slices.Equal(input,[]int{3,3,5}) {t.Errorf("input got %v want [3 3 5]",input)}`

Realistic required defect models:

- odd-indexed
- range-value-copy
- copy-instead
- skip-last
- empty-placeholder

Retained source behavior: bird-watcher selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| even indexes corrected | 01–03 | odd-indexed; range-value-copy; skip-last |
| odd values/order preserved | 01 | odd-indexed |
| empty output | 04 | empty-placeholder |
| caller input corrected | 05 | copy-instead |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
