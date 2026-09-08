# Author review notes

Independent content review pending. Estimates are uncalibrated.

Requirements:

- Input contains complete groups of seven nonnegative values.
- week is one-based and names an existing group.
- Return the sum of exactly those seven values; exclude both neighboring groups.
- Preserve all input values; the result fits int.

Public/reference assertions (actual/expected diagnostics):

- Case01: `if got := subject.WeekSum([]int{3,0,5,1,0,4,1,0,3,4,3,0,8,0},1); got != 14 { t.Errorf("WeekSum: got %v, want %v", got, 14) }`
- Case02: `if got := subject.WeekSum([]int{2,5,0,7,4,1,3,0,2,5,0,1,3,1},2); got != 12 { t.Errorf("WeekSum: got %v, want %v", got, 12) }`
- Case03: `if got := subject.WeekSum([]int{3,0,3,3,2,1,0},1); got != 12 { t.Errorf("WeekSum: got %v, want %v", got, 12) }`
- Case04: `input:=[]int{1,2,3,4,5,6,7};subject.WeekSum(input,1);if !slices.Equal(input,[]int{1,2,3,4,5,6,7}) {t.Errorf("input changed: %v",input)}`

Realistic required defect models:

- always-first
- all-groups
- six-values
- skip-group-first
- clear-group

Retained source behavior: bird-watcher selected operation; original narrative removed, public example values or test structure retained, new edge cases and defect models authored independently. All execution uses disposable task-check. Automated fixtures are not learner calibration.


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| selected group only | 01–02 | always-first; all-groups |
| all seven values | 01–03 | six-values; skip-group-first |
| input unchanged | 04 | clear-group |

For test-writing tasks these are reference learner-submission cases; private mutants require the learner to supply the missing checks. For implementation/debugging every basic observable requirement is publicly asserted. Constraints define the accepted input domain and do not create hidden checks outside it.
