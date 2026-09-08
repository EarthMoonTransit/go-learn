# Independent content review

Reviewed by `/root/ui_goal_reviewer` on 2026-09-09. Content approved for normal lifecycle promotion after required batch acceptance.

Original task: ADR 0008 project identity, creation record, content digest and empty external sources verified. No invented source commit; release binding remains deferred.

## Requirement and feedback mapping

For this test-writing profile, the public run checks learner tests against the correct implementation; full verification challenges those tests with defects. The assertions below are the strong learner reference suite. The starter is intentionally insufficient.

| Stated requirement | Public/reference assertion | Required defect or acceptance probe | Useful diagnostic |
|---|---|---|---|
| Exact six-byte big-endian representation and full unsigned ranges | TestContractCase01–02 | little-endian, drop-high-count | wire hex; decoded record; boundary |
| Successful unmarshal replaces an already populated receiver | TestReplacesPopulatedReceiver | merge-receiver | replacement want{1 2},nil |
| Every invalid length preserves receiver and original sentinel | TestContractCase03 | accept-trailing, clear-on-error, fresh-error | invalid length with record/error |
| Input preservation and independent marshal buffers | TestContractCase02 and Case04 | mutate-input, shared-buffer | decoder changed input; outputs share storage |

## Instructional and technical evidence

Primary skill: `encoding.binary-marshaling`. Declared prerequisites resolve before this task at its actual numbered route position. New concepts: `encoding.binary-marshaling`. Short technical English prompt, minimal implementation starter or intentionally weak learner tests, no decorative narrative. Tests assess observable behavior rather than allocation, helper naming or table organization.

Required defects reviewed: `little-endian`, `drop-high-count`, `accept-trailing`, `clear-on-error`, `mutate-input`, `fresh-error`, `shared-buffer`, `merge-receiver`. The author-provided `.tmp/checks/binary-record-tests-001.yaml` records technical PASS, reference PASS and 8/8 detected required defects. For mutation-tests, this count describes the strong reference suite; starter-full PASS is the checker's insufficiency gate, not evidence that the starter kills every defect.

No authored code or task tests were run from the checkout by this reviewer. No real learner home was accessed. Full collection, ordered acceptance gates, reports, PR validation and lifecycle movement remain the author's separate responsibilities. Automated fixtures and estimates do not establish human calibration or learner mastery.
