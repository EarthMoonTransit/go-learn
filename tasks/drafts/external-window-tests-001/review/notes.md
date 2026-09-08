# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| External package uses exported constructor and method | package subject_test strong/starter suites and mutation-profile source inspection | Private-field access cannot compile from the external package | normal external-package compiler diagnostics |
| Half-open bounds including negative and empty windows | TestContractCase01–02 | include-upper, exclude-lower, wrong-lower, empty-is-point, swapped-bounds | window[min,max) value got/want |
| Independent windows retain bounds | TestContractCase03 | Explicit interleaved-constructor acceptance assertions | independent windows; empty constructor changed existing window |

## Instructional and technical evidence

Primary skill: `packages.test-package-boundaries`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `packages.test-package-boundaries`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `include-upper`, `exclude-lower`, `wrong-lower`, `empty-is-point`, `swapped-bounds`. The author-provided `.tmp/checks/external-window-tests-001.yaml` records technical PASS, reference PASS and 5/5 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
