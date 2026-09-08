# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Exactly one object with permitted fields and no extra value/suffix | TestPublicCase01–02 | allow-unknown, ignore-trailing, null-as-zero | decode/reject quoted input |
| Rejection returns zero record and original non-nil sentinel | TestPublicCase02; TestDiscardPartialObject | fresh-error, partial-on-reject | wantzero,original; partial object |
| Standard duplicate keys, case-folded fields and field-level null, while rejecting top-level null | TestPublicCase01; TestStandardFieldMatchingAndNull | null-as-zero and valid-alternative acceptance cases | standard fields got/want |
| Caller owns reader on success and rejection | TestPublicCase03 | close-reader | success/failure closes |

## Instructional and technical evidence

Primary skill: `encoding.streaming-strictness`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `encoding.streaming-strictness`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `allow-unknown`, `ignore-trailing`, `null-as-zero`, `partial-on-reject`, `fresh-error`, `close-reader`. The author-provided `.tmp/checks/single-json-value-001.yaml` records technical PASS, reference PASS and 6/6 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
