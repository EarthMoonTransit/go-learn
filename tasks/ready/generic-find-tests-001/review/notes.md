# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-08. Content approved for normal lifecycle promotion after required batch acceptance.

Substantive adaptation from `learn-go-with-tests` at `4675d96ad50b815d3e9fdc2122ab3dbc14636252`. Reviewed the exact declared source files and checked every SHA256 against the pinned checkout: `arrays/v8/collection_fun.go` (adapted), `LICENSE.md` (copied). Registry permits adaptation; complete MIT notice and modification boundary are retained. No reference-only content or generated canonical cases used.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| First match, including matched zero; no match returns zero/false | TestContractCase01–03 | last-match, first-unchecked, false-success, zero-on-success, true-miss | first even / zero / no match |
| Ordered, once-only predicate calls; empty input and preservation | TestContractCase02–03 | double-predicate, erase-input, first-unchecked | zero/visits/preservation; empty calls |
| Any element type, including strings and non-comparable structs | TestContractCase04 and matching-string assertion | strings-skipped, structs-skipped | noncomparable / matching string |

## Instructional and technical evidence

Primary skill: `generics.type-parameters-inference`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `generics.type-parameters-inference`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `last-match`, `first-unchecked`, `false-success`, `zero-on-success`, `true-miss`, `erase-input`, `double-predicate`, `strings-skipped`, `structs-skipped`. The author-provided `.tmp/checks/generic-find-tests-001.yaml` records technical PASS, reference PASS and 9/9 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
