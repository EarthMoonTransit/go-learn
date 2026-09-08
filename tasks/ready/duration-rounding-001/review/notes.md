# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Truncate positive remainder and use whole-minute units | TestPublicCase01 | round-nearest, seconds-as-minutes | duration got/want |
| Inclusive zero/24h range; reject negative and excessive duration | TestPublicCase01–02 | exclude-day, negative-accepted, swallow-invalid | invalid duration; endpoint values |
| Preserve original non-nil error sentinel | TestPublicCase02 with package-level originalSentinel | fresh-error | want0,original |

## Instructional and technical evidence

Primary skill: `time.duration-arithmetic`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `time.duration-arithmetic`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `round-nearest`, `seconds-as-minutes`, `exclude-day`, `negative-accepted`, `fresh-error`, `swallow-invalid`. The author-provided `.tmp/checks/duration-rounding-001.yaml` records technical PASS, reference PASS and 6/6 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
