# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Real named child tests receive distinct *testing.T identities | TestPublicCase01–02 inspect child names and parent identity | parent-callback, constant-name | child first/second/single name |
| Serial table order, once-only callbacks and unchanged input names | TestPublicCase01 | skip-last, reverse-order, double-check | serial cases calls/names |
| Empty and single tables; independent later invocations | TestPublicCase02 and grader fresh call | skip-last, double-check | single/fresh callback count |

## Instructional and technical evidence

Primary skill: `testing.subtests`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `testing.subtests`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `parent-callback`, `constant-name`, `skip-last`, `reverse-order`, `double-check`. The author-provided `.tmp/checks/named-cases-001.yaml` records technical PASS, reference PASS and 5/5 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
