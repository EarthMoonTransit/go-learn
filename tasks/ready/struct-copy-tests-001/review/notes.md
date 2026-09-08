# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-08. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Ordinary scalar value copy and preserved input values | TestContractCase01 | clear-count, trim-label, destructive-items | copy; scalar values changed |
| Shared slice backing storage in both directions | TestContractCase01 | deep-copy | slice backing storage not shared; reverse sharing missing |
| Preserve nil/allocated-empty, length and capacity | TestContractCase02 | drop-capacity, nil-to-empty | slice header got nil/len/cap … want … |

## Instructional and technical evidence

Primary skill: `structs.copy-equality`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: none. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `deep-copy`, `clear-count`, `trim-label`, `drop-capacity`, `nil-to-empty`, `destructive-items`. The author-provided `.tmp/checks/struct-copy-tests-001.yaml` records technical PASS, reference PASS and 6/6 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
