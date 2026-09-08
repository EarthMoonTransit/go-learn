# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

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


## Requirements → public assertions → defects

| Requirement | Case | Required defect |
|---|---|---|
| one result per group/order | 01–03 | skip-empty; skip-singleton; shared-total |
| tail excludes only head | 01–02 | include-head |
| negative values included | 05 | negative-ignored |
| zero groups | 04 | no-groups-one |
| input preserved | 06 | consume-input |

For test-writing these are reference learner-submission cases; private mutants require learner assertions. Fixed-test public diagnostics identify actual/expected values. No forbidden aliasing or input outside the stated domain is graded.


## Independent review

Public checks cover distinct ordered tail sums, empty/nil/singleton groups, zero groups, negative tail elements and preserved inputs. The LGWT variadic tail-sum operation is substantively retained.

Reviewed the actual task contract, starter, public/reference assertions, all 7 required defects and route dependencies. The declared operations are substantive permitted adaptations. Exact source artifact hashes and verbatim MIT notices match the pinned local source checkouts, including Exercism authors/contributors where supplied and the literal LGWT copyright text. Private reference/mutant/review material stays outside learner payloads.

The author-provided targeted report passes, with the strong reference detecting all 7 required defects. The reviewer inspected files and reports without executing authored code or accessing the production learner home. This approves content for the lifecycle process; full batch gates, allocation reconciliation and T07 acceptance remain separate. Automated fixtures are not learner calibration.
