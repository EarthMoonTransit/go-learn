# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-08. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Extract original pointer, not a reconstruction | TestContractCase01–02 | copy-detail, false-found | direct; first cause |
| Wrapped/joined first-match traversal and custom As | TestContractCase02 and Case04 | direct-only, joined-order-reversed, custom-as-ignored | first cause; custom As |
| No match/nil differs from typed-nil match | TestContractCase03 | fabricate-miss, nil-is-miss, false-found | miss; typed nil |

## Instructional and technical evidence

Primary skill: `errors.as`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `errors.as`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `direct-only`, `copy-detail`, `false-found`, `nil-is-miss`, `fabricate-miss`, `custom-as-ignored`, `joined-order-reversed`. The author-provided `.tmp/checks/error-detail-tests-001.yaml` records technical PASS, reference PASS and 7/7 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
