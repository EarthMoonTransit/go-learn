# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Return one sum per input slice, in the order supplied.
- Exclude only the first element of each slice; include every remaining integer, including negatives.
- Nil, empty and one-element groups produce zero; no groups produce zero results.
- Preserve all input elements. Every tail sum fits int; nil or empty zero-length output is acceptable.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.SumTails([]int{1,2},[]int{0,9}); !slices.Equal(got,[]int{2,9}) {t.Errorf("SumTails: got %v, want %v",got,[]int{2,9})}`
- Case02: `if got := subject.SumTails([]int{},[]int{3,4,5}); !slices.Equal(got,[]int{0,9}) {t.Errorf("SumTails: got %v, want %v",got,[]int{0,9})}`
- Case03: `if got := subject.SumTails(nil,[]int{7}); !slices.Equal(got,[]int{0,0}) {t.Errorf("SumTails: got %v, want %v",got,[]int{0,0})}`
- Case04: `if got := subject.SumTails(); !slices.Equal(got,[]int{}) {t.Errorf("SumTails: got %v, want %v",got,[]int{})}`
- Case05: `if got := subject.SumTails([]int{9,-2,0,3}); !slices.Equal(got,[]int{1}) {t.Errorf("SumTails: got %v, want %v",got,[]int{1})}`
- Case06: `a:=[]int{1,2,3};subject.SumTails(a);if !slices.Equal(a,[]int{1,2,3}){t.Errorf("input changed: %v",a)}`

Realistic required defect models:

- include-head
- skip-empty
- skip-singleton
- shared-total
- negative-ignored
- consume-input
- no-groups-one

Retained source behavior: lgwt:arrays/v7/sum.go selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.
