# Author review notes

Independent content review approved by `/root/ui_goal_reviewer` on 2026-09-08. Estimates remain uncalibrated.

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


## Independent review

Required defects distinguish prefix order, multiple prefix arguments, preserved original order, no-prefix input and empty original input.

Reviewed the English prompt, prerequisite placement, public/reference assertions, all 6 required defects, stable API and learner payload boundaries. The adaptation preserves a substantive permitted source operation rather than only its topic; exact pinned artifact hashes, authors/contributors and MIT notice contents match the local Exercism checkout. No reference-only or generated source artifact is claimed as adapted lineage.

The author-provided targeted task-check report is technically passing with all 6 required defects detected. The reviewer inspected files and reports without executing authored code. This approves content for the lifecycle process; full batch gates remain the author's responsibility. Automated checks are not learner calibration.


## E02 variadic bridge, 2026-10-17

The starter now supplies TestStartingCase01 with separate arguments and
TestPrependExpandedPrefix with prefix... expansion. Both named examples assert
prefix-before-original order and multiple prefix values with actual/want slices.
The starter passes the reference but intentionally leaves no-prefix-clears and
empty-input-ignored alive; learners must still add missing contract cases.
Hints explain both calls without giving an implementation or new ownership,
allocation or nilness guarantees. TASK, API, reference and six mutants are
unchanged. Slice header/slicing prerequisites describe already visible syntax.

The exact old→new digest is admitted in ADR 0027. Compatibility tooling tests
paused/completed, generation/legacy marker, numbered/legacy paths; old editable
and read-only files and history remain byte-identical. Disposable integration
checks require both named example diagnostics, old/new starter incompleteness,
reference detection of all six defects and unchanged primary completion credit.
Independent content review of this actual package is pending.

Independent content/provenance approval by /root/clamp_e01_content_review on 2026-10-17 covers the final package, including the GCD rejection/tolerance fix where applicable. No further content blockers. Normal lifecycle promotion, final targeted checks, full ordered repository acceptance and exact-head CI remain separate gates.
