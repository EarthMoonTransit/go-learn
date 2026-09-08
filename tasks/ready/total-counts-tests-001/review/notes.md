# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

Requirements:

- Write tests for the sum of all values, including the first and final elements.
- Nil and empty inputs return zero; zero-valued elements do not stop processing.
- A single element returns its value.
- Do not change input elements; values are nonnegative and the total fits int.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.TotalCounts([]int{9,0,8,4,5,1,3}); got != 30 { t.Errorf("TotalCounts: got %v, want %v", got, 30) }`
- Case02: `if got := subject.TotalCounts(nil); got != 0 { t.Errorf("TotalCounts: got %v, want %v", got, 0) }`
- Case03: `if got := subject.TotalCounts([]int{}); got != 0 { t.Errorf("TotalCounts: got %v, want %v", got, 0) }`
- Case04: `if got := subject.TotalCounts([]int{2}); got != 2 { t.Errorf("TotalCounts: got %v, want %v", got, 2) }`
- Case05: `if got := subject.TotalCounts([]int{4,0,7}); got != 11 { t.Errorf("TotalCounts: got %v, want %v", got, 11) }`
- Case06: `input:=[]int{3,1,4};subject.TotalCounts(input);if !slices.Equal(input,[]int{3,1,4}){t.Errorf("input changed: got %v want [3 1 4]",input)}`

Realistic required defect models:

- skip-first
- skip-last
- stop-at-zero
- empty-one
- singleton-zero
- consume-input

Retained source behavior: bird-watcher selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| all elements including ends | 01 | skip-first; skip-last |
| empty/nil zero | 02–03 | empty-one |
| singleton value | 04 | singleton-zero |
| zero does not terminate | 05 | stop-at-zero |
| input unchanged | 06 | consume-input |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.


## Independent review

Required defects distinguish empty and singleton results, first/final elements, processing past zero and input preservation; the supplied narrow starter does not establish a complete suite.

Reviewed the English prompt, prerequisite placement, public/reference assertions, all 6 required defects, stable API and learner payload boundaries. The adaptation preserves a substantive permitted source operation rather than only its topic; exact pinned artifact hashes, authors/contributors and MIT notice contents match the local Exercism checkout. No reference-only or generated source artifact is claimed as adapted lineage.

The author-provided targeted task-check report is technically passing with all 6 required defects detected. The reviewer inspected files and reports without executing authored code. This approves content for the lifecycle process; full batch gates remain the author's responsibility. Automated checks are not learner calibration.
