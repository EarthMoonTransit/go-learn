# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- For a valid index, replace exactly that element and return the resulting slice.
- Valid-index replacement must modify the input storage so existing views observe the new element.
- For either an invalid negative index or an index at or beyond length, append value to the end.
- Preserve the order and values of other elements; nil and empty input accept an appended value.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.SetItem([]int{1,2,4,1},2,6); !slices.Equal(got,[]int{1,2,6,1}) {t.Errorf("SetItem: got %v, want %v",got,[]int{1,2,6,1})}`
- Case02: `if got := subject.SetItem([]int{1,2,4,1},-1,6); !slices.Equal(got,[]int{1,2,4,1,6}) {t.Errorf("SetItem: got %v, want %v",got,[]int{1,2,4,1,6})}`
- Case03: `if got := subject.SetItem([]int{1,2},2,7); !slices.Equal(got,[]int{1,2,7}) {t.Errorf("SetItem: got %v, want %v",got,[]int{1,2,7})}`
- Case04: `if got := subject.SetItem(nil,0,7); !slices.Equal(got,[]int{7}) {t.Errorf("SetItem: got %v, want %v",got,[]int{7})}`
- Case05: `if got := subject.SetItem([]int{1},0,9); !slices.Equal(got,[]int{9}) {t.Errorf("SetItem: got %v, want %v",got,[]int{9})}`
- Case06: `input:=[]int{1,2,3};view:=input[1:];subject.SetItem(input,1,9);if view[0]!=9 {t.Errorf("shared element got %d want 9",view[0])}`

Realistic required defect models:

- copy-valid
- ignore-negative
- ignore-upper
- always-append
- empty-ignored

Retained source behavior: card-tricks selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| valid replacement, other values/order | 01,05 | always-append |
| negative append | 02 | ignore-negative |
| upper append | 03 | ignore-upper |
| nil append | 04 | empty-ignored |
| shared storage mutation | 06 | copy-valid |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.


## Independent review

The public cases distinguish valid replacement through shared storage, invalid-index append, order preservation and nil input.

Reviewed the English prompt, prerequisite placement, public/reference assertions, all 5 required defects, stable API and learner payload boundaries. The adaptation preserves a substantive permitted source operation rather than only its topic; exact pinned artifact hashes, authors/contributors and MIT notice contents match the local Exercism checkout. No reference-only or generated source artifact is claimed as adapted lineage.

The author-provided targeted task-check report is technically passing with all 5 required defects detected. The reviewer inspected files and reports without executing authored code. This approves content for the lifecycle process; full batch gates remain the author's responsibility. Automated checks are not learner calibration.
